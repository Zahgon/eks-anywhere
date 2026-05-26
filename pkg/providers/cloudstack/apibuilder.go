package cloudstack

import (
	cloudstackv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// CloudStackMachineTemplateKind defines the K8s Kind corresponding with the MachineTemplate.
const CloudStackMachineTemplateKind = "CloudStackMachineTemplate"

func machineDeployment(clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration, kubeadmConfigTemplate *bootstrapv1beta2.KubeadmConfigTemplate, cloudstackMachineTemplate *cloudstackv1.CloudStackMachineTemplate) *clusterv1beta2.MachineDeployment {
	_ = "STUB: not implemented"
	return nil
}

// MachineDeployments returns generated CAPI MachineDeployment objects for a given cluster spec.
func MachineDeployments(clusterSpec *cluster.Spec, kubeadmConfigTemplates map[string]*bootstrapv1beta2.KubeadmConfigTemplate, machineTemplates map[string]*cloudstackv1.CloudStackMachineTemplate) map[string]*clusterv1beta2.MachineDeployment {
	_ = "STUB: not implemented"
	return nil
}

func generateMachineTemplateAnnotations(machineConfig *v1alpha1.CloudStackMachineConfigSpec) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// sorting for unit test determinism

func setDiskOffering(machineConfig *v1alpha1.CloudStackMachineConfigSpec, template *cloudstackv1.CloudStackMachineTemplate) {
	_ = "STUB: not implemented"
	return
}

// MachineTemplate returns a generated CloudStackMachineTemplate object for a given EKS-A CloudStackMachineConfig.
func MachineTemplate(name string, machineConfig *v1alpha1.CloudStackMachineConfigSpec) *cloudstackv1.CloudStackMachineTemplate {
	_ = "STUB: not implemented"
	return nil
}
