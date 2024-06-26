package model

import "time"

type GrafanaAlertPayload struct {
	Title       string      `json:"title" binding:"required"`
	Message     string      `json:"message" binding:"required"`
	OrgID       int64       `json:"orgId" binding:"required"`
	DashboardID int64       `json:"dashboardId"`
	PanelID     int64       `json:"panelId"`
	Tags        []string    `json:"tags"`
	RuleID      int64       `json:"ruleId" binding:"required"`
	RuleName    string      `json:"ruleName" binding:"required"`
	RuleURL     string      `json:"ruleUrl" binding:"required,url"`
	State       string      `json:"state" binding:"required,oneof=ok alerting no_data"`
	EvalMatches []EvalMatch `json:"evalMatches"`
	ImageURL    string      `json:"imageUrl" binding:"omitempty,url"`
	Timestamp   time.Time   `json:"timestamp" binding:"required"`
}

type EvalMatch struct {
	Value  float64           `json:"value" binding:"required"`
	Metric string            `json:"metric" binding:"required"`
	Tags   map[string]string `json:"tags"`
}
