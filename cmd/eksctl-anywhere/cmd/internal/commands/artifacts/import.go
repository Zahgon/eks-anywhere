package artifacts

import (
	"context"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type Import struct {
	Reader             Reader
	Bundles            *releasev1.Bundles
	ImageMover         ImageMover
	ChartImporter      ChartImporter
	TmpArtifactsFolder string
	FileImporter       FileImporter
}

type ChartImporter interface {
	Import(ctx context.Context, charts ...string) error
}

type FileImporter interface {
	Push(ctx context.Context, bundles *releasev1.Bundles)
}

func (i Import) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Filter out CSI component images as they're not used by EKS Anywhere
// but are still referenced in the EKS-D release manifest

// Skip CSI component images
