package test

import (
	"embed"
	"testing"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type ClusterSpecOpt func(*cluster.Spec)

//go:embed testdata
var configFS embed.FS

// DevEksaVersion can be used in tests.
func DevEksaVersion() v1alpha1.EksaVersion {
	_ = "STUB: not implemented"
	return *new(v1alpha1.EksaVersion)
}

func NewClusterSpec(opts ...ClusterSpecOpt) *cluster.Spec { _ = "STUB: not implemented"; return nil }

func NewFullClusterSpec(t *testing.T, clusterConfigFile string) *cluster.Spec {
	_ = "STUB: not implemented"
	return nil
}

// NewClusterSpecForCluster builds a compliant [cluster.Spec] from a Cluster using a test
// Bundles and EKS-D Release.
func NewClusterSpecForCluster(tb testing.TB, c *v1alpha1.Cluster) *cluster.Spec {
	_ = "STUB: not implemented"
	return nil
}

// NewClusterSpecForConfig builds a compliant [cluster.Spec] from a [cluster.Config] using a test
// Bundles and EKS-D Release.
func NewClusterSpecForConfig(tb testing.TB, config *cluster.Config) *cluster.Spec {
	_ = "STUB: not implemented"
	return nil
}

// Bundles returs a test Bundles. All the paths to referenced manifests are valid and can be read.
func Bundles(tb testing.TB) *releasev1alpha1.Bundles { _ = "STUB: not implemented"; return nil }

// EksdReleaseFromTestData returns a test release struct for unit testing from a testdata file.
// See EksdRelease() for a static struct to test with.
func EksdReleaseFromTestData(t *testing.T) *eksdv1alpha1.Release {
	_ = "STUB: not implemented"
	return nil
}

func SetTag(image *releasev1alpha1.Image, tag string) { _ = "STUB: not implemented"; return }

// RegistryMirrorEndpoint returns the address of the registry mirror configured on the Cluster if any. Just the host and the port.
func RegistryMirrorEndpoint(cluster *v1alpha1.Cluster) string { _ = "STUB: not implemented"; return "" }
