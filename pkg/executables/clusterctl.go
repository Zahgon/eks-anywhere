package executables

import (
	"context"
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/cluster"
	anywherecluster "github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	clusterCtlPath                = "clusterctl"
	clusterctlConfigFile          = "clusterctl_tmp.yaml"
	capiPrefix                    = "/generated/overrides"
	etcdadmBootstrapProviderName  = "etcdadm-bootstrap"
	etcdadmControllerProviderName = "etcdadm-controller"
	kubeadmBootstrapProviderName  = "kubeadm"
)

//go:embed config/clusterctl.yaml
var clusterctlConfigTemplate string

type Clusterctl struct {
	Executable
	writer filewriter.FileWriter
	reader manifests.FileReader
}

type clusterctlConfiguration struct {
	coreVersion              string
	bootstrapVersion         string
	controlPlaneVersion      string
	configFile               string
	etcdadmBootstrapVersion  string
	etcdadmControllerVersion string
}

// NewClusterctl builds a new [Clusterctl].
func NewClusterctl(executable Executable, writer filewriter.FileWriter, reader manifests.FileReader) *Clusterctl {
	_ = "STUB: not implemented"
	return nil
}

func imageRepository(image v1alpha1.Image) string { _ = "STUB: not implemented"; return "" }

// This method will write the configuration files
// used by cluster api to install components.
// See: https://cluster-api.sigs.k8s.io/clusterctl/configuration.html
func (c *Clusterctl) buildOverridesLayer(managementComponents *cluster.ManagementComponents, clusterName string, provider providers.Provider) error {
	_ = "STUB: not implemented"
	// Adding cluster name to path temporarily following suggestion.
	//
	// This adds an implicit dependency between this method
	// and the writer passed to NewClusterctl
	// Ideally the writer implementation should be modified to
	// accept a path and file name and it should create the path in case it
	// does not exists.
	return nil
}

func (c *Clusterctl) writeInfrastructureBundle(rootFolder string, bundle *types.InfrastructureBundle) error {
	_ = "STUB: not implemented"
	return nil
}

// BackupManagement saves the CAPI resources of a cluster to the provided path. This will overwrite any existing contents
// in the path if the backup succeeds. If `clusterName` is provided, it filters and backs up only the provided cluster.
func (c *Clusterctl) BackupManagement(ctx context.Context, cluster *types.Cluster, managementStatePath, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

// MoveManagement moves management components `from` cluster `to` cluster
// If `clusterName` is provided, it filters and moves only the provided cluster.
func (c *Clusterctl) MoveManagement(ctx context.Context, from, to *types.Cluster, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Clusterctl) GetWorkloadKubeconfig(ctx context.Context, clusterName string, cluster *types.Cluster) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InitInfrastructure initializes the infrastructure for the cluster using clusterctl.
func (c *Clusterctl) InitInfrastructure(ctx context.Context, managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec, cluster *types.Cluster, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Clusterctl) buildConfig(managementComponents *anywherecluster.ManagementComponents, clusterName string, provider providers.Provider) (*clusterctlConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var providerNamespaces = map[string]string{
	constants.VSphereProviderName:    constants.CapvSystemNamespace,
	constants.DockerProviderName:     constants.CapdSystemNamespace,
	constants.CloudStackProviderName: constants.CapcSystemNamespace,
	constants.AWSProviderName:        constants.CapaSystemNamespace,
	constants.SnowProviderName:       constants.CapasSystemNamespace,
	constants.NutanixProviderName:    constants.CapxSystemNamespace,
	constants.TinkerbellProviderName: constants.CaptSystemNamespace,
	etcdadmBootstrapProviderName:     constants.EtcdAdmBootstrapProviderSystemNamespace,
	etcdadmControllerProviderName:    constants.EtcdAdmControllerSystemNamespace,
	kubeadmBootstrapProviderName:     constants.CapiKubeadmBootstrapSystemNamespace,
}

// Upgrade executes an upgrade of the cluster to the new management components and the spec.
func (c *Clusterctl) Upgrade(ctx context.Context, managementCluster *types.Cluster, provider providers.Provider, managementComponents *cluster.ManagementComponents, newSpec *cluster.Spec, changeDiff *clusterapi.CAPIChangeDiff) error {
	_ = "STUB: not implemented"
	return nil
}

// InstallEtcdadmProviders installs the etcdadm providers for the cluster using clusterctl.
func (c *Clusterctl) InstallEtcdadmProviders(ctx context.Context, managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec, cluster *types.Cluster, infraProvider providers.Provider, installProviders []string) error {
	_ = "STUB: not implemented"
	return nil
}
