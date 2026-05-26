package cluster

import eksav1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"

func BuildMapForWorkerNodeGroupsByName(workerNodeGroups []eksav1alpha1.WorkerNodeGroupConfiguration) map[string]eksav1alpha1.WorkerNodeGroupConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func NodeGroupsToDelete(currentSpec, newSpec *Spec) []eksav1alpha1.WorkerNodeGroupConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// Current spec doesn't have the default name since we never set the defaults at the api server level
