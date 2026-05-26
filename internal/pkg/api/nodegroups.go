package api

import (
	corev1 "k8s.io/api/core/v1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type WorkerNodeGroupFiller func(w *anywherev1.WorkerNodeGroupConfiguration)

func FillWorkerNodeGroup(w *anywherev1.WorkerNodeGroupConfiguration, fillers ...WorkerNodeGroupFiller) {
	_ = "STUB: not implemented"
	return
}

func WithTaint(taint corev1.Taint) WorkerNodeGroupFiller {
	_ = "STUB: not implemented"
	return *new(WorkerNodeGroupFiller)
}

func WithNoTaints() WorkerNodeGroupFiller {
	_ = "STUB: not implemented"
	return *new(WorkerNodeGroupFiller)
}

func WithLabel(key, value string) WorkerNodeGroupFiller {
	_ = "STUB: not implemented"
	return *new(WorkerNodeGroupFiller)
}

func WithCount(count int) WorkerNodeGroupFiller {
	_ = "STUB: not implemented"
	return *new(WorkerNodeGroupFiller)
}

func WithMachineGroupRef(name, kind string) WorkerNodeGroupFiller {
	_ = "STUB: not implemented"
	return *new(WorkerNodeGroupFiller)
}
