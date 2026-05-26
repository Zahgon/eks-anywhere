package eksd

import (
	eksdv1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
)

type Reader interface {
	ReadFile(url string) ([]byte, error)
}

func ReadManifest(reader Reader, url string) (*eksdv1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
