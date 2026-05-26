package awsiamauth

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/kubeconfig"
	"github.com/aws/eks-anywhere/pkg/types"
)

// KubernetesClient provides Kubernetes API access.
type KubernetesClient interface {
	Apply(ctx context.Context, cluster *types.Cluster, data []byte) error
	GetAPIServerURL(ctx context.Context, cluster *types.Cluster) (string, error)
	GetClusterCACert(ctx context.Context, cluster *types.Cluster, clusterName string) ([]byte, error)
	GetAWSIAMKubeconfigSecretValue(ctx context.Context, cluster *types.Cluster, clusterName string) ([]byte, error)
}

// Installer provides the necessary behavior for installing the AWS IAM Authenticator.
type Installer struct {
	k8s              KubernetesClient
	writer           filewriter.FileWriter
	kubeconfigWriter kubeconfig.Writer
}

// NewInstaller creates a new installer instance.
func NewInstaller(
	k8s KubernetesClient,
	writer filewriter.FileWriter,
	kubeconfigWriter kubeconfig.Writer,
) *Installer {
	_ = "STUB: not implemented"
	return nil
}

// GenerateWorkloadKubeconfig generates the AWS IAM auth kubeconfig.
func (i *Installer) GenerateWorkloadKubeconfig(
	ctx context.Context,
	management, workload *types.Cluster,
	spec *cluster.Spec,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateManagementKubeconfig generates the AWS IAM auth kubeconfig.
func (i *Installer) GenerateManagementKubeconfig(
	ctx context.Context,
	cluster *types.Cluster,
) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanupKubeconfig removes an existing AWS IAM kubeconfig file when AWS IAM is removed from cluster.
func (i *Installer) CleanupKubeconfig(clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}
