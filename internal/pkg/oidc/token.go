package oidc

type oidcTokenClaim struct {
	Issuer           string `json:"iss,omitempty"`
	Subject          string `json:"sub,omitempty"`
	kid              string
	keyFile          string
	Role             string   `json:"role,omitempty"`
	Email            string   `json:"email,omitempty"`
	Audience         []string `json:"aud,omitempty"`
	Groups           []string `json:"groups,omitempty"`
	ExpiresAt        int64    `json:"exp,omitempty"`
	IssuedAt         int64    `json:"iat,omitempty"`
	NotBefore        int64    `json:"nbf,omitempty"`
	KubernetesAccess string   `json:"kubernetesAccess,omitempty"`
}

type JWTOpt func(*oidcTokenClaim)

func NewJWT(issuerName, kid, keyFile string, opts ...JWTOpt) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func WithEmail(email string) JWTOpt { _ = "STUB: not implemented"; return *new(JWTOpt) }

func WithGroup(group string) JWTOpt { _ = "STUB: not implemented"; return *new(JWTOpt) }

func WithRole(role string) JWTOpt { _ = "STUB: not implemented"; return *new(JWTOpt) }

func WithKubernetesAccess(access bool) JWTOpt { _ = "STUB: not implemented"; return *new(JWTOpt) }

func WithAudience(audience string) JWTOpt { _ = "STUB: not implemented"; return *new(JWTOpt) }

func WithSubject(subject string) JWTOpt { _ = "STUB: not implemented"; return *new(JWTOpt) }

func (o *oidcTokenClaim) generateToken() (string, error) { _ = "STUB: not implemented"; return "", nil }
