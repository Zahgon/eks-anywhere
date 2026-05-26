package conformance

const (
	destinationFile   = "sonobuoy"
	sonobouyGitHubAPI = "https://api.github.com/repos/vmware-tanzu/sonobuoy/releases/latest"
)

type githubRelease struct {
	Assets []asset `json:"assets"`
}

type asset struct {
	BrowserDownloadURL string `json:"browser_download_url"`
}

func Download() error { _ = "STUB: not implemented"; return nil }
