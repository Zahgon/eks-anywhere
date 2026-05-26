package nutanix

import (
	"encoding/json"
	"net"

	capxv1beta1 "github.com/nutanix-cloud-native/cluster-api-provider-nutanix/api/v1beta1"
	"github.com/nutanix-cloud-native/prism-go-client/environment/credentials"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

var jsonMarshal = json.Marshal

// TemplateBuilder builds templates for nutanix.
type TemplateBuilder struct {
	datacenterSpec              *v1alpha1.NutanixDatacenterConfigSpec
	controlPlaneMachineSpec     *v1alpha1.NutanixMachineConfigSpec
	etcdMachineSpec             *v1alpha1.NutanixMachineConfigSpec
	workerNodeGroupMachineSpecs map[string]v1alpha1.NutanixMachineConfigSpec
	creds                       credentials.BasicAuthCredential
	now                         types.NowFunc
}

var _ providers.TemplateBuilder = &TemplateBuilder{}

func NewNutanixTemplateBuilder(
	datacenterSpec *v1alpha1.NutanixDatacenterConfigSpec,
	controlPlaneMachineSpec,
	etcdMachineSpec *v1alpha1.NutanixMachineConfigSpec,
	workerNodeGroupMachineSpecs map[string]v1alpha1.NutanixMachineConfigSpec,
	creds credentials.BasicAuthCredential,
	now types.NowFunc,
) *TemplateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ntb *TemplateBuilder) GenerateCAPISpecControlPlane(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ntb *TemplateBuilder) GenerateCAPISpecWorkers(clusterSpec *cluster.Spec, workloadTemplateNames, kubeadmconfigTemplateNames map[string]string) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateCAPISpecSecret generates the secret containing the credentials for the nutanix prism central and is used by the
// CAPX controller. The secret is named after the cluster name.
func (ntb *TemplateBuilder) GenerateCAPISpecSecret(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CAPXSecretName returns the name of the secret containing the credentials for the nutanix prism central and is used by the
// CAPX controller.
func CAPXSecretName(spec *cluster.Spec) string { _ = "STUB: not implemented"; return "" }

// GenerateEKSASpecSecret generates the secret containing the credentials for the nutanix prism central and is used by the
// EKS-A controller. The secret is named nutanix-credentials.
func (ntb *TemplateBuilder) GenerateEKSASpecSecret(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EKSASecretName returns the name of the secret containing the credentials for the nutanix prism central and is used by the
// EKS-Anywhere controller.
func EKSASecretName(spec *cluster.Spec) string { _ = "STUB: not implemented"; return "" }

func (ntb *TemplateBuilder) generateSpecSecret(clusterSpec *cluster.Spec, secretName string, creds credentials.BasicAuthCredential, buildOptions ...providers.BuildMapOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func machineDeploymentName(clusterName, nodeGroupName string) string {
	_ = "STUB: not implemented"
	return ""
}

func buildTemplateMapCP(
	datacenterSpec *v1alpha1.NutanixDatacenterConfigSpec,
	clusterSpec *cluster.Spec,
	controlPlaneMachineSpec v1alpha1.NutanixMachineConfigSpec,
	etcdMachineSpec v1alpha1.NutanixMachineConfigSpec,
	creds credentials.BasicAuthCredential,
) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calcFailureDomainReplicas(workerNodeGroupConfiguration v1alpha1.WorkerNodeGroupConfiguration, failureDomains []v1alpha1.NutanixDatacenterFailureDomain) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func getFailureDomainsForWorkerNodeGroup(allFailureDomains []v1alpha1.NutanixDatacenterFailureDomain, workerNodeGroupConfigurationName string) []v1alpha1.NutanixDatacenterFailureDomain {
	_ = "STUB: not implemented"
	return nil
}

func buildTemplateMapMD(clusterSpec *cluster.Spec, workerNodeGroupMachineSpec v1alpha1.NutanixMachineConfigSpec, workerNodeGroupConfiguration v1alpha1.WorkerNodeGroupConfiguration) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildTemplateMapSecret(clusterSpec *cluster.Spec, secretName string, creds credentials.BasicAuthCredential) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateNoProxyList(clusterSpec *cluster.Spec) []string { _ = "STUB: not implemented"; return nil }

// Add no-proxy defaults

func generateNutanixFailureDomains(eksNutanixFailureDomains []v1alpha1.NutanixDatacenterFailureDomain) []capxv1beta1.NutanixFailureDomain {
	_ = "STUB: not implemented"
	return nil
}

func incrementIP(ip net.IP) { _ = "STUB: not implemented"; return }

func compareIP(ip1, ip2 net.IP) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func addCIDRToIgnoredNodeIPsList(cidr string, result []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Add all ip addresses in the range to the list

func addIPRangeToIgnoredNodeIPsList(ipRangeStr string, result []string) []string {
	_ = "STUB: not implemented"
	// Parse the range
	return nil
}

// Parse the start and end of the range

// swap start and end if start is greater than end

// Add all ip addresses in the range to the list

func addIPAddressToIgnoredNodeIPsList(ipAddrStr string, result []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func generateCcmIgnoredNodeIPsList(clusterSpec *cluster.Spec) []string {
	_ = "STUB: not implemented"
	// Add the kube-vip IP address to the list
	return nil
}
