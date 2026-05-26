package framework

import (
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
)

type MulticlusterE2ETest struct {
	T                 *testing.T
	ManagementCluster *ClusterE2ETest
	WorkloadClusters  WorkloadClusters
	// MaxConcurrentWorkers defines the max number of workers for concurrent operations.
	// If it's -1, it will use one worker per job.
	MaxConcurrentWorkers     int
	workloadClusterNameCount int
}

func NewMulticlusterE2ETest(t *testing.T, managementCluster *ClusterE2ETest, workloadClusters ...*ClusterE2ETest) *MulticlusterE2ETest {
	_ = "STUB: not implemented"
	return nil
}

// WithWorkloadClusters adds ClusterE2ETest's as workload clusters to the test.
func (m *MulticlusterE2ETest) WithWorkloadClusters(workloadClusters ...*ClusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

// NewWorkloadClusterName returns a new unique name for a workload cluster based on the management cluster name.
// This is not thread safe.
func (m *MulticlusterE2ETest) NewWorkloadClusterName() string { _ = "STUB: not implemented"; return "" }

func (m *MulticlusterE2ETest) RunInWorkloadClusters(flow func(*WorkloadCluster)) {
	_ = "STUB: not implemented"
	return
}

// RunConcurrentlyInWorkloadClusters executes the given flow concurrently for all workload
// clusters. It respects MaxConcurrentWorkers.
func (m *MulticlusterE2ETest) RunConcurrentlyInWorkloadClusters(flow func(*WorkloadCluster)) {
	_ = "STUB: not implemented"
	return
}

// RunConcurrently runs the given jobs concurrently using no more than MaxConcurrentWorkers workers.
// If MaxConcurrentWorkers is -1, it will use one worker per job.
func (m *MulticlusterE2ETest) RunConcurrently(flows ...func()) { _ = "STUB: not implemented"; return }

func (m *MulticlusterE2ETest) CreateManagementClusterForVersion(eksaVersion string, opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// CreateManagementClusterWithConfig first generates a cluster config based on the management cluster test's
// previous configuration and proceeds to create a management cluster with the CLI.
func (m *MulticlusterE2ETest) CreateManagementClusterWithConfig(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (m *MulticlusterE2ETest) CreateManagementCluster(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// CreateTinkerbellManagementCluster runs tinkerbell related steps for cluster creation.
func (m *MulticlusterE2ETest) CreateTinkerbellManagementCluster(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (m *MulticlusterE2ETest) DeleteManagementCluster() { _ = "STUB: not implemented"; return }

// DeleteTinkerbellManagementCluster runs tinkerbell related steps for cluster deletion.
func (m *MulticlusterE2ETest) DeleteTinkerbellManagementCluster() {
	_ = "STUB: not implemented"
	return
}

// PushWorkloadClusterToGit builds the workload cluster config file for git and pushing changes to git.
func (m *MulticlusterE2ETest) PushWorkloadClusterToGit(w *WorkloadCluster, opts ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}

// DeleteWorkloadClusterFromGit deletes a workload cluster config file and pushes the changes to git.
func (m *MulticlusterE2ETest) DeleteWorkloadClusterFromGit(w *WorkloadCluster) {
	_ = "STUB: not implemented"
	return
}
