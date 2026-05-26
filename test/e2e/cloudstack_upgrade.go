//go:build e2e

package e2e

import (
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/test/framework"
)

type cloudStackAPIUpgradeTestStep struct {
	name         string
	configFiller api.ClusterConfigFiller
}

type cloudStackAPIUpgradeTest struct {
	name string
	// steps is a list are grouped updates to be applied during a test synchronously.
	steps []cloudStackAPIUpgradeTestStep
}

func clusterPrefix(value, prefix string) string { _ = "STUB: not implemented"; return "" }

func cloudStackAPIUpdateTestBaseStep(e *framework.ClusterE2ETest, cloudstack *framework.CloudStack) cloudStackAPIUpgradeTestStep {
	_ = "STUB: not implemented"
	return *new(cloudStackAPIUpgradeTestStep)
}

// This gives us a blank slate

// Add new WorkerNodeGroups

func cloudstackAPIManagementClusterUpgradeTests(e *framework.ClusterE2ETest, cloudstack *framework.CloudStack) []cloudStackAPIUpgradeTest {
	_ = "STUB: not implemented"
	return nil
}

// This gives us a blank slate

// This gives us a blank slate

// Add new WorkerNodeGroups

func cloudStackAPIWorkloadUpgradeTests(wc *framework.WorkloadCluster, cloudstack *framework.CloudStack) []cloudStackAPIUpgradeTest {
	_ = "STUB: not implemented"
	return nil
}

// This gives us a blank slate

// This gives us a blank slate

// Add new WorkerNodeGroups

func runCloudStackAPIUpgradeTest(t *testing.T, test *framework.ClusterE2ETest, ut cloudStackAPIUpgradeTest) {
	_ = "STUB: not implemented"
	return
}

func runCloudStackAPIWorkloadUpgradeTest(t *testing.T, wc *framework.WorkloadCluster, ut cloudStackAPIUpgradeTest) {
	_ = "STUB: not implemented"
	return
}

func runCloudStackAPIWorkloadUpgradeTestWithFlux(t *testing.T, test *framework.MulticlusterE2ETest, wc *framework.WorkloadCluster, ut cloudStackAPIUpgradeTest) {
	_ = "STUB: not implemented"
	return
}
