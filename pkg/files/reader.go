package files

import (
	"crypto/x509"
	"embed"
	"net/http"
	"net/url"
)

const (
	httpsScheme = "https"
	embedScheme = "embed"
)

type Reader struct {
	embedFS    embed.FS
	httpClient *http.Client
	userAgent  string
}

type ReaderOpt func(*Reader)

func WithEmbedFS(embedFS embed.FS) ReaderOpt { _ = "STUB: not implemented"; return *new(ReaderOpt) }

func WithUserAgent(userAgent string) ReaderOpt { _ = "STUB: not implemented"; return *new(ReaderOpt) }

// WithEKSAUserAgent sets the user agent for a particular eks-a component and version.
// component should be something like "cli", "controller", "e2e", etc.
// version should generally be a semver, but when not available, any string is valid.
func WithEKSAUserAgent(eksAComponent, version string) ReaderOpt {
	_ = "STUB: not implemented"
	return *new(ReaderOpt)
}

// WithRootCACerts configures the HTTP client's trusted CAs. Note that this will overwrite
// the defaults so the host's trust will be ignored. This option is only for testing.
func WithRootCACerts(certs []*x509.Certificate) ReaderOpt {
	_ = "STUB: not implemented"
	return *new(ReaderOpt)
}

// WithNonCachedProxyConfig configures the HTTP client to read the Proxy configuration
// from the environment on every request instead of relying on the default package
// level cache (implemented in the http package with envProxyFuncValue), which is only
// read once. If Proxy is not configured in the client's transport, nothing is changed.
// This is only for testing.
func WithNonCachedProxyConfig() ReaderOpt { _ = "STUB: not implemented"; return *new(ReaderOpt) }

func NewReader(opts ...ReaderOpt) *Reader {
	_ = "STUB: not implemented"
	// In order to modify the TLSHandshakeTimeout we first clone the default transport.
	// It has some defaults that we want to preserve. In particular Proxy, which is set
	// to http.ProxyFromEnvironment. This will make the client honor the HTTP_PROXY,
	// HTTPS_PROXY and NO_PROXY env variables.
	return nil
}

func (r *Reader) ReadFile(uri string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reader) readHttpFile(uri string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) readEmbedFile(url *url.URL) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readLocalFile(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func eksaUserAgent(eksAComponent, version string) string { _ = "STUB: not implemented"; return "" }
