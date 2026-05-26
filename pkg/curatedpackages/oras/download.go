package oras

import (
	"context"

	"github.com/go-logr/logr"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type BundleDownloader struct {
	dstFolder string
	log       logr.Logger
}

// NewBundleDownloader returns a new BundleDownloader.
func NewBundleDownloader(log logr.Logger, dstFolder string) *BundleDownloader {
	_ = "STUB: not implemented"
	return nil
}

func (bd *BundleDownloader) Download(ctx context.Context, bundles *releasev1.Bundles) {
	_ = "STUB: not implemented"
	return
}

func UniqueCharts(charts []string) []string { _ = "STUB: not implemented"; return nil }

// If the key(values of the slice) is not equal
// to the already present value in new slice (list)
// then we append it. else we jump on another element.

func writeToFile(dir string, packageName string, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadFilesFromBundles(bundles *releasev1.Bundles) []string {
	_ = "STUB: not implemented"
	return nil
}
