package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

// Create is the workflow that creates a workload clusters.
type Create struct {
	clientFactory    interfaces.ClientFactory
	provider         providers.Provider
	clusterManager   interfaces.ClusterManager
	gitOpsManager    interfaces.GitOpsManager
	writer           filewriter.FileWriter
	eksdInstaller    interfaces.EksdInstaller
	clusterCreator   interfaces.ClusterCreator
	packageInstaller interfaces.PackageManager
	iamAuth          interfaces.AwsIamAuth
}

// NewCreate builds a new create construct.
func NewCreate(provider providers.Provider,
	clusterManager interfaces.ClusterManager, gitOpsManager interfaces.GitOpsManager,
	writer filewriter.FileWriter,
	eksdInstaller interfaces.EksdInstaller,
	packageInstaller interfaces.PackageManager,
	clusterCreator interfaces.ClusterCreator,
	clientFactory interfaces.ClientFactory,
	iamAuth interfaces.AwsIamAuth,
) *Create {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the tasks to create a workload cluster.
func (c *Create) Run(ctx context.Context, clusterSpec *cluster.Spec, validator interfaces.Validator) error {
	_ = "STUB: not implemented"
	return nil
}
