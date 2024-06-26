package api

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/Juxsta/grafana-consumer/model" // Corrected import path

	"github.com/gin-gonic/gin"
)

func GrafanaAlertHandler(c *gin.Context) {
	var payload model.GrafanaWebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Iterate through alerts
	for _, alert := range payload.Alerts {
		if alert.Labels["alertname"] == "Container Health" && strings.Contains(alert.Annotations["description"], "qbittorrent") {
			// Restart the qbittorrent container
			cmd := exec.Command("docker", "restart", "qbittorrent")
			if err := cmd.Run(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restart qbittorrent"})
				return
			}
		}
	}

	// Process the valid payload
	c.JSON(http.StatusOK, gin.H{"message": "Processed successfully"})
}
