package e2e

const (
	tinkerbellInventoryCsvFilePathEnvVar       = "T_TINKERBELL_INVENTORY_CSV"
	tinkerbellControlPlaneNetworkCidrEnvVar    = "T_TINKERBELL_CP_NETWORK_CIDR"
	tinkerbellHardwareS3FileKeyEnvVar          = "T_TINKERBELL_S3_INVENTORY_CSV_KEY"
	tinkerbellAirgappedHardwareS3FileKeyEnvVar = "T_TINKERBELL_S3_AG_INVENTORY_CSV_KEY"
	tinkerbellTestsRe                          = `^.*Tinkerbell.*$`
	e2eHardwareCsvFilePath                     = "e2e-inventory.csv"
	e2eAirgappedHardwareCsvFilePath            = "e2e-ag-inventory.csv"
	maxHardwarePerE2ETestEnvVar                = "T_TINKERBELL_MAX_HARDWARE_PER_TEST"
	tinkerbellDefaultMaxHardwarePerE2ETest     = 4
	tinkerbellBootstrapInterfaceEnvVar         = "T_TINKERBELL_BOOTSTRAP_INTERFACE"
	tinkerbellCIEnvironmentEnvVar              = "T_TINKERBELL_CI_ENVIRONMENT"
)

// TinkerbellTest maps each Tinkbell test with the hardware count needed for the test.
type TinkerbellTest struct {
	Name  string `yaml:"name"`
	Count int    `yaml:"count"`
}

func (e *E2ESession) setupTinkerbellEnv(testRegex string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) setTinkerbellBootstrapIPInInstance(tinkInterface string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get non airgapped, normal tinkerbell tests.
func getTinkerbellNonAirgappedTests(tests []string) []string { _ = "STUB: not implemented"; return nil }

func getTinkerbellAirgappedTests(tests []string) []string { _ = "STUB: not implemented"; return nil }
