package bundles

import (
	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

// GetNutanixBundle returns the bundle for Nutanix.
func GetNutanixBundle(r *releasetypes.ReleaseConfig, imageDigests releasetypes.ImageDigestsTable) (anywherev1alpha1.NutanixBundle, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.NutanixBundle), nil
}
