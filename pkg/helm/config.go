package helm

import "github.com/aws/eks-anywhere/pkg/registrymirror"

// Config contains configuration options for Helm.
type Config struct {
	RegistryMirror *registrymirror.RegistryMirror
	ProxyConfig    map[string]string
	Insecure       bool
	ExtraFlags     []string
}

// NewConfig retuns a new helm Config.
func NewConfig(opts ...Opt) *Config { _ = "STUB: not implemented"; return nil }

// Opt is a functional option for configuring Helm behavior.
type Opt func(*Config)

// WithRegistryMirror sets up registry mirror for helm.
func WithRegistryMirror(mirror *registrymirror.RegistryMirror) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithInsecure configures helm to skip validating TLS certificates when
// communicating with the Kubernetes API.
func WithInsecure() Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithProxyConfig sets the proxy configurations on the Config.
// proxyConfig contains configurations for the proxy settings used by the Helm client.
// The valid keys are as follows:
// "HTTPS_PROXY": Specifies the HTTPS proxy URL.
// "HTTP_PROXY": Specifies the HTTP proxy URL.
// "NO_PROXY": Specifies a comma-separated a list of destination domain names, domains, IP addresses, or other network CIDRs to exclude proxying.
func WithProxyConfig(proxyConfig map[string]string) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithExtraFlags configures the extra flags to pass to helm.
func WithExtraFlags(extraFlags []string) Opt { _ = "STUB: not implemented"; return *new(Opt) }
