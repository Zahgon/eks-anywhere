package executables

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/helm"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
)

const defaultEksaImage = "public.ecr.aws/l0g8r8j6/eks-anywhere-cli-tools:v0.18.4-eks-a-v0.0.0-dev-build.8100"

type ExecutableBuilder interface {
	Init(ctx context.Context) (Closer, error)
	Build(binaryPath string) Executable
}

type ExecutablesBuilder struct {
	executableBuilder ExecutableBuilder
}

func NewExecutablesBuilder(executableBuilder ExecutableBuilder) *ExecutablesBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ExecutablesBuilder) BuildKindExecutable(writer filewriter.FileWriter) *Kind {
	_ = "STUB: not implemented"
	return nil
}

func (b *ExecutablesBuilder) BuildClusterAwsAdmExecutable() *Clusterawsadm {
	_ = "STUB: not implemented"
	return nil
}

// BuildClusterCtlExecutable builds a new Clusterctl executable.
func (b *ExecutablesBuilder) BuildClusterCtlExecutable(writer filewriter.FileWriter, reader manifests.FileReader) *Clusterctl {
	_ = "STUB: not implemented"
	return nil
}

func (b *ExecutablesBuilder) BuildKubectlExecutable() *Kubectl {
	_ = "STUB: not implemented"
	return nil
}

func (b *ExecutablesBuilder) BuildGovcExecutable(writer filewriter.FileWriter, opts ...GovcOpt) *Govc {
	_ = "STUB: not implemented"
	return nil
}

// BuildCmkExecutable initializes a Cmk object and returns it.
func (b *ExecutablesBuilder) BuildCmkExecutable(writer filewriter.FileWriter, config *decoder.CloudStackExecConfig) (*Cmk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *ExecutablesBuilder) BuildAwsCli() *AwsCli { _ = "STUB: not implemented"; return nil }

func (b *ExecutablesBuilder) BuildFluxExecutable() *Flux { _ = "STUB: not implemented"; return nil }

func (b *ExecutablesBuilder) BuildTroubleshootExecutable() *Troubleshoot {
	_ = "STUB: not implemented"
	return nil
}

// BuildHelmExecutable initializes a helm executable and returns it.
func (b *ExecutablesBuilder) BuildHelmExecutable(opts ...helm.Opt) *Helm {
	_ = "STUB: not implemented"
	return nil
}

// BuildHelm initializes a helm executable and returns it.
func (b *ExecutablesBuilder) BuildHelm(opts ...helm.Opt) helm.Client {
	_ = "STUB: not implemented"
	return *new(helm.Client)
}

// BuildDockerExecutable initializes a docker executable and returns it.
func (b *ExecutablesBuilder) BuildDockerExecutable() *Docker { _ = "STUB: not implemented"; return nil }

// BuildSSHExecutable initializes a SSH executable and returns it.
func (b *ExecutablesBuilder) BuildSSHExecutable() *SSH { _ = "STUB: not implemented"; return nil }

// Init initializes the executable builder and returns a Closer
// that needs to be called once the executables are not in used anymore
// The closer will cleanup and free all internal resources.
func (b *ExecutablesBuilder) Init(ctx context.Context) (Closer, error) {
	_ = "STUB: not implemented"
	return *new(Closer), nil
}

func BuildSonobuoyExecutable() *Sonobuoy { _ = "STUB: not implemented"; return nil }

func BuildDockerExecutable() *Docker { _ = "STUB: not implemented"; return nil }

// RunExecutablesInDocker determines if binary executables should be ran
// from a docker container or native binaries from the host path
// It reads MR_TOOLS_DISABLE variable.
func ExecutablesInDocker() bool { _ = "STUB: not implemented"; return false }

// InitInDockerExecutablesBuilder builds and inits a default ExecutablesBuilder to run executables in a docker container
// that will make use of a long running docker container.
func InitInDockerExecutablesBuilder(ctx context.Context, image string, mountDirs ...string) (*ExecutablesBuilder, Closer, error) {
	_ = "STUB: not implemented"
	return nil, *new(Closer), nil
}

// NewInDockerExecutablesBuilder builds an executables builder for docker.
func NewInDockerExecutablesBuilder(dockerClient DockerClient, image string, mountDirs ...string) (*ExecutablesBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLocalExecutablesBuilder() *ExecutablesBuilder { _ = "STUB: not implemented"; return nil }

func DefaultEksaImage() string { _ = "STUB: not implemented"; return "" }

type Closer func(ctx context.Context) error

// Close implements interface types.Closer.
func (c Closer) Close(ctx context.Context) error {
	_ = "STUB: not implemented"

	// CheckErr just calls the closer and logs an error if present
	// It's mostly a helper for defering the close in a oneliner without ignoring the error.
	return nil
}

func (c Closer) CheckErr(ctx context.Context) { _ = "STUB: not implemented"; return }
