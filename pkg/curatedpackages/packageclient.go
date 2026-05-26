package curatedpackages

import (
	"context"
	"io"

	packagesv1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
)

const (
	CustomName = "generated-"
	kind       = "Package"
)

type PackageClientOpt func(*PackageClient)

type PackageClient struct {
	bundle         *packagesv1.PackageBundle
	customPackages []string
	kubectl        KubectlRunner
	customConfigs  []string
}

func NewPackageClient(kubectl KubectlRunner, options ...PackageClientOpt) *PackageClient {
	_ = "STUB: not implemented"
	return nil
}

// sourceWithVersions is a wrapper to help get package versions.
//
// This should be pushed upstream to eks-anywhere-packages, then this
// implementation can be removed.
type sourceWithVersions packagesv1.BundlePackageSource

func (s sourceWithVersions) VersionsSlice() []string { _ = "STUB: not implemented"; return nil }

// DisplayPackages pretty-prints a table of available packages.
func (pc *PackageClient) DisplayPackages(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// packagesHeaderLines pretties-up a table of curated packages info.
var packagesHeaderLines = [][]string{
	{"Package", "Version(s)"},
	{"-------", "----------"},
}

func (pc *PackageClient) GeneratePackages(clusterName string) ([]packagesv1.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PackageClient) WritePackagesToStdOut(packages []packagesv1.Package) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PackageClient) GetPackageFromBundle(packageName string) (*packagesv1.BundlePackage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PackageClient) packageMap() map[string]packagesv1.BundlePackage {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PackageClient) InstallPackage(ctx context.Context, bp *packagesv1.BundlePackage, customName string, clusterName string, kubeConfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PackageClient) getInstallConfigurations() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (pc *PackageClient) ApplyPackages(ctx context.Context, fileName string, kubeConfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PackageClient) CreatePackages(ctx context.Context, fileName string, kubeConfig string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PackageClient) DeletePackages(ctx context.Context, packages []string, kubeConfig string, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PackageClient) DescribePackages(ctx context.Context, packages []string, kubeConfig string, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func convertBundlePackageToPackage(bp packagesv1.BundlePackage, name string, clusterName string, apiVersion string, config string) packagesv1.Package {
	_ = "STUB: not implemented"
	return *new(packagesv1.Package)
}

func WithBundle(bundle *packagesv1.PackageBundle) func(*PackageClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithCustomPackages(customPackages []string) func(*PackageClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithCustomConfigs(customConfigs []string) func(*PackageClient) {
	_ = "STUB: not implemented"
	return nil
}
