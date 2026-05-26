package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VSphereMachineConfigSpec defines the desired state of VSphereMachineConfig.
type VSphereMachineConfigSpec struct {
	DiskGiB           int      `json:"diskGiB,omitempty"`
	Datastore         string   `json:"datastore"`
	Folder            string   `json:"folder"`
	NumCPUs           int      `json:"numCPUs"`
	MemoryMiB         int      `json:"memoryMiB"`
	OSFamily          OSFamily `json:"osFamily"`
	ResourcePool      string   `json:"resourcePool"`
	StoragePolicyName string   `json:"storagePolicyName,omitempty"`
	// The field Networks is for configuring custom networks for worker nodes. This can be used to configure upto 2 networks for a worker node group.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxItems=2
	Networks []string `json:"networks,omitempty"`
	// Template field is the template to use for provisioning the VM. It must include the Kubernetes
	// version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.
	Template            string               `json:"template,omitempty"`
	Users               []UserConfiguration  `json:"users,omitempty"`
	TagIDs              []string             `json:"tags,omitempty"`
	CloneMode           CloneMode            `json:"cloneMode,omitempty"`
	HostOSConfiguration *HostOSConfiguration `json:"hostOSConfiguration,omitempty"`
}

// ResourcePaths returns a map of vSphere resource paths defined in the VSphereMachineConfig.
// It collects the Template, ResourcePool, Datastore, and Folder paths
// into a structured map for easier access and validation during cluster operations.
func (c *VSphereMachineConfig) ResourcePaths() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *VSphereMachineConfig) PauseReconcile() { _ = "STUB: not implemented"; return }

func (c *VSphereMachineConfig) IsReconcilePaused() bool { _ = "STUB: not implemented"; return false }

func (c *VSphereMachineConfig) SetControlPlane() { _ = "STUB: not implemented"; return }

func (c *VSphereMachineConfig) IsControlPlane() bool { _ = "STUB: not implemented"; return false }

func (c *VSphereMachineConfig) SetEtcd() { _ = "STUB: not implemented"; return }

func (c *VSphereMachineConfig) IsEtcd() bool { _ = "STUB: not implemented"; return false }

func (c *VSphereMachineConfig) SetManagedBy(clusterName string) { _ = "STUB: not implemented"; return }

// IsManaged returns true if the vspheremachineconfig is associated with a workload cluster.
func (c *VSphereMachineConfig) IsManaged() bool { _ = "STUB: not implemented"; return false }

func (c *VSphereMachineConfig) OSFamily() OSFamily {
	_ = "STUB: not implemented"
	return *

	// Users returns a list of configuration for OS users.
	new(OSFamily)
}

func (c *VSphereMachineConfig) Users() []UserConfiguration { _ = "STUB: not implemented"; return nil }

func (c *VSphereMachineConfig) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (c *VSphereMachineConfig) GetName() string {
	_ = "STUB: not implemented"

	// VSphereMachineConfigStatus defines the observed state of VSphereMachineConfig.
	return ""
}

type VSphereMachineConfigStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VSphereMachineConfig is the Schema for the vspheremachineconfigs API.
type VSphereMachineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VSphereMachineConfigSpec   `json:"spec,omitempty"`
	Status VSphereMachineConfigStatus `json:"status,omitempty"`
}

func (c *VSphereMachineConfig) ConvertConfigToConfigGenerateStruct() *VSphereMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *VSphereMachineConfig) Marshallable() Marshallable {
	_ = "STUB: not implemented"
	return *new(Marshallable)
}

func (c *VSphereMachineConfig) SetDefaults() { _ = "STUB: not implemented"; return }

// SetUserDefaults initializes Spec.Users for the VSphereMachineConfig with default values.
// This only runs in the CLI, as we support do support user defaults through the webhook.
func (c *VSphereMachineConfig) SetUserDefaults() { _ = "STUB: not implemented"; return }

func (c *VSphereMachineConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// ValidateUsers verifies a VSphereMachineConfig object must have a users with ssh authorized keys.
// This validation only runs in VSphereMachineConfig validation webhook, as we support
// auto-generate and import ssh key when creating a cluster via CLI.
func (c *VSphereMachineConfig) ValidateUsers() error { _ = "STUB: not implemented"; return nil }

func validateVSphereMachineConfigOSFamilyUser(machineConfig *VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateHasTemplate verifies that a VSphereMachineConfig object has a template.
// Specifying a template is required when submitting an object via webhook,
// as we only support auto-importing templates when creating a cluster via CLI.
func (c *VSphereMachineConfig) ValidateHasTemplate() error { _ = "STUB: not implemented"; return nil }

// +kubebuilder:object:generate=false

// VSphereMachineConfigGenerate Same as VSphereMachineConfig except stripped down for generation of yaml file during generate clusterconfig.
type VSphereMachineConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec VSphereMachineConfigSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// VSphereMachineConfigList contains a list of VSphereMachineConfig.
type VSphereMachineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereMachineConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VSphereMachineConfig{}, &VSphereMachineConfigList{})
}
