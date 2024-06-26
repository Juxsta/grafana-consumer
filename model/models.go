package model

import "time"

type GrafanaWebhookPayload struct {
	Receiver          string            `json:"receiver" validate:"required"`
	Status            string            `json:"status" validate:"required,oneof=firing resolved"`
	OrgID             int64             `json:"orgId" validate:"required"`
	Alerts            []GrafanaAlert    `json:"alerts" validate:"required,dive"`
	GroupLabels       map[string]string `json:"groupLabels" validate:"required"`
	CommonLabels      map[string]string `json:"commonLabels" validate:"required"`
	CommonAnnotations map[string]string `json:"commonAnnotations" validate:"required"`
	ExternalURL       string            `json:"externalURL" validate:"required,url"`
	Version           string            `json:"version" validate:"required"`
	GroupKey          string            `json:"groupKey" validate:"required"`
	TruncatedAlerts   int               `json:"truncatedAlerts" validate:"min=0"`
	Title             string            `json:"title" validate:"required"`
	State             string            `json:"state" validate:"required"`
	Message           string            `json:"message" validate:"required"`
}

// GrafanaAlert struct with validation tags
type GrafanaAlert struct {
	Status       string             `json:"status" validate:"required,oneof=firing resolved"`
	Labels       map[string]string  `json:"labels" validate:"required"`
	Annotations  map[string]string  `json:"annotations" validate:"required"`
	StartsAt     time.Time          `json:"startsAt" validate:"required"`
	EndsAt       time.Time          `json:"endsAt" validate:"required"`
	GeneratorURL string             `json:"generatorURL" validate:"required,url"`
	Fingerprint  string             `json:"fingerprint" validate:"required"`
	SilenceURL   string             `json:"silenceURL" validate:"required,url"`
	DashboardURL string             `json:"dashboardURL" validate:"omitempty,url"`
	PanelURL     string             `json:"panelURL" validate:"omitempty,url"`
	ImageURL     string             `json:"imageURL" validate:"omitempty,url"`
	Values       map[string]float64 `json:"values" validate:"required"`
}
