package validation

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/errors"
)

// Validatable is anything that can be validated.
type Validatable[O any] interface {
	DeepCopy() O
}

// Validation is the logic for a validation of a type O.
type Validation[O Validatable[O]] func(ctx context.Context, obj O) error

// Runner allows to compose and run validations.
type Runner[O Validatable[O]] struct {
	validations []Validation[O]
	config      *RunnerConfig
}

// RunnerConfig contains the configuration for a Runner.
type RunnerConfig struct {
	maxJobs int
}

// RunnerOpt allows to configure a Runner with optional parameters.
type RunnerOpt func(*RunnerConfig)

// WithMaxJobs sets the maximun number of concurrent routines the runner will use.
func WithMaxJobs(m int) RunnerOpt { _ = "STUB: not implemented"; return *new(RunnerOpt) }

// NewRunner constructs a new Runner.
func NewRunner[O Validatable[O]](opts ...RunnerOpt) *Runner[O] {
	_ = "STUB: not implemented"
	return nil
}

// Register adds validations to the Runner.
func (r *Runner[O]) Register(validations ...Validation[O]) { _ = "STUB: not implemented"; return }

// RunAll runs all validations concurrently and waits until they all finish,
// aggregating the errors if present. obj must not be modified. If it is, this
// indicates a programming error and the method will panic.
func (r *Runner[O]) RunAll(ctx context.Context, obj O) errors.Aggregate {
	_ = "STUB: not implemented"
	return *new(errors.Aggregate)
}

func (r *Runner[O]) run(ctx context.Context, obj O) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

// Sequentially composes a set of validations into one which will run them sequentially and in order.
func Sequentially[O Validatable[O]](validations ...Validation[O]) Validation[O] {
	_ = "STUB: not implemented"
	return nil
}

// flatten unfolds and flattens errors inside a errors.Aggregate. If err is not
// a errors.Aggregate, it just returns a slice with one single error.
func flatten(err error) []error { _ = "STUB: not implemented"; return nil }
