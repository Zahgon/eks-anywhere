package framework

import (
	"github.com/aws/eks-anywhere/pkg/semver"
	"github.com/aws/eks-anywhere/pkg/version"
)

func newVersion(version string) *semver.Version { _ = "STUB: not implemented"; return nil }

// localEKSAVersion returns the version of eks-anywhere installed locally.
func localEKSAVersion() (string, error) { _ = "STUB: not implemented"; return "", nil }

// localEKSAVersionCommand returns the output of the eks-anywhere version command.
func localEKSAVersionCommand() (version.Info, error) {
	_ = "STUB: not implemented"
	return *new(version.Info), nil
}
