package e2e

import (
	"os"
	"strings"
	"sync"
)

const (
	snowCredentialsS3Path  = "T_SNOW_CREDENTIALS_S3_PATH"
	snowCertificatesS3Path = "T_SNOW_CERTIFICATES_S3_PATH"
	snowDevices            = "T_SNOW_DEVICES"
	snowCPCidr             = "T_SNOW_CONTROL_PLANE_CIDR"
	snowCPCidrs            = "T_SNOW_CONTROL_PLANE_CIDRS"
	snowCredsFile          = "EKSA_AWS_CREDENTIALS_FILE"
	snowCertsFile          = "EKSA_AWS_CA_BUNDLES_FILE"

	snowTestsRe       = `^.*Snow.*$`
	snowCredsFilename = "snow_creds"
	snowCertsFilename = "snow_certs"
)

var (
	snowCPCidrArray  []string
	snowCPCidrArrayM sync.Mutex
)

func init() {
	snowCPCidrArray = strings.Split(os.Getenv(snowCPCidrs), ",")
}

// Note that this function cannot be called more than the the number of cidrs in the list.
func getSnowCPCidr() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *E2ESession) setupSnowEnv(testRegex string) error { _ = "STUB: not implemented"; return nil }

func sendFileViaS3(e *E2ESession, s3Path string, filename string) error {
	_ = "STUB: not implemented"
	return nil
}
