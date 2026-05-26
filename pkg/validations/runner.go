package validations

import (
	"errors"
)

var errRunnerValidation = errors.New("validations failed")

type Validation func() *ValidationResult

type Runner struct {
	validations []Validation
}

func NewRunner() *Runner { _ = "STUB: not implemented"; return nil }

func (r *Runner) Register(validations ...Validation) { _ = "STUB: not implemented"; return }

func (r *Runner) Run() error { _ = "STUB: not implemented"; return nil }
