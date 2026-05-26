package eksd

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/types"
	releavev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type Upgrader struct {
	*Installer
}

// UpgraderOpt allows to customize an eksd upgrader on construction.
type UpgraderOpt = InstallerOpt

// NewUpgrader constructs a new eks-d upgrader.
func NewUpgrader(client EksdInstallerClient, reader Reader, opts ...UpgraderOpt) *Upgrader {
	_ = "STUB: not implemented"
	return nil
}

// Upgrade checks for EKS-D updates, and if there are updates the EKS-D CRDs in the cluster.
func (u *Upgrader) Upgrade(ctx context.Context, cluster *types.Cluster, currentSpec, newSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ChangeDiff returns the change diff between the current and new EKS-D versions.
func ChangeDiff(currentSpec, newSpec *cluster.Spec) *types.ChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func eksdKubernetesVersionTag(eksd releavev1alpha1.EksDRelease) string {
	_ = "STUB: not implemented"
	return ""
}
