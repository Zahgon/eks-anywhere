//go:build e2e
// +build e2e

package e2e

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func runOIDCFlow(test *framework.ClusterE2ETest) { _ = "STUB: not implemented"; return }

func runTinkerbellOIDCFlow(test *framework.ClusterE2ETest) { _ = "STUB: not implemented"; return }

func runUpgradeFlowWithOIDC(test *framework.ClusterE2ETest, updateVersion v1alpha1.KubernetesVersion, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}
