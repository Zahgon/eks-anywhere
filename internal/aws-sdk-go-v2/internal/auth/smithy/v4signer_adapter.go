package smithy

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/logging"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// V4SignerAdapter adapts v4.HTTPSigner to smithy http.Signer.
type V4SignerAdapter struct {
	Signer     v4.HTTPSigner
	Logger     logging.Logger
	LogSigning bool
}

var _ (smithyhttp.Signer) = (*V4SignerAdapter)(nil)

// SignRequest signs the request with the provided identity.
func (v *V4SignerAdapter) SignRequest(ctx context.Context, r *smithyhttp.Request, identity auth.Identity, props smithy.Properties) error {
	_ = "STUB: not implemented"
	return nil
}
