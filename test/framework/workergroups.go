package framework

import "github.com/aws/eks-anywhere/internal/pkg/api"

type WorkerNodeGroup struct {
	Name                                 string
	Fillers                              []api.WorkerNodeGroupFiller
	MachineConfigKind, MachineConfigName string
}

func WithWorkerNodeGroup(name string, fillers ...api.WorkerNodeGroupFiller) *WorkerNodeGroup {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkerNodeGroup) ClusterFiller() api.ClusterFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterFiller)
}
