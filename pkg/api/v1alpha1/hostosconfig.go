package v1alpha1

import (
	"sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
)

func validateHostOSConfig(config *HostOSConfiguration, osFamily OSFamily) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNTPServers(config *NTPConfiguration) error { _ = "STUB: not implemented"; return nil }

// ParseRequestURI expects a scheme but ntp servers generally don't have one
// Prepending a scheme here so it doesn't fail because of missing scheme

func addNTPScheme(server string) string { _ = "STUB: not implemented"; return "" }

func validateCertBundles(config *certBundle, osFamily OSFamily) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBotterocketConfig(config *BottlerocketConfiguration, osFamily OSFamily) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBottlerocketKubernetesConfig(config *v1beta2.BottlerocketKubernetesSettings) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBottlerocketKernelConfiguration(config *v1beta2.BottlerocketKernelSettings) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBottlerocketBootSettingsConfiguration(config *v1beta2.BottlerocketBootSettings) error {
	_ = "STUB: not implemented"
	return nil
}

// validateTrustedCertBundle validates that the cert is valid.
func validateTrustedCertBundle(certBundle string) error { _ = "STUB: not implemented"; return nil }

// cert bundles could contain more than one certificate

// no more PEM structed objects
