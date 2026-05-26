package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// OIDCConfigOpt updates an OIDC config.
type OIDCConfigOpt func(o *anywherev1.OIDCConfig)

// WithOIDCConfig builds a ClusterConfigFiller that adds a OIDCConfig with the
// given name and spec to the cluster config.
func WithOIDCConfig(name string, opts ...OIDCConfigOpt) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

func WithOIDCClientId(id string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithOIDCIssuerUrl(url string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithOIDCUsernameClaim(claim string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithOIDCUsernamePrefix(prefix string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithOIDCGroupsClaim(claim string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithOIDCGroupsPrefix(prefix string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithOIDCRequiredClaims(claim, value string) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}

func WithStringFromEnvVarOIDCConfig(envVar string, opt func(string) OIDCConfigOpt) OIDCConfigOpt {
	_ = "STUB: not implemented"
	return *new(OIDCConfigOpt)
}
