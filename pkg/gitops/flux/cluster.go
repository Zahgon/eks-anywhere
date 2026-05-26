package flux

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/git"
	"github.com/aws/eks-anywhere/pkg/providers"
)

// fluxForCluster bundles the Flux struct with a specific clusterSpec, so that all the git and file write
// operations for the clusterSpec can be done in each structure method.
type fluxForCluster struct {
	*Flux
	clusterSpec      *cluster.Spec
	datacenterConfig providers.DatacenterConfig
	machineConfigs   []providers.MachineConfig
}

func newFluxForCluster(flux *Flux, clusterSpec *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig) *fluxForCluster {
	_ = "STUB: not implemented"
	return nil
}

// commitFluxAndClusterConfigToGit commits the cluster configuration file to the flux-managed git repository.
// If the remote repository does not exist it will initialize a local repository and push it to the configured remote.
// It will generate the kustomization file and marshal the cluster configuration file to the required locations in the repo.
// These will later be used by Flux and our controllers to reconcile the repository contents and the cluster configuration.
func (fc *fluxForCluster) commitFluxAndClusterConfigToGit(ctx context.Context, managementComponents *cluster.ManagementComponents) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *fluxForCluster) syncGitRepo(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure the local git repo is on the branch specified in config and up-to-date with the remote

func (fc *fluxForCluster) initializeProviderRepositoryIfNotExists(ctx context.Context) (*git.Repository, error) {
	_ = "STUB: not implemented"
	// If git provider, the repository should be pre-initialized by the user.
	return nil, nil
}

// setupRepository will set up the repository which will house the GitOps configuration for the cluster.
// if the repository exists and is not empty, it will be cloned.
// if the repository exists but is empty, it will be initialized locally, as a bare repository cannot be cloned.
// if the repository does not exist, it will be created and then initialized locally.
func (fc *fluxForCluster) setupRepository(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (fc *fluxForCluster) clone(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// createRemoteRepository will create a repository in the remote git provider with the user-provided configuration.
func (fc *fluxForCluster) createRemoteRepository(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// initializeLocalRepository will git init the local repository directory, initialize a git repository.
// it will then change branches to the branch specified in the GitOps configuration.
func (fc *fluxForCluster) initializeLocalRepository() error { _ = "STUB: not implemented"; return nil }

// git requires at least one commit in the repo to branch from

// validateLocalConfigPathDoesNotExist returns an exception if the cluster configuration file exists.
// This is done so that we avoid clobbering existing cluster configurations in the user-provided git repository.
func (fc *fluxForCluster) validateLocalConfigPathDoesNotExist() error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *fluxForCluster) validateRemoteConfigPathDoesNotExist(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *fluxForCluster) namespace() string { _ = "STUB: not implemented"; return "" }

func (fc *fluxForCluster) repository() string { _ = "STUB: not implemented"; return "" }

func (fc *fluxForCluster) owner() string { _ = "STUB: not implemented"; return "" }

func (fc *fluxForCluster) branch() string { _ = "STUB: not implemented"; return "" }

func (fc *fluxForCluster) personal() bool { _ = "STUB: not implemented"; return false }

func (fc *fluxForCluster) path() string { _ = "STUB: not implemented"; return "" }

func (fc *fluxForCluster) eksaSystemDir() string { _ = "STUB: not implemented"; return "" }

func (fc *fluxForCluster) fluxSystemDir() string { _ = "STUB: not implemented"; return "" }
