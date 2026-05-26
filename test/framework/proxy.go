package framework

const (
	vsphereHttpProxyVar     = "T_HTTP_PROXY_VSPHERE"
	vsphereHttpsProxyVar    = "T_HTTPS_PROXY_VSPHERE"
	vsphereNoProxyVar       = "T_NO_PROXY_VSPHERE"
	cloudstackHttpProxyVar  = "T_HTTP_PROXY_CLOUDSTACK"
	cloudstackHttpsProxyVar = "T_HTTPS_PROXY_CLOUDSTACK"
	cloudstackNoProxyVar    = "T_NO_PROXY_CLOUDSTACK"
	tinkerbellHTTPProxyVar  = "T_HTTP_PROXY_TINKERBELL"
	tinkerbellHTTPSProxyVar = "T_HTTPS_PROXY_TINKERBELL"
	tinkerbellNoProxyVar    = "T_NO_PROXY_TINKERBELL"
)

var VsphereProxyRequiredEnvVars = ProxyRequiredEnvVars{
	HttpProxy:  vsphereHttpProxyVar,
	HttpsProxy: vsphereHttpsProxyVar,
	NoProxy:    vsphereNoProxyVar,
}

var CloudstackProxyRequiredEnvVars = ProxyRequiredEnvVars{
	HttpProxy:  cloudstackHttpProxyVar,
	HttpsProxy: cloudstackHttpsProxyVar,
	NoProxy:    cloudstackNoProxyVar,
}

// TinkerbellProxyRequiredEnvVars is for proxy related variables for tinkerbell.
var TinkerbellProxyRequiredEnvVars = ProxyRequiredEnvVars{
	HttpProxy:  tinkerbellHTTPProxyVar,
	HttpsProxy: tinkerbellHTTPSProxyVar,
	NoProxy:    tinkerbellNoProxyVar,
}

type ProxyRequiredEnvVars struct {
	HttpProxy  string
	HttpsProxy string
	NoProxy    string
}

func WithProxy(requiredEnvVars ProxyRequiredEnvVars) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}
