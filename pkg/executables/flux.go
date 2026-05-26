package executables

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	fluxPath                   = "flux"
	eksaGithubTokenEnv         = "EKSA_GITHUB_TOKEN"
	githubTokenEnv             = "GITHUB_TOKEN"
	githubProvider             = "github"
	gitProvider                = "git"
	defaultPrivateKeyAlgorithm = "ecdsa"
)

type Flux struct {
	Executable
}

func NewFlux(executable Executable) *Flux { _ = "STUB: not implemented"; return nil }

// BootstrapGithub creates the GitHub repository if it doesn’t exist, and commits the toolkit
// components manifests to the main branch. Then it configures the target cluster to synchronize with the repository.
// If the toolkit components are present on the cluster, the bootstrap command will perform an upgrade if needed.
func (f *Flux) BootstrapGithub(ctx context.Context, cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// BootstrapGit commits the toolkit components manifests to the branch of a Git repository.
// It then configures the target cluster to synchronize with the repository. If the toolkit components are present on the cluster, the
// bootstrap command will perform an upgrade if needed.
func (f *Flux) BootstrapGit(ctx context.Context, cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig, cliConfig *config.CliConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func setUpCommonParamsBootstrap(cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig, params []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flux) Uninstall(ctx context.Context, cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flux) SuspendKustomization(ctx context.Context, cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flux) ResumeKustomization(ctx context.Context, cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flux) Reconcile(ctx context.Context, cluster *types.Cluster, fluxConfig *v1alpha1.FluxConfig) error {
	_ = "STUB: not implemented"
	return nil
}
