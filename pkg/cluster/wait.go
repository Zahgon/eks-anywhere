package cluster

import (
	"context"

	"github.com/go-logr/logr"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/retrier"
)

// WaitForCondition blocks until either the cluster has this condition as True
// or the retrier timeouts. If observedGeneration is not equal to generation,
// the condition is considered false regardless of the status value.
// total field is to check the total number of times the given condition is met for consistency.
func WaitForCondition(ctx context.Context, log logr.Logger, client kubernetes.Reader, cluster *anywherev1.Cluster, total int, retrier *retrier.Retrier, conditionType anywherev1.ConditionType) error {
	_ = "STUB: not implemented"
	return nil
}

// Matcher matches the given condition.
type Matcher func(*anywherev1.Cluster) error

// WaitFor gets the cluster object from the client
// checks for generation and observedGeneration condition
// matches condition and returns error if the condition is not met.
func WaitFor(ctx context.Context, log logr.Logger, client kubernetes.Reader, cluster *anywherev1.Cluster, total int, retrier *retrier.Retrier, matcher Matcher) error {
	_ = "STUB: not implemented"
	return nil
}

// total field is to check the total number of times the given condition is met.
// for ex, the total is set to 5 and we want to check certain condition is met
// we check the condition is met for 5 times to make sure the given behavior is consistent.

// when the count matches total number it returns without error
