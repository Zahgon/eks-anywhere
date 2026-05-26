package yaml

import (
	"github.com/go-logr/logr"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

const machineDeploymentKind = "MachineDeployment"

// WorkersBuilder implements yamlutil.Builder
// It's a wrapper around Workers to provide yaml parsing functionality.
type WorkersBuilder[M clusterapi.Object[M]] struct {
	Workers *clusterapi.Workers[M]
}

// NewWorkersBuilder builds a WorkersBuilder.
func NewWorkersBuilder[M clusterapi.Object[M]]() *WorkersBuilder[M] {
	_ = "STUB: not implemented"
	return nil
}

// BuildFromParsed reads parsed objects in ObjectLookup and sets them in the Workers.
func (cp *WorkersBuilder[M]) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

// NewWorkersParserAndBuilder builds a Parser and a Builder for a particular provider Workers
// It registers the basic shared mappings plus another one for the provider machine template
// For worker specs that need to include more objects, wrap around the provider builder and
// implement BuildFromParsed.
// Any extra mappings will need to be registered manually in the Parser.
func NewWorkersParserAndBuilder[M clusterapi.Object[M]](
	logger logr.Logger,
	machineTemplateMapping yamlutil.Mapping[M],
) (*yamlutil.Parser, *WorkersBuilder[M], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RegisterWorkerMappings records the basic mappings for CAPI MachineDeployment
// and KubeadmConfigTemplate in a Parser.
func RegisterWorkerMappings(parser *yamlutil.Parser) error { _ = "STUB: not implemented"; return nil }

// ProcessWorkerObjects finds all necessary objects in the parsed objects and sets them in Workers.
func ProcessWorkerObjects[M clusterapi.Object[M]](w *clusterapi.Workers[M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// ProcessWorkerGroupObjects looks in the parsed objects for the KubeadmConfigTemplate and
// the provider machine template referenced in the MachineDeployment and sets them in the WorkerGroup.
// MachineDeployment needs to be already set in the WorkerGroup.
func ProcessWorkerGroupObjects[M clusterapi.Object[M]](g *clusterapi.WorkerGroup[M], lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// sortKCTExtraArgs sorts all ExtraArgs slices in a KubeadmConfigTemplate for deterministic comparison.
func sortKCTExtraArgs(kct *bootstrapv1beta2.KubeadmConfigTemplate) {
	_ = "STUB: not implemented"
	return
}
