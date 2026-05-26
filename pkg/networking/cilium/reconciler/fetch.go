package reconciler

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type preflightInstallation struct {
	daemonSet  *appsv1.DaemonSet
	deployment *appsv1.Deployment
}

func (p *preflightInstallation) installed() bool { _ = "STUB: not implemented"; return false }

func getPreflightInstallation(ctx context.Context, client client.Client) (*preflightInstallation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDeployment(ctx context.Context, client client.Client, name string) (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDaemonSet(ctx context.Context, client client.Client, name string) (*appsv1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
