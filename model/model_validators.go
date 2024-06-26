package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateGrafanaAlertPayload validates the fields of GrafanaAlertPayload
func (g GrafanaAlertPayload) Validate() error {
	// Validate struct based on tags
	err := validate.Struct(g)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

// ValidateEvalMatch validates the fields of EvalMatch
func (e EvalMatch) Validate() error {
	if e.Metric == "" {
		return fmt.Errorf("metric field is required")
	}
	return nil
}
