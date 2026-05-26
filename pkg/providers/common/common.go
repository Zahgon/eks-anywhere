package common

import (
	_ "embed"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/types"
)

//go:embed config/audit-policy.yaml
var auditPolicy string

// TODO: Split out common into separate packages to avoid becoming a dumping ground

const (
	privateKeyFileName = "eks-a-id_rsa"
	publicKeyFileName  = "eks-a-id_rsa.pub"
)

func BootstrapClusterOpts(clusterConfig *v1alpha1.Cluster, serverEndpoints ...string) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StripSshAuthorizedKeyComment(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ssh.MarshalAuthorizedKey returns the key with a trailing newline, which we want to remove

func GenerateSSHAuthKey(writer filewriter.FileWriter) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CPMachineTemplateBase(clusterName string) string { _ = "STUB: not implemented"; return "" }

func EtcdMachineTemplateBase(clusterName string) string { _ = "STUB: not implemented"; return "" }

func WorkerMachineTemplateBase(clusterName, workerNodeGroupName string) string {
	_ = "STUB: not implemented"
	return ""
}

func CPMachineTemplateName(clusterName string, now types.NowFunc) string {
	_ = "STUB: not implemented"
	return ""
}

func EtcdMachineTemplateName(clusterName string, now types.NowFunc) string {
	_ = "STUB: not implemented"
	return ""
}

func WorkerMachineTemplateName(clusterName, workerNodeGroupName string, now types.NowFunc) string {
	_ = "STUB: not implemented"
	return ""
}

func KubeadmConfigTemplateName(clusterName, workerNodeGroupName string, now types.NowFunc) string {
	_ = "STUB: not implemented"
	return ""
}

// GetCAPIBottlerocketSettingsConfig returns the formatted CAPI Bottlerocket settings config as a YAML marshaled string.
func GetCAPIBottlerocketSettingsConfig(config *v1alpha1.HostOSConfiguration, brKubeSettings *bootstrapv1beta2.BottlerocketKubernetesSettings) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCAPIConfig(b *bootstrapv1beta2.BottlerocketSettings) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetExternalEtcdReleaseURL returns a valid etcd URL  from version bundles if the eksaVersion is greater than
// MinEksAVersionWithEtcdURL. Return "" if eksaVersion < MinEksAVersionWithEtcdURL to prevent etcd node rolled out.
func GetExternalEtcdReleaseURL(clusterVersion *v1alpha1.EksaVersion, versionBundle *cluster.VersionsBundle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ConvertToBottlerocketKubernetesSettings converts an unstructured object into a Bottlerocket
// Kubernetes settings object.
func ConvertToBottlerocketKubernetesSettings(kubeletConfig *unstructured.Unstructured) (*bootstrapv1beta2.BottlerocketKubernetesSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyObject(kubeletConfig *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyBottlerocketKubernetesSettings(config *bootstrapv1beta2.BottlerocketKubernetesSettings) *bootstrapv1beta2.BottlerocketKubernetesSettings {
	_ = "STUB: not implemented"
	return nil
}

func copyBottlerocketMaps(source map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
