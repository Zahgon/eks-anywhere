package curatedpackages

import (
	"context"

	packagesv1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	ImageRepositoryName = "eks-anywhere-packages-bundles"
)

type Reader interface {
	ReadBundlesForVersion(eksaVersion string) (*releasev1.Bundles, error)
}

type BundleRegistry interface {
	GetRegistryBaseRef(ctx context.Context) (string, error)
}

type BundleReader struct {
	kubeConfig    string
	clusterName   string
	kubectl       KubectlRunner
	bundleManager Manager
	registry      BundleRegistry
}

func NewBundleReader(kubeConfig string, clusterName string, k KubectlRunner, bm Manager, reg BundleRegistry) *BundleReader {
	_ = "STUB: not implemented"
	return nil
}

func (b *BundleReader) GetLatestBundle(ctx context.Context, kubeVersion string) (*packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BundleReader) getLatestBundleFromRegistry(ctx context.Context, kubeVersion string) (*packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BundleReader) getActiveBundleFromCluster(ctx context.Context) (*packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	// Active BundleReader is set at the bundle Controller
	return nil, nil
}

func (b *BundleReader) getPackageBundle(ctx context.Context, bundleName string) (*packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BundleReader) GetActiveController(ctx context.Context) (*packagesv1.PackageBundleController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BundleReader) UpgradeBundle(ctx context.Context, controller *packagesv1.PackageBundleController, newBundleVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetPackageBundleRef(vb releasev1.VersionsBundle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Use package controller registry to fetch packageBundles.
// Format of controller image is: <uri>/<env_type>/<repository_name>
