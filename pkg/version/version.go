package version

var gitVersion string

type Info struct {
	GitVersion         string `json:"version"`
	BundleManifestURL  string `json:"bundleManifestURL,omitempty"`
	ReleaseManifestURL string `json:"releaseManifestURL,omitempty"`
}

func Get() Info { _ = "STUB: not implemented"; return *new(Info) }

// GetFullVersionInfo returns the complete version information for the
// EKS Anywhere, including Git version and bundle manifest URL
// associated with this release.
func GetFullVersionInfo() (Info, error) { _ = "STUB: not implemented"; return *new(Info), nil }
