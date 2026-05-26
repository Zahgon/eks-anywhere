package framework

import (
	"context"
)

const (
	kubectlDeleteTimeout = "20m"
)

type WorkloadCluster struct {
	*ClusterE2ETest
	ManagementClusterKubeconfigFile func() string
}

type WorkloadClusters map[string]*WorkloadCluster

func (w *WorkloadCluster) CreateCluster(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

func (w *WorkloadCluster) UpgradeCluster(clusterOpts []ClusterE2ETestOpt, commandOpts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (w *WorkloadCluster) DeleteCluster(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

// ApplyClusterManifest uses client-side logic to create/update objects defined in a cluster yaml manifest.
func (w *WorkloadCluster) ApplyClusterManifest() { _ = "STUB: not implemented"; return }

// DeleteClusterWithKubectl uses client-side logic to delete a cluster.
func (w *WorkloadCluster) DeleteClusterWithKubectl() { _ = "STUB: not implemented"; return }

// WaitForAvailableHardware waits for workload cluster hardware to be available.
func (w *WorkloadCluster) WaitForAvailableHardware() { _ = "STUB: not implemented"; return }

// WaitForKubeconfig waits for the kubeconfig for the workload cluster to be available and then writes it to disk.
func (w *WorkloadCluster) WaitForKubeconfig() { _ = "STUB: not implemented"; return }

// ValidateClusterDelete verifies the cluster has been deleted.
func (w *WorkloadCluster) ValidateClusterDelete() { _ = "STUB: not implemented"; return }

func (w *WorkloadCluster) writeKubeconfigToDisk(ctx context.Context, secretName string, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkloadCluster) availableHardware(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
