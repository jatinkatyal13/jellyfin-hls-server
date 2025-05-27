package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	authHandler "jellyfin-hls-server/internal/api/auth"
	contentHandler "jellyfin-hls-server/internal/api/content"
	systemhandler "jellyfin-hls-server/internal/api/system"
	"jellyfin-hls-server/internal/config"
	"jellyfin-hls-server/internal/db"
	repoContent "jellyfin-hls-server/internal/repo/content"
	repoUsers "jellyfin-hls-server/internal/repo/users"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the database connection
	dbConn, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer dbConn.Close()

	// Initialize repositories
	contentRepoInstance := repoContent.NewContentRepo(dbConn.Db.DB)
	userRepoInstance := repoUsers.NewUserRepo(dbConn.Db.DB)

	// Setup router and routes
	router := initRouter(userRepoInstance, contentRepoInstance, cfg)

	// Start the server with graceful shutdown
	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func initRouter(userRepoInstance *repoUsers.UserRepo, contentRepoInstance *repoContent.ContentRepo, cfg *config.Config) *gin.Engine {
	var router *gin.Engine
	if cfg.Debug {
		router = gin.Default() // includes logger and recovery middleware
	} else {
		router = gin.New() // no logger by default
	}

	// System Info
	system := systemhandler.NewSystemHandler()
	systemGroup := router.Group("/system")
	{
		systemGroup.GET("/info/public", system.GetSystemInfoHandler)
	}

	// Authentication
	authGroup := router.Group("/users")
	authGroupUpper := router.Group("/Users") // Uppercase group for compatibility
	auth := authHandler.NewAuthHandler(userRepoInstance, cfg)
	{
		// Pass userRepo to auth handlers
		authGroup.GET("/public", auth.UserPublicHandler)
		authGroup.POST("/authenticatebyname", auth.AuthenticateUser)
		authGroup.GET("/me", auth.GetCurrentUser)

		authGroupUpper.GET("/public", auth.UserPublicHandler)
		authGroupUpper.POST("/authenticatebyname", auth.AuthenticateUser)
		authGroupUpper.GET("/me", auth.GetCurrentUser)
	}

	content := contentHandler.NewContentHandler(contentRepoInstance)

	// Pass contentRepo to content handlers
	// User Items
	// usersGroup := router.Group("/users")
	{
		// usersGroup.GET("/:userId/items", content.GetUserItemsHandler)
		// usersGroup.GET("/:userId/items/:itemId", content.GetUserItemHandler)
	}

	// Item Playback and Images
	itemsGroup := router.Group("/items")
	{
		// itemsGroup.GET("/:itemId/playbackinfo", content.GetItemPlaybackInfoHandler)
		itemsGroup.GET("/:itemId/images/:imageType", content.GetItemImageHandler)
	}

	// Search (No group needed for now)
	// router.GET("/search/hints", content.GetSearchHintsHandler)

	return router
}
