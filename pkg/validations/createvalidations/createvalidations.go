package createvalidations

import (
	"github.com/aws/eks-anywhere/pkg/validations"
)

// SkippableValidations represents all the validations we offer for users to skip.
var SkippableValidations = []string{
	validations.VSphereUserPriv,
}

func New(opts *validations.Opts) *CreateValidations { _ = "STUB: not implemented"; return nil }

type CreateValidations struct {
	Opts *validations.Opts
}
