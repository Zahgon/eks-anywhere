package reconciler

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/networking/cilium"
)

const (
	defaultRequeueTime = time.Second * 10
)

var (
	serviceKind    = corev1.SchemeGroupVersion.WithKind("Service").GroupKind()
	daemonSetKind  = appsv1.SchemeGroupVersion.WithKind("DaemonSet").GroupKind()
	deploymentKind = appsv1.SchemeGroupVersion.WithKind("Deployment").GroupKind()
)

type Templater interface {
	GenerateUpgradePreflightManifest(ctx context.Context, spec *cluster.Spec) ([]byte, error)
	GenerateManifest(ctx context.Context, spec *cluster.Spec, opts ...cilium.ManifestOpt) ([]byte, error)
}

// Reconciler allows to reconcile a Cilium CNI.
type Reconciler struct {
	templater          Templater
	providerNamespaces []string
}

// New creates a new cilium reconciler object with a templater and providerNamespaces to generate manifests.
func New(templater Templater, providerNamespaces []string) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile takes the Cilium CNI in a cluster to the desired state defined in a cluster Spec.
// It uses a controller.Result to indicate when requeues are needed. client is connected to the
// target Kubernetes cluster, not the management cluster.
// nolint:gocyclo
// TODO: reduce cyclomatic complexity - https://github.com/aws/eks-anywhere-internal/issues/1461
func (r *Reconciler) Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *cluster.Spec) (res controller.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// We use a marker to detect if EKS-A Cilium has ever been installed. If it has never been
// installed and isn't currently installed we always attempt to install it regardless of whether
// the user is skipping EKS-A Cilium management. This satsifies criteria for successful cluster
// creation.
//
// If EKS-A Cilium was previously installed, as denoted by the marker, we only want to
// manage it if its still installed and the user still wants us to manage the installation (as
// denoted by the API skip flag).
//
// In the event a user uninstalls EKS-A Cilium, updates the cluster spec to skip EKS-A Cilium
// management, then tries to upgrade, we will attempt to install EKS-A Cilium. This is because
// reconciliation has no operational context (create vs upgrade) and can only observe that no
// installation is present and there is no marker indicating it was ever present which is
// equivilent to a typical create scenario where we must install a CNI to satisfy cluster
// create success criteria.

// To accommodate upgrades of cluster created prior to introducing markers, we check for
// an existing installation and try to mark the cluster as having already had EKS-A
// Cilium installed.

// Upgrade process has run its course, and so we can now mark that the default cni has been configured.

func (r *Reconciler) install(ctx context.Context, log logr.Logger, client client.Client, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) upgrade(ctx context.Context, logger logr.Logger, client client.Client, installation *cilium.Installation, spec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// When upgrading from Cilium v1.11.x --> Cilium v1.12.x, the port number for a few components changed but the port names remained the same.
// This caused "duplicate value" errors when upgrading to the new Cilium version using server-side apply because the merge key for these fields is the port number, not the name.
// To alleviate this issue, we will use the client.Update strategy to update the yaml to be compatible with the new version of Cilium.
// We are only doing this for the Cilium Service, DaemonSet, and Deployment since those are the only objects affected.
// The rest of the objects in the Cilium upgrade manifest will continue to be applied using server-side apply.

func (r *Reconciler) updateConfig(ctx context.Context, client client.Client, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) applyFullManifest(ctx context.Context, client client.Client, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) deletePreflightIfExists(ctx context.Context, client client.Client, spec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (r *Reconciler) installPreflight(ctx context.Context, client client.Client, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) reconcileSpecialCases(ctx context.Context, c client.Client, yaml []byte) ([]client.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
