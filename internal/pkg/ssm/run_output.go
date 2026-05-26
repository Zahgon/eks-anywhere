package ssm

import "github.com/aws/aws-sdk-go/service/ssm"

type RunOutput struct {
	commandOut     *ssm.GetCommandInvocationOutput
	CommandId      string
	StdOut, StdErr []byte
}

func buildRunOutput(commandOut *ssm.GetCommandInvocationOutput) *RunOutput {
	_ = "STUB: not implemented"
	return nil
}

func (r *RunOutput) Successful() bool { _ = "STUB: not implemented"; return false }

// StatusDetails returns the status details of the ssm command.
func (r *RunOutput) StatusDetails() string {
	_ = "STUB: not implemented"
	// handle nil pointer
	return ""
}
