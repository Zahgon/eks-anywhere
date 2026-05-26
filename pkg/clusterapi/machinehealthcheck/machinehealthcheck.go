package machinehealthcheck

import (
	"context"

	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// GetControlPlaneMachineHealthCheck checks if machine health checks already exist on the cluster and returns it.
func GetControlPlaneMachineHealthCheck(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) (*clusterv1beta2.MachineHealthCheck, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we get only kcp machine health check as it had the same timeouts as worker machine health checks
