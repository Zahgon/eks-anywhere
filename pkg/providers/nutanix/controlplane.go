package nutanix

import (
	"context"

	"github.com/go-logr/logr"
	nutanixv1 "github.com/nutanix-cloud-native/cluster-api-provider-nutanix/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	addonsv1 "sigs.k8s.io/cluster-api/api/addons/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	yamlcapi "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// BaseControlPlane represents a CAPI Nutanix control plane.
type BaseControlPlane = clusterapi.ControlPlane[*nutanixv1.NutanixCluster, *nutanixv1.NutanixMachineTemplate]

// ControlPlane holds the Nutanix specific objects for a CAPI Nutanix control plane.
type ControlPlane struct {
	BaseControlPlane
	ConfigMaps          []*corev1.ConfigMap
	ClusterResourceSets []*addonsv1.ClusterResourceSet
	Secrets             []*corev1.Secret
}

// Objects returns the control plane objects associated with the Nutanix cluster.
func (p ControlPlane) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// ControlPlaneBuilder defines the builder for all objects in the CAPI Nutanix control plane.
type ControlPlaneBuilder struct {
	BaseBuilder  *yamlcapi.ControlPlaneBuilder[*nutanixv1.NutanixCluster, *nutanixv1.NutanixMachineTemplate]
	ControlPlane *ControlPlane
}

// BuildFromParsed implements the base yamlcapi.BuildFromParsed and processes any additional objects for the Nutanix control plane.
func (b *ControlPlaneBuilder) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

// ControlPlaneSpec builds a nutanix ControlPlane definition based on an eks-a cluster spec.
func ControlPlaneSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, spec *cluster.Spec) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getControlPlaneMachineSpecs(machineConfigs map[string]*v1alpha1.NutanixMachineConfig, controlPlaneConfig *v1alpha1.ControlPlaneConfiguration, externalEtcdConfig *v1alpha1.ExternalEtcdConfiguration) (*v1alpha1.NutanixMachineConfigSpec, *v1alpha1.NutanixMachineConfigSpec) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateControlPlaneYAML(templateBuilder *TemplateBuilder, spec *cluster.Spec) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseControlPlaneYAML(logger logr.Logger, controlPlaneYAML []byte) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newControlPlaneParser(logger logr.Logger) (*yamlutil.Parser, *ControlPlaneBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func appendKubeObjects[V kubernetes.Object](objList []kubernetes.Object, objToAdd []V) []kubernetes.Object {
	_ = "STUB: not implemented"
	return nil
}

func buildObjects(cp *ControlPlane, lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}
