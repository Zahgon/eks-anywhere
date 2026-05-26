package auth

import (
	"github.com/aws/smithy-go/auth"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// HTTPAuthScheme is the SDK's internal implementation of smithyhttp.AuthScheme
// for pre-existing implementations where the signer was added to client
// config. SDK clients will key off of this type and ensure per-operation
// updates to those signers persist on the scheme itself.
type HTTPAuthScheme struct {
	schemeID string
	signer   smithyhttp.Signer
}

var _ smithyhttp.AuthScheme = (*HTTPAuthScheme)(nil)

// NewHTTPAuthScheme returns an auth scheme instance with the given config.
func NewHTTPAuthScheme(schemeID string, signer smithyhttp.Signer) *HTTPAuthScheme {
	_ = "STUB: not implemented"
	return nil
}

// SchemeID identifies the auth scheme.
func (s *HTTPAuthScheme) SchemeID() string {
	_ = "STUB: not implemented"

	// IdentityResolver gets the identity resolver for the auth scheme.
	return ""
}

func (s *HTTPAuthScheme) IdentityResolver(o auth.IdentityResolverOptions) auth.IdentityResolver {
	_ = "STUB: not implemented"
	return *new(auth.IdentityResolver)
}

// Signer gets the signer for the auth scheme.
func (s *HTTPAuthScheme) Signer() smithyhttp.Signer {
	_ = "STUB: not implemented"

	// WithSigner returns a new instance of the auth scheme with the updated signer.
	return *new(smithyhttp.Signer)
}

func (s *HTTPAuthScheme) WithSigner(signer smithyhttp.Signer) *HTTPAuthScheme {
	_ = "STUB: not implemented"
	return nil
}
