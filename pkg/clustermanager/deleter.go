package clustermanager

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/types"
)

// DeleterOpt allows to customize a Deleter on construction.
type DeleterOpt func(*Deleter)

// Deleter deletes the cluster from the management cluster and waits
// until the deletions are fully reconciled.
type Deleter struct {
	log                         logr.Logger
	clientFactory               ClientFactory
	deleteClusterTimeout        time.Duration
	retryBackOff                time.Duration
	conditionCheckoutTotalCount int
}

const deleteClusterSpecTimeout = 5 * time.Minute

// NewDeleter builds an Deleter.
func NewDeleter(log logr.Logger, clientFactory ClientFactory, opts ...DeleterOpt) Deleter {
	_ = "STUB: not implemented"
	return *new(Deleter)
}

// WithDeleterNoTimeouts disables the timeout for all the waits and retries in management upgrader.
func WithDeleterNoTimeouts() DeleterOpt { _ = "STUB: not implemented"; return *new(DeleterOpt) }

// WithDeleterApplyClusterTimeout allows to configure how long the deleter retries
// to delete the objects in case of failure.
// Generally only used in tests.
func WithDeleterApplyClusterTimeout(timeout time.Duration) DeleterOpt {
	_ = "STUB: not implemented"
	return *new(DeleterOpt)
}

// WithDeleterRetryBackOff allows to configure how long the deleter waits between requests
// to update the cluster spec objects and check the status of the Cluster.
// Generally only used in tests.
func WithDeleterRetryBackOff(backOff time.Duration) DeleterOpt {
	_ = "STUB: not implemented"
	return *new(DeleterOpt)
}

// Run deletes the cluster's spec in the management cluster and waits
// until the changes are fully reconciled.
func (a Deleter) Run(ctx context.Context, spec *cluster.Spec, managementCluster types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
