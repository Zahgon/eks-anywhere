package providers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// ValidateSSHKeyPresentForUpgrade checks that all machine configs in the cluster spec
// contain at least one SSH key.
func ValidateSSHKeyPresentForUpgrade(_ context.Context, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't add snow since SnowMachineConfig's don't have []User

type machineWithUsers interface {
	metav1.Object
	runtime.Object
	Users() []v1alpha1.UserConfiguration
}

func validateAtLeastOneSSHKey(machines []machineWithUsers) error {
	_ = "STUB: not implemented"
	return nil
}

func appendMachinesWithUsers[O machineWithUsers](m []machineWithUsers, addMap map[string]O) []machineWithUsers {
	_ = "STUB: not implemented"
	return nil
}
