package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

// Create is a schema for create cluster.
type Create struct {
	bootstrapper   interfaces.Bootstrapper
	clientFactory  interfaces.ClientFactory
	provider       providers.Provider
	clusterManager interfaces.ClusterManager
	gitOpsManager  interfaces.GitOpsManager
	writer         filewriter.FileWriter
	eksdInstaller  interfaces.EksdInstaller
	packageManager interfaces.PackageManager
	clusterCreator interfaces.ClusterCreator
	eksaInstaller  interfaces.EksaInstaller
	clusterMover   interfaces.ClusterMover
	iamAuth        interfaces.AwsIamAuth
}

// NewCreate builds a new create construct.
func NewCreate(bootstrapper interfaces.Bootstrapper,
	clientFactory interfaces.ClientFactory, provider providers.Provider,
	clusterManager interfaces.ClusterManager, gitOpsManager interfaces.GitOpsManager,
	writer filewriter.FileWriter, eksdInstaller interfaces.EksdInstaller,
	packageManager interfaces.PackageManager,
	clusterCreator interfaces.ClusterCreator,
	eksaInstaller interfaces.EksaInstaller,
	mover interfaces.ClusterMover,
	iamAuth interfaces.AwsIamAuth,
) *Create {
	_ = "STUB: not implemented"
	return nil
}

// Run runs all the create management cluster tasks.
func (c *Create) Run(ctx context.Context, clusterSpec *cluster.Spec, validator interfaces.Validator) error {
	_ = "STUB: not implemented"
	return nil
}
