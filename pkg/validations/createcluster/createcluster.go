package createcluster

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/gitops/flux"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/validations"
)

type ValidationManager struct {
	clusterSpec       *cluster.Spec
	provider          providers.Provider
	gitOpsFlux        *flux.Flux
	createValidations Validator
	dockerExec        validations.DockerExecutable
}

type Validator interface {
	PreflightValidations(ctx context.Context) []validations.Validation
}

func NewValidations(clusterSpec *cluster.Spec, provider providers.Provider, gitOpsFlux *flux.Flux, createValidations Validator, dockerExec validations.DockerExecutable) *ValidationManager {
	_ = "STUB: not implemented"
	return nil
}

func (v *ValidationManager) Validate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *ValidationManager) generateCreateValidations(ctx context.Context) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}
