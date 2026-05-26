package bundles

import (
	eksdv1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type Reader interface {
	ReadFile(url string) ([]byte, error)
}

func Read(reader Reader, url string) (*releasev1.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadEKSD(reader Reader, versionsBundle releasev1.VersionsBundle) (*eksdv1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
