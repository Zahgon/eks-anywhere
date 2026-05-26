package oras

import (
	"context"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type FileRegistryImporter struct {
	registry           string
	username, password string
	srcFolder          string
}

func NewFileRegistryImporter(registry, username, password, srcFolder string) *FileRegistryImporter {
	_ = "STUB: not implemented"
	return nil
}

func (fr *FileRegistryImporter) Push(ctx context.Context, bundles *releasev1.Bundles) {
	_ = "STUB: not implemented"
	return
}

func ChartFileName(chart string) string { _ = "STUB: not implemented"; return "" }
