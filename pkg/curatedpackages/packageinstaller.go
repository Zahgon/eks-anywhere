package curatedpackages

import (
	"context"
	"time"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

type PackageController interface {
	// Enable curated packages support.
	Enable(ctx context.Context) error
	IsInstalled(ctx context.Context) bool
}

type PackageHandler interface {
	CreatePackages(ctx context.Context, fileName string, kubeConfig string) error
}

type Installer struct {
	packageController PackageController
	spec              *cluster.Spec
	packageClient     PackageHandler
	kubectl           KubectlRunner
	packagesLocation  string
	mgmtKubeconfig    string
	installMaxRetries int
	installBackoff    time.Duration
}

// IsPackageControllerDisabled detect if the package controller is disabled.
func IsPackageControllerDisabled(cluster *anywherev1.Cluster) bool {
	_ = "STUB: not implemented"
	return false
}

// NewInstaller installs packageController and packages during cluster creation.
func NewInstaller(runner KubectlRunner, pc PackageHandler, pcc PackageController, spec *cluster.Spec, packagesLocation, mgmtKubeconfig string) *Installer {
	_ = "STUB: not implemented"
	return nil
}

// WithRetries sets the retry parameters for the package controller installation.
func (pi *Installer) WithRetries(maxRetries int, backoff time.Duration) *Installer {
	_ = "STUB: not implemented"
	return nil
}

// InstallCuratedPackages installs curated packages as part of the cluster creation.
func (pi *Installer) InstallCuratedPackages(ctx context.Context) { _ = "STUB: not implemented"; return }

// There is an ask from customers to avoid considering the failure of installing curated packages
// controller as an error but rather a warning

// There is an ask from customers to avoid considering the failure of the installation of curated packages
// as an error but rather a warning

// UpgradeCuratedPackages upgrades curated packages as part of the cluster upgrade.
func (pi *Installer) UpgradeCuratedPackages(ctx context.Context) { _ = "STUB: not implemented"; return }

func (pi *Installer) installPackagesController(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (pi *Installer) installPackages(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
