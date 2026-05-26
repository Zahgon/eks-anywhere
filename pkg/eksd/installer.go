package eksd

import (
	"context"
	"time"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	maxRetries    = 5
	backOffPeriod = 5 * time.Second
)

type EksdInstallerClient interface {
	ApplyKubeSpecFromBytesWithNamespace(ctx context.Context, cluster *types.Cluster, data []byte, namespace string) error
}

type Reader interface {
	ReadFile(url string) ([]byte, error)
}

// InstallerOpt allows to customize an eksd installer
// on construction.
type InstallerOpt func(*Installer)

type Installer struct {
	client  EksdInstallerClient
	retrier *retrier.Retrier
	reader  Reader
}

// NewEksdInstaller constructs a new eks-d installer.
func NewEksdInstaller(client EksdInstallerClient, reader Reader, opts ...InstallerOpt) *Installer {
	_ = "STUB: not implemented"
	return nil
}

// WithRetrier allows to use a custom retrier.
func WithRetrier(retrier *retrier.Retrier) InstallerOpt {
	_ = "STUB: not implemented"
	return *new(InstallerOpt)
}

func (i *Installer) InstallEksdCRDs(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// SetRetrier allows to modify the internal retrier
// For unit testing purposes only. It is not thread safe.
func (i *Installer) SetRetrier(retrier *retrier.Retrier) { _ = "STUB: not implemented"; return }

func (i *Installer) InstallEksdManifest(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
