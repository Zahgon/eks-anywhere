package auth

import (
	"context"

	smithy "github.com/aws/smithy-go"
)

// SigV4 is a constant representing
// Authentication Scheme Signature Version 4
const SigV4 = "sigv4"

// SigV4A is a constant representing
// Authentication Scheme Signature Version 4A
const SigV4A = "sigv4a"

// SigV4S3Express identifies the S3 S3Express auth scheme.
const SigV4S3Express = "sigv4-s3express"

// None is a constant representing the
// None Authentication Scheme
const None = "none"

// SupportedSchemes is a data structure
// that indicates the list of supported AWS
// authentication schemes
var SupportedSchemes = map[string]bool{
	SigV4:          true,
	SigV4A:         true,
	SigV4S3Express: true,
	None:           true,
}

// AuthenticationScheme is a representation of
// AWS authentication schemes
type AuthenticationScheme interface {
	isAuthenticationScheme()
}

// AuthenticationSchemeV4 is a AWS SigV4 representation
type AuthenticationSchemeV4 struct {
	Name                  string
	SigningName           *string
	SigningRegion         *string
	DisableDoubleEncoding *bool
}

func (a *AuthenticationSchemeV4) isAuthenticationScheme() {
	_ = "STUB: not implemented"

	// AuthenticationSchemeV4A is a AWS SigV4A representation
	return
}

type AuthenticationSchemeV4A struct {
	Name                  string
	SigningName           *string
	SigningRegionSet      []string
	DisableDoubleEncoding *bool
}

func (a *AuthenticationSchemeV4A) isAuthenticationScheme() {
	_ = "STUB: not implemented"

	// AuthenticationSchemeNone is a representation for the none auth scheme
	return
}

type AuthenticationSchemeNone struct{}

func (a *AuthenticationSchemeNone) isAuthenticationScheme() {
	_ = "STUB: not implemented"

	// NoAuthenticationSchemesFoundError is used in signaling
	// that no authentication schemes have been specified.
	return
}

type NoAuthenticationSchemesFoundError struct{}

func (e *NoAuthenticationSchemesFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// UnSupportedAuthenticationSchemeSpecifiedError is used in
// signaling that only unsupported authentication schemes
// were specified.
type UnSupportedAuthenticationSchemeSpecifiedError struct {
	UnsupportedSchemes []string
}

func (e *UnSupportedAuthenticationSchemeSpecifiedError) Error() string {
	_ = "STUB: not implemented"
	return ""
}

// GetAuthenticationSchemes extracts the relevant authentication scheme data
// into a custom strongly typed Go data structure.
func GetAuthenticationSchemes(p *smithy.Properties) ([]AuthenticationScheme, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type disableDoubleEncoding struct{}

// SetDisableDoubleEncoding sets or modifies the disable double encoding option
// on the context.
//
// Scoped to stack values. Use github.com/aws/smithy-go/middleware#ClearStackValues
// to clear all stack values.
func SetDisableDoubleEncoding(ctx context.Context, value bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetDisableDoubleEncoding retrieves the disable double encoding option
// from the context.
//
// Scoped to stack values. Use github.com/aws/smithy-go/middleware#ClearStackValues
// to clear all stack values.
func GetDisableDoubleEncoding(ctx context.Context) (value bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func getSigningName(authScheme map[string]interface{}) *string {
	_ = "STUB: not implemented"
	return nil
}

func getSigningRegionSet(authScheme map[string]interface{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func getSigningRegion(authScheme map[string]interface{}) *string {
	_ = "STUB: not implemented"
	return nil
}

func getDisableDoubleEncoding(authScheme map[string]interface{}) *bool {
	_ = "STUB: not implemented"
	return nil
}
