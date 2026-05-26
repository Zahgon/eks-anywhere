// Important: Run "make generate" to regenerate code after modifying this file
// json tags are required; new fields must have json tags for the fields to be serialized

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capiv1 "sigs.k8s.io/cluster-api/api/core/v1beta1"
)

// NutanixMachineConfigSpec defines the desired state of NutanixMachineConfig.
type NutanixMachineConfigSpec struct {
	OSFamily OSFamily            `json:"osFamily"`
	Users    []UserConfiguration `json:"users,omitempty"`
	// vcpusPerSocket is the number of vCPUs per socket of the VM
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	VCPUsPerSocket int32 `json:"vcpusPerSocket"`
	// vcpuSockets is the number of vCPU sockets of the VM
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	VCPUSockets int32 `json:"vcpuSockets"`
	// memorySize is the memory size (in Quantity format) of the VM
	// The minimum memorySize is 2Gi bytes
	// +kubebuilder:validation:Required
	MemorySize resource.Quantity `json:"memorySize"`
	// image is to identify the OS image uploaded to the Prism Central (PC)
	// The image identifier (uuid or name) can be obtained from the Prism Central console
	// or using the Prism Central API.
	// It must include the Kubernetes version(s). For example, a template used for
	// Kubernetes 1.27 could be ubuntu-2204-1.27.
	// +kubebuilder:validation:Required
	Image NutanixResourceIdentifier `json:"image"`
	// cluster is to identify the cluster (the Prism Element under management
	// of the Prism Central), in which the Machine's VM will be created.
	// The cluster identifier (uuid or name) can be obtained from the Prism Central console
	// or using the prism_central API.
	// +kubebuilder:validation:Required
	Cluster NutanixResourceIdentifier `json:"cluster"`
	// subnet is to identify the cluster's network subnet to use for the Machine's VM
	// The cluster identifier (uuid or name) can be obtained from the Prism Central console
	// or using the Prism Central API.
	// +kubebuilder:validation:Required
	Subnet NutanixResourceIdentifier `json:"subnet"`
	// Project is an optional property that specifies the Prism Central project so that machine resources
	// can be linked to it. The project identifier (uuid or name) can be obtained from the Prism Central console
	// or using the Prism Central API.
	// +optional
	Project *NutanixResourceIdentifier `json:"project,omitempty"`

	// systemDiskSize is size (in Quantity format) of the system disk of the VM
	// The minimum systemDiskSize is 20Gi bytes
	// +kubebuilder:validation:Required
	SystemDiskSize resource.Quantity `json:"systemDiskSize"`

	// additionalCategories is a list of optional categories to be added to the VM.
	// Categories must be created in Prism Central before they can be used.
	// +kubebuilder:validation:Optional
	AdditionalCategories []NutanixCategoryIdentifier `json:"additionalCategories,omitempty"`

	// List of GPU devices that should be added to the VMs.
	// +kubebuilder:validation:Optional
	GPUs []NutanixGPUIdentifier `json:"gpus,omitempty"`

	// BootType defines the boot type of the VM. Allowed values: legacy, uefi
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=legacy;uefi
	BootType NutanixBootType `json:"bootType,omitempty"`
}

// SetDefaults sets defaults to NutanixMachineConfig if user has not provided.
func (in *NutanixMachineConfig) SetDefaults() { _ = "STUB: not implemented"; return }

// PauseReconcile pauses the reconciliation of the NutanixMachineConfig.
func (in *NutanixMachineConfig) PauseReconcile() { _ = "STUB: not implemented"; return }

// IsReconcilePaused returns true if the NutanixMachineConfig is paused.
func (in *NutanixMachineConfig) IsReconcilePaused() bool { _ = "STUB: not implemented"; return false }

// SetControlPlane sets the NutanixMachineConfig as a control plane node.
func (in *NutanixMachineConfig) SetControlPlane() { _ = "STUB: not implemented"; return }

