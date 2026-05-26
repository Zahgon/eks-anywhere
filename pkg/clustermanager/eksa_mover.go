package clustermanager

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// MoverOpt allows to customize a Mover on construction.
type MoverOpt func(*Mover)

// Mover applies the cluster spec to the management cluster and waits
// until the changes are fully reconciled.
type Mover struct {
	log                logr.Logger
	clientFactory      ClientFactory
	moveClusterTimeout time.Duration
	retryBackOff       time.Duration
}

// NewMover builds an Mover.
func NewMover(log logr.Logger, clientFactory ClientFactory, opts ...MoverOpt) *Mover {
	_ = "STUB: not implemented"
	return nil
}

// WithMoverNoTimeouts disables the timeout for all the waits and retries in management upgrader.
func WithMoverNoTimeouts() MoverOpt { _ = "STUB: not implemented"; return *new(MoverOpt) }

// WithMoverApplyClusterTimeout allows to configure how long the mover retries
// to apply the objects in case of failure.
// Generally only used in tests.
func WithMoverApplyClusterTimeout(timeout time.Duration) MoverOpt {
	_ = "STUB: not implemented"
	return *new(MoverOpt)
}

// WithMoverRetryBackOff allows to configure how long the mover waits between requests
// to update the cluster spec objects and check the status of the Cluster.
// Generally only used in tests.
func WithMoverRetryBackOff(backOff time.Duration) MoverOpt {
	_ = "STUB: not implemented"
	return *new(MoverOpt)
}

// Move applies the cluster's namespace and spec without checking for reconcile conditions.
func (m *Mover) Move(ctx context.Context, spec *cluster.Spec, fromClient, toClient kubernetes.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// read the cluster from bootstrap

// pause cluster on bootstrap

// For baremetal provider we need to clear this annotation once we move to management cluster
// at this point all the hardware is provisioned and tink stack is up on the management cluster.
// We don't need the Bootstrap IP anymore.
// ideally this logic should be outside of move but the current code structure needs refactor
// to be able to do that.

func moveClusterResource(ctx context.Context, cluster *v1alpha1.Cluster, client kubernetes.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// move eksa cluster

func moveChildObjects(ctx context.Context, spec *cluster.Spec, fromClient, toClient kubernetes.Client) error {
	_ = "STUB: not implemented"
	// read and move child objects
	return nil
}
