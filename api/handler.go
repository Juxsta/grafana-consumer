package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

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
	// Start building the message
	var messageBuilder strings.Builder

	// Add a title and a link to the external URL if available
	messageBuilder.WriteString(fmt.Sprintf("**%s**\n", payload.Title))
	if payload.ExternalURL != "" {
		messageBuilder.WriteString(fmt.Sprintf("[View Details](%s)\n", payload.ExternalURL))
	}

	// Add status and message
	messageBuilder.WriteString(fmt.Sprintf("Status: **%s**\n", payload.Status))
	messageBuilder.WriteString(fmt.Sprintf("Message: **%s**\n", payload.Message))

	// Loop through each alert and add detailed information
	for _, alert := range payload.Alerts {
		messageBuilder.WriteString("\n---\n")
		messageBuilder.WriteString(fmt.Sprintf("Alert: **%s**\n", alert.Labels["alertname"]))
		messageBuilder.WriteString(fmt.Sprintf("Severity: **%s**\n", alert.Labels["severity"]))
		messageBuilder.WriteString(fmt.Sprintf("Starts At: **%s**\n", alert.StartsAt.Format(time.RFC1123)))
		messageBuilder.WriteString(fmt.Sprintf("Ends At: **%s**\n", alert.EndsAt.Format(time.RFC1123)))
		messageBuilder.WriteString(fmt.Sprintf("Description: **%s**\n", alert.Annotations["description"]))

		// Include URLs if available
		if alert.DashboardURL != "" {
			messageBuilder.WriteString(fmt.Sprintf("[Dashboard](%s)\n", alert.DashboardURL))
		}
		if alert.PanelURL != "" {
			messageBuilder.WriteString(fmt.Sprintf("[Panel](%s)\n", alert.PanelURL))
		}
		if alert.ImageURL != "" {
			messageBuilder.WriteString(fmt.Sprintf("![Image](%s)\n", alert.ImageURL))
		}
	}

	// Return the constructed message
	return messageBuilder.String()
}
