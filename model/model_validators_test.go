package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Define execCommand as a variable that can be overridden for testing

func TestValidateGrafanaAlertPayload(t *testing.T) {
	tests := []struct {
		name    string
		payload GrafanaWebhookPayload
		wantErr bool
	}{
		{
			name: "valid payload",
			payload: GrafanaWebhookPayload{
				Receiver:          "My Super Webhook",
				Status:            "firing",
				OrgID:             1,
				Alerts:            []GrafanaAlert{{Status: "firing", Labels: map[string]string{"alertname": "High memory usage"}, Annotations: map[string]string{"description": "High memory usage detected"}, StartsAt: time.Now(), EndsAt: time.Now().Add(1 * time.Hour), GeneratorURL: "http://example.com", Fingerprint: "abc123", SilenceURL: "http://example.com/silence", DashboardURL: "http://example.com/dashboard", PanelURL: "http://example.com/panel", ImageURL: "http://example.com/image", Values: map[string]float64{"A": 100}}},
				GroupLabels:       map[string]string{"team": "blue"},
				CommonLabels:      map[string]string{"severity": "high"},
				CommonAnnotations: map[string]string{"summary": "Memory usage alert"},
				ExternalURL:       "http://example.com",
				Version:           "1",
				GroupKey:          "alertgroup123",
				TruncatedAlerts:   0,
				Title:             "Alert Title",
				State:             "alerting",
				Message:           "Alert Message",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGrafanaAlertPayload(tt.payload)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
