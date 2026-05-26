package workload

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
	clientFactory    interfaces.ClientFactory
	provider         providers.Provider
	clusterManager   interfaces.ClusterManager
	gitOpsManager    interfaces.GitOpsManager
	writer           filewriter.FileWriter
	eksdInstaller    interfaces.EksdInstaller
	clusterUpgrader  interfaces.ClusterUpgrader
	packageInstaller interfaces.PackageManager
	iamAuth          interfaces.AwsIamAuth
}

// NewUpgrade builds a new upgrade construct.
func NewUpgrade(clientFactory interfaces.ClientFactory,
	provider providers.Provider,
	clusterManager interfaces.ClusterManager, gitOpsManager interfaces.GitOpsManager,
	writer filewriter.FileWriter,
	clusterUpgrader interfaces.ClusterUpgrader,
	eksdInstaller interfaces.EksdInstaller,
	packageInstaller interfaces.PackageManager,
	iamAuth interfaces.AwsIamAuth,
) *Upgrade {
	_ = "STUB: not implemented"
	return nil
}

// Run Upgrade implements upgrade functionality for workload cluster's upgrade operation.
func (c *Upgrade) Run(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec, validator interfaces.Validator) error {
	_ = "STUB: not implemented"
	return nil
}
