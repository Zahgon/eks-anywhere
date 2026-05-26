package test

import (
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// RegistryMirrorInsecureSkipVerifyEnabled returns a test RegistryMirrorConfiguration with InsecureSkipVerify enabled.
func RegistryMirrorInsecureSkipVerifyEnabled() *anywherev1.RegistryMirrorConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// RegistryMirrorInsecureSkipVerifyEnabledAndCACert returns a test RegistryMirrorConfiguration with a CACert specified and InsecureSkipVerify enabled.
func RegistryMirrorInsecureSkipVerifyEnabledAndCACert() *anywherev1.RegistryMirrorConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// CACertContent returns a test string representing a cacert contents.
func CACertContent() string { _ = "STUB: not implemented"; return "" }

// RegistryMirrorConfigFilesInsecureSkipVerify returns cluster-api bootstrap files that configure containerd
// to use a registry mirror with the insecure_skip_verify flag enabled.
func RegistryMirrorConfigFilesInsecureSkipVerify() []bootstrapv1beta2.File {
	_ = "STUB: not implemented"
	return nil
}

// RegistryMirrorConfigFilesInsecureSkipVerifyAndCACert returns cluster-api bootstrap files that configure containerd
// to use a registry mirror with a cacert file and insecure_skip_verify flag enabled.
func RegistryMirrorConfigFilesInsecureSkipVerifyAndCACert() []bootstrapv1beta2.File {
	_ = "STUB: not implemented"
	return nil
}

// RegistryMirrorPreKubeadmCommands returns a list of commands to writes a config_append.toml file
// to configure the registry mirror and restart containerd.
func RegistryMirrorPreKubeadmCommands() []string { _ = "STUB: not implemented"; return nil }

// RegistryMirrorSudoPreKubeadmCommands returns a list of commands that writes a config_append.toml file
// to configure the registry mirror and restart containerd with sudo permissions.
func RegistryMirrorSudoPreKubeadmCommands() []string { _ = "STUB: not implemented"; return nil }
