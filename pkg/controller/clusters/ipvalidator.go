package clusters

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
)

// IPUniquenessValidator defines an interface for the methods to validate the control plane IP.
type IPUniquenessValidator interface {
	ValidateControlPlaneIPUniqueness(cluster *anywherev1.Cluster) error
}

// IPValidator validates control plane IP.
type IPValidator struct {
	ipUniquenessValidator IPUniquenessValidator
	client                client.Client
}

// NewIPValidator returns a new NewIPValidator.
func NewIPValidator(ipUniquenessValidator IPUniquenessValidator, client client.Client) *IPValidator {
	_ = "STUB: not implemented"
	return nil
}

// ValidateControlPlaneIP only validates IP on cluster creation.
func (i *IPValidator) ValidateControlPlaneIP(ctx context.Context, log logr.Logger, spec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// If CAPI cluster exists, the control plane IP has already been validated,
// and it's possibly already in use so no need to validate it again
