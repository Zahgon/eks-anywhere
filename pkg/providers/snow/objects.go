package snow

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
)

// ControlPlaneObjects generates the control plane objects for snow provider from clusterSpec.
func ControlPlaneObjects(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec, kubeClient kubernetes.Client) ([]kubernetes.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WorkersObjects generates all the objects that compose a Snow specific CAPI spec for the worker nodes of an eks-a cluster.
func WorkersObjects(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec, kubeClient kubernetes.Client) ([]kubernetes.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MachineTemplateDeepDerivative compares two awssnowmachinetemplates to determine if their spec fields are equal.
// DeepDerivative is used so that unset fields in new object are not compared. Although DeepDerivative treats
// new subset slice equal to the original slice. i.e. DeepDerivative([]int{1}, []int{1, 2}) returns true.
// Custom logic is added to justify this usecase since removing a device from the devices list shall trigger machine
// rollout and recreate or the snow cluster goes into a state where the machines on the removed device can’t be deleted.
func MachineTemplateDeepDerivative(new, old *snowv1.AWSSnowMachineTemplate) bool {
	_ = "STUB: not implemented"
	return false
}
