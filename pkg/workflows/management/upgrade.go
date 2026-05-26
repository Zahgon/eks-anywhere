package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

// Upgrade is a schema for upgrade cluster.
type Upgrade struct {
	clientFactory     interfaces.ClientFactory
	provider          providers.Provider
	clusterManager    interfaces.ClusterManager
	gitOpsManager     interfaces.GitOpsManager
	writer            filewriter.FileWriter
	capiManager       interfaces.CAPIManager
	eksdInstaller     interfaces.EksdInstaller
	eksdUpgrader      interfaces.EksdUpgrader
	upgradeChangeDiff *types.ChangeDiff
	clusterUpgrader   interfaces.ClusterUpgrader
	packageManager    interfaces.PackageManager
	iamAuth           interfaces.AwsIamAuth
}

// NewUpgrade builds a new upgrade construct.
func NewUpgrade(clientFactory interfaces.ClientFactory, provider providers.Provider,
	capiManager interfaces.CAPIManager,
	clusterManager interfaces.ClusterManager,
	gitOpsManager interfaces.GitOpsManager,
	writer filewriter.FileWriter,
	eksdUpgrader interfaces.EksdUpgrader,
	eksdInstaller interfaces.EksdInstaller,
	clusterUpgrade interfaces.ClusterUpgrader,
	packageManager interfaces.PackageManager,
	iamAuth interfaces.AwsIamAuth,
) *Upgrade {
	_ = "STUB: not implemented"
	return nil
}

// Run Upgrade implements upgrade functionality for management cluster's upgrade operation.
func (c *Upgrade) Run(ctx context.Context, clusterSpec *cluster.Spec, managementCluster *types.Cluster, validator interfaces.Validator) error {
	_ = "STUB: not implemented"
	return nil
}
