package bundles

import (
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// ReadImages returns a list of all images included in the Bundles and the referenced
// EKS-D Releases, all of them filtered by kubernetes version. If not kubernetes versions
// all provided, all images are returned.
func ReadImages(reader Reader, bundles *releasev1.Bundles, kubeVersions ...string) ([]releasev1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
