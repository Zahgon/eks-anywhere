package snow

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
)

type (
	// BaseWorkers represents the Snow specific CAPI spec for worker nodes.
	BaseWorkers     = clusterapi.Workers[*snowv1.AWSSnowMachineTemplate]
	baseWorkerGroup = clusterapi.WorkerGroup[*snowv1.AWSSnowMachineTemplate]
)

// Workers holds the Snow specific objects for CAPI snow worker groups.
type Workers struct {
	BaseWorkers
	CAPASIPPools CAPASIPPools
}

// Objects returns the worker nodes objects associated with the snow cluster.
func (w Workers) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// WorkersSpec generates a Snow specific CAPI spec for an eks-a cluster worker nodes.
// It talks to the cluster with a client to detect changes in immutable objects and generates new
// names for them.
func WorkersSpec(ctx context.Context, log logr.Logger, spec *cluster.Spec, client kubernetes.Client) (*Workers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
