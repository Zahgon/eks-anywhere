package executables

import (
	"context"
)

// SSH is an executable for running SSH commands.
type SSH struct {
	Executable
}

const (
	sshPath             = "ssh"
	strictHostCheckFlag = "StrictHostKeyChecking=no"
)

// NewSSH returns a new instance of SSH client.
func NewSSH(executable Executable) *SSH { _ = "STUB: not implemented"; return nil }

// RunCommand runs a command on the host using SSH.
func (s *SSH) RunCommand(ctx context.Context, privateKeyPath, username, IP string, command ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
