package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// GrafanaWebhookPayload struct with validation tags

// ValidateGrafanaAlertPayload validates the fields of GrafanaWebhookPayload using struct tags
func ValidateGrafanaAlertPayload(g GrafanaWebhookPayload) error {
	err := validate.Struct(g)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}
