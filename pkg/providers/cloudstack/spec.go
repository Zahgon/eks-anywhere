package cloudstack

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

func etcdMachineConfig(s *cluster.Spec) *anywherev1.CloudStackMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func controlPlaneMachineConfig(s *cluster.Spec) *anywherev1.CloudStackMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func workerMachineConfig(s *cluster.Spec, workers anywherev1.WorkerNodeGroupConfiguration) *anywherev1.CloudStackMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func controlPlaneEndpointHost(clusterSpec *cluster.Spec) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getValidControlPlaneHostPort(pHost string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
