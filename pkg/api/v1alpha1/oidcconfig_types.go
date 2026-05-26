package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// OIDCConfig defines an OpenID Connect (OIDCConfigSpec) identity provider configuration

// OIDCConfigSpec defines the desired state of OIDCConfig.
type OIDCConfigSpec struct {
	// ClientId defines the client ID for the OpenID Connect client
	ClientId string `json:"clientId,omitempty"`
	// +kubebuilder:validation:Optional
	// GroupsClaim defines the name of a custom OpenID Connect claim for specifying user groups
	GroupsClaim string `json:"groupsClaim,omitempty"`
	// +kubebuilder:validation:Optional
	// GroupsPrefix defines a string to be prefixed to all groups to prevent conflicts with other authentication strategies
	GroupsPrefix string `json:"groupsPrefix,omitempty"`
	// IssuerUrl defines the URL of the OpenID issuer, only HTTPS scheme will be accepted
	IssuerUrl string `json:"issuerUrl,omitempty"`
	// +kubebuilder:validation:Optional
	// RequiredClaims defines a key=value pair that describes a required claim in the ID Token
	RequiredClaims []OIDCConfigRequiredClaim `json:"requiredClaims,omitempty"`
	// +kubebuilder:validation:Optional
	// UsernameClaim defines the OpenID claim to use as the user name. Note that claims other than the default ('sub') is not guaranteed to be unique and immutable
	UsernameClaim string `json:"usernameClaim,omitempty"`
	// +kubebuilder:validation:Optional
	// UsernamePrefix defines a string to prefixed to all usernames. If not provided, username claims other than 'email' are prefixed by the issuer URL to avoid clashes. To skip any prefixing, provide the value '-'.
	UsernamePrefix string `json:"usernamePrefix,omitempty"`
}

func (e *OIDCConfigSpec) Equal(n *OIDCConfigSpec) bool { _ = "STUB: not implemented"; return false }

func RequiredClaimsSliceEqual(a, b []OIDCConfigRequiredClaim) bool {
	_ = "STUB: not implemented"
	return false
}

// IsManaged returns true if the oidcconfig is associated with a workload cluster.
func (c *OIDCConfig) IsManaged() bool { _ = "STUB: not implemented"; return false }

func (c *OIDCConfig) SetManagedBy(clusterName string) { _ = "STUB: not implemented"; return }

type OIDCConfigRequiredClaim struct {
	Claim string `json:"claim,omitempty"`
	Value string `json:"value,omitempty"`
}

// OIDCConfigStatus defines the observed state of OIDCConfig.
type OIDCConfigStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OIDCConfig is the Schema for the oidcconfigs API.
type OIDCConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OIDCConfigSpec   `json:"spec,omitempty"`
	Status OIDCConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:generate=false
// Same as OIDCConfig except stripped down for generation of yaml file while writing to github repo when flux is enabled.
type OIDCConfigGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec OIDCConfigSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// OIDCConfigList contains a list of OIDCConfig.
type OIDCConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OIDCConfig `json:"items"`
}

func (c *OIDCConfig) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *OIDCConfig) ExpectedKind() string { _ = "STUB: not implemented"; return "" }

func (c *OIDCConfig) Validate() field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (c *OIDCConfig) ConvertConfigToConfigGenerateStruct() *OIDCConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	SchemeBuilder.Register(&OIDCConfig{}, &OIDCConfigList{})
}
