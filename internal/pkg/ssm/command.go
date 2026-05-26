package ssm

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/go-logr/logr"
)

const (
	ssmLogGroup               = "/eks-anywhere/test/e2e"
	defaultSSMDeliveryTimeout = 300
)

var initE2EDirCommand = "mkdir -p /home/e2e/bin && cd /home/e2e"

// WaitForSSMReady waits for the SSM command to be ready.
func WaitForSSMReady(session *session.Session, instanceID string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

type CommandOpt func(c *ssm.SendCommandInput)

func WithOutputToS3(bucket, dir string) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

func WithOutputToCloudwatch() CommandOpt { _ = "STUB: not implemented"; return *new(CommandOpt) }

var nonFinalStatuses = map[string]struct{}{
	ssm.CommandInvocationStatusInProgress: {}, ssm.CommandInvocationStatusDelayed: {}, ssm.CommandInvocationStatusPending: {},
}

// Run runs the command using SSM on the instance corresponding to the instanceID.
func Run(session *session.Session, logger logr.Logger, instanceID, command string, timeout time.Duration, opts ...CommandOpt) error {
	_ = "STUB: not implemented"
	return nil
}

// RunCommand runs the command using SSM on the instance corresponding to the instanceID.
func RunCommand(session *session.Session, logger logr.Logger, instanceID, command string, timeout time.Duration, opts ...CommandOpt) (*RunOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make sure ssm send command is registered

// Making the retrier wait for longer than the provided SSM timeout to make sure
// we always get the output results.

func sendCommand(service *ssm.SSM, logger logr.Logger, instanceID, command string, timeout time.Duration, opts ...CommandOpt) (*ssm.SendCommandOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isFinalStatus(status string) bool { _ = "STUB: not implemented"; return false }
