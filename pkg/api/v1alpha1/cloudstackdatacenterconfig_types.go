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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const DefaultCloudStackAZPrefix = "default-az"

// CloudStackDatacenterConfigSpec defines the desired state of CloudStackDatacenterConfig.
type CloudStackDatacenterConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Domain contains a grouping of accounts. Domains usually contain multiple accounts that have some logical relationship to each other and a set of delegated administrators with some authority over the domain and its subdomains
	// This field is considered as a fully qualified domain name which is the same as the domain path without "ROOT/" prefix. For example, if "foo" is specified then a domain with "ROOT/foo" domain path is picked.
	// The value "ROOT" is a special case that points to "the" ROOT domain of the CloudStack. That is, a domain with a path "ROOT/ROOT" is not allowed.
	// +optional
	// Deprecated: Please use AvailabilityZones instead
	Domain string `json:"domain,omitempty"`
	// Zones is a list of one or more zones that are managed by a single CloudStack management endpoint.
	// +optional
	// Deprecated: Please use AvailabilityZones instead
	Zones []CloudStackZone `json:"zones,omitempty"`
	// Account typically represents a customer of the service provider or a department in a large organization. Multiple users can exist in an account, and all CloudStack resources belong to an account. Accounts have users and users have credentials to operate on resources within that account. If an account name is provided, a domain must also be provided.
	// +optional
	// Deprecated: Please use AvailabilityZones instead
	Account string `json:"account,omitempty"`
	// CloudStack Management API endpoint's IP. It is added to VM's noproxy list
	// +optional
	// Deprecated: Please use AvailabilityZones instead
	ManagementApiEndpoint string `json:"managementApiEndpoint,omitempty"`
	// AvailabilityZones list of different partitions to distribute VMs across - corresponds to a list of CAPI failure domains
	AvailabilityZones []CloudStackAvailabilityZone `json:"availabilityZones,omitempty"`
}

type CloudStackResourceIdentifier struct {
	// Id of a resource in the CloudStack environment. Mutually exclusive with Name
	// +optional
	Id string `json:"id,omitempty"`
	// Name of a resource in the CloudStack environment. Mutually exclusive with Id
	// +optional
	Name string `json:"name,omitempty"`
}

func (r *CloudStackResourceIdentifier) Equal(o *CloudStackResourceIdentifier) bool {
	_ = "STUB: not implemented"
	return false
}

// CloudStackZone is an organizational construct typically used to represent a single datacenter, and all its physical and virtual resources exist inside that zone. It can either be specified as a UUID or name.
type CloudStackZone struct {
	// Zone is the name or UUID of the CloudStack zone in which clusters should be created. Zones should be managed by a single CloudStack Management endpoint.
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Network is the name or UUID of the CloudStack network in which clusters should be created. It can either be an isolated or shared network. If it doesn’t already exist in CloudStack, it’ll automatically be created by CAPC as an isolated network. It can either be specified as a UUID or name
	// In multiple-zones situation, only 'Shared' network is supported.
	Network CloudStackResourceIdentifier `json:"network"`
}

// CloudStackAvailabilityZone maps to a CAPI failure domain to distribute machines across Cloudstack infrastructure.
type CloudStackAvailabilityZone struct {
	// Name is used as a unique identifier for each availability zone
	Name string `json:"name"`
	// CredentialRef is used to reference a secret in the eksa-system namespace
	CredentialsRef string `json:"credentialsRef"`
	// Zone represents the properties of the CloudStack zone in which clusters should be created, like the network.
	Zone CloudStackZone `json:"zone"`
	// Domain contains a grouping of accounts. Domains usually contain multiple accounts that have some logical relationship to each other and a set of delegated administrators with some authority over the domain and its subdomains
	// This field is considered as a fully qualified domain name which is the same as the domain path without "ROOT/" prefix. For example, if "foo" is specified then a domain with "ROOT/foo" domain path is picked.
	// The value "ROOT" is a special case that points to "the" ROOT domain of the CloudStack. That is, a domain with a path "ROOT/ROOT" is not allowed.
	Domain string `json:"domain"`
	// Account typically represents a customer of the service provider or a department in a large organization. Multiple users can exist in an account, and all CloudStack resources belong to an account. Accounts have users and users have credentials to operate on resources within that account. If an account name is provided, a domain must also be provided.
	Account string `json:"account,omitempty"`
	// CloudStack Management API endpoint's IP. It is added to VM's noproxy list
	ManagementApiEndpoint string `json:"managementApiEndpoint"`
}

// CloudStackDatacenterConfigStatus defines the observed state of CloudStackDatacenterConfig.
type CloudStackDatacenterConfigStatus struct { // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// SpecValid is set to true if cloudstackdatacenterconfig is validated.
	SpecValid bool `json:"specValid,omitempty"`

	// ObservedGeneration is the latest generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// FailureMessage indicates that there is a fatal problem reconciling the
	// state, and will be set to a descriptive error message.
	FailureMessage *string `json:"failureMessage,omitempty"`
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CloudStackDatacenterConfig is the Schema for the cloudstackdatacenterconfigs API.
type CloudStackDatacenterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudStackDatacenterConfigSpec   `json:"spec,omitempty"`
	Status CloudStackDatacenterConfigStatus `json:"status,omitempty"`
}

func (v *CloudStackDatacenterConfig) Kind() string { _ = "STUB: not implemented"; return "" }

func (v *CloudStackDatacenterConfig) ExpectedKind() string { _ = "STUB: not implemented"; return "" }

func (v *CloudStackDatacenterConfig) PauseReconcile() { _ = "STUB: not implemented"; return }

func (v *CloudStackDatacenterConfig) IsReconcilePaused() bool {
	_ = "STUB: not implemented"
	return false
}

func (v *CloudStackDatacenterConfig) ClearPauseAnnotation() { _ = "STUB: not implemented"; return }

func (v *CloudStackDatacenterConfig) ConvertConfigToConfigGenerateStruct() *CloudStackDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (v *CloudStackDatacenterConfig) Marshallable() Marshallable {
	_ = "STUB: not implemented"
	return *new(Marshallable)
}

func (v *CloudStackDatacenterConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (v *CloudStackDatacenterConfig) SetDefaults() { _ = "STUB: not implemented"; return }

func (s *CloudStackDatacenterConfigSpec) Equal(o *CloudStackDatacenterConfigSpec) bool {
	_ = "STUB: not implemented"
	return false
}

func (z *CloudStackZone) Equal(o *CloudStackZone) bool { _ = "STUB: not implemented"; return false }

func (az *CloudStackAvailabilityZone) Equal(o *CloudStackAvailabilityZone) bool {
	_ = "STUB: not implemented"
	return false
}

// +kubebuilder:object:generate=false

// Same as CloudStackDatacenterConfig except stripped down for generation of yaml file during generate clusterconfig.
type CloudStackDatacenterConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec CloudStackDatacenterConfigSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// CloudStackDatacenterConfigList contains a list of CloudStackDatacenterConfig.
type CloudStackDatacenterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudStackDatacenterConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudStackDatacenterConfig{}, &CloudStackDatacenterConfigList{})
}
