package v1alpha1

const (
	GitOpsConfigKind     = "GitOpsConfig"
	FluxDefaultNamespace = "flux-system"
	FluxDefaultBranch    = "main"
)

func validateGitOpsConfig(config *GitOpsConfig) error { _ = "STUB: not implemented"; return nil }

func validateGitBranchName(branchName string) error { _ = "STUB: not implemented"; return nil }

func validateGitRepoName(repoName string) error { _ = "STUB: not implemented"; return nil }

func setGitOpsConfigDefaults(gitops *GitOpsConfig) { _ = "STUB: not implemented"; return }
