package cloudstack

import (
	"context"

	"github.com/go-logr/logr"
	cloudstackv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	yamlcapi "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// BaseControlPlane represents a CAPI CloudStack control plane.
type BaseControlPlane = clusterapi.ControlPlane[*cloudstackv1.CloudStackCluster, *cloudstackv1.CloudStackMachineTemplate]

// ControlPlane holds the CloudStack specific objects for a CAPI CloudStack control plane.
type ControlPlane struct {
	BaseControlPlane
}

// Objects returns the control plane objects associated with the CloudStack cluster.
func (p ControlPlane) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// ControlPlaneSpec builds a CloudStack ControlPlane definition based on an eks-a cluster spec.
func ControlPlaneSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, clusterSpec *cluster.Spec) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ControlPlaneBuilder defines the builder for all objects in the CAPI CloudStack control plane.
type controlPlaneBuilder struct {
	BaseBuilder  *yamlcapi.ControlPlaneBuilder[*cloudstackv1.CloudStackCluster, *cloudstackv1.CloudStackMachineTemplate]
	ControlPlane *ControlPlane
}

// BuildFromParsed implements the base yamlcapi.BuildFromParsed and processes any additional objects for the CloudStack control plane.
func (b *controlPlaneBuilder) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

func newControlPlaneParser(logger logr.Logger) (*yamlutil.Parser, *controlPlaneBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
