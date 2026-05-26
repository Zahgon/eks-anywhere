// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudStackMachineConfigSpec defines the desired state of CloudStackMachineConfig.
type CloudStackMachineConfigSpec struct {
	// Template refers to a VM image template which has been previously registered in CloudStack.
	// It can either be specified as a UUID or name.
	// When using a template name it must include the Kubernetes version(s). For example,
	// a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.
	Template CloudStackResourceIdentifier `json:"template"`
	// ComputeOffering refers to a compute offering which has been previously registered in
	// CloudStack. It represents a VM’s instance size including number of CPU’s, memory, and CPU
	// speed. It can either be specified as a UUID or name
	ComputeOffering CloudStackResourceIdentifier `json:"computeOffering"`
	// DiskOffering refers to a disk offering which has been previously registered in CloudStack.
	// It represents a disk offering with pre-defined size or custom specified disk size. It can
	// either be specified as a UUID or name
	DiskOffering *CloudStackResourceDiskOffering `json:"diskOffering,omitempty"`
	// Users consists of an array of objects containing the username, as well as a list of their
	// public keys. These users will be authorized to ssh into the machines
	Users []UserConfiguration `json:"users,omitempty"`
	// Defaults to `no`. Can be `pro` or `anti`. If set to `pro` or `anti`, will create an affinity
	// group per machine set of the corresponding type
	Affinity string `json:"affinity,omitempty"`
	// AffinityGroupIds allows users to pass in a list of UUIDs for previously-created Affinity
	// Groups. Any VM’s created with this spec will be added to the affinity group, which will
	// dictate which physical host(s) they can be placed on. Affinity groups can be type “affinity”
	// or “anti-affinity” in CloudStack. If they are type “anti-affinity”, all VM’s in the group
	// must be on separate physical hosts for high availability. If they are type “affinity”, all
	// VM’s in the group must be on the same physical host for improved performance
	AffinityGroupIds []string `json:"affinityGroupIds,omitempty"`
	// UserCustomDetails allows users to pass in non-standard key value inputs, outside those
	// defined [here](https://github.com/shapeblue/cloudstack/blob/main/api/src/main/java/com/cloud/vm/VmDetailConstants.java)
	UserCustomDetails map[string]string `json:"userCustomDetails,omitempty"`
	// Symlinks create soft symbolic links folders. One use case is to use data disk to store logs
	Symlinks SymlinkMaps `json:"symlinks,omitempty"`
}

type SymlinkMaps map[string]string

type CloudStackResourceDiskOffering struct {
	CloudStackResourceIdentifier `json:",inline"`
	// disk size in GB, > 0 for customized disk offering; = 0 for non-customized disk offering
	// +optional
	CustomSize int64 `json:"customSizeInGB,omitempty"`
	// path the filesystem will use to mount in VM
	MountPath string `json:"mountPath"`
	// device name of the disk offering in VM, shows up in lsblk command
	Device string `json:"device"`
	// filesystem used to mkfs in disk offering partition
	Filesystem string `json:"filesystem"`
	// disk label used to label disk partition
	Label string `json:"label"`
}

func (r *CloudStackResourceDiskOffering) Equal(o *CloudStackResourceDiskOffering) bool {
	_ = "STUB: not implemented"
	return false
}

// IsEmpty Introduced for backwards compatibility purposes. When CloudStackResourceDiskOffering
// was initially added to the CloudStackMachineConfig type, it was added with omitempty at the top
// level, but the subtypes were *not* optional, so we have old clusters today with zero value
// fields for the CloudStackResourceDiskOffering.
//
// Since then, we have made DiskOffering an optional pointer, with everything inside it as optional.
// Functionally, setting DiskOffering=nil is equivalent to a CloudStackResourceDiskOffering with
// zero values. Introducing this check should help prevent unintended RollingUpgrades when
// upgrading a cluster which has this "empty" DiskOffering in it.
func (r *CloudStackResourceDiskOffering) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (r *CloudStackResourceDiskOffering) Validate() (err error, field string, value string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func (r SymlinkMaps) Validate() (err error, field string, value string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

// ValidateUsers verifies a CloudStackMachineConfig object must have a users with ssh authorized keys.
// This validation only runs in CloudStackMachineConfig validation webhook, as we support
// auto-generate and import ssh key when creating a cluster via CLI.
func (c *CloudStackMachineConfig) ValidateUsers() error { _ = "STUB: not implemented"; return nil }

func (c *CloudStackMachineConfig) PauseReconcile() { _ = "STUB: not implemented"; return }

func (c *CloudStackMachineConfig) IsReconcilePaused() bool { _ = "STUB: not implemented"; return false }

func (c *CloudStackMachineConfig) SetControlPlane() { _ = "STUB: not implemented"; return }

func (c *CloudStackMachineConfig) IsControlPlane() bool { _ = "STUB: not implemented"; return false }

func (c *CloudStackMachineConfig) SetEtcd() { _ = "STUB: not implemented"; return }

func (c *CloudStackMachineConfig) IsEtcd() bool { _ = "STUB: not implemented"; return false }

func (c *CloudStackMachineConfig) SetManagement(clusterName string) {
	_ = "STUB: not implemented"
	return
}

func (c *CloudStackMachineConfig) IsManagement() bool { _ = "STUB: not implemented"; return false }

func (c *CloudStackMachineConfig) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (c *CloudStackMachineConfig) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CloudStackMachineConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// SetUserDefaults initializes Spec.Users for the CloudStackMachineConfig with default values.
// This only runs in the CLI, as we don't support user defaults through the webhook.
func (c *CloudStackMachineConfig) SetUserDefaults() { _ = "STUB: not implemented"; return }

// CloudStackMachineConfigStatus defines the observed state of CloudStackMachineConfig.
type CloudStackMachineConfigStatus struct {
	// SpecValid is set to true if cloudstackmachineconfig is validated.
	SpecValid bool `json:"specValid,omitempty"`

	// FailureMessage indicates that there is a fatal problem reconciling the
	// state, and will be set to a descriptive error message.
	// +optional
	FailureMessage *string `json:"failureMessage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CloudStackMachineConfig is the Schema for the cloudstackmachineconfigs API.
type CloudStackMachineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudStackMachineConfigSpec   `json:"spec,omitempty"`
	Status CloudStackMachineConfigStatus `json:"status,omitempty"`
}

func (c *CloudStackMachineConfig) OSFamily() OSFamily {
	_ = "STUB: not implemented"
	// This method must be defined to implement the providers.MachineConfig interface, but it's not actually used
	return *new(OSFamily)
}

// Users returns a list of configuration for OS users.
func (c *CloudStackMachineConfig) Users() []UserConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudStackMachineConfigSpec) Equal(o *CloudStackMachineConfigSpec) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CloudStackMachineConfig) ConvertConfigToConfigGenerateStruct() *CloudStackMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudStackMachineConfig) Marshallable() Marshallable {
	_ = "STUB: not implemented"
	return *new(Marshallable)
}

// +kubebuilder:object:generate=false

// CloudStackMachineConfigGenerate the same as CloudStackMachineConfig except stripped down for
// generation of yaml file during generate  clusterconfig.
type CloudStackMachineConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec CloudStackMachineConfigSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// CloudStackMachineConfigList contains a list of CloudStackMachineConfig.
type CloudStackMachineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudStackMachineConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudStackMachineConfig{}, &CloudStackMachineConfigList{})
}
