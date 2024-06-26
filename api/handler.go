package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
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

	// Forward alert to Discord asynchronously
	go func() {
		if err := sendAlertToDiscord(payload); err != nil {
			// Log the error instead of modifying the response
			log.Printf("Failed to send alert to Discord: %v", err)
		}
	}()

	// Process the valid payload
	c.JSON(http.StatusOK, gin.H{"message": "Processed successfully"})
}

func sendAlertToDiscord(payload model.GrafanaWebhookPayload) error {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	message := formatMessageForDiscord(payload)
	content := map[string]string{"content": message}
	jsonData, err := json.Marshal(content)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func formatMessageForDiscord(payload model.GrafanaWebhookPayload) string {
	// Customize this function based on how you want the message to appear in Discord
	return "Alert received: " + payload.Title
}
