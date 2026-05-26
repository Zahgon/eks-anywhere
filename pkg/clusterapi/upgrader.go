package clusterapi

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

type Upgrader struct {
	*clients
}

func NewUpgrader(capiClient CAPIClient, kubectlClient KubectlClient) *Upgrader {
	_ = "STUB: not implemented"
	return nil
}

// Upgrade checks whether upgrading the CAPI components is necessary and, if so, upgrades them the new versions.
func (u *Upgrader) Upgrade(ctx context.Context, managementCluster *types.Cluster, provider providers.Provider, currentManagementComponents, newManagementComponents *cluster.ManagementComponents, newSpec *cluster.Spec) (*types.ChangeDiff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CAPIChangeDiff struct {
	CertManager            *types.ComponentChangeDiff
	Core                   *types.ComponentChangeDiff
	ControlPlane           *types.ComponentChangeDiff
	BootstrapProviders     []types.ComponentChangeDiff
	InfrastructureProvider *types.ComponentChangeDiff
}

func (c *CAPIChangeDiff) toChangeDiff() *types.ChangeDiff { _ = "STUB: not implemented"; return nil }

// ChangeDiff generates a version change diff for the CAPI components.
func ChangeDiff(currentManagementComponents, newManagementComponents *cluster.ManagementComponents, provider providers.Provider) *types.ChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func capiChangeDiff(currentManagementComponents, newManagementComponents *cluster.ManagementComponents, provider providers.Provider) *CAPIChangeDiff {
	_ = "STUB: not implemented"
	return nil
}
