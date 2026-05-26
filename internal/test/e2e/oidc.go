package e2e

const (
	openIDConfPath = "oidc/.well-known/openid-configuration"
	keysPath       = "oidc/keys.json"
	saSignerPath   = "oidc/sa-signer.key"
)

func (e *E2ESession) setupOIDC(testRegex string) error { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) createOIDCFiles(issuerURL, folder string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) getKeyID(folder string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *E2ESession) downloadSignerKeyInInstance(folder string) (pathInInstance string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
