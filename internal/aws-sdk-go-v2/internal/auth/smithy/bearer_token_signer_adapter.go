package smithy

import (
	"context"

	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/auth/bearer"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// BearerTokenSignerAdapter adapts smithy bearer.Signer to smithy http
// auth.Signer.
type BearerTokenSignerAdapter struct {
	Signer bearer.Signer
}

var _ (smithyhttp.Signer) = (*BearerTokenSignerAdapter)(nil)

// SignRequest signs the request with the provided bearer token.
func (v *BearerTokenSignerAdapter) SignRequest(ctx context.Context, r *smithyhttp.Request, identity auth.Identity, _ smithy.Properties) error {
	_ = "STUB: not implemented"
	return nil
}
