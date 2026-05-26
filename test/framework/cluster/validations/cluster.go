package validations

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

// ValidateClusterReady gets the CAPICluster from the client then validates that it is in a ready state. Also check if CAPI objects are in expected state for InPlace Upgrades.
func ValidateClusterReady(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateEKSAObjects retrieves all the child objects from the cluster.Spec and validates that they exist in the clusterf.
func ValidateEKSAObjects(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCAPIobjectsForInPlace(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateKCP(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMDs(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateInPlaceCRsDoesNotExist(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCPUDeleted(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMDUsDeleted(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNUsAndPodsDeleted(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateControlPlaneNodes retrieves the control plane nodes from the cluster and checks them against the cluster.Spec.
func ValidateControlPlaneNodes(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateWorkerNodes retries the worker nodes from the cluster and checks them against the cluster.Spec.
func ValidateWorkerNodes(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// deduce the worker node group configuration to node mapping via the machine deployment and machine set

// ValidateClusterDoesNotExist checks that the cluster does not exist by attempting to retrieve the CAPI cluster.
func ValidateClusterDoesNotExist(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateCilium gets the cilium-config from the cluster and checks that the cilium
// policy in cluster.Spec matches the enabled policy in the config.
func ValidateCilium(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// It would be nice if we could log something here given we're skipping the validation.

func validateNodeReady(node corev1.Node, kubeVersion v1alpha1.KubernetesVersion) error {
	_ = "STUB: not implemented"
	return nil
}

func validateControlPlaneTaints(cluster *v1alpha1.Cluster, node corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func filterWorkerNodes(nodes []corev1.Node, ms []clusterv1beta2.MachineSet, w v1alpha1.WorkerNodeGroupConfiguration) []corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// there will be multiple machineSets present on a cluster following an upgrade.
// find the one that is associated with this worker node, and execute the validations.

func getWorkerNodeMachineSets(ctx context.Context, vc clusterf.StateValidationConfig, w v1alpha1.WorkerNodeGroupConfiguration) ([]clusterv1beta2.MachineSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateKCPForAPIServerExtraArgs(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Build a lookup map from []Arg for comparison
