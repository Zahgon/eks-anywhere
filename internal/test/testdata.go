package test

import (
	eksdv1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// Namespace returns a test namespace struct for unit testing.
func Namespace(name string) *corev1.Namespace { _ = "STUB: not implemented"; return nil }

// EKSARelease returns a test eksaRelease struct for unit testing.
func EKSARelease() *releasev1.EKSARelease { _ = "STUB: not implemented"; return nil }

// EksdReleases returns a test release slice for unit testing.
func EksdReleases() []eksdv1.Release { _ = "STUB: not implemented"; return nil }

// VersionsBundlesMap returns a test VersionsBundle map for unit testing.
func VersionsBundlesMap() map[anywherev1.KubernetesVersion]*cluster.VersionsBundle {
	_ = "STUB: not implemented"
	return nil
}

// VersionBundle returns a test VersionsBundle struct for unit testing.
func VersionBundle() *cluster.VersionsBundle { _ = "STUB: not implemented"; return nil }

// EksdRelease returns a test release struct for unit testing.
func EksdRelease(channel string) *eksdv1.Release { _ = "STUB: not implemented"; return nil }

// Bundle returns a test bundle struct for unit testing.
func Bundle() *releasev1.Bundles { _ = "STUB: not implemented"; return nil }

// CAPIClusterOpt represents a function where a capi cluster (v1beta2) is passed as an argument.
type CAPIClusterOpt func(*clusterv1beta2.Cluster)

// CAPICluster returns a capi v1beta2 cluster which can be configured by passing in opts arguments.
func CAPICluster(opts ...CAPIClusterOpt) *clusterv1beta2.Cluster {
	_ = "STUB: not implemented"
	return nil
}

// KubeadmControlPlaneOpt represents an function where a kubeadmcontrolplane is passed as an argument.
type KubeadmControlPlaneOpt func(kcp *controlplanev1beta2.KubeadmControlPlane)

// KubeadmControlPlane returns a kubeadm controlplane which can be configured by passing in opts arguments.
func KubeadmControlPlane(opts ...KubeadmControlPlaneOpt) *controlplanev1beta2.KubeadmControlPlane {
	_ = "STUB: not implemented"
	return nil
}

// MachineDeploymentOpt represents a function where a MachineDeployment is passed as an argument.
type MachineDeploymentOpt func(md *clusterv1beta2.MachineDeployment)

// MachineDeployment returns a machinedeployment which can be configured by passing in opts arguments.
func MachineDeployment(opts ...MachineDeploymentOpt) *clusterv1beta2.MachineDeployment {
	_ = "STUB: not implemented"
	return nil
}
