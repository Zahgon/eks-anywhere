package upgradevalidations

import (
	"github.com/aws/eks-anywhere/pkg/validations"
)

// SkippableValidations represents all the validations we offer for users to skip.
var SkippableValidations = []string{
	validations.PDB,
	validations.VSphereUserPriv,
	validations.EksaVersionSkew,
}

func New(opts *validations.Opts) *UpgradeValidations { _ = "STUB: not implemented"; return nil }

type UpgradeValidations struct {
	Opts *validations.Opts
}
