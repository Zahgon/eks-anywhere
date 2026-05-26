package framework

import (
	"os/exec"

	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type CommandOpt func(*string, *[]string) (err error)

func appendOpt(new ...string) CommandOpt { _ = "STUB: not implemented"; return *new(CommandOpt) }

func withKubeconfig(kubeconfigFile string) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

func WithControlPlaneWaitTimeout(timeout string) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

func WithExternalEtcdWaitTimeout(timeout string) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

func WithPerMachineWaitTimeout(timeout string) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

func ExecuteWithEksaRelease(release *releasev1alpha1.EksARelease) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

// PackagedBinary represents a binary that can be extracted
// executed from local disk.
type PackagedBinary interface {
	// BinaryPath returns the local disk path to the binary.
	BinaryPath() (string, error)
}

// ExecuteWithBinary executes the command with a binary from an specific path.
func ExecuteWithBinary(eksa PackagedBinary) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

// WithSudo add prefix "sudo" to the command. And preserve PATH.
func WithSudo(user string) CommandOpt { _ = "STUB: not implemented"; return *new(CommandOpt) }

// WithBundlesOverride modify bundles-override.
func WithBundlesOverride(bundles string) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

type binaryFetcher func() (binaryPath string, err error)

func executeWithBinaryCommandOpt(fetcher binaryFetcher) CommandOpt {
	_ = "STUB: not implemented"
	return *new(CommandOpt)
}

// When bundles override is present, the manifest belongs to the current
// build of the CLI and it's intended to be used only with that version

func removeFlag(flag string, args *[]string) { _ = "STUB: not implemented"; return }

// If it's not the last arg and next arg is not a flag,
// that means it's the value for the current flag, remove it as well

// DefaultLocalEKSABinaryPath returns the full path for the local eks-a binary being tested.
func DefaultLocalEKSABinaryPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// DefaultLocalEKSABinDir returns the full path for the local directory where
// the tested eks-a binary lives.
func DefaultLocalEKSABinDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func prepareCommand(name string, args ...string) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
