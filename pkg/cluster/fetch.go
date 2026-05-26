package cluster

import (
	"context"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	v1alpha1release "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// ManagementComponents bundles the resource definitions of all EKS-A management components.
type ManagementComponents struct {
	EksD                   v1alpha1release.EksDRelease
	CertManager            v1alpha1release.CertManagerBundle
	ClusterAPI             v1alpha1release.CoreClusterAPI
	Bootstrap              v1alpha1release.KubeadmBootstrapBundle
	ControlPlane           v1alpha1release.KubeadmControlPlaneBundle
	VSphere                v1alpha1release.VSphereBundle
	CloudStack             v1alpha1release.CloudStackBundle
	Docker                 v1alpha1release.DockerBundle
	Eksa                   v1alpha1release.EksaBundle
	Flux                   v1alpha1release.FluxBundle
	ExternalEtcdBootstrap  v1alpha1release.EtcdadmBootstrapBundle
	ExternalEtcdController v1alpha1release.EtcdadmControllerBundle
	Tinkerbell             v1alpha1release.TinkerbellBundle
	Snow                   v1alpha1release.SnowBundle
	Nutanix                v1alpha1release.NutanixBundle
}

// ManagementComponentsFromBundles returns ManagementComponents built from a VersionsBundle.
//
// For decoupled component upgrades, the management components can be upgraded to the new EKS-A version
// separately from the Cluster. So, here we have the management components bundles for that new version,
// but, there are still multiple Kubernetes versions to choose from within the bundle to get the
// components information. However, because management component images are the same for every Kubernetes
// version within the same bundle manifest, it's OK to use the first bundle. If there are is differences between
// the management components on this first versions bundle, and the new cluster specs first versions bundle,
// that indicates an upgrade is required. In the future, we might change the bundles API to remove the assumption
// and make this explicit. When that happens, this method will need to change.
func ManagementComponentsFromBundles(bundles *v1alpha1release.Bundles) *ManagementComponents {
	_ = "STUB: not implemented"
	return nil
}

// newManagementComponents returns a ManagementComponents object built from a VersionsBundle.
func newManagementComponents(vb *v1alpha1release.VersionsBundle) *ManagementComponents {
	_ = "STUB: not implemented"
	return nil
}

func bundlesNamespacedKey(cluster *v1alpha1.Cluster, release *v1alpha1release.EKSARelease) (name, namespace string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Handles old clusters that don't contain a reference yet to the Bundles
// For those clusters, the Bundles was created with the same name as the cluster
// and in the same namespace

// GetManagementComponents returns the first VersionsBundle from the Bundles object for cluster's the management components version.
func GetManagementComponents(ctx context.Context, client Client, cluster *v1alpha1.Cluster) (*ManagementComponents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVersionsBundle gets the VersionsBundle that corresponds to KubernetesVersion.
func GetVersionsBundle(version v1alpha1.KubernetesVersion, bundles *v1alpha1release.Bundles) (*v1alpha1release.VersionsBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVersionsBundleForKubernetesVersion(kubernetesVersion v1alpha1.KubernetesVersion, bundles *v1alpha1release.Bundles) (*v1alpha1release.VersionsBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildSpec constructs a cluster.Spec for an eks-a cluster by retrieving all
// necessary objects from the cluster using a kubernetes client.
func BuildSpec(ctx context.Context, client Client, cluster *v1alpha1.Cluster) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildSpecFromConfig constructs a cluster.Spec for an eks-a cluster config by retrieving all dependencies objects from the cluster using a kubernetes client.
func BuildSpecFromConfig(ctx context.Context, client Client, config *Config) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchAllEksdReleases(ctx context.Context, client Client, cluster *v1alpha1.Cluster, bundles *v1alpha1release.Bundles) ([]eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEksdRelease(ctx context.Context, client Client, version v1alpha1.KubernetesVersion, bundles *v1alpha1release.Bundles) (*eksdv1alpha1.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ideally we would use the same namespace as the Bundles, but Bundles can be in any namespace and
// the eksd release is always in eksa-system

// BundlesForCluster returns a bundles resource for the cluster.
func BundlesForCluster(ctx context.Context, client Client, cluster *v1alpha1.Cluster) (*v1alpha1release.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func eksaReleaseForCluster(ctx context.Context, client Client, cluster *v1alpha1.Cluster) (*v1alpha1release.EKSARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bundlesForEksaRelease(ctx context.Context, client Client, cluster *v1alpha1.Cluster, eksaRelease *v1alpha1release.EKSARelease) (*v1alpha1release.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
