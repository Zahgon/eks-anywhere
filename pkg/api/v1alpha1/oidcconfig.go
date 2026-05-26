package v1alpha1

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
)

const OIDCConfigKind = "OIDCConfig"

func GetAndValidateOIDCConfig(fileName string, refName string, clusterConfig *Cluster) (*OIDCConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getOIDCConfig(fileName string) (*OIDCConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the name is empty, we can assume that they didn't configure their OIDC configuration, so return nil

func validateOIDCConfig(config *OIDCConfig) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateOIDCRefName(config *OIDCConfig, refName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOIDCNamespace(config *OIDCConfig, clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
