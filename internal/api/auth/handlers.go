package auth

import (
	"net/http"
	"time"

	"jellyfin-hls-server/internal/config"
	usersrepo "jellyfin-hls-server/internal/repo/users"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthHandler struct {
	config   *config.Config
	userRepo *usersrepo.UserRepo
}

func NewAuthHandler(userRepo *usersrepo.UserRepo, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
		config:   cfg,
	}
}

func (h *AuthHandler) UserPublicHandler(c *gin.Context) {
	// This handler is for public user information, such as registration or public profile.
	// It can be used to register a new user or fetch public user details.
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, "[]")
}

// AuthenticateUser handles the POST /users/authenticate endpoint.
func (h *AuthHandler) AuthenticateUser(c *gin.Context) {
	var req struct {
		Username string `json:"Username"`
		Pw       string `json:"Pw"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Username == "admin" && req.Pw == "admin" {
		userID := uuid.New().String()

		// Create JWT token with username in payload
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, err := token.SignedString([]byte(h.config.JwtSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, CreateAuthenticateUserResponse(tokenString, userID, h.config))
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// GetCurrentUser handles the GET /users/me endpoint.
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	// TODO: Implement logic to get current authenticated user
	// - Extract access token from Authorization header
	// - Validate access token
	// - Retrieve user information based on token

	c.JSON(http.StatusOK, gin.H{
		"id":           "placeholder-user-id",
		"name":         "Placeholder User",
		"primaryImage": "placeholder-image-url",
	})
}
