package yaml

import (
	"github.com/go-logr/logr"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// NewControlPlaneParserAndBuilder builds a Parser and a Builder for a particular provider ControlPlane
// It registers the basic shared mappings plus another two for the provider cluster and machine template
// For ControlPlane that need to include more objects, wrap around the provider builder and implement BuildFromParsed
// Any extra mappings will need to be registered manually in the Parser.
func NewControlPlaneParserAndBuilder[C clusterapi.Object[C], M clusterapi.Object[M]](logger logr.Logger, clusterMapping yamlutil.Mapping[C], machineTemplateMapping yamlutil.Mapping[M]) (*yamlutil.Parser, *ControlPlaneBuilder[C, M], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RegisterControlPlaneMappings records the basic mappings for CAPI cluster, kubeadmcontrolplane
// and etcdadm cluster in a Parser.
func RegisterControlPlaneMappings(parser *yamlutil.Parser) error {
	_ = "STUB: not implemented"
	return nil
}

// ControlPlaneBuilder implements yamlutil.Builder
// It's a wrapper around ControlPlane to provide yaml parsing functionality.
type ControlPlaneBuilder[C clusterapi.Object[C], M clusterapi.Object[M]] struct {
	ControlPlane *clusterapi.ControlPlane[C, M]
}

// NewControlPlaneBuilder builds a ControlPlaneBuilder.
func NewControlPlaneBuilder[C clusterapi.Object[C], M clusterapi.Object[M]]() *ControlPlaneBuilder[C, M] {
	_ = "STUB: not implemented"
	return nil
}

// BuildFromParsed reads parsed objects in ObjectLookup and sets them in the ControlPlane.
func (cp *ControlPlaneBuilder[C, M]) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessControlPlaneObjects finds all necessary objects in the parsed objects and sets them in the ControlPlane.
func ProcessControlPlaneObjects[C clusterapi.Object[C], M clusterapi.Object[M]](cp *clusterapi.ControlPlane[C, M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// ProcessCluster finds the CAPI cluster in the parsed objects and sets it in ControlPlane.
func ProcessCluster[C clusterapi.Object[C], M clusterapi.Object[M]](cp *clusterapi.ControlPlane[C, M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// ProcessProviderCluster finds the provider cluster in the parsed objects and sets it in ControlPlane.
func ProcessProviderCluster[C clusterapi.Object[C], M clusterapi.Object[M]](cp *clusterapi.ControlPlane[C, M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// ProcessKubeadmControlPlane finds the CAPI kubeadm control plane and the kubeadm control plane machine template
// in the parsed objects and sets it in ControlPlane.
func ProcessKubeadmControlPlane[C clusterapi.Object[C], M clusterapi.Object[M]](cp *clusterapi.ControlPlane[C, M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// ProcessEtcdCluster finds the CAPI etcdadm cluster (for unstacked clusters) in the parsed objects and sets it in ControlPlane.
func ProcessEtcdCluster[C clusterapi.Object[C], M clusterapi.Object[M]](cp *clusterapi.ControlPlane[C, M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// sortKCPExtraArgs sorts all ExtraArgs slices in a KubeadmControlPlane for deterministic comparison.
// YAML-parsed KCP objects have args in template order, while test helpers use alphabetical order via ToArgs().
func sortKCPExtraArgs(kcp *controlplanev1beta2.KubeadmControlPlane) {
	_ = "STUB: not implemented"
	return
}
