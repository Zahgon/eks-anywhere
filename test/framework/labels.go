package framework

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const LabelPrefix = "eksa.e2e"

func ValidateControlPlaneLabels(controlPlane v1alpha1.ControlPlaneConfiguration, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateControlPlaneFailureDomainLabels validate if Cloudstack provider replaces ds.meta_data.failuredomain with proper failuredomain name
// in control plane node label 'cluster.x-k8s.io/failure-domain'.
func ValidateControlPlaneFailureDomainLabels(controlPlane v1alpha1.ControlPlaneConfiguration, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateWorkerNodeLabels(w v1alpha1.WorkerNodeGroupConfiguration, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateWorkerNodeFailureDomainLabels validate if Cloudstack provider replaces ds.meta_data.failuredomain with proper failuredomain name
// in worker group node label 'cluster.x-k8s.io/failure-domain'.
func ValidateWorkerNodeFailureDomainLabels(w v1alpha1.WorkerNodeGroupConfiguration, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLabels(expectedLabels map[string]string, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func retrieveTestNodeLabels(nodeLabels map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func validateFailureDomainLabel(expectedLabels map[string]string, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}
