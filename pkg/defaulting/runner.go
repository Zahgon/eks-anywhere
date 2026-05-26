package defaulting

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/errors"
)

// Default is the logic for a default for a type O. It should return a value of O
// whether it updates it or not. When there is an error, return the zero value of O
// and the error.
type Default[O any] func(ctx context.Context, obj O) (O, error)

// Runner allows to compose and run validations/defaults.
type Runner[O any] struct {
	defaults []Default[O]
}

// NewRunner constructs a new Runner.
func NewRunner[O any]() *Runner[O] { _ = "STUB: not implemented"; return nil }

// Register adds defaults to the Runner.
func (r *Runner[O]) Register(defaults ...Default[O]) { _ = "STUB: not implemented"; return }

// RunAll runs all defaults sequentially and returns the updated O. When there are errors,
// it returns the zero value of O and the aggregated errors.
func (r *Runner[O]) RunAll(ctx context.Context, obj O) (O, errors.Aggregate) {
	_ = "STUB: not implemented"
	return *new(O), *new(errors.Aggregate)
}

// flatten unfolds and flattens errors inside a errors.Aggregate. If err is not
// a errors.Aggregate, it just returns a slice with one single error.
func flatten(err error) []error { _ = "STUB: not implemented"; return nil }
