package tinkerbell

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"

	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	yamlcapi "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// BaseControlPlane represents a CAPI Tinkerbell control plane.
type BaseControlPlane = clusterapi.ControlPlane[*tinkerbellv1.TinkerbellCluster, *tinkerbellv1.TinkerbellMachineTemplate]

// ControlPlane holds the Tinkerbell specific objects for a CAPI Tinkerbell control plane.
type ControlPlane struct {
	BaseControlPlane
	Secrets *corev1.Secret
}

// Objects returns the control plane objects associated with the Tinkerbell cluster.
func (p ControlPlane) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// ControlPlaneSpec builds a Tinkerbell ControlPlane definition based on an eks-a cluster spec.
func ControlPlaneSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, clusterSpec *cluster.Spec) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ControlPlaneBuilder defines the builder for all objects in the CAPI Tinkerbell control plane.
type controlPlaneBuilder struct {
	BaseBuilder  *yamlcapi.ControlPlaneBuilder[*tinkerbellv1.TinkerbellCluster, *tinkerbellv1.TinkerbellMachineTemplate]
	ControlPlane *ControlPlane
}

// BuildFromParsed implements the base yamlcapi.BuildFromParsed and processes any additional objects (secrets) for the Tinkerbell control plane.
func (b *controlPlaneBuilder) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

func newControlPlaneParser(logger logr.Logger) (*yamlutil.Parser, *controlPlaneBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
