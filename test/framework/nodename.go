package framework

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// ValidateControlPlaneNodeNameMatchCAPIMachineName validate if node name is same as CAPI machine name.
func ValidateControlPlaneNodeNameMatchCAPIMachineName(controlPlane v1alpha1.ControlPlaneConfiguration, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateWorkerNodeNameMatchCAPIMachineName validate if node name is same as CAPI machine name.
func ValidateWorkerNodeNameMatchCAPIMachineName(w v1alpha1.WorkerNodeGroupConfiguration, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNodeNameMatchCAPIMachineName(node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}
