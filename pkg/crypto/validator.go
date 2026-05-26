package crypto

type DefaultTlsValidator struct{}

type TlsValidator interface {
	ValidateCert(host, port, caCertContent string) error
	IsSignedByUnknownAuthority(host, port string) (bool, error)
}

func NewTlsValidator() TlsValidator { _ = "STUB: not implemented"; return *new(TlsValidator) }

// IsSignedByUnknownAuthority determines if the url is signed by an unknown authority.
func (tv *DefaultTlsValidator) IsSignedByUnknownAuthority(host, port string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ValidateCert parses the cert, ensures that the cert format is valid and verifies that the cert is valid for the url.
func (tv *DefaultTlsValidator) ValidateCert(host, port, caCertContent string) error {
	_ = "STUB: not implemented"
	// Validates that the cert format is valid
	return nil
}

// Verifies that the cert is valid by making a connection to the endpoint
