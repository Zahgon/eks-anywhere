//go:build e2e
// +build e2e

package e2e

import (
	"testing"
	"time"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/test/framework"
)

type eksaPackagedBinary interface {
	framework.PackagedBinary
	// Version returns the EKS-A version.
	Version() string
}

// runFlowUpgradeManagementClusterCheckForSideEffects creates management and workload cluster
// with a specific eks-a version then upgrades the management cluster with another CLI version
// and checks that this doesn't cause any side effects (machine rollout) in the workload clusters.
func runFlowUpgradeManagementClusterCheckForSideEffects(test *framework.MulticlusterE2ETest, currentEKSA, newEKSA eksaPackagedBinary, clusterOpts ...framework.ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

// Since we do not generate cluster config from the new binary and update the existing cluster config,
// we will need to manually add the license token for upgrade tests as long as the latest minor is 'v0.21.*'.

type clusterWithMachines struct {
	name     string
	machines clusterMachines
}

type clusterMachines map[string]types.Machine

func anyMachinesChanged(original, current clusterMachines) (changed bool, reason string) {
	_ = "STUB: not implemented"
	return false, ""
}

func printStateOfMachines(managementCluster *anywherev1.Cluster, clusters []clusterWithMachines) {
	_ = "STUB: not implemented"
	return
}

func buildClusterWithMachines(managementCluster *framework.ClusterE2ETest, clusterName string) clusterWithMachines {
	_ = "STUB: not implemented"
	return *new(clusterWithMachines)
}

func buildWorkloadClustersWithMachines(managementCluster *framework.ClusterE2ETest, workloadClusters framework.WorkloadClusters) []clusterWithMachines {
	_ = "STUB: not implemented"
	return nil
}

func waitForMachineDeploymentReady(managementCluster *framework.ClusterE2ETest, cluster *anywherev1.Cluster, workerNodeGroup anywherev1.WorkerNodeGroupConfiguration) {
	_ = "STUB: not implemented"
	return
}

func waitForClusterMachineDeploymentsReady(managementCluster *framework.ClusterE2ETest, cluster *anywherev1.Cluster) {
	_ = "STUB: not implemented"
	return
}

func waitForWorkloadCustersMachineDeploymentsReady(managementCluster *framework.ClusterE2ETest, workloadClusters framework.WorkloadClusters) {
	_ = "STUB: not implemented"
	return
}

type machineSideEffectChecker struct {
	tb                                testing.TB
	checkDuration, waitInBetweenTries time.Duration
}

func (m machineSideEffectChecker) haveMachinesChanged(managementCluster *framework.ClusterE2ETest, preUpgradeWorkloadClustersState []clusterWithMachines) (changed bool, changeReason string) {
	_ = "STUB: not implemented"
	return false, ""
}

// eksaLocalPackagedBinary implements eksaPackagedBinary using the local eks-a binary
// being tested by this suite.
type eksaLocalPackagedBinary struct {
	path, version string
}

func (b eksaLocalPackagedBinary) BinaryPath() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b eksaLocalPackagedBinary) Version() string { _ = "STUB: not implemented"; return "" }

func newEKSAPackagedBinaryForLocalBinary(tb testing.TB) eksaPackagedBinary {
	_ = "STUB: not implemented"
	return *new(eksaPackagedBinary)
}

func runTestManagementClusterUpgradeSideEffects(t *testing.T, provider framework.Provider, os framework.OS, kubeVersion anywherev1.KubernetesVersion, configFillers ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}
