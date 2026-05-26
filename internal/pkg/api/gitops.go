package api

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type GitOpsConfigOpt func(o *v1alpha1.GitOpsConfig)

func NewGitOpsConfig(name string, opts ...GitOpsConfigOpt) *v1alpha1.GitOpsConfig {
	_ = "STUB: not implemented"
	return nil
}

func WithGitOpsNamespace(ns string) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithFluxOwner(username string) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithFluxRepository(repository string) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithFluxConfigurationPath(configPath string) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithFluxNamespace(namespace string) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithFluxBranch(branch string) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithPersonalFluxRepository(personal bool) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}

func WithStringFromEnvVarGitOpsConfig(envVar string, opt func(string) GitOpsConfigOpt) GitOpsConfigOpt {
	_ = "STUB: not implemented"
	return *new(GitOpsConfigOpt)
}
