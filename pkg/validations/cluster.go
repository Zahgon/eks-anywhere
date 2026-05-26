package validations

import (
	"context"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/semver"
	"github.com/aws/eks-anywhere/pkg/types"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	supportedManagementComponentsMinorVersionIncrement int64 = 1
	releaseV022                                              = "v0.22.0"
)

// ValidateOSForRegistryMirror checks if the OS is valid for the provided registry mirror configuration.
func ValidateOSForRegistryMirror(clusterSpec *cluster.Spec, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateCertForRegistryMirror(clusterSpec *cluster.Spec, tlsValidator TlsValidator) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateAuthenticationForRegistryMirror checks if REGISTRY_USERNAME and REGISTRY_PASSWORD is set if authenticated registry mirrors are used.
func ValidateAuthenticationForRegistryMirror(clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateManagementClusterName checks if the management cluster specified in the workload cluster spec is valid.
func ValidateManagementClusterName(ctx context.Context, k KubectlClient, mgmtCluster *types.Cluster, mgmtClusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateEksaVersion ensures that the version matches EKS-A CLI.
func ValidateEksaVersion(ctx context.Context, cliVersion string, workload *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateEksaVersionSkew ensures that upgrades are sequential by CLI minor versions.
func ValidateEksaVersionSkew(ctx context.Context, k KubectlClient, mgmtCluster *types.Cluster, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateManagementClusterEksaVersion ensures workload cluster isn't created by a newer version than management cluster.
func ValidateManagementClusterEksaVersion(ctx context.Context, k KubectlClient, mgmtCluster *types.Cluster, workload *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateManagementEksaVersion ensures a workload cluster's EksaVersion is not greater than a management cluster's version.
func ValidateManagementEksaVersion(mgmtCluster, cluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// reset failure message if old matches this validation

func clustersHaveEksaVersion(mgmtCluster, cluster *v1alpha1.Cluster) bool {
	_ = "STUB: not implemented"
	return false
}

func parseClusterEksaVersion(mgmtCluster, cluster *v1alpha1.Cluster) (*semver.Version, *semver.Version, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ValidateEksaReleaseExistOnManagement checks if there is a corresponding eksareleases CR for workload's eksaVersion on the mgmt cluster.
func ValidateEksaReleaseExistOnManagement(ctx context.Context, k kubernetes.Client, workload *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidatePauseAnnotation checks if the target cluster has annotation anywhere.eks.amazonaws.com/paused set to true or not.
func ValidatePauseAnnotation(ctx context.Context, k KubectlClient, cluster *types.Cluster, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateManagementComponentsVersionSkew checks if the management components version is only one minor version greater than the cluster version.
func ValidateManagementComponentsVersionSkew(ctx context.Context, k KubectlClient, mgmtCluster *types.Cluster, eksaRelease *releasev1alpha1.EKSARelease) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBottlerocketKubeletConfig validates bottlerocket settings for Kubelet Configuration.
func ValidateBottlerocketKubeletConfig(spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateExtendedKubernetesVersionSupport validates the extended kubernetes version support for create and upgrade operations.
func ValidateExtendedKubernetesVersionSupport(ctx context.Context, clusterSpec v1alpha1.Cluster, reader *manifests.Reader, k kubernetes.Client, bundlesOverride string) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip the signature validation for those versions prior to 'v0.22.0'

// Get the release manifest from the bundle

// getReleaseManifestFromBundle retrieves the EKS Distro release manifest from the bundle.
// For airgapped clusters, it reads from a local file path.
// For non-airgapped clusters, it fetches from the public URL.
func getReleaseManifestFromBundle(clusterSpec v1alpha1.Cluster, bundle *releasev1alpha1.Bundles) (*eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if this is an airgapped cluster (EksDReleaseUrl points to local path)

// Airgapped case: read from local file path

// Non-airgapped case: fetch from public URL
