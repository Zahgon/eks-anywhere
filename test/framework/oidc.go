package framework

import (
	"github.com/aws/eks-anywhere/internal/pkg/api"
)

const (
	OIDCIssuerUrlVar = "T_OIDC_ISSUER_URL"
	OIDCClientIdVar  = "T_OIDC_CLIENT_ID"
	OIDCKidVar       = "T_OIDC_KID"
	OIDCKeyFileVar   = "T_OIDC_KEY_FILE"
)

var oidcRequiredEnvVars = []string{
	OIDCIssuerUrlVar,
	OIDCClientIdVar,
	OIDCKidVar,
	OIDCKeyFileVar,
}

func WithOIDC() ClusterE2ETestOpt { _ = "STUB: not implemented"; return *new(ClusterE2ETestOpt) }

// WithOIDCClusterConfig returns a ClusterConfigFiller that adds the default
// OIDCConfig for E2E tests to the cluster Config and links it by name in the
// Cluster resource.
func WithOIDCClusterConfig(t T) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func (e *ClusterE2ETest) ValidateOIDC() { _ = "STUB: not implemented"; return }

// WithOIDCEnvVarCheck returns a ClusterE2ETestOpt that checks for the required env vars.
func WithOIDCEnvVarCheck() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}
