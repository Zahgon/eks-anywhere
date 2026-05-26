package createvalidations

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/validations"
)

// PreflightValidations returns the validations required before creating a cluster.
func (v *CreateValidations) PreflightValidations(ctx context.Context) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}
