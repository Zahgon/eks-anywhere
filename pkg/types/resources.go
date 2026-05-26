package types

import "time"

type Deployment struct {
	Namespace string
	Name      string
	Container string
}

type Machine struct {
	Metadata MachineMetadata `json:"metadata"`
	Status   MachineStatus   `json:"status"`
}

func (m *Machine) HasAnyLabel(labels []string) bool { _ = "STUB: not implemented"; return false }

type MachineStatus struct {
	NodeRef    *ResourceRef `json:"nodeRef,omitempty"`
	Conditions Conditions
}

type MachineMetadata struct {
	Name   string            `json:"name,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

type ResourceRef struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"Name"`
}

type Conditions []Condition

type ConditionType string

type ConditionStatus string

type Condition struct {
	Type   ConditionType   `json:"type"`
	Status ConditionStatus `json:"status"`
}

type CAPICluster struct {
	Metadata Metadata
	Status   ClusterStatus
}

type ClusterStatus struct {
	Phase      string
	Conditions Conditions
}

type Metadata struct {
	Name string
}

type Datastores struct {
	Info Info `json:"Info"`
}

type Info struct {
	FreeSpace float64 `json:"FreeSpace"`
}

type NowFunc func() time.Time

type NodeReadyChecker func(status MachineStatus) bool

func WithNodeRef() NodeReadyChecker { _ = "STUB: not implemented"; return *new(NodeReadyChecker) }

func WithNodeHealthy() NodeReadyChecker { _ = "STUB: not implemented"; return *new(NodeReadyChecker) }
