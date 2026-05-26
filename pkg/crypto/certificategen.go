package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"math/big"
	"time"
)

type certificategenerator struct{}

type CertificateGenerator interface {
	GenerateIamAuthSelfSignCertKeyPair() ([]byte, []byte, error)
}

func NewCertificateGenerator() CertificateGenerator {
	_ = "STUB: not implemented"
	return *new(CertificateGenerator)
}

func (cg *certificategenerator) GenerateIamAuthSelfSignCertKeyPair() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (cg *certificategenerator) generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	// Private Key generation
	return nil, nil
}

// Validate Private Key

func (cg *certificategenerator) getCertLifeTime() (time.Time, time.Time) {
	_ = "STUB: not implemented"
	// lifetime of the CA certificate
	return *new(time.Time), *new(time.Time)
}

func (cg *certificategenerator) generateCertSerialNumber() (*big.Int, error) {
	_ = "STUB: not implemented"
	// choose a random 128 bit serial number
	return nil, nil
}

func (cg *certificategenerator) generateAwsIamAuthCertTemplate(serialNumber *big.Int, notBefore, notAfter time.Time) x509.Certificate {
	_ = "STUB: not implemented"
	return *new(x509.Certificate)
}

func (cg *certificategenerator) generateSelfSignCertificate(template x509.Certificate, privateKey *rsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cg *certificategenerator) encodePrivateKey(privateKey *rsa.PrivateKey) []byte {
	_ = "STUB: not implemented"
	// ASN.1 DER format
	return nil
}

func (cg *certificategenerator) encodeToPEM(bytes []byte, blockType string) []byte {
	_ = "STUB: not implemented"
	return nil
}
