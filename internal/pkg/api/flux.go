package api

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type FluxConfigOpt func(o *v1alpha1.FluxConfig)

func NewFluxConfig(name string, opts ...FluxConfigOpt) *v1alpha1.FluxConfig {
	_ = "STUB: not implemented"
	return nil
}

func WithFluxConfigName(n string) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

func WithFluxConfigNamespace(ns string) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

func WithBranch(branch string) FluxConfigOpt { _ = "STUB: not implemented"; return *new(FluxConfigOpt) }

func WithClusterConfigPath(configPath string) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

func WithSystemNamespace(namespace string) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

func WithStringFromEnvVarFluxConfig(envVar string, opt func(string) FluxConfigOpt) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

type GitProviderOpt func(o *v1alpha1.GitProviderConfig)

func WithGenericGitProvider(opts ...GitProviderOpt) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

func WithGitRepositoryUrl(url string) GitProviderOpt {
	_ = "STUB: not implemented"
	return *new(GitProviderOpt)
}

func WithStringFromEnvVarGenericGitProviderConfig(envVar string, opt func(string) GitProviderOpt) GitProviderOpt {
	_ = "STUB: not implemented"
	return *new(GitProviderOpt)
}

type GithubProviderOpt func(o *v1alpha1.GithubProviderConfig)

func WithGithubProvider(opts ...GithubProviderOpt) FluxConfigOpt {
	_ = "STUB: not implemented"
	return *new(FluxConfigOpt)
}

func WithGithubOwner(owner string) GithubProviderOpt {
	_ = "STUB: not implemented"
	return *new(GithubProviderOpt)
}

func WithGithubRepository(repository string) GithubProviderOpt {
	_ = "STUB: not implemented"
	return *new(GithubProviderOpt)
}

func WithPersonalGithubRepository(personal bool) GithubProviderOpt {
	_ = "STUB: not implemented"
	return *new(GithubProviderOpt)
}

func WithStringFromEnvVarGithubProviderConfig(envVar string, opt func(string) GithubProviderOpt) GithubProviderOpt {
	_ = "STUB: not implemented"
	return *new(GithubProviderOpt)
}
