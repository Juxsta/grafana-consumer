package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidateGrafanaAlertPayload(t *testing.T) {
	tests := []struct {
		name    string
		payload GrafanaAlertPayload
		wantErr bool
	}{
		{
			name: "valid payload",
			payload: GrafanaAlertPayload{
				Title:       "Alert Title",
				Message:     "Alert Message",
				OrgID:       1,
				DashboardID: 1,
				PanelID:     1,
				Tags:        []string{"tag1", "tag2"},
				RuleID:      1,
				RuleName:    "Rule Name",
				RuleURL:     "https://example.com/rule",
				State:       "alerting",
				EvalMatches: []EvalMatch{{Value: 10.5, Metric: "cpu_usage"}},
				ImageURL:    "https://example.com/image.jpg",
				Timestamp:   time.Now(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.payload.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateEvalMatch(t *testing.T) {
	tests := []struct {
		name    string
		match   EvalMatch
		wantErr bool
	}{
		{
			name: "valid match",
			match: EvalMatch{
				Value:  10.5,
				Metric: "cpu_usage",
			},
			wantErr: false,
		},
		{
			name: "missing metric",
			match: EvalMatch{
				Value: 10.5,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.match.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
