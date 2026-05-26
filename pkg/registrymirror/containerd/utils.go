package containerd

import (
	"net/url"
)

// ToAPIEndpoint turns URL to a valid API endpoint used in
// a containerd config file for a local registry.
// Original input is returned in case of malformed inputs.
func ToAPIEndpoint(url string) string { _ = "STUB: not implemented"; return "" }

func parseURL(in string) (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

// ToAPIEndpoints utilizes ToAPIEndpoint to turn all URLs from a
// map to valid API endpoints for a local registry.
func ToAPIEndpoints(URLs map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
