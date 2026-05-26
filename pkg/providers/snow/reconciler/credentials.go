package reconciler

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const BoostrapSecretName = "capas-manager-bootstrap-credentials"

func getSnowCredentials(ctx context.Context, cli client.Client) (credentials, caBundle []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
