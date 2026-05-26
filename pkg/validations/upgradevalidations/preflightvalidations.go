package upgradevalidations

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/validations"
)

// PreflightValidations returns the validations required before upgrading a cluster.
func (u *UpgradeValidations) PreflightValidations(ctx context.Context) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}

func resultForRemediableValidation(name string, err error) *validations.ValidationResult {
	_ = "STUB: not implemented"
	return nil
}
