package framework

import (
	"testing"

	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	// VSphereMultiTemplateUbuntu127 is to test multiple vSphere templates.
	VSphereMultiTemplateUbuntu127 = "T_VSPHERE_TEMPLATE_UBUNTU_2204_1_27"
	// ControlPlaneMachineLabel is to get control plane vSphere machines from machine list.
	ControlPlaneMachineLabel = "cluster.x-k8s.io/control-plane"
	// EtcdMachineLabel is to get etcd vSphere machines from machine list.
	EtcdMachineLabel = "cluster.x-k8s.io/etcd-cluster"
	// workerMachineLabel is to get worker vSphere machines from machine list.
	workerMachineLabel = "cluster.x-k8s.io/deployment-name"
)

// VSphereMultiTemplateUbuntu127RequiredEnvVars is required for TestVSphereMultipleTemplatesUbuntu127.
var VSphereMultiTemplateUbuntu127RequiredEnvVars = []string{
	VSphereMultiTemplateUbuntu127,
}

// RequiredVsphereMultiTemplateUbuntu127EnvVars return required env vars for TestVSphereMultipleTemplatesUbuntu127.
func RequiredVsphereMultiTemplateUbuntu127EnvVars() []string { _ = "STUB: not implemented"; return nil }

// CheckVsphereMultiTemplateUbuntu127EnvVars checks is required env vars are present.
func CheckVsphereMultiTemplateUbuntu127EnvVars(t *testing.T) { _ = "STUB: not implemented"; return }

// VsphereMachineValidation should return an error if either an error is encountered during execution or the validation logically fails.
// This validation function will be executed by ValidateVsphereMachine and ValidateWorkerNodeVsphereMachine
// with a vSphere machine config and a corresponding vSphere machine.
type VsphereMachineValidation func(machineConfig *v1alpha1.VSphereMachineConfig, machine vspherev1.VSphereMachine) (err error)

// ValidateVsphereMachine deduces the control plane or etcd configuration to machine mapping
// and for each configuration/machine pair executes the provided validation functions.
func (e *ClusterE2ETest) ValidateVsphereMachine(selector string, machineConfig *v1alpha1.VSphereMachineConfig, validations ...VsphereMachineValidation) {
	_ = "STUB: not implemented"
	return
}

// ValidateWorkerNodeVsphereMachine deduces the worker node group configuration to machine mapping
// and for each configuration/machine pair executes the provided validation functions.
func (e *ClusterE2ETest) ValidateWorkerNodeVsphereMachine(validations ...VsphereMachineValidation) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) getMachineConfigToMachine() map[*v1alpha1.VSphereMachineConfig][]vspherev1.VSphereMachine {
	_ = "STUB: not implemented"
	return nil
}

func getWngNameToWng(wngConfigs []v1alpha1.WorkerNodeGroupConfiguration) map[string]v1alpha1.WorkerNodeGroupConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// getWngNameFromMachine gets worker node group name from machine name by trimming cluster name prefix and two unix nano time suffix.
func getWngNameFromMachine(machineName string, clusterName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ValidateMachineTemplate validates if template configured in machine config matches the vSphere machine.
func ValidateMachineTemplate(machineConfig *v1alpha1.VSphereMachineConfig, machine vspherev1.VSphereMachine) (err error) {
	_ = "STUB: not implemented"
	return nil
}
