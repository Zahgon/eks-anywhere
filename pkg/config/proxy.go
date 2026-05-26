package config

const (
	HttpsProxyKey = "HTTPS_PROXY"
	HttpProxyKey  = "HTTP_PROXY"
	NoProxyKey    = "NO_PROXY"
)

func GetProxyConfigFromEnv() map[string]string { _ = "STUB: not implemented"; return nil }
