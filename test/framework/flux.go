package framework

import (
	"context"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	gitfactory "github.com/aws/eks-anywhere/pkg/git/factory"
	"github.com/aws/eks-anywhere/pkg/providers"
)

const (
	eksaConfigFileName    = "eksa-cluster.yaml"
	fluxSystemNamespace   = "flux-system"
	GitRepositoryVar      = "T_GIT_REPOSITORY"
	GitRepoSshUrl         = "T_GIT_SSH_REPO_URL"
	GithubUserVar         = "T_GITHUB_USER"
	GithubTokenVar        = "EKSA_GITHUB_TOKEN"
	GitKnownHosts         = "EKSA_GIT_KNOWN_HOSTS"
	GitPrivateKeyFile     = "EKSA_GIT_PRIVATE_KEY"
	DefaultFluxConfigName = "eksa-test"
)

var fluxGithubRequiredEnvVars = []string{
	GitRepositoryVar,
	GithubUserVar,
	GithubTokenVar,
}

var fluxGitRequiredEnvVars = []string{
	GitKnownHosts,
	GitPrivateKeyFile,
	GitRepoSshUrl,
}

var fluxGitCreateGenerateRepoEnvVars = []string{
	GitKnownHosts,
	GitPrivateKeyFile,
	GithubUserVar,
	GithubTokenVar,
}

func getJobIDFromEnv() string { _ = "STUB: not implemented"; return "" }

func WithFluxGit(opts ...api.FluxConfigOpt) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// apply the rest of the opts passed into the function

func WithFluxGithub(opts ...api.FluxConfigOpt) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// apply the rest of the opts passed into the function

// Adding Job ID suffix to repo name
// e2e test jobs have Job Id with a ":", replacing with "-"

// Setting GitRepo cleanup since GitOps configured

// WithFluxGithubConfig returns ClusterConfigFiller that adds FluxConfig using the Github provider to the cluster config.
func WithFluxGithubConfig(opts ...api.FluxConfigOpt) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// apply the rest of the opts passed into the function

// Adding Job ID suffix to repo name
// e2e test jobs have Job Id with a ":", replacing with "-"

// WithFluxGithubEnvVarCheck returns a ClusterE2ETestOpt that checks for the required env vars.
func WithFluxGithubEnvVarCheck() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithFluxGithubCleanup returns a ClusterE2ETestOpt that registers the git repository cleanup operation.
func WithFluxGithubCleanup() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithClusterUpgradeGit(fillers ...api.ClusterFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// TODO: e.ClusterConfig.GitOpsConfig is defined from api.NewGitOpsConfig in WithFluxLegacy()
// instead of marshalling from the actual file in git repo.
// By default it does not include the namespace field. But Flux requires namespace always
// exist for all the objects managed by its kustomization controller.
// Need to refactor this to read gitopsconfig directly from file in git repo
// which always has the namespace field.

func withFluxRepositorySuffix(suffix string) api.FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(api.FluxConfigOpt)
}

func fluxConfigName() string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) UpgradeWithGitOps(clusterOpts ...ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) upgradeWithGitOps(clusterOpts []ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

// Read the cluster config we just pulled into e.ClusterConfig

// Apply the options, these are most of the times fillers, so they will update the
// cluster config we just read from the repo. This has to happen after we parse the cluster
// config from the repo or we might be updating a different version of the config.

// Marshall e.ClusterConfig and write it to the repo path

func (e *ClusterE2ETest) initGit(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) workloadClusterConfigPath(w *WorkloadCluster) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *ClusterE2ETest) workloadClusterConfigGitPath(w *WorkloadCluster) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *ClusterE2ETest) buildWorkloadClusterConfigFileForGit(w *WorkloadCluster) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) addWorkloadClusterConfigToGit(ctx context.Context, w *WorkloadCluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) deleteWorkloadClusterConfigFromGit(ctx context.Context, w *WorkloadCluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) pushStagedChanges(ctx context.Context, commitMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) pushWorkloadClusterToGit(w *WorkloadCluster, opts ...api.ClusterConfigFiller) error {
	_ = "STUB: not implemented"
	return nil
}

// Pull remote config using managment cluster

// Read the cluster config we just pulled into w.ClusterConfig

// Update the cluster config with the provided api.ClusterConfigFillers

// Marshall w.ClusterConfig and write it to the repo path

func (e *ClusterE2ETest) deleteWorkloadClusterFromGit(w *WorkloadCluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) parseClusterConfigFromLocalGitRepo() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) buildClusterConfigFileForGit() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) ValidateFlux() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) CleanUpGitRepo() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) CleanUpGithubRepo() { _ = "STUB: not implemented"; return }

type providerConfig struct {
	datacenterConfig providers.DatacenterConfig
	machineConfigs   []providers.MachineConfig
}

func (e *ClusterE2ETest) validateInitialFluxState(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) validateWorkerNodeMultiConfigUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// update workernode specs

// update replica

func (e *ClusterE2ETest) validateGitopsRepoContentPath(repoName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *ClusterE2ETest) validateGitopsRepoContent(gitTools *gitfactory.GitTools) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) convertVSphereMachineConfigs(cpName, workerName, etcdName string, vsphereMachineConfigs map[string]*v1alpha1.VSphereMachineConfig) []providers.MachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) convertCloudstackMachineConfigs(cpName, workerName, etcdName string, cloudstackMachineConfigs map[string]*v1alpha1.CloudStackMachineConfig) []providers.MachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) validateWorkerNodeReplicaUpdates(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) validateWorkerNodeUpdates(ctx context.Context, opts ...CommandOpt) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) machineTemplateName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *ClusterE2ETest) validateFluxDeployments(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) validateEksaSystemDeployments(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) validateDeploymentsInManagementCluster(ctx context.Context, namespace string, expectedeployments map[string]int) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) updateWorkerNodeCountValue(ctx context.Context, newValue int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *ClusterE2ETest) providerConfig(clusterConfGitPath string) (*providerConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ClusterE2ETest) waitForWorkerNodeValidation() error { _ = "STUB: not implemented"; return nil }

func (e *ClusterE2ETest) validateWorkerNodeMachineSpec(ctx context.Context, clusterConfGitPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) waitForWorkerScaling(name string, targetvalue int) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) updateEKSASpecInGit(ctx context.Context, s *cluster.Spec, providersConfig providerConfig) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *ClusterE2ETest) pushConfigChanges(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) pullRemoteConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// todo: reuse logic in clustermanager to template resources
func (e *ClusterE2ETest) writeEKSASpec(s *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig) (path string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *ClusterE2ETest) gitRepoName() string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) gitBranch() string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) clusterConfigPathFromName(clusterName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *ClusterE2ETest) clusterConfGitPath() string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) clusterConfigGitPath() string { _ = "STUB: not implemented"; return "" }

func (e *ClusterE2ETest) clusterSpecFromGit() (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This makes sure that the cluster.Spec uses the same Bundles we pass to the CLI
// It avoids the budlesRef getting overwritten with whatever default Bundles the
// e2e test build is configured to use

func RequiredFluxGithubEnvVars() []string { _ = "STUB: not implemented"; return nil }

func RequiredFluxGitCreateRepoEnvVars() []string { _ = "STUB: not implemented"; return nil }
