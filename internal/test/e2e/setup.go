package e2e

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/networkutils"
)

var requiredFiles = []string{cliBinary, e2eBinary}

const (
	cliBinary                  = "eksctl-anywhere"
	e2eBinary                  = "e2e.test"
	bundlesReleaseManifestFile = "local-bundle-release.yaml"
	eksAComponentsManifestFile = "local-eksa-components.yaml"
	testNameFile               = "e2e-test-name"
	maxUserWatches             = 524288
	maxUserInstances           = 512
	key                        = "Integration-Test"
	tag                        = "EKSA-E2E"
)

type E2ESession struct {
	session             *session.Session
	instanceProfileName string
	storageBucket       string
	jobId               string
	instanceId          string
	ipPool              networkutils.IPPool
	testEnvVars         map[string]string
	bundlesOverride     bool
	cleanup             bool
	requiredFiles       []string
	branchName          string
	hardware            []*api.Hardware
	logger              logr.Logger
	stage               string
}

func newE2ESession(instanceId string, conf instanceRunConf) (*E2ESession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *E2ESession) setup(regex string) error { _ = "STUB: not implemented"; return nil }

// Adding JobId to Test Env variables

func (e *E2ESession) updateFSInotifyResources() error { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) uploadRequiredFile(file string) error { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) uploadRequiredFiles() error { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) downloadRequiredFileInInstance(file string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *E2ESession) downloadRequiredFilesInInstance() error { _ = "STUB: not implemented"; return nil }

func (e *E2ESession) createTestNameFile(testName string) error {
	_ = "STUB: not implemented"
	return nil
}

func clusterPrefix(branch, instanceId string) (clusterPrefix string) {
	_ = "STUB: not implemented"
	return ""
}

func (e *E2ESession) clusterName(branch, instanceId, testName string) (clusterName string) {
	_ = "STUB: not implemented"
	return ""
}
