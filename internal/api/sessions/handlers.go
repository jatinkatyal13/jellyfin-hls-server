package sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PlayingHandler handles the POST /sessions/playing endpoint.
func PlayingHandler(c *gin.Context) {
	// TODO: Implement logic to handle playback progress and session keepalive.
	// This will likely involve receiving a request body with item ID, position, etc.
	// For now, return a placeholder success response.

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}