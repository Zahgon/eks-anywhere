package cloudstack

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

// TemplateBuilder is responsible for building the CAPI templates.
type TemplateBuilder struct {
	now types.NowFunc
}

// NewTemplateBuilder creates a new TemplateBuilder.
func NewTemplateBuilder(now types.NowFunc) *TemplateBuilder { _ = "STUB: not implemented"; return nil }

// GenerateCAPISpecControlPlane builds the CAPI controlplane template containing the CAPI objects for the control plane configuration defined in the cluster.Spec.
func (cs *TemplateBuilder) GenerateCAPISpecControlPlane(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateCAPISpecWorkers builds the CAPI worker template containing the CAPI objects for the worker node groups configuration defined in the cluster.Spec.
func (cs *TemplateBuilder) GenerateCAPISpecWorkers(clusterSpec *cluster.Spec, workloadTemplateNames, kubeadmconfigTemplateNames map[string]string) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Extract out worker MachineDeployments from templates to use apibuilder instead

// nolint:gocyclo
func buildTemplateMapCP(clusterSpec *cluster.Spec) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildControlPlaneTemplate(machineSpec *v1alpha1.CloudStackMachineConfigSpec, values map[string]interface{}) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildEtcdTemplate(machineSpec *v1alpha1.CloudStackMachineConfigSpec, values map[string]interface{}) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fillDiskOffering(values map[string]interface{}, diskOffering *v1alpha1.CloudStackResourceDiskOffering, machineType string) {
	_ = "STUB: not implemented"
	return
}

func fillProxyConfigurations(values map[string]interface{}, clusterSpec *cluster.Spec, controlPlaneEndpoint string) {
	_ = "STUB: not implemented"
	return
}

func buildTemplateMapMD(clusterSpec *cluster.Spec, workerNodeGroupConfiguration v1alpha1.WorkerNodeGroupConfiguration) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
