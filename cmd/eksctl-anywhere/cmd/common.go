package cmd

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// getImages returns all the images in the Bundle for the cluster kubernetes versions.
// This is deprecated. It builds a file reader in line, prefer using the dependency factory.
func getImages(clusterSpecPath, bundlesOverride string) ([]v1alpha1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getKubeconfigPath returns an EKS-A kubeconfig path. The return van be overriden using override
// to give preference to a user specified kubeconfig.
func getKubeconfigPath(clusterName, override string) string { _ = "STUB: not implemented"; return "" }

func NewDependenciesForPackages(ctx context.Context, opts ...PackageOpt) (*dependencies.Dependencies, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PackageOpt func(*PackageConfig)

type PackageConfig struct {
	registryName    string
	kubeVersion     string
	kubeConfig      string
	mountPaths      []string
	spec            *cluster.Spec
	bundlesOverride string
	cluster         *anywherev1.Cluster
}

func New(options ...PackageOpt) *PackageConfig { _ = "STUB: not implemented"; return nil }

func WithRegistryName(registryName string) func(*PackageConfig) {
	_ = "STUB: not implemented"
	return nil
}

func WithKubeVersion(kubeVersion string) func(*PackageConfig) {
	_ = "STUB: not implemented"
	return nil
}

func WithMountPaths(mountPaths ...string) func(*PackageConfig) {
	_ = "STUB: not implemented"
	return nil
}

func WithClusterSpec(spec *cluster.Spec) func(config *PackageConfig) {
	_ = "STUB: not implemented"
	return nil
}

func WithKubeConfig(kubeConfig string) func(*PackageConfig) { _ = "STUB: not implemented"; return nil }

// WithBundlesOverride sets bundlesOverride in the config with incoming value.
func WithBundlesOverride(bundlesOverride string) func(*PackageConfig) {
	_ = "STUB: not implemented"
	return nil
}

// WithCluster sets cluster in the config with incoming value.
func WithCluster(cluster *anywherev1.Cluster) func(config *PackageConfig) {
	_ = "STUB: not implemented"
	return nil
}
