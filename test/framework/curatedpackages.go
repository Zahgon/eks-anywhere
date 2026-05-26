package framework

import (
	"testing"
)

const (
	eksaPackagesSecretKey       = "EKSA_AWS_SECRET_ACCESS_KEY"
	eksaPackagesAccessKey       = "EKSA_AWS_ACCESS_KEY_ID"
	eksaPackagesSessionTokenKey = "EKSA_AWS_SESSION_TOKEN"
	eksaPackagesRegion          = "EKSA_AWS_REGION"
	route53AccessKey            = "ROUTE53_ACCESS_KEY_ID"
	route53SecretKey            = "ROUTE53_SECRET_ACCESS_KEY"
	route53SessionToken         = "ROUTE53_SESSION_TOKEN"
	route53Region               = "ROUTE53_REGION"
	route53ZoneID               = "ROUTE53_ZONEID"
)

var requiredPackagesEnvVars = []string{
	eksaPackagesRegion,
	eksaPackagesAccessKey,
	eksaPackagesSecretKey,
	eksaPackagesSessionTokenKey,
}

var requiredCertManagerEnvVars = []string{
	route53Region,
	route53AccessKey,
	route53SecretKey,
	route53SessionToken,
	route53ZoneID,
}

// RequiredPackagesEnvVars returns the list of packages env vars.
func RequiredPackagesEnvVars() []string { _ = "STUB: not implemented"; return nil }

// RequiredCertManagerEnvVars returns the list of cert manager env vars.
func RequiredCertManagerEnvVars() []string { _ = "STUB: not implemented"; return nil }

// CheckCuratedPackagesCredentials will exit out if the Curated Packages environment variables are not set.
func CheckCuratedPackagesCredentials(t *testing.T) { _ = "STUB: not implemented"; return }

// CheckCertManagerCredentials will exit if route53 credentials are not set.
func CheckCertManagerCredentials(t *testing.T) { _ = "STUB: not implemented"; return }

// GetRoute53Configs returns route53 configurations for cert-manager.
func GetRoute53Configs() (string, string) { _ = "STUB: not implemented"; return "", "" }
