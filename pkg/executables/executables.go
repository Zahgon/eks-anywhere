package executables

import (
	"bytes"
	"context"

	"github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
)

const (
	redactMask = "*****"
)

var redactedEnvKeys = []string{
	constants.VSphereUsernameKey,
	constants.VSpherePasswordKey,
	constants.GovcUsernameKey,
	constants.GovcPasswordKey,
	decoder.CloudStackCloudConfigB64SecretKey,
	eksaGithubTokenEnv,
	githubTokenEnv,
	config.EksaAccessKeyIdEnv,
	config.EksaSecretAccessKeyEnv,
	config.EksaSessionTokenKeyEnv,
	config.AwsAccessKeyIdEnv,
	config.AwsSecretAccessKeyEnv,
	constants.SnowCredentialsKey,
	constants.SnowCertsKey,
	constants.NutanixUsernameKey,
	constants.NutanixPasswordKey,
	constants.RegistryUsername,
	constants.RegistryPassword,
}

type executable struct {
	cli string
}

type Executable interface {
	Execute(ctx context.Context, args ...string) (stdout bytes.Buffer, err error)
	ExecuteWithEnv(ctx context.Context, envs map[string]string, args ...string) (stdout bytes.Buffer, err error) // TODO: remove this from interface in favor of Command
	ExecuteWithStdin(ctx context.Context, in []byte, args ...string) (stdout bytes.Buffer, err error)            // TODO: remove this from interface in favor of Command
	Command(ctx context.Context, args ...string) *Command
	Run(cmd *Command) (stdout bytes.Buffer, err error)
}

// this should only be called through the executables.builder.
func NewExecutable(cli string) Executable { _ = "STUB: not implemented"; return *new(Executable) }

func (e *executable) Execute(ctx context.Context, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *executable) ExecuteWithStdin(ctx context.Context, in []byte, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *executable) ExecuteWithEnv(ctx context.Context, envs map[string]string, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *executable) Command(ctx context.Context, args ...string) *Command {
	_ = "STUB: not implemented"
	return nil
}

func (e *executable) Run(cmd *Command) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (e *executable) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func RedactCreds(cmd string, envMap map[string]string) string { _ = "STUB: not implemented"; return "" }

func execute(ctx context.Context, cli string, in []byte, envVars map[string]string, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}
