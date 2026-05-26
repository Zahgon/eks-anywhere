package e2e

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/internal/pkg/ssm"
	"github.com/aws/eks-anywhere/pkg/networkutils"
)

const (
	testResultPass       = "pass"
	testResultFail       = "fail"
	testResultError      = "error"
	nonAirgappedHardware = "nonAirgappedHardware"
	airgappedHardware    = "AirgappedHardware"
	tinkerbellIPPoolSize = 2

	// Default timeout for E2E test instance.
	e2eTimeout           = 150 * time.Minute
	e2eSSMTimeoutPadding = 10 * time.Minute

	// Default timeout used for all SSM commands besides running the actual E2E test.
	ssmTimeout = 10 * time.Minute
)

type ParallelRunConf struct {
	TestInstanceConfigFile string
	MaxConcurrentTests     int
	InstanceProfileName    string
	StorageBucket          string
	JobId                  string
	Regex                  string
	TestsToSkip            []string
	BundlesOverride        bool
	CleanupResources       bool
	TestReportFolder       string
	BranchName             string
	Logger                 logr.Logger
	Stage                  string
}

type (
	testCommandResult    = ssm.RunOutput
	instanceTestsResults struct {
		conf              instanceRunConf
		testCommandResult *testCommandResult
		err               error
	}
)

// RunTestsInParallel Run Tests in parallel by spawning multiple admin machines.
func RunTestsInParallel(conf ParallelRunConf) error { _ = "STUB: not implemented"; return nil }

// For Tinkerbell tests, get hardware inventory pool

// This variable can be used in cloudwatch log insights query for e2e test success rate

// TODO: keeping the old logs temporarily for compatibility with the test tool
// Once the tool is updated to support the unified message, remove them

type instanceRunConf struct {
	InstanceProfileName     string
	StorageBucket           string
	JobID                   string
	ParentJobID             string
	Regex                   string
	InstanceID              string
	TestReportFolder        string
	BranchName              string
	IPPool                  networkutils.IPPool
	Hardware                []*api.Hardware
	HardwareCount           int
	TinkerbellAirgappedTest bool
	BundlesOverride         bool
	TestRunnerType          TestRunnerType
	TestRunnerConfig        TestInfraConfig
	CleanupResources        bool
	Logger                  logr.Logger
	Session                 *session.Session
	Stage                   string
}

//nolint:gocyclo, revive // RunTests responsible launching test runner to run tests is complex.
func RunTests(conf instanceRunConf, inventoryCatalogue map[string]*hardwareCatalogue) (testInstanceID string, testCommandResult *testCommandResult, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// shuffle hardware to introduce randomness during hardware reservation.
// we do not want quick e2e runs to always pick the first few available hardware from the list and over-populate the boot entries
// this will quickly break the booting process as the hardware runs out of boot space to store these entries.
// randomly picking the hardware will distribute the boot entries across these hardware during each run
// ideally for long term we want a clear cleanup of the boot entries in the hardware

// Release hardware back to inventory for Tinkerbell Tests

// Tagging only successful e2e test instances.
// The aws cleanup periodic job deletes the tagged EC2 instances and long lived instances.

func (e *E2ESession) runTests(regex string) (testCommandResult *testCommandResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c instanceRunConf) runPostTestsProcessing(e *E2ESession, testCommandResult *testCommandResult) error {
	_ = "STUB: not implemented"
	return nil
}

// For Tinkerbell tests we run multiple tests on the same instance.
// Hence upload fails for passed tests within the instance.
// TODO (pokearu): Find a way to only upload for failed tests within the instance.

func (e *E2ESession) commandWithEnvVars(command string) string {
	_ = "STUB: not implemented"
	return ""
}

func splitTests(testsList []string, conf ParallelRunConf) ([]instanceRunConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load IP pool size requirements from YAML config

//nolint:gocyclo // This legacy function is complex but the team too busy to simplify it
func appendNonAirgappedTinkerbellRunConfs(awsSession *session.Session, testsList []string, conf ParallelRunConf, testRunnerConfig *TestInfraConfig, runConfs []instanceRunConf, ipManager *E2EIPManager) ([]instanceRunConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pop from both ends to run a longer count tests and shorter count tests together
// to efficiently use the available hardware.

func appendAirgappedTinkerbellRunConfs(awsSession *session.Session, testsList []string, conf ParallelRunConf, testRunnerConfig *TestInfraConfig, runConfs []instanceRunConf, ipManager *E2EIPManager) ([]instanceRunConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTinkerbellTestsWithCount(tinkerbellTests []string, conf ParallelRunConf) ([]TinkerbellTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort tests by Hardware count, to enable running larger tests first for Tinkerbell Provider

func newInstanceRunConf(awsSession *session.Session, conf ParallelRunConf, jobNumber int, testRegex string, ipPool networkutils.IPPool, hardware []*api.Hardware, hardwareCount int, tinkerbellAirgappedTest bool, testRunnerType TestRunnerType, testRunnerConfig *TestInfraConfig) instanceRunConf {
	_ = "STUB: not implemented"
	return *new(instanceRunConf)
}

func logTestGroups(logger logr.Logger, instancesConf []instanceRunConf) {
	_ = "STUB: not implemented"
	return
}

func getNonAirgappedHardwarePool(storageBucket string) ([]*api.Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Airgapped tinkerbell tests have special hardware requirements that doesn't have internet connectivity.
func getAirgappedHardwarePool(storageBucket string) ([]*api.Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func reserveTinkerbellHardware(conf *instanceRunConf, invCatalogue *hardwareCatalogue) error {
	_ = "STUB: not implemented"
	return nil
}

func releaseTinkerbellHardware(conf *instanceRunConf, invCatalogue *hardwareCatalogue) {
	_ = "STUB: not implemented"
	return
}

func logTinkerbellTestHardwareInfo(conf *instanceRunConf, action string) {
	_ = "STUB: not implemented"
	return
}

func containsTinkerbellTest(testsList []string) bool { _ = "STUB: not implemented"; return false }
