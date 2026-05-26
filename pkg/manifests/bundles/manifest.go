package bundles

import (
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// Manifest holds the data of a manifest referenced in the Bundles.
type Manifest struct {
	Filename string
	Content  []byte
}

// ReadManifest reads the content of a [releasev1.Manifest].
func ReadManifest(reader Reader, manifest releasev1.Manifest) (*Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
