package framework

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	// NTP configuration environment variables.
	ntpServersVar = "T_NTP_SERVERS"

	// Bottlerocket configuration environment variables.
	maxPodsVar              = "T_BR_K8S_SETTINGS_MAX_PODS"
	clusterDNSIPSVar        = "T_BR_K8S_SETTINGS_CLUSTER_DNS_IPS"
	allowedUnsafeSysctlsVar = "T_BR_K8S_SETTINGS_ALLOWED_UNSAFE_SYSCTLS"

	// other constants.
	defaultSSHUsername = "ec2-user"
	privateKeyFileName = "eks-a-id_rsa"
)

var (
	ntpServersRequiredVar   = []string{ntpServersVar}
	brKubernetesRequiredVar = []string{maxPodsVar, clusterDNSIPSVar, allowedUnsafeSysctlsVar}
)

// RequiredNTPServersEnvVars returns a slice of environment variables required for NTP tests.
func RequiredNTPServersEnvVars() []string { _ = "STUB: not implemented"; return nil }

// RequiredBottlerocketKubernetesSettingsEnvVars returns a slice of environment variables required for Bottlerocket Kubernetes tests.
func RequiredBottlerocketKubernetesSettingsEnvVars() []string {
	_ = "STUB: not implemented"
	return nil
}

// GetNTPServersFromEnv returns a slice of NTP servers read from the NTP environment veriables.
func GetNTPServersFromEnv() []string { _ = "STUB: not implemented"; return nil }

// GetBottlerocketKubernetesSettingsFromEnv returns a Bottlerocket Kubernetes settings read from the environment variables.
func GetBottlerocketKubernetesSettingsFromEnv() (allowedUnsafeSysclts, clusterDNSIPS []string, maxPods int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// ValidateNTPConfig validates NTP servers are configured properly on all cluster nodes using SSH.
func (e *ClusterE2ETest) ValidateNTPConfig(osFamily v1alpha1.OSFamily) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) validateNTP(ctx context.Context, osFamily v1alpha1.OSFamily, IP string) {
	_ = "STUB: not implemented"
	return
}

// ValidateBottlerocketKubernetesSettings validates Bottlerocket Kubernetes settings are configured properly on all cluster nodes using SSH.
func (e *ClusterE2ETest) ValidateBottlerocketKubernetesSettings() {
	_ = "STUB: not implemented"
	return
}

// nolint:gocyclo
func (e *ClusterE2ETest) validateBottlerocketKubernetesSettings(ctx context.Context, IP string) {
	_ = "STUB: not implemented"
	return
}
