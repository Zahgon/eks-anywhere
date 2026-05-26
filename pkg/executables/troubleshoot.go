package executables

import (
	"context"
	"time"
)

const (
	troubleshootPath          = "support-bundle"
	supportBundleArchiveRegex = `support-bundle-([0-9]+(-[0-9]+)+)T([0-9]+(_[0-9]+)+)\.tar\.gz`
)

type Troubleshoot struct {
	Executable
}

func NewTroubleshoot(executable Executable) *Troubleshoot { _ = "STUB: not implemented"; return nil }

func (t *Troubleshoot) Collect(ctx context.Context, bundlePath string, sinceTime *time.Time, kubeconfig string) (archivePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *Troubleshoot) Analyze(ctx context.Context, bundleSpecPath string, archivePath string) ([]*SupportBundleAnalysis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseArchivePathFromCollectOutput(tsLogs string) (archivePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

type SupportBundleAnalysis struct {
	Title   string `json:"title"`
	IsPass  bool   `json:"isPass"`
	IsFail  bool   `json:"isFail"`
	IsWarn  bool   `json:"isWarn"`
	Message string `json:"message"`
	Uri     string `json:"URI"`
}
