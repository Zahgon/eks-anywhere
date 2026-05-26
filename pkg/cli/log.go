package cli

import (
	"github.com/go-logr/logr"
)

const (
	markPass   = "✅ "
	markFailed = "❌ "
)

// ValidationPassed logs a success message for a validation.
func ValidationPassed(log logr.Logger, msg string) { _ = "STUB: not implemented"; return }

// ValidationFailed logs an error message for a validation.
func ValidationFailed(log logr.Logger, msg string) { _ = "STUB: not implemented"; return }
