package oidc

import (
	jose "github.com/go-jose/go-jose/v3"
)

type keyResponse struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

func GetKeyID(keysBytes []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseKeys(bytes []byte) (*keyResponse, error) { _ = "STUB: not implemented"; return nil, nil }
