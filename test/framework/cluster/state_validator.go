package cluster

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/retrier"
)

// StateValidation defines a validation that can be registered to the StateValidator.
type StateValidation = func(ctx context.Context, vc StateValidationConfig) error

// RetriableStateValidation returns a StateValidation that is executed with the provided retrier.
func RetriableStateValidation(retrier *retrier.Retrier, validation StateValidation) StateValidation {
	_ = "STUB: not implemented"
	return *new(StateValidation)
}

// StateValidator is responsible for checking if a cluster is valid against the spec that is provided.
type StateValidator struct {
	Config      StateValidationConfig
	validations []StateValidation
}

// WithValidations registers multiple validations to the StateValidator that will be run when Validate is called.
func (c *StateValidator) WithValidations(validations ...StateValidation) {
	_ = "STUB: not implemented"
	return
}

// Validate runs through the set registered validations and returns an error if any of them fail after a number of retries.
func (c *StateValidator) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Opt represents is a function that represents an option to configure a StateValidator.
type Opt = func(cv *StateValidator)

// NewStateValidator returns a cluster validator which can be configured by passing Opt arguments.
func NewStateValidator(config StateValidationConfig, opts ...Opt) *StateValidator {
	_ = "STUB: not implemented"
	return nil
}

// StateValidationConfig represents the input for the performing validations on the cluster.
type StateValidationConfig struct {
	ClusterClient           client.Client // the client for the cluster
	ManagementClusterClient client.Client // the client for the management cluster
	ClusterSpec             *cluster.Spec // the cluster spec
}
