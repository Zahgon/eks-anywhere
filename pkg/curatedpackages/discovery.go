package curatedpackages

import (
	"k8s.io/apimachinery/pkg/version"
)

// Discovery
/**
Implements ServerVersionInterface to provide the Kubernetes client version to be used.
*/
type Discovery struct {
	kubeVersion *KubeVersion
}

type KubeVersion struct {
	major string
	minor string
}

func NewDiscovery(kubeVersion *KubeVersion) *Discovery { _ = "STUB: not implemented"; return nil }

func NewKubeVersion(major string, minor string) *KubeVersion { _ = "STUB: not implemented"; return nil }

func (d *Discovery) ServerVersion() (*version.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
