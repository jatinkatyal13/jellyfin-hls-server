package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jellyfin-hls-server/internal/repo/users"
)

type AuthHandler struct {
	userRepo *usersrepo.UserRepo
}

func NewAuthHandler (userRepo *usersrepo.UserRepo) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

// AuthenticateUser handles the POST /users/authenticate endpoint.
func (h *AuthHandler) AuthenticateUser(c *gin.Context) {
	// TODO: Implement user authentication logic here
	// - Parse request body for username and password
	// - Validate credentials
	// - Generate an access token

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":   "placeholder-user-id",
			"name": "Placeholder User",
		},
		"accessToken": "placeholder-access-token",
	})
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
