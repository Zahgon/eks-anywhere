package v1alpha1

const (
	FluxConfigKind   = "FluxConfig"
	RsaAlgorithm     = "rsa"
	EcdsaAlgorithm   = "ecdsa"
	Ed25519Algorithm = "ed25519"
)

func validateFluxConfig(config *FluxConfig) error { _ = "STUB: not implemented"; return nil }

func validateGitProviderConfig(gitProviderConfig GitProviderConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGithubProviderConfig(config GithubProviderConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRepositoryUrl(repositoryUrl string) error { _ = "STUB: not implemented"; return nil }

func validateSshKeyAlgorithm(sshKeyAlgorithm string) error { _ = "STUB: not implemented"; return nil }

func setFluxConfigDefaults(flux *FluxConfig) { _ = "STUB: not implemented"; return }
