package clusterapi

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

const (
	machineHealthCheckKind = "MachineHealthCheck"
)

// durationToSeconds converts a *metav1.Duration to *int32 seconds for v1beta2.
func durationToSeconds(d *metav1.Duration) *int32 { _ = "STUB: not implemented"; return nil }

func machineHealthCheck(clusterName string, unhealthyTimeout, nodeStartupTimeout *metav1.Duration) *clusterv1beta2.MachineHealthCheck {
	_ = "STUB: not implemented"
	return nil
}

// MachineHealthCheckForControlPlane creates MachineHealthCheck resources for the control plane.
func MachineHealthCheckForControlPlane(cluster *v1alpha1.Cluster) *clusterv1beta2.MachineHealthCheck {
	_ = "STUB: not implemented"
	return nil
}

// MachineHealthCheckForWorkers creates MachineHealthCheck resources for the workers.
func MachineHealthCheckForWorkers(cluster *v1alpha1.Cluster) []*clusterv1beta2.MachineHealthCheck {
	_ = "STUB: not implemented"
	return nil
}

func machineHealthCheckForWorker(cluster *v1alpha1.Cluster, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) *clusterv1beta2.MachineHealthCheck {
	_ = "STUB: not implemented"
	return nil
}

// MachineHealthCheckObjects creates MachineHealthCheck resources for control plane and all the worker node groups.
func MachineHealthCheckObjects(cluster *v1alpha1.Cluster) []kubernetes.Object {
	_ = "STUB: not implemented"
	return nil
}
