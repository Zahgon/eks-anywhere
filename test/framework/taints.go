package framework

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const ownerAnnotation = "cluster.x-k8s.io/owner-name"

// ValidateControlPlaneTaints will validate that a controlPlane node has the expected taints.
func ValidateControlPlaneTaints(controlPlane v1alpha1.ControlPlaneConfiguration, node corev1.Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ValidateControlPlaneNoTaints will validate that a controlPlane has no taints, for example in the case of a single node cluster.
func ValidateControlPlaneNoTaints(controlPlane v1alpha1.ControlPlaneConfiguration, node corev1.Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ValidateWorkerNodeTaints will validate that a worker node has the expected taints in the worker node group configuration.
func ValidateWorkerNodeTaints(w v1alpha1.WorkerNodeGroupConfiguration, node corev1.Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func NoExecuteTaint() corev1.Taint { _ = "STUB: not implemented"; return *new(corev1.Taint) }

func NoScheduleTaint() corev1.Taint { _ = "STUB: not implemented"; return *new(corev1.Taint) }

func PreferNoScheduleTaint() corev1.Taint { _ = "STUB: not implemented"; return *new(corev1.Taint) }

func NoScheduleWorkerNodeGroup(name string, count int) *WorkerNodeGroup {
	_ = "STUB: not implemented"
	return nil
}

func PreferNoScheduleWorkerNodeGroup(name string, count int) *WorkerNodeGroup {
	_ = "STUB: not implemented"
	return nil
}
