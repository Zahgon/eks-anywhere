package api

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// MasterTaint will be deprecated from kubernetes version 1.25 onwards.
func MasterTaint() corev1.Taint { _ = "STUB: not implemented"; return *new(corev1.Taint) }

// ControlPlaneTaint has been added from 1.24 onwards.
func ControlPlaneTaint() corev1.Taint { _ = "STUB: not implemented"; return *new(corev1.Taint) }

// ValidateControlPlaneTaints will validate that a controlPlane node has the expected taints.
func ValidateControlPlaneTaints(controlPlane v1alpha1.ControlPlaneConfiguration, node corev1.Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// if no taints are specified, kubeadm defaults it to a well-known control plane taint.
// so, we make sure to check for that well-known taint if no taints are provided in the spec.

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

func validateDefaultControlPlaneTaints(node corev1.Node) bool {
	_ = "STUB: not implemented"
	// Due to the transition from "master" to "control-plane", CP nodes can have one or both
	// of these taints, depending on the k8s version. So checking that the node has at least one
	// of them.
	return false
}

func taintEqual(a, b corev1.Taint) bool { _ = "STUB: not implemented"; return false }
