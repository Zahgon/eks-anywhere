package artifacts

import (
	"context"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type ImportToolsImage struct {
	Bundles            *releasev1.Bundles
	ImageMover         ImageMover
	UnPackager         UnPackager
	InputFile          string
	TmpArtifactsFolder string
}

type UnPackager interface {
	UnPackage(orgFile, dstFolder string) error
}

func (i ImportToolsImage) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
