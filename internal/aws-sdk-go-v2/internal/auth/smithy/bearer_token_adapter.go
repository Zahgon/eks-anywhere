package smithy

import (
	"context"
	"time"

	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/auth/bearer"
)

// BearerTokenAdapter adapts smithy bearer.Token to smithy auth.Identity.
type BearerTokenAdapter struct {
	Token bearer.Token
}

var _ auth.Identity = (*BearerTokenAdapter)(nil)

// Expiration returns the time of expiration for the token.
func (v *BearerTokenAdapter) Expiration() time.Time {
	_ = "STUB: not implemented"
	return *

	// BearerTokenProviderAdapter adapts smithy bearer.TokenProvider to smithy
	// auth.IdentityResolver.
	new(time.Time)
}

type BearerTokenProviderAdapter struct {
	Provider bearer.TokenProvider
}

var _ (auth.IdentityResolver) = (*BearerTokenProviderAdapter)(nil)

// GetIdentity retrieves a bearer token using the underlying provider.
func (v *BearerTokenProviderAdapter) GetIdentity(ctx context.Context, _ smithy.Properties) (
	auth.Identity, error,
) {
	_ = "STUB: not implemented"
	return *new(auth.Identity), nil
}
