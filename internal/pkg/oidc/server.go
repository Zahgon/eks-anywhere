package oidc

import (
	"crypto/rsa"
)

type MinimalProvider struct {
	Discovery, Keys, PrivateKey []byte
	KeyID                       string
}

type discoveryResponse struct {
	Issuer                           string   `json:"issuer"`
	JwksUri                          string   `json:"jwks_uri"`
	AuthorizationEndpoint            string   `json:"authorization_endpoint"`
	ResponseTypesSupported           []string `json:"response_types_supported"`
	SubjectTypesSupported            []string `json:"subject_types_supported"`
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported                  []string `json:"claims_supported"`
}

func GenerateMinimalProvider(issuerURL string) (*MinimalProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set the key id to the sha1 of the pubkey

func marshalPubKey(pubKey *rsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalPrivateKey(k *rsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
