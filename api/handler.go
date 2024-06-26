package api

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/Juxsta/grafana-consumer/model" // Corrected import path

	"github.com/gin-gonic/gin"
)

func GrafanaAlertHandler(c *gin.Context) {
	var payload model.GrafanaAlertPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for specific rule and container
	if payload.RuleName == "Container Health" && strings.Contains(payload.Message, "container_name=qbittorrent") {
		// Restart the qbittorrent container
		cmd := exec.Command("docker", "restart", "qbittorrent")
		if err := cmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restart qbittorrent"})
			return
		}
	}

	// Process the valid payload
	// ...
	c.JSON(http.StatusOK, gin.H{"message": "Processed successfully"})
}
