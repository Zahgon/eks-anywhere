package tinkerbell

import (
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

//go:embed config/template-cp.yaml
var defaultCAPIConfigCP string

//go:embed config/template-md.yaml
var defaultClusterConfigMD string

const (
	TinkerbellMachineTemplateKind = "TinkerbellMachineTemplate"
	defaultRegistry               = "public.ecr.aws"
	netbootMode                   = "netboot"
	isobootMode                   = "iso"
	// SmeeHTTPPort is the port in Smee that serves HTTP requests.
	SmeeHTTPPort = "7171"
)

type TemplateBuilder struct {
	controlPlaneMachineSpec     *v1alpha1.TinkerbellMachineConfigSpec
	datacenterSpec              *v1alpha1.TinkerbellDatacenterConfigSpec
	WorkerNodeGroupMachineSpecs map[string]v1alpha1.TinkerbellMachineConfigSpec
	etcdMachineSpec             *v1alpha1.TinkerbellMachineConfigSpec
	tinkerbellIP                string
	now                         types.NowFunc
}

// NewTemplateBuilder creates a new TemplateBuilder instance.
func NewTemplateBuilder(datacenterSpec *v1alpha1.TinkerbellDatacenterConfigSpec, controlPlaneMachineSpec, etcdMachineSpec *v1alpha1.TinkerbellMachineConfigSpec, workerNodeGroupMachineSpecs map[string]v1alpha1.TinkerbellMachineConfigSpec, tinkerbellIP string, now types.NowFunc) providers.TemplateBuilder {
	_ = "STUB: not implemented"
	return *new(providers.TemplateBuilder)
}

func (tb *TemplateBuilder) GenerateCAPISpecControlPlane(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tb *TemplateBuilder) GenerateCAPISpecWorkers(clusterSpec *cluster.Spec, workloadTemplateNames, kubeadmconfigTemplateNames map[string]string) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func machineDeploymentName(clusterName, nodeGroupName string) string {
	_ = "STUB: not implemented"
	return ""
}

// nolint:gocyclo
func buildTemplateMapCP(
	clusterSpec *cluster.Spec,
	controlPlaneMachineSpec,
	etcdMachineSpec v1alpha1.TinkerbellMachineConfigSpec,
	cpTemplateOverride,
	etcdTemplateOverride string,
	datacenterSpec v1alpha1.TinkerbellDatacenterConfigSpec,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: need to get this values for creating template IMAGE_URL
// TODO: need to get this values for creating template IMAGE_URL
// TODO: need to get this values for creating template IMAGE_URL

// Replace public.ecr.aws endpoint with the endpoint given in the cluster config file

// isoURL path is only served in the top level /iso path.

func buildTemplateMapMD(
	clusterSpec *cluster.Spec,
	workerNodeGroupMachineSpec v1alpha1.TinkerbellMachineConfigSpec,
	workerNodeGroupConfiguration v1alpha1.WorkerNodeGroupConfiguration,
	workerTemplateOverride string,
	datacenterSpec v1alpha1.TinkerbellDatacenterConfigSpec,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Replace public.ecr.aws endpoint with the endpoint given in the cluster config file

// isoURL path is only served in the top level /iso path.

// omitTinkerbellMachineTemplate removes TinkerbellMachineTemplate API objects from yml. yml is
// typically an EKSA cluster configuration.
func omitTinkerbellMachineTemplate(yml []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Omit TinkerbellMachineTemplate kind.

func populateRegistryMirrorValues(clusterSpec *cluster.Spec, values map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getControlPlaneMachineSpec(clusterSpec *cluster.Spec) (*v1alpha1.TinkerbellMachineConfigSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getWorkerNodeGroupMachineSpec(clusterSpec *cluster.Spec) (map[string]v1alpha1.TinkerbellMachineConfigSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEtcdMachineSpec(clusterSpec *cluster.Spec) (*v1alpha1.TinkerbellMachineConfigSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateTemplateBuilder(clusterSpec *cluster.Spec) (providers.TemplateBuilder, error) {
	_ = "STUB: not implemented"
	return *new(providers.TemplateBuilder), nil
}

// GenerateNoProxyList generates NOPROXY list for tinkerbell provider based on HTTP_PROXY, HTTPS_PROXY, NOPROXY and tinkerbellIP.
func generateNoProxyList(clusterSpec *v1alpha1.Cluster, datacenterSpec v1alpha1.TinkerbellDatacenterConfigSpec) []string {
	_ = "STUB: not implemented"
	return nil
}
