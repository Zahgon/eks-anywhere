package e2e

const (
	packagesRegex            = `^.*CuratedPackages.*$`
	nonRegionalPackagesRegex = `^.*NonRegionalCuratedPackages.*$`
	certManagerRegex         = "^.*CuratedPackagesCertManager.*$"
)

// assumeRoleAndGetCredentials assumes an IAM role using the role ARN from env and
// returns the temporary credentials (access key, secret key, session token) or an error.
func assumeRoleAndGetCredentials(roleArnEnvVar, sessionName string) (accessKey, secretKey, sessionToken string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// (max for role chaining)

func (e *E2ESession) setupPackagesEnv(testRegex string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) setupCertManagerEnv(testRegex string) error {
	_ = "STUB: not implemented"
	return nil
}
