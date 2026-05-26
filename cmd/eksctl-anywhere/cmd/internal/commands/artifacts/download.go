package artifacts

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/files"
	"github.com/aws/eks-anywhere/pkg/version"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type Reader interface {
	ReadBundlesForVersion(eksaVersion string) (*releasev1.Bundles, error)
	ReadImagesFromBundles(ctx context.Context, bundles *releasev1.Bundles) ([]releasev1.Image, error)
	ReadChartsFromBundles(ctx context.Context, bundles *releasev1.Bundles) []releasev1.Image
}

type ImageMover interface {
	Move(ctx context.Context, artifacts ...string) error
}

type ChartDownloader interface {
	Download(ctx context.Context, artifacts ...string) error
}

type ManifestDownloader interface {
	Download(ctx context.Context, bundles *releasev1.Bundles)
}

type Packager interface {
	Package(folder string, dstFile string) error
}

type Download struct {
	Reader                   Reader
	FileReader               *files.Reader
	Version                  version.Info
	BundlesImagesDownloader  ImageMover
	EksaToolsImageDownloader ImageMover
	ChartDownloader          ChartDownloader
	Packager                 Packager
	TmpDowloadFolder         string
	DstFile                  string
	ManifestDownloader       ManifestDownloader
	BundlesOverride          string
}

func (d Download) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func artifactNames(artifacts []releasev1.Image) []string { _ = "STUB: not implemented"; return nil }

func removeFromSlice(s []string, toRemove string) []string { _ = "STUB: not implemented"; return nil }
