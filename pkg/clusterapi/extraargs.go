package clusterapi

import (
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type ExtraArgs map[string]string

// ToArgs converts ExtraArgs (map[string]string) to the v1beta2 []bootstrapv1.Arg format.
// The output is sorted by name for deterministic ordering.
// SortArgs sorts a slice of bootstrapv1beta2.Arg by name for deterministic ordering.
func SortArgs(args []bootstrapv1beta2.Arg) { _ = "STUB: not implemented"; return }

func (e ExtraArgs) ToArgs() []bootstrapv1beta2.Arg { _ = "STUB: not implemented"; return nil }

// copy for pointer

func OIDCToExtraArgs(oidc *v1alpha1.OIDCConfig) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func AwsIamAuthExtraArgs(awsiam *v1alpha1.AWSIamConfig) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

// EtcdEncryptionExtraArgs takes a list of EtcdEncryption configs and returns the relevant API server extra args if it's not nil or empty.
func EtcdEncryptionExtraArgs(config *[]v1alpha1.EtcdEncryption) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

// APIServerExtraArgs takes a map of API Server extra args and returns the relevant API server extra args if it's not nil or empty.
func APIServerExtraArgs(apiServerExtraArgs map[string]string) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func PodIAMAuthExtraArgs(podIAMConfig *v1alpha1.PodIAMConfig) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func NodeCIDRMaskExtraArgs(clusterNetwork *v1alpha1.ClusterNetwork) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func ResolvConfExtraArgs(resolvConf *v1alpha1.ResolvConf) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

// We don't need to add these once the Kubernetes components default to using the secure cipher suites.
func SecureTlsCipherSuitesExtraArgs() ExtraArgs { _ = "STUB: not implemented"; return *new(ExtraArgs) }

func SecureEtcdTlsCipherSuitesExtraArgs() ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func WorkerNodeLabelsExtraArgs(wnc v1alpha1.WorkerNodeGroupConfiguration) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func ControlPlaneNodeLabelsExtraArgs(cpc v1alpha1.ControlPlaneConfiguration) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

// CgroupDriverExtraArgs args added for kube versions below 1.24.
func CgroupDriverCgroupfsExtraArgs() ExtraArgs { _ = "STUB: not implemented"; return *new(ExtraArgs) }

// CgroupDriverSystemdExtraArgs args added for kube versions 1.24 and above.
func CgroupDriverSystemdExtraArgs() ExtraArgs { _ = "STUB: not implemented"; return *new(ExtraArgs) }

func nodeLabelsExtraArgs(labels map[string]string) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

func (e ExtraArgs) AddIfNotEmpty(k, v string) { _ = "STUB: not implemented"; return }

func (e ExtraArgs) Append(args ExtraArgs) ExtraArgs {
	_ = "STUB: not implemented"
	return *new(ExtraArgs)
}

// SetPodIAMAuthExtraArgs sets the api server extra args for the podIAMConfig.
func SetPodIAMAuthExtraArgs(podIAMConfig *v1alpha1.PodIAMConfig, apiServerExtraArgs map[string]string) {
	_ = "STUB: not implemented"
	return
}

// SetPodIAMAuthInArgs merges PodIAM service-account-issuer into an existing []Arg slice.
// If service-account-issuer already exists, it concatenates the values with a comma
// (matching the behavior of SetPodIAMAuthExtraArgs for map[string]string).
func SetPodIAMAuthInArgs(podIAMConfig *v1alpha1.PodIAMConfig, args []bootstrapv1beta2.Arg) []bootstrapv1beta2.Arg {
	_ = "STUB: not implemented"
	return nil
}

// ToYaml outputs ExtraArgs as v1beta2 []Arg YAML format (list of name/value pairs).
// Output is sorted by name for deterministic ordering.
func (e ExtraArgs) ToYaml() string { _ = "STUB: not implemented"; return "" }

// Sort keys for deterministic output

func requiredClaimToArg(r *v1alpha1.OIDCConfigRequiredClaim) string {
	_ = "STUB: not implemented"
	return ""
}

func labelsMapToArg(m map[string]string) string { _ = "STUB: not implemented"; return "" }
