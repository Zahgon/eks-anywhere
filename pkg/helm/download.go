package helm

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/retrier"
)

type ChartRegistryDownloader struct {
	client    Client
	dstFolder string

	Retrier retrier.Retrier
}

func NewChartRegistryDownloader(client Client, dstFolder string) *ChartRegistryDownloader {
	_ = "STUB: not implemented"
	return nil
}

func (d *ChartRegistryDownloader) Download(ctx context.Context, charts ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func uniqueCharts(charts []string) []string { _ = "STUB: not implemented"; return nil }

// TODO: maybe optimize this, avoiding the sort and just following the same order as the original slice
