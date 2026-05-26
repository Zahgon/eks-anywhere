package cluster

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/util/intstr"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

func SetConfigDefaults(c *Config) error { _ = "STUB: not implemented"; return nil }

// ControlPlaneIPCheckAnnotationDefaulter is the defaulter created to set the skip ip value.
type ControlPlaneIPCheckAnnotationDefaulter struct {
	skipCPIPCheck bool
}

// NewControlPlaneIPCheckAnnotationDefaulter allows to create a new ControlPlaneIPCheckAnnotationDefaulter.
func NewControlPlaneIPCheckAnnotationDefaulter(skipIPCheck bool) ControlPlaneIPCheckAnnotationDefaulter {
	_ = "STUB: not implemented"
	return *new(ControlPlaneIPCheckAnnotationDefaulter)
}

// ControlPlaneIPCheckDefault sets the annotation for control plane skip ip check if the flag is set to true.
func (d ControlPlaneIPCheckAnnotationDefaulter) ControlPlaneIPCheckDefault(ctx context.Context, spec *Spec) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MachineHealthCheckDefaulter is the defaulter created to configure the machine health check timeouts.
type MachineHealthCheckDefaulter struct {
	NodeStartupTimeout      time.Duration
	UnhealthyMachineTimeout time.Duration
	MaxUnhealthy            intstr.IntOrString
	WorkerMaxUnhealthy      intstr.IntOrString
}

// NewMachineHealthCheckDefaulter allows to create a new MachineHealthCheckDefaulter.
func NewMachineHealthCheckDefaulter(nodeStartupTimeout, unhealthyMachineTimeout time.Duration, globalMaxUnhealthy, workerMaxUnhealthy intstr.IntOrString) MachineHealthCheckDefaulter {
	_ = "STUB: not implemented"
	return *new(MachineHealthCheckDefaulter)
}

// MachineHealthCheckDefault sets the defaults for machine health check timeouts and maxUnhealthy.
func (d MachineHealthCheckDefaulter) MachineHealthCheckDefault(ctx context.Context, spec *Spec) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetMachineHealthCheckTimeoutDefaults sets default timeouts for MHCs in the EKSA cluster object based on the input.
func SetMachineHealthCheckTimeoutDefaults(cluster *anywherev1.Cluster, nodeStartupTimeout, unhealthyMachineTimeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// SetMachineHealthCheckMaxUnhealthyDefaults sets defaults maxUnhealthy for MHCs in the EKSA cluster object based on the input.
func SetMachineHealthCheckMaxUnhealthyDefaults(cluster *anywherev1.Cluster, globalMaxUnhealthy, workerMaxUnhealthy intstr.IntOrString) {
	_ = "STUB: not implemented"
	return
}

// setMachineHealthCheckTimeoutDefaults sets default timeout values for cluster's machine health checks.
func setMachineHealthCheckTimeoutDefaults(cluster *anywherev1.Cluster, nodeStartupTimeout, unhealthyMachineTimeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// setMachineHealthCheckMaxUnhealthyDefaults sets default maxUnhealthy values for cluster's machine health checks.
func setMachineHealthCheckMaxUnhealthyDefaults(cluster *anywherev1.Cluster, globalMaxUnhealthy, workerMaxUnhealthy intstr.IntOrString) {
	_ = "STUB: not implemented"
	return
}

// NamespaceDefaulter is the defaulter created to configure the cluster's namespace.
type NamespaceDefaulter struct {
	defaultClusterNamespace string
}

// NewNamespaceDefaulter allows to create a new ClusterNamespaceDefaulter.
func NewNamespaceDefaulter(namespace string) NamespaceDefaulter {
	_ = "STUB: not implemented"
	return *new(NamespaceDefaulter)
}

// NamespaceDefault sets the defaults for cluster's namespace.
func (c NamespaceDefaulter) NamespaceDefault(ctx context.Context, spec *Spec) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
