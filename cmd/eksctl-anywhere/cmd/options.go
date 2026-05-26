package cmd

import (
	"time"

	"github.com/spf13/pflag"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/version"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const defaultTinkerbellNodeStartupTimeout = 20 * time.Minute

const timeoutErrorTemplate = "failed to parse timeout %s: %v"

type timeoutOptions struct {
	cpWaitTimeout           string
	externalEtcdWaitTimeout string
	perMachineWaitTimeout   string
	noTimeouts              bool
}

func applyTimeoutFlags(flagSet *pflag.FlagSet, t *timeoutOptions) {
	_ = "STUB: not implemented"
	return
}

// buildClusterManagerOpts builds options for constructing a ClusterManager from CLI flags.
// datacenterKind is an API kind such as v1alpha1.TinkerbellDatacenterKind.
func buildClusterManagerOpts(t timeoutOptions, datacenterKind string) (*dependencies.ClusterManagerTimeoutOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type clusterOptions struct {
	fileName             string
	bundlesOverride      string
	managementKubeconfig string
}

func (c clusterOptions) mountDirs() []string { _ = "STUB: not implemented"; return nil }

func readClusterSpec(clusterConfigPath string, cliVersion version.Info, opts ...cluster.FileSpecBuilderOpt) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readAndValidateClusterSpec(clusterConfigPath string, cliVersion version.Info, opts ...cluster.FileSpecBuilderOpt) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newClusterSpec(options clusterOptions) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBundles(cliVersion version.Info, bundlesManifestURL string) (*releasev1.Bundles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEksaRelease(cliVersion version.Info) (*releasev1.EksARelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getConfig(clusterConfigPath string, cliVersion version.Info) (*cluster.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newBasicSpec creates a new cluster.Spec with the given Config, Bundles, and EKSARelease.
// This was created as a short term fix to management upgrades when the Cluster object is using an
// unsupported version of Kubernetes.
//
// When building the full cluster.Spec definition, we fetch the eksdReleases for
// the KubernetesVersions for all unique k8s versions specified in the Cluster for both CP and workers.
// If the Cluster object is using an unsupported version of Kubernetes, an error thrown
// because it does not exist in the Bundles file. This method allows to build a cluster.Spec without
// encountering this problem when performing only a management component upgrade.
func newBasicSpec(config *cluster.Config, bundles *releasev1.Bundles, eksaRelease *releasev1.EKSARelease) *cluster.Spec {
	_ = "STUB: not implemented"
	return nil
}

func readBasicClusterSpec(clusterConfigPath string, cliVersion version.Info, options clusterOptions) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newBasicClusterSpec(options clusterOptions) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildCliConfig(clusterSpec *cluster.Spec) *config.CliConfig {
	_ = "STUB: not implemented"
	return nil
}

func buildCreateCliConfig(clusterOptions *createClusterOptions) (*config.CreateClusterCLIConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildUpgradeCliConfig(clusterOptions *upgradeClusterOptions) (*config.UpgradeClusterCLIConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildDeleteCliConfig() *config.DeleteClusterCLIConfig { _ = "STUB: not implemented"; return nil }

func getManagementClusterKubeconfig(clusterName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// check if kubeconfig for management cluster exists locally

func getManagementCluster(clusterSpec *cluster.Spec) *types.Cluster {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterOptions) directoriesToMount(clusterSpec *cluster.Spec, cliConfig *config.CliConfig, addDirs ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterOptions) cloudStackDirectoriesToMount() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
