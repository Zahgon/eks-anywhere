package framework

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	RegistryEndpointVar                  = "T_REGISTRY_MIRROR_ENDPOINT"
	RegistryPortVar                      = "T_REGISTRY_MIRROR_PORT"
	RegistryUsernameVar                  = "T_REGISTRY_MIRROR_USERNAME"
	RegistryPasswordVar                  = "T_REGISTRY_MIRROR_PASSWORD"
	RegistryCACertVar                    = "T_REGISTRY_MIRROR_CA_CERT"
	RegistryEndpointTinkerbellVar        = "T_REGISTRY_MIRROR_ENDPOINT_TINKERBELL"
	RegistryPortTinkerbellVar            = "T_REGISTRY_MIRROR_PORT_TINKERBELL"
	RegistryUsernameTinkerbellVar        = "T_REGISTRY_MIRROR_USERNAME_TINKERBELL"
	RegistryPasswordTinkerbellVar        = "T_REGISTRY_MIRROR_PASSWORD_TINKERBELL"
	RegistryCACertTinkerbellVar          = "T_REGISTRY_MIRROR_CA_CERT_TINKERBELL"
	RegistryMirrorDefaultSecurityGroup   = "T_REGISTRY_MIRROR_DEFAULT_SECURITY_GROUP"
	RegistryMirrorAirgappedSecurityGroup = "T_REGISTRY_MIRROR_AIRGAPPED_SECURITY_GROUP"
	PrivateRegistryEndpointVar           = "T_PRIVATE_REGISTRY_MIRROR_ENDPOINT"
	PrivateRegistryPortVar               = "T_PRIVATE_REGISTRY_MIRROR_PORT"
	PrivateRegistryUsernameVar           = "T_PRIVATE_REGISTRY_MIRROR_USERNAME"
	PrivateRegistryPasswordVar           = "T_PRIVATE_REGISTRY_MIRROR_PASSWORD"
	PrivateRegistryCACertVar             = "T_PRIVATE_REGISTRY_MIRROR_CA_CERT"
	PrivateRegistryEndpointTinkerbellVar = "T_PRIVATE_REGISTRY_MIRROR_ENDPOINT_TINKERBELL"
	PrivateRegistryPortTinkerbellVar     = "T_PRIVATE_REGISTRY_MIRROR_PORT_TINKERBELL"
	PrivateRegistryUsernameTinkerbellVar = "T_PRIVATE_REGISTRY_MIRROR_USERNAME_TINKERBELL"
	PrivateRegistryPasswordTinkerbellVar = "T_PRIVATE_REGISTRY_MIRROR_PASSWORD_TINKERBELL"
	PrivateRegistryCACertTinkerbellVar   = "T_PRIVATE_REGISTRY_MIRROR_CA_CERT_TINKERBELL"

	RegistryMirrorOciNamespacesRegistry1Var  = "T_REGISTRY_MIRROR_OCINAMESPACES_REGISTRY1"
	RegistryMirrorOciNamespacesNamespace1Var = "T_REGISTRY_MIRROR_OCINAMESPACES_NAMESPACE1"
	RegistryMirrorOciNamespacesRegistry2Var  = "T_REGISTRY_MIRROR_OCINAMESPACES_REGISTRY2"
	RegistryMirrorOciNamespacesNamespace2Var = "T_REGISTRY_MIRROR_OCINAMESPACES_NAMESPACE2"
)

var (
	registryMirrorRequiredEnvVars                  = []string{RegistryEndpointVar, RegistryPortVar, RegistryUsernameVar, RegistryPasswordVar, RegistryCACertVar}
	registryMirrorTinkerbellRequiredEnvVars        = []string{RegistryEndpointTinkerbellVar, RegistryPortTinkerbellVar, RegistryUsernameTinkerbellVar, RegistryPasswordTinkerbellVar, RegistryCACertTinkerbellVar}
	registryMirrorDockerAirgappedRequiredEnvVars   = []string{RegistryMirrorDefaultSecurityGroup, RegistryMirrorAirgappedSecurityGroup}
	privateRegistryMirrorRequiredEnvVars           = []string{PrivateRegistryEndpointVar, PrivateRegistryPortVar, PrivateRegistryUsernameVar, PrivateRegistryPasswordVar, PrivateRegistryCACertVar}
	privateRegistryMirrorTinkerbellRequiredEnvVars = []string{PrivateRegistryEndpointTinkerbellVar, PrivateRegistryPortTinkerbellVar, PrivateRegistryUsernameTinkerbellVar, PrivateRegistryPasswordTinkerbellVar, PrivateRegistryCACertTinkerbellVar}
	registryMirrorOciNamespacesRequiredEnvVars     = []string{RegistryMirrorOciNamespacesRegistry1Var, RegistryMirrorOciNamespacesNamespace1Var}
)

// WithRegistryMirrorInsecureSkipVerify sets up e2e for registry mirrors with InsecureSkipVerify option.
func WithRegistryMirrorInsecureSkipVerify(providerName string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithRegistryMirrorEndpointAndCert sets up e2e for registry mirrors.
func WithRegistryMirrorEndpointAndCert(providerName string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithRegistryMirrorOciNamespaces sets up e2e for registry mirrors with ocinamespaces.
func WithRegistryMirrorOciNamespaces(providerName string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// DefaultOciNamespaces returns the default OCI namespaces from environment variables.
func DefaultOciNamespaces(e *ClusterE2ETest) []v1alpha1.OCINamespace {
	_ = "STUB: not implemented"
	return nil
}

// WithAuthenticatedRegistryMirror sets up e2e for authenticated registry mirrors.
func WithAuthenticatedRegistryMirror(providerName string, optNamespaces ...v1alpha1.OCINamespace) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// Set env vars for helm login/push

func RequiredRegistryMirrorEnvVars() []string { _ = "STUB: not implemented"; return nil }

// RequiredOciNamespacesEnvVars returns the Env variables to set for OCI Namespaces tests.
func RequiredOciNamespacesEnvVars() []string { _ = "STUB: not implemented"; return nil }

func setupRegistryMirrorEndpointAndCert(e *ClusterE2ETest, providerName string, insecureSkipVerify bool, ociNamespaces ...v1alpha1.OCINamespace) {
	_ = "STUB: not implemented"
	return
}

// Set env vars for helm login/push

// SetRegistryMirrorDefaultInstanceSecurityGroupOnCleanup sets the instance security group to the registry mirror default security group on cleanup.
func (e *ClusterE2ETest) SetRegistryMirrorDefaultInstanceSecurityGroupOnCleanup(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}
