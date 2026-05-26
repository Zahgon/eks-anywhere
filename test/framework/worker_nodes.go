package framework

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// WorkerNodeValidation should return an error if either an error is encountered during execution or the validation logically fails.
// This validation function will be executed by ValidateWorkerNodes with a worker node group configuration and a corresponding node
// which was created as a part of that worker node group configuration.
type WorkerNodeValidation func(configuration v1alpha1.WorkerNodeGroupConfiguration, node corev1.Node) (err error)

// ValidateWorkerNodes deduces the worker node group configuration to node mapping
// and for each configuration/node pair executes the provided validation functions.
func (e *ClusterE2ETest) ValidateWorkerNodes(workerNodeValidations ...WorkerNodeValidation) {
	_ = "STUB: not implemented"
	return
}

// deduce the worker node group configuration to node mapping via the machine deployment and machine set

// there will be multiple machineSets present on a cluster following an upgrade.
// find the one that is associated with this worker node, and execute the validations.