// IsControlPlane returns true if the NutanixMachineConfig is a control plane node.
func (in *NutanixMachineConfig) IsControlPlane() bool { _ = "STUB: not implemented"; return false }

// SetEtcd sets the NutanixMachineConfig as an etcd node.
func (in *NutanixMachineConfig) SetEtcd() { _ = "STUB: not implemented"; return }

// IsEtcd returns true if the NutanixMachineConfig is an etcd node.
func (in *NutanixMachineConfig) IsEtcd() bool { _ = "STUB: not implemented"; return false }

// SetManagedBy sets the cluster name that manages the NutanixMachineConfig.
func (in *NutanixMachineConfig) SetManagedBy(clusterName string) { _ = "STUB: not implemented"; return }

// IsManaged returns true if the NutanixMachineConfig is managed by a cluster.
func (in *NutanixMachineConfig) IsManaged() bool { _ = "STUB: not implemented"; return false }

// OSFamily returns the OSFamily of the NutanixMachineConfig.
func (in *NutanixMachineConfig) OSFamily() OSFamily {
	_ = "STUB: not implemented"
	return *

	// Users returns a list of configuration for OS users.
	new(OSFamily)
}

func (in *NutanixMachineConfig) Users() []UserConfiguration { _ = "STUB: not implemented"; return nil }

// GetNamespace returns the namespace of the NutanixMachineConfig.
func (in *NutanixMachineConfig) GetNamespace() string { _ = "STUB: not implemented"; return "" }

// GetName returns the name of the NutanixMachineConfig.
func (in *NutanixMachineConfig) GetName() string {
	_ = "STUB: not implemented"

	// NutanixMachineConfigStatus defines the observed state of NutanixMachineConfig.
	return ""
}

type NutanixMachineConfigStatus struct {
	// Ready is true when the provider resource is ready.
	// +optional
	Ready bool `json:"ready"`

	// Addresses contains the Nutanix VM associated addresses.
	// Address type is one of Hostname, ExternalIP, InternalIP, ExternalDNS, InternalDNS
	Addresses []capiv1.MachineAddress `json:"addresses,omitempty"`

	// The Nutanix VM's UUID
	// +optional
	VmUUID *string `json:"vmUUID,omitempty"`

	// NodeRef is a reference to the corresponding workload cluster Node if it exists.
	// +optional
	NodeRef *corev1.ObjectReference `json:"nodeRef,omitempty"`

	// Conditions defines current service state of the NutanixMachine.
	// +optional
	Conditions capiv1.Conditions `json:"conditions,omitempty"`
}

// NutanixMachineConfig is the Schema for the nutanix machine configs API
//
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type NutanixMachineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NutanixMachineConfigSpec   `json:"spec,omitempty"`
	Status NutanixMachineConfigStatus `json:"status,omitempty"`
}

// ConvertConfigToConfigGenerateStruct converts the NutanixMachineConfig to NutanixMachineConfigGenerate.
func (in *NutanixMachineConfig) ConvertConfigToConfigGenerateStruct() *NutanixMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

// Marshallable returns a Marshallable version of the NutanixMachineConfig.
func (in *NutanixMachineConfig) Marshallable() Marshallable {
	_ = "STUB: not implemented"
	return *new(Marshallable)
}

// Validate validates the NutanixMachineConfig.
func (in *NutanixMachineConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// NutanixMachineConfigGenerate is same as NutanixMachineConfig except stripped down for generation of yaml file during
// generate clusterconfig
//
// +kubebuilder:object:generate=false
type NutanixMachineConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec NutanixMachineConfigSpec `json:"spec,omitempty"`
}

// NutanixMachineConfigList contains a list of NutanixMachineConfig
//
// +kubebuilder:object:root=true
type NutanixMachineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NutanixMachineConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NutanixMachineConfig{}, &NutanixMachineConfigList{})
}
