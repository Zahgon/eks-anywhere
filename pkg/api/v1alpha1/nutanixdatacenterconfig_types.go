// Important: Run "make generate" to regenerate code after modifying this file
// json tags are required; new fields must have json tags for the fields to be serialized

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NutanixDatacenterConfigSpec defines the desired state of NutanixDatacenterConfig.
type NutanixDatacenterConfigSpec struct {
	// Endpoint is the Endpoint of Nutanix Prism Central
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`

	// Port is the Port of Nutanix Prism Central
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Default=9440
	Port int `json:"port"`

	// AdditionalTrustBundle is the optional PEM-encoded certificate bundle for
	// users that configured their Prism Central with certificates from non-publicly
	// trusted CAs
	AdditionalTrustBundle string `json:"additionalTrustBundle,omitempty"`

	// Insecure is the optional flag to skip TLS verification. Nutanix Prism
	// Central installation by default ships with a self-signed certificate
	// that will fail TLS verification because the certificate is not issued by
	// a public CA and does not have the IP SANs with the Prism Central endpoint.
	// To accommodate the scenario where the user has not changed the default
	// Certificate that ships with Prism Central, we allow the user to skip TLS
	// verification. This is not recommended for production use.
	Insecure bool `json:"insecure,omitempty"`

	// CredentialRef is the reference to the secret name that contains the credentials
	// for the Nutanix Prism Central. The namespace for the secret is assumed to be a constant i.e. eksa-system.
	// +optional
	CredentialRef *Ref `json:"credentialRef,omitempty"`

	// FailureDomains is the optional list of failure domains for the Nutanix Datacenter.
	// +optional
	FailureDomains []NutanixDatacenterFailureDomain `json:"failureDomains,omitempty"`

	// CcmExcludeIPs is the optional list of IP addresses that should be excluded from the CCM IP pool for nodes.
	// List should be valid IP addresses and IP address ranges.
	// +optional
	CcmExcludeNodeIPs []string `json:"ccmExcludeNodeIPs,omitempty"`
}

// NutanixDatacenterFailureDomain defines the failure domain for the Nutanix Datacenter.
type NutanixDatacenterFailureDomain struct {
	// Name is the unique name of the failure domain.
	// Name must be between 1 and 64 characters long.
	// It must consist of only lower case alphanumeric characters and hyphens (-).
	// It must start and end with an alphanumeric character.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Name string `json:"name"`

	// Cluster is the Prism Element cluster name or uuid that is connected to the Prism Central.
	// +kubebuilder:validation:Required
	Cluster NutanixResourceIdentifier `json:"cluster,omitempty"`

	// Subnets holds the list of subnets identifiers cluster's network subnets.
	// +kubebuilder:validation:Required
	Subnets []NutanixResourceIdentifier `json:"subnets,omitempty"`

	// Worker Machine Groups holds the list of worker machine group names that will use this failure domain.
	// +optional
	WorkerMachineGroups []string `json:"workerMachineGroups,omitempty"`
}

// NutanixDatacenterConfigStatus defines the observed state of NutanixDatacenterConfig.
type NutanixDatacenterConfigStatus struct{}

// NutanixDatacenterConfig is the Schema for the NutanixDatacenterConfigs API
//
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type NutanixDatacenterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NutanixDatacenterConfigSpec   `json:"spec,omitempty"`
	Status NutanixDatacenterConfigStatus `json:"status,omitempty"`
}

func (in *NutanixDatacenterConfig) Kind() string { _ = "STUB: not implemented"; return "" }

func (in *NutanixDatacenterConfig) ExpectedKind() string { _ = "STUB: not implemented"; return "" }

func (in *NutanixDatacenterConfig) PauseReconcile() { _ = "STUB: not implemented"; return }

func (in *NutanixDatacenterConfig) IsReconcilePaused() bool {
	_ = "STUB: not implemented"
	return false
}

func (in *NutanixDatacenterConfig) ClearPauseAnnotation() { _ = "STUB: not implemented"; return }

func (in *NutanixDatacenterConfig) ConvertConfigToConfigGenerateStruct() *NutanixDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (in *NutanixDatacenterConfig) Marshallable() Marshallable {
	_ = "STUB: not implemented"
	return *new(Marshallable)
}

func (in *NutanixDatacenterConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func createValidateNutanixResourceFunc(msgPrefix, entityName, mfstName string) func(*NutanixResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

// SetDefaults sets default values for the NutanixDatacenterConfig object.
func (in *NutanixDatacenterConfig) SetDefaults() { _ = "STUB: not implemented"; return }

// NutanixDatacenterConfigGenerate is same as NutanixDatacenterConfig except stripped down for generation of yaml file during generate clusterconfig
//
// +kubebuilder:object:generate=false
type NutanixDatacenterConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec NutanixDatacenterConfigSpec `json:"spec,omitempty"`
}

// NutanixDatacenterConfigList contains a list of NutanixDatacenterConfig
//
// +kubebuilder:object:root=true
type NutanixDatacenterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NutanixDatacenterConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NutanixDatacenterConfig{}, &NutanixDatacenterConfigList{})
}
