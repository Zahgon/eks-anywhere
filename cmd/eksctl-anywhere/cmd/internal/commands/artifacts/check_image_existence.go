package artifacts

import (
	"context"
)

const (
	realmKey   = "realm="
	serviceKey = "service="
	scopeKey   = "scope="
)

type CheckImageExistence struct {
	ImageUri   string
	AuthHeader string
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (d CheckImageExistence) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func getRegistryToken(realm, service, scope string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func splitImageUri(imageUri string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}
