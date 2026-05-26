package validations

import (
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

type Opts struct {
	Kubectl            KubectlClient
	Spec               *cluster.Spec
	WorkloadCluster    *types.Cluster
	ManagementCluster  *types.Cluster
	Provider           providers.Provider
	TLSValidator       TlsValidator
	CliConfig          *config.CliConfig
	SkippedValidations map[string]bool
	CliVersion         string
	KubeClient         kubernetes.Client
	ManifestReader     *manifests.Reader
	BundlesOverride    string
}

func (o *Opts) SetDefaults() { _ = "STUB: not implemented"; return }
