package clusterapi

import (
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// Autoscaler annotation constants.
const (
	NodeGroupMinSizeAnnotation = "cluster.x-k8s.io/cluster-api-autoscaler-node-group-min-size"
	NodeGroupMaxSizeAnnotation = "cluster.x-k8s.io/cluster-api-autoscaler-node-group-max-size"
)

func ConfigureAutoscalingInMachineDeployment(md *clusterv1beta2.MachineDeployment, autoscalingConfig *anywherev1.AutoScalingConfiguration) {
	_ = "STUB: not implemented"
	return
}
