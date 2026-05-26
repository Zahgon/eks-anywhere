package clusterapi

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

type Installer struct {
	*clients
}

func NewInstaller(capiClient CAPIClient, kubectlClient KubectlClient) *Installer {
	_ = "STUB: not implemented"
	return nil
}

// EnsureEtcdProvidersInstallation ensures that the CAPI etcd providers are installed in the management cluster.
func (i *Installer) EnsureEtcdProvidersInstallation(ctx context.Context, managementCluster *types.Cluster, provider providers.Provider, managementComponents *cluster.ManagementComponents, currSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}
