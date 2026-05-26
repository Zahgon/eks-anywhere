package certificates

import (
	"context"

	"golang.org/x/crypto/ssh"
)

// sshClient interface and sshDialer type remain the same.
type sshClient interface {
	Close() error
	NewSession() (*ssh.Session, error)
}

// sshDialer is a function type for dialing SSH connections.
type sshDialer func(network, addr string, config *ssh.ClientConfig) (sshClient, error)

// SSHRunner provides methods for running commands over SSH.
type SSHRunner interface {
	// RunCommand runs a command on the remote host and returns the output
	RunCommand(ctx context.Context, node string, cmd string, opts ...SSHOption) (string, error)
}

// DefaultSSHRunner is the default implementation of SSHRunner.
type DefaultSSHRunner struct {
	sshConfig  *ssh.ClientConfig
	sshDialer  sshDialer
	sshKeyPath string
	sshPasswd  string
}

// SSHOption represents a configuration option for SSH operations.
type SSHOption func(*sshConfigOption)

type sshConfigOption struct {
	displayLogs bool
}

func defaultSSHConfig() *sshConfigOption { _ = "STUB: not implemented"; return nil }

// WithSSHLogging configures whether SSH command output should be displayed in logs.
func WithSSHLogging(display bool) SSHOption { _ = "STUB: not implemented"; return *new(SSHOption) }

// NewSSHRunner creates a new SSH runner with the given configuration.
func NewSSHRunner(cfg SSHConfig) (*DefaultSSHRunner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parsePrivateKey only get password from enviroment variables.
func (r *DefaultSSHRunner) parsePrivateKey(key []byte) (ssh.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ssh.Signer), nil
}

// RunCommand executes a command on the remote node via SSH and returns the output.
func (r *DefaultSSHRunner) RunCommand(ctx context.Context, node string, cmd string, opts ...SSHOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
