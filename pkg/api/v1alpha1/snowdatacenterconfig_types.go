package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SnowIdentityKind    = "Secret"
	SnowCredentialsKey  = "credentials"
	SnowCertificatesKey = "ca-bundle"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SnowDatacenterConfigSpec defines the desired state of SnowDatacenterConfig.
type SnowDatacenterConfigSpec struct { // Important: Run "make generate" to regenerate code after modifying this file

	// IdentityRef is a reference to an identity for the Snow API to be used when reconciling this cluster
	IdentityRef Ref `json:"identityRef,omitempty"`
}

// SnowDatacenterConfigStatus defines the observed state of SnowDatacenterConfig.
type SnowDatacenterConfigStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SnowDatacenterConfig is the Schema for the SnowDatacenterConfigs API.
type SnowDatacenterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnowDatacenterConfigSpec   `json:"spec,omitempty"`
	Status SnowDatacenterConfigStatus `json:"status,omitempty"`
}

func (s *SnowDatacenterConfig) Kind() string { _ = "STUB: not implemented"; return "" }

func (s *SnowDatacenterConfig) ExpectedKind() string { _ = "STUB: not implemented"; return "" }

func (s *SnowDatacenterConfig) PauseReconcile() { _ = "STUB: not implemented"; return }

func (s *SnowDatacenterConfig) ClearPauseAnnotation() { _ = "STUB: not implemented"; return }

func (s *SnowDatacenterConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (s *SnowDatacenterConfig) ConvertConfigToConfigGenerateStruct() *SnowDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (s *SnowDatacenterConfig) Marshallable() Marshallable {
	_ = "STUB: not implemented"
	return *new(Marshallable)
}

// +kubebuilder:object:generate=false

// Same as SnowDatacenterConfig except stripped down for generation of yaml file during generate clusterconfig.
type SnowDatacenterConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec SnowDatacenterConfigSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// SnowDatacenterConfigList contains a list of SnowDatacenterConfig.
type SnowDatacenterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnowDatacenterConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SnowDatacenterConfig{}, &SnowDatacenterConfigList{})
}
