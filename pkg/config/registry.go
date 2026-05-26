package config

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const registryAuthSecretName = "registry-credentials"

func ReadCredentials() (username, password string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ReadCredentialsFromSecret reads from Kubernetes secret registry-credentials.
// Returns the username and password, or error.
func ReadCredentialsFromSecret(ctx context.Context, client client.Client) (username, password string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// SetCredentialsEnv sets the registry username and password env variables.
func SetCredentialsEnv(username, password string) error { _ = "STUB: not implemented"; return nil }
