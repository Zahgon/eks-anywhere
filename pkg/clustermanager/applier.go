package clustermanager

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	applyClusterSpecTimeout           = 2 * time.Minute
	waitForClusterReconcileTimeout    = time.Hour
	retryBackOff                      = time.Second
	waitForFailureMessageErrorTimeout = 10 * time.Minute
	defaultFieldManager               = "eks-a-cli"
	defaultConditionCheckTotalCount   = 20
)

// ApplierOpt allows to customize a Applier on construction.
type ApplierOpt func(*Applier)

// Applier applies the cluster spec to the management cluster and waits
// until the changes are fully reconciled.
type Applier struct {
	log                                                                 logr.Logger
	clientFactory                                                       ClientFactory
	applyClusterTimeout, waitForClusterReconcile, waitForFailureMessage time.Duration
	retryBackOff                                                        time.Duration
	conditionCheckoutTotalCount                                         int
}

// NewApplier builds an Applier.
func NewApplier(log logr.Logger, clientFactory ClientFactory, opts ...ApplierOpt) Applier {
	_ = "STUB: not implemented"
	return *new(Applier)
}

// WithApplierNoTimeouts disables the timeout for all the waits and retries in management upgrader.
func WithApplierNoTimeouts() ApplierOpt { _ = "STUB: not implemented"; return *new(ApplierOpt) }

// WithApplierApplyClusterTimeout allows to configure how long the applier retries
// to apply the objects in case of failure.
// Generally only used in tests.
func WithApplierApplyClusterTimeout(timeout time.Duration) ApplierOpt {
	_ = "STUB: not implemented"
	return *new(ApplierOpt)
}

// WithApplierWaitForClusterReconcile allows to configure how long the applier waits
// for the cluster to reach the Ready state after applying changes.
// Generally only used in tests.
func WithApplierWaitForClusterReconcile(timeout time.Duration) ApplierOpt {
	_ = "STUB: not implemented"
	return *new(ApplierOpt)
}

// WithApplierRetryBackOff allows to configure how long the applier waits between requests
// to update the cluster spec objects and check the status of the Cluster.
// Generally only used in tests.
func WithApplierRetryBackOff(backOff time.Duration) ApplierOpt {
	_ = "STUB: not implemented"
	return *new(ApplierOpt)
}

// WithApplierWaitForFailureMessage allows to configure how long the applier waits for failure message
// to be empty and check the status of the Cluster.
// Generally only used in tests.
func WithApplierWaitForFailureMessage(timeout time.Duration) ApplierOpt {
	_ = "STUB: not implemented"
	return *new(ApplierOpt)
}

// Run applies the cluster's spec in the management cluster and waits
// until the changes are fully reconciled.
func (a Applier) Run(ctx context.Context, spec *cluster.Spec, managementCluster types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// We build the client inside the retrier to take advantage of the configurable timeout.
// The only time when a client can fail to init is when there is a transient error contacting
// the api-server (there might be some validations checks during init). These are the same
// transient errors we want to protect the apply against, so it makes sense to group them together.
// The drawback here is we rebuild the client on each retry if one of the objects fails to be
// created/updated. It's a fair tradeoff given this is already handling an edge case.

// We use this start time to compute the leftover time on each condition wait

func (a Applier) retrierForWait(waitStartTime time.Time) *retrier.Retrier {
	_ = "STUB: not implemented"
	return nil
}

func (a Applier) retrierForFailureMessage() *retrier.Retrier { _ = "STUB: not implemented"; return nil }
