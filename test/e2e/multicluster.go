//go:build e2e
// +build e2e

package e2e

import (
	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func runWorkloadClusterFlow(test *framework.MulticlusterE2ETest) { _ = "STUB: not implemented"; return }

func runWorkloadClusterExistingConfigFlow(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runWorkloadClusterPrevVersionCreateFlow(test *framework.MulticlusterE2ETest, latestMinorRelease *releasev1.EksARelease) {
	_ = "STUB: not implemented"
	return
}

func runWorkloadClusterFlowWithGitOps(test *framework.MulticlusterE2ETest, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

func runWorkloadClusterGitOpsAPIFlowForBareMetal(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runWorkloadClusterGitOpsAPIUpgradeFlowForBareMetal(test *framework.MulticlusterE2ETest, filler ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}

func runTinkerbellWorkloadClusterFlow(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runWorkloadClusterWithAPIFlowForBareMetal(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runSimpleWorkloadUpgradeFlowForBareMetal(test *framework.MulticlusterE2ETest, updateVersion v1alpha1.KubernetesVersion, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

func runWorkloadClusterUpgradeFlowWithAPIForBareMetal(test *framework.MulticlusterE2ETest, filler ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}

func runInPlaceWorkloadUpgradeFlow(test *framework.MulticlusterE2ETest, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}
