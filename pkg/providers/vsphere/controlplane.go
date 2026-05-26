package vsphere

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
	addonsv1 "sigs.k8s.io/cluster-api/api/addons/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	yamlcapi "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// BaseControlPlane represents a CAPI VSphere control plane.
type BaseControlPlane = clusterapi.ControlPlane[*vspherev1.VSphereCluster, *vspherev1.VSphereMachineTemplate]

// ControlPlane holds the VSphere specific objects for a CAPI VSphere control plane.
type ControlPlane struct {
	BaseControlPlane
	Secrets             []*corev1.Secret
	ConfigMaps          []*corev1.ConfigMap
	ClusterResourceSets []*addonsv1.ClusterResourceSet
}

// Objects returns the control plane objects associated with the VSphere cluster.
func (p ControlPlane) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// ControlPlaneBuilder defines the builder for all objects in the CAPI VSphere control plane.
type ControlPlaneBuilder struct {
	BaseBuilder  *yamlcapi.ControlPlaneBuilder[*vspherev1.VSphereCluster, *vspherev1.VSphereMachineTemplate]
	ControlPlane *ControlPlane
}

// BuildFromParsed implements the base yamlcapi.BuildFromParsed and processes any additional objects for the VSphere control plane.
func (b *ControlPlaneBuilder) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

// ControlPlaneSpec builds a vsphere ControlPlane definition based on an eks-a cluster spec.
func ControlPlaneSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, spec *cluster.Spec) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newControlPlaneParser(logger logr.Logger) (*yamlutil.Parser, *ControlPlaneBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func processObjects(c *ControlPlane, lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

func getSecrets(o []kubernetes.Object, secrets []*corev1.Secret) []kubernetes.Object {
	_ = "STUB: not implemented"
	return nil
}

func getConfigMaps(o []kubernetes.Object, configMaps []*corev1.ConfigMap) []kubernetes.Object {
	_ = "STUB: not implemented"
	return nil
}

func getClusterResourceSets(o []kubernetes.Object, clusterResourceSets []*addonsv1.ClusterResourceSet) []kubernetes.Object {
	_ = "STUB: not implemented"
	return nil
}
