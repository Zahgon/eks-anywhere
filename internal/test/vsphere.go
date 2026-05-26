package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// VSphereClusterSpec builds a complete and valid cluster spec for a vSphere cluster.
func VSphereClusterSpec(tb testing.TB, namespace string, opts ...ClusterSpecOpt) *cluster.Spec {
	_ = "STUB: not implemented"
	return nil
}

// VSphereMachineOpt allows to customize a VSphereMachineConfig.
type VSphereMachineOpt func(config *anywherev1.VSphereMachineConfig)

// VSphereMachineConfig builds a VSphereMachineConfig with some basic defaults.
func VSphereMachineConfig(opts ...VSphereMachineOpt) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

// VSphereDatacenterConfigOpt allows to customize a VSphereDatacenterConfig.
type VSphereDatacenterConfigOpt func(config *anywherev1.VSphereDatacenterConfig)

// VSphereDatacenter builds a VSphereDatacenterConfig for tests with some basix defaults.
func VSphereDatacenter(opts ...VSphereDatacenterConfigOpt) *anywherev1.VSphereDatacenterConfig {
	_ = "STUB: not implemented"
	return nil
}

// VSphereCredentialsSecret builds a new Secret follwoing the format expected for vSphere credentials.
func VSphereCredentialsSecret() *corev1.Secret { _ = "STUB: not implemented"; return nil }

// ClusterOpt allows to customize a Cluster.
type ClusterOpt func(*anywherev1.Cluster)

// Cluster builds a Cluster for tests with some basic defaults.
func Cluster(opts ...ClusterOpt) *anywherev1.Cluster { _ = "STUB: not implemented"; return nil }
