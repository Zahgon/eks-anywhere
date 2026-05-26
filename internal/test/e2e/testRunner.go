package e2e

import (
	"github.com/go-logr/logr"
)

const (
	testRunnerVCUserEnvVar     string = "TEST_RUNNER_GOVC_USERNAME"
	testRunnerVCPasswordEnvVar string = "TEST_RUNNER_GOVC_PASSWORD"
	govcUsernameKey            string = "GOVC_USERNAME"
	govcPasswordKey            string = "GOVC_PASSWORD"
	govcURLKey                 string = "GOVC_URL"
	govcInsecure               string = "GOVC_INSECURE"
	govcDatacenterKey          string = "GOVC_DATACENTER"
	ssmActivationCodeKey       string = "ssm_activation_code"
	ssmActivationIdKey         string = "ssm_activation_id"
	ssmActivationRegionKey     string = "ssm_activation_region"
)

type TestRunner interface {
	createInstance(instanceConf instanceRunConf) (string, error)
	tagInstance(instanceConf instanceRunConf, key, value string) error
	decommInstance(instanceRunConf) error
}

type TestRunnerType string

const (
	Ec2TestRunnerType     TestRunnerType = "ec2"
	VSphereTestRunnerType TestRunnerType = "vSphere"
)

func newTestRunner(runnerType TestRunnerType, config TestInfraConfig) (TestRunner, error) {
	_ = "STUB: not implemented"
	return *new(TestRunner), nil
}

type TestInfraConfig struct {
	Ec2TestRunner     `yaml:"ec2,omitempty"`
	VSphereTestRunner `yaml:"vSphere,omitempty"`
}

func NewTestRunnerConfigFromFile(logger logr.Logger, configFile string) (*TestInfraConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type testRunner struct {
	InstanceID string
	logger     logr.Logger
}

type Ec2TestRunner struct {
	testRunner
	AmiID    string `yaml:"amiId"`
	SubnetID string `yaml:"subnetId"`
}

type VSphereTestRunner struct {
	testRunner
	ActivationId string
	envMap       map[string]string
	Url          string `yaml:"url"`
	Insecure     bool   `yaml:"insecure"`
	Library      string `yaml:"library"`
	Template     string `yaml:"template"`
	Datacenter   string `yaml:"datacenter"`
	Datastore    string `yaml:"datastore"`
	ResourcePool string `yaml:"resourcePool"`
	Network      string `yaml:"network"`
	Folder       string `yaml:"folder"`
}

func (v *VSphereTestRunner) setEnvironment() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VSphereTestRunner) createInstance(c instanceRunConf) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: import ova template from url if not exist

// deploy template

func (e *Ec2TestRunner) createInstance(c instanceRunConf) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v *VSphereTestRunner) tagInstance(c instanceRunConf, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Ec2TestRunner) tagInstance(c instanceRunConf, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *VSphereTestRunner) decommInstance(c instanceRunConf) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Ec2TestRunner) decommInstance(c instanceRunConf) error {
	_ = "STUB: not implemented"
	return nil
}

func getTestRunnerName(logger logr.Logger, jobId string) string {
	_ = "STUB: not implemented"
	return ""
}
