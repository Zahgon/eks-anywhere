package vsphere

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

type Spec struct {
	*cluster.Spec
}

// NewSpec constructs a new vSphere cluster Spec.
func NewSpec(clusterSpec *cluster.Spec) *Spec { _ = "STUB: not implemented"; return nil }

func (s *Spec) controlPlaneMachineConfig() *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *Spec) workerMachineConfig(c anywherev1.WorkerNodeGroupConfiguration) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *Spec) etcdMachineConfig() *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *Spec) machineConfigs() []*anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

// MachineConfigCount represents a machineConfig with it's associated count.
type MachineConfigCount struct {
	*anywherev1.VSphereMachineConfig
	Count int
}

func (s *Spec) machineConfigsWithCount() []MachineConfigCount {
	_ = "STUB: not implemented"
	return nil
}

func etcdMachineConfig(s *cluster.Spec) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func controlPlaneMachineConfig(s *cluster.Spec) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func workerMachineConfig(s *cluster.Spec, workers anywherev1.WorkerNodeGroupConfiguration) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}
