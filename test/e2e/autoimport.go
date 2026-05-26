//go:build e2e
// +build e2e

package e2e

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func runAutoImportFlow(test *framework.ClusterE2ETest, provider *framework.VSphere) {
	_ = "STUB: not implemented"
	return
}

func getMachineConfigs(test *framework.ClusterE2ETest) map[string]v1alpha1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func deleteTemplates(test *framework.ClusterE2ETest, provider *framework.VSphere, machineConfigs map[string]v1alpha1.VSphereMachineConfig) {
	_ = "STUB: not implemented"
	return
}
