package framework

import (
	"context"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// WithSkipCiliumUpgrade returns an E2E test option that configures the Cluster object to
// skip Cilium upgrades.
func WithSkipCiliumUpgrade() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// UninstallCilium uninstalls the workload clusters Cilium.
func (e *ClusterE2ETest) UninstallCilium() { _ = "STUB: not implemented"; return }

// ValidateCiliumCLIAvailable ensures the Cilium CLI can be found on the PATH.
func (e *ClusterE2ETest) ValidateCiliumCLIAvailable() { _ = "STUB: not implemented"; return }

// InstallOSSCilium installs an open source version of Cilium. The version is dependent on the
// Cilium CLI version available on the PATH.
func (e *ClusterE2ETest) InstallOSSCilium() { _ = "STUB: not implemented"; return }

// ReplaceCiliumWithOSSCilium replaces the current Cilium installation in the workload cluster
// with an open source version. See InstallOSSCilium().
func (e *ClusterE2ETest) ReplaceCiliumWithOSSCilium() { _ = "STUB: not implemented"; return }

// ValidateEKSACiliumNotInstalled inspects the workload cluster for an EKSA Cilium installation
// erroring if one is found.
func (e *ClusterE2ETest) ValidateEKSACiliumNotInstalled() { _ = "STUB: not implemented"; return }

// ValidateEKSACiliumInstalled inspects the workload cluster for an EKSA Cilium installation
// erroring if one is not found.
func (e *ClusterE2ETest) ValidateEKSACiliumInstalled() { _ = "STUB: not implemented"; return }

// AwaitCiliumDaemonSetReady awaits the Cilium daemonset to be ready in the cluster represented by client.
// It is ready when the DaemonSet's .Status.NumberUnavailable is 0.
func AwaitCiliumDaemonSetReady(ctx context.Context, client client.Client, retries int, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
