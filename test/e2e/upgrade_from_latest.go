//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func latestMinorRelease(t testing.TB) *releasev1.EksARelease { _ = "STUB: not implemented"; return nil }

func prevLatestMinorRelease(t testing.TB) *releasev1.EksARelease {
	_ = "STUB: not implemented"

	// Fetch the previous latest minor release for workload creation For ex. curr latest release 15.x prev latest minor release: 14.x
	return nil
}

func runUpgradeFromReleaseFlow(test *framework.ClusterE2ETest, latestRelease *releasev1.EksARelease, wantVersion anywherev1.KubernetesVersion, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

// Adding this manual wait because old versions of the cli don't wait long enough
// after creation, which makes the upgrade preflight validations fail

func runUpgradeWithFluxFromReleaseFlow(test *framework.ClusterE2ETest, latestRelease *releasev1.EksARelease, wantVersion anywherev1.KubernetesVersion, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

// Adding this manual wait because old versions of the cli don't wait long enough
// after creation, which makes the upgrade preflight validations fail

func runInPlaceUpgradeFromReleaseFlow(test *framework.ClusterE2ETest, latestRelease *releasev1.EksARelease, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

// runMulticlusterUpgradeFromReleaseFlowAPI tests the ability to create workload clusters with an old Bundle in a management cluster
// that has been updated to a new Bundle. It follows the following steps:
//  1. Create a management cluster with the old Bundle.
//  2. Create workload clusters with the old Bundle.
//  3. Upgrade the management cluster to the new Bundle and new Kubernetes version (newVersion).
//  4. Upgrade the workload clusters to the new Bundle and new Kubernetes version (newVersion).
//  5. Delete the workload clusters.
//  6. Re-create the workload clusters with the old Bundle and previous Kubernetes version (oldVersion). It's necessary to sometimes
//     use a different kube version because the old Bundle might not support the new kubernetes version.
//  7. Delete the workload clusters.
//  8. Delete the management cluster.
func runMulticlusterUpgradeFromReleaseFlowAPI(test *framework.MulticlusterE2ETest, release *releasev1.EksARelease, oldVersion, newVersion anywherev1.KubernetesVersion, os framework.OS) {
	_ = "STUB: not implemented"
	return
}

// 1. Create management cluster

// 2. Create workload clusters with the old Bundle

// 3. Upgrade management cluster to new Bundle and new Kubernetes version

// 4. Upgrade the workload clusters to the new Bundle and new Kubernetes version (newVersion).

// 5. Delete the workload clusters.

// 6. Re-create the workload clusters with the old Bundle and previous Kubernetes version (oldVersion).

// 7. Delete the workload clusters.

// It's necessary to call stop here because if any of the workload clusters failed,
// their panic was thrown in a go routine, which doesn't stop the main test routine.

// 8. Delete the management cluster.

func runMulticlusterUpgradeFromReleaseFlowAPIWithFlux(test *framework.MulticlusterE2ETest, release *releasev1.EksARelease, kubeVersion anywherev1.KubernetesVersion, os framework.OS) {
	_ = "STUB: not implemented"
	return
}

// Upgrade bundle workload clusters now using the new EksaVersion

// Create workload cluster with the old EksaVersion

func runUpgradeManagementComponentsFlow(t *testing.T, release *releasev1.EksARelease, provider framework.Provider, kubeVersion anywherev1.KubernetesVersion, os framework.OS) {
	_ = "STUB: not implemented"
	return
}

// create cluster with old eksa

// upgrade management-components with new eksa
