//go:build e2e
// +build e2e

package e2e

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func runAWSIamAuthFlow(test *framework.ClusterE2ETest) { _ = "STUB: not implemented"; return }

func runUpgradeFlowWithAWSIamAuth(test *framework.ClusterE2ETest, updateVersion v1alpha1.KubernetesVersion, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

func runTinkerbellAWSIamAuthFlow(test *framework.ClusterE2ETest) { _ = "STUB: not implemented"; return }

func runAWSIamAuthFlowWorkload(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runUpgradeFlowAddAWSIamAuth(test *framework.ClusterE2ETest, updateVersion v1alpha1.KubernetesVersion) {
	_ = "STUB: not implemented"
	return
}
