package vsphere

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

func NewVsphereTemplateBuilder(
	now types.NowFunc,
) *VsphereTemplateBuilder {
	_ = "STUB: not implemented"
	return nil
}

type VsphereTemplateBuilder struct {
	now types.NowFunc
}

func (vs *VsphereTemplateBuilder) GenerateCAPISpecControlPlane(
	clusterSpec *cluster.Spec,
	buildOptions ...providers.BuildMapOption,
) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vs *VsphereTemplateBuilder) isCgroupDriverSystemd(clusterSpec *cluster.Spec, worker anywherev1.WorkerNodeGroupConfiguration) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CAPIWorkersSpecWithInitialNames generates a yaml spec with the CAPI objects representing the worker
// nodes for a particular eks-a cluster. It uses default initial names (ended in '-1') for the vsphere
// machine templates and kubeadm config templates.
func (vs *VsphereTemplateBuilder) CAPIWorkersSpecWithInitialNames(spec *cluster.Spec) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vs *VsphereTemplateBuilder) GenerateCAPISpecWorkers(
	clusterSpec *cluster.Spec,
	workloadTemplateNames,
	kubeadmconfigTemplateNames map[string]string,
) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pin cgroupDriver to systemd for k8s >= 1.21 when generating template in controller
// remove this check once the controller supports order upgrade.
// i.e. control plane, etcd upgrade before worker nodes.

// GenerateVsphereFailureDomainsSpec generates a yaml spec with the Vsphere failure domains objects.
// It uses the provided template names for the VsphereFailureDomain and VsphereDeploymentZone.
func (vs *VsphereTemplateBuilder) GenerateVsphereFailureDomainsSpec(spec *cluster.Spec, templateNames map[string]string) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildTemplateMapCP(
	clusterSpec *cluster.Spec,
	datacenterSpec anywherev1.VSphereDatacenterConfigSpec,
	controlPlaneMachineSpec, etcdMachineSpec anywherev1.VSphereMachineConfigSpec,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add no-proxy defaults

func buildTemplateMapMD(
	clusterSpec *cluster.Spec,
	datacenterSpec anywherev1.VSphereDatacenterConfigSpec,
	workerNodeGroupMachineSpec anywherev1.VSphereMachineConfigSpec,
	workerNodeGroupConfiguration anywherev1.WorkerNodeGroupConfiguration,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add no-proxy defaults

func buildTemplateMapFailureDomain(
	clusterSpec *cluster.Spec,
	failureDomain anywherev1.FailureDomain,
) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Currently, we only support compute cluster topology in failure domain
// In future, when we add supports for other topologies, update this get region type and name based on topology type.
// For example, if topology type is host group, region will be one level above host group i.e ComputeCluster.
func getFailureDomainRegionTypeAndName(datacenterSpec anywherev1.VSphereDatacenterConfigSpec) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Currently, we only support compute cluster topology in failure domain
// In future, when we add supports for other topologies, update this get zone type and name based on topology type.
// For example, if topology type is host group, zone type = HostGroup and name = host group name.
func getFailureDomainZoneTypeAndName(failureDomain anywherev1.FailureDomain) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}
