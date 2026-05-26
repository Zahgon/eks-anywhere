package validator

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/networkutils"
)

// IPValidator defines the struct for control plane IP validations.
type IPValidator struct {
	netClient networkutils.NetClient
}

// IPValidatorOpt is the type for optional IPValidator configurations.
type IPValidatorOpt func(e *IPValidator)

// CustomNetClient passes in a custom net client to the IPValidator.
func CustomNetClient(netClient networkutils.NetClient) IPValidatorOpt {
	_ = "STUB: not implemented"
	return *new(IPValidatorOpt)
}

// NewIPValidator initializes a new IPValidator object.
func NewIPValidator(opts ...IPValidatorOpt) *IPValidator { _ = "STUB: not implemented"; return nil }

// ValidateControlPlaneIPUniqueness checks whether or not the control plane endpoint defined
// in the cluster spec is available.
func (v *IPValidator) ValidateControlPlaneIPUniqueness(cluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
