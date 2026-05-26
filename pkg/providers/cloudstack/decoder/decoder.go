package decoder

import (
	"gopkg.in/ini.v1"
	apiv1 "k8s.io/api/core/v1"
)

const (
	EksacloudStackCloudConfigB64SecretKey = "EKSA_CLOUDSTACK_B64ENCODED_SECRET"
	CloudStackCloudConfigB64SecretKey     = "CLOUDSTACK_B64ENCODED_SECRET"
	EksaCloudStackHostPathToMount         = "EKSA_CLOUDSTACK_HOST_PATHS_TO_MOUNT"
	defaultVerifySslValue                 = "true"
	CloudStackGlobalAZ                    = "global"

	APIKeyKey    = "api-key"
	SecretKeyKey = "secret-key"
	APIUrlKey    = "api-url"
	VerifySslKey = "verify-ssl"
)

// ParseCloudStackCredsFromSecrets parses a list of secrets to extract out the api keys, secret keys, and urls.
func ParseCloudStackCredsFromSecrets(secrets []apiv1.Secret) (*CloudStackExecConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseCloudStackCredsFromEnv parses the input b64 string into the ini object to extract out the api key, secret key, and url.
func ParseCloudStackCredsFromEnv() (*CloudStackExecConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseCloudStackProfileSection(section *ini.Section) (*CloudStackProfileConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CloudStackExecConfig struct {
	Profiles []CloudStackProfileConfig
}

type CloudStackProfileConfig struct {
	Name          string
	ApiKey        string
	SecretKey     string
	ManagementUrl string
	VerifySsl     string
	Timeout       string
}
