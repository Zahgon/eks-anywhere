package aws

import (
	"io"
)

const (
	EksaAwsCredentialsFileKey = "EKSA_AWS_CREDENTIALS_FILE"
	EksaAwsCABundlesFileKey   = "EKSA_AWS_CA_BUNDLES_FILE"
)

func AwsCredentialsFile() (filePath string, err error) { _ = "STUB: not implemented"; return "", nil }

func AwsCABundlesFile() (filePath string, err error) { _ = "STUB: not implemented"; return "", nil }

func validateFileFromEnv(envKey string) (filePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func EncodeFileFromEnv(envKey string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseDeviceIPsFromFile(filePath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseDeviceIPs(r io.Reader) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func subtractProfileName(input string) string { _ = "STUB: not implemented"; return "" }
