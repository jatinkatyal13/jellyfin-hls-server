package systemhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SystemHandler handles system-related endpoints
type SystemHandler struct{}

// NewSystemHandler creates a new SystemHandler
func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

// GetSystemInfoHandler handles the GET /system/info endpoint
func (h *SystemHandler) GetSystemInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"LocalAddress":           "http://127.0.0.1:8096",
		"ServerName":             "lima-rancher-desktop",
		"Version":                "10.10.7",
		"ProductName":            "Jellyfin Server",
		"OperatingSystem":        "",
		"Id":                     "7632a9ef559f4496bbc3fda569d320c7",
		"StartupWizardCompleted": false,
	})
}
