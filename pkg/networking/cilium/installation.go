package cilium

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DaemonSetName is the default name for the Cilium DS installed in EKS-A clusters.
	DaemonSetName = "cilium"
	// PreflightDaemonSetName is the default name for the Cilium preflight DS installed
	// in EKS-A clusters during Cilium upgrades.
	PreflightDaemonSetName = "cilium-pre-flight-check"
	// DeploymentName is the default name for the Cilium Operator installed in EKS-A clusters.
	DeploymentName = "cilium-operator"
	// PreflightDeploymentName is the default name for the Cilium preflight deployment used during upgrades.
	PreflightDeploymentName = "cilium-pre-flight-check"
	// ConfigMapName is the default name for the Cilium ConfigMap
	// containing Cilium's configuration.
	ConfigMapName = "cilium-config"
	// ServiceName is the default name for the Cilium Service installed in EKS-A clusters.
	ServiceName = "cilium-agent"

	ciliumConfigMapName   = "cilium-config"
	ciliumConfigNamespace = "kube-system"
)

// Installation is an installation of EKSA Cilium components.
type Installation struct {
	DaemonSet *appsv1.DaemonSet
	Operator  *appsv1.Deployment
	ConfigMap *corev1.ConfigMap
}

// Installed determines if all EKS-A Embedded Cilium components are present. It identifies
// EKS-A Embedded Cilium by the image name. If the ConfigMap doesn't exist we still considered
// Cilium is installed. The installation might not be complete but it can be functional.
func (i Installation) Installed() bool { _ = "STUB: not implemented"; return false }

// GetInstallation creates a new Installation instance. The returned installation's DaemonSet,
// Operator and ConfigMap fields will be nil if they could not be found within the target cluster.
func GetInstallation(ctx context.Context, client client.Client) (*Installation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDaemonSet(ctx context.Context, client client.Client) (*appsv1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getConfigMap(ctx context.Context, client client.Client, name string, namespace string) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDeployment(ctx context.Context, client client.Client) (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
