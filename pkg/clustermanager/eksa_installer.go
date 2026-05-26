package clustermanager

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// EKSAInstallerOpt updates an EKSAInstaller.
type EKSAInstallerOpt func(*EKSAInstaller)

// EKSAInstaller allows to install eks-a components in a cluster.
type EKSAInstaller struct {
	client                KubernetesClient
	reader                manifests.FileReader
	deploymentWaitTimeout time.Duration
}

// NewEKSAInstaller constructs a new EKSAInstaller.
func NewEKSAInstaller(client KubernetesClient, reader manifests.FileReader, opts ...EKSAInstallerOpt) *EKSAInstaller {
	_ = "STUB: not implemented"
	return nil
}

// WithEKSAInstallerNoTimeouts disables the timeout when waiting for a deployment to be ready.
func WithEKSAInstallerNoTimeouts() EKSAInstallerOpt {
	_ = "STUB: not implemented"
	return *new(EKSAInstallerOpt)
}

// Install configures and applies eks-a components in a cluster accordingly to a spec.
func (i *EKSAInstaller) Install(ctx context.Context, log logr.Logger, cluster *types.Cluster, managementComponents *cluster.ManagementComponents, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to update this config map with the new upgrader images whenever we
// apply a new Bundles object to the cluster in order to support in-place upgrades.

func (i *EKSAInstaller) getUpgraderImagesFromBundle(ctx context.Context, cluster *types.Cluster, cl *cluster.Spec) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Upgrade re-installs the eksa components in a cluster if the VersionBundle defined in the
// new spec has a different eks-a components version. Workload clusters are ignored.
func (i *EKSAInstaller) Upgrade(ctx context.Context, log logr.Logger, c *types.Cluster, currentManagementComponents, newManagementComponents *cluster.ManagementComponents, newSpec *cluster.Spec) (*types.ChangeDiff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createEKSAComponents creates eksa components and applies the objects to the cluster.
func (i *EKSAInstaller) createEKSAComponents(ctx context.Context, log logr.Logger, cluster *types.Cluster, managementComponents *cluster.ManagementComponents, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// applyBundles applies the bundles to the cluster.
func (i *EKSAInstaller) applyBundles(ctx context.Context, log logr.Logger, cluster *types.Cluster, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// applyReleases applies the releases to the cluster.
func (i *EKSAInstaller) applyReleases(ctx context.Context, log logr.Logger, cluster *types.Cluster, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// EKSAComponentGenerator generates and configures eks-a components.
type EKSAComponentGenerator struct {
	log    logr.Logger
	reader manifests.FileReader
}

func (g *EKSAComponentGenerator) buildEKSAComponentsSpec(managamentComponents *cluster.ManagementComponents, spec *cluster.Spec) (*eksaComponents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *EKSAComponentGenerator) configureEKSAComponents(c *eksaComponents, spec *cluster.Spec) {
	_ = "STUB: not implemented"
	// TODO(g-gaston): we should do this with a custom ControllerManagerConfig.
	// This requires wider changes in the controller manager setup and config manifest,
	// so leaving this for later.
	return
}

func setManagerFlags(d *appsv1.Deployment, spec *cluster.Spec) { _ = "STUB: not implemented"; return }

func setManagerEnvVars(d *appsv1.Deployment, spec *cluster.Spec) { _ = "STUB: not implemented"; return }

// TODO: remove this feature flag if we decide to support in-place upgrades for vSphere provider.

// TODO: remove this feature flag when we support API server flags.

func managerEnabledGates(spec *cluster.Spec) []string { _ = "STUB: not implemented"; return nil }

func (g *EKSAComponentGenerator) parseEKSAComponentsSpec(managementComponents *cluster.ManagementComponents) (*eksaComponents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type eksaComponents struct {
	deployment *appsv1.Deployment
	rest       []*unstructured.Unstructured
}

func (c *eksaComponents) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

// EksaChangeDiff computes the version diff in eksa components between two specs.
func EksaChangeDiff(currentManagementComponents, newManagementComponents *cluster.ManagementComponents) *types.ChangeDiff {
	_ = "STUB: not implemented"
	return nil
}
