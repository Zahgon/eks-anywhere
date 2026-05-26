package vsphere

import (
	"github.com/go-logr/logr"
	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/cluster"
)

// FailureDomains represents the list of failure domain groups.
type FailureDomains struct {
	Groups []FailureDomainGroup
}

// FailureDomainGroup represents the Vsphere failure domains objects group.
type FailureDomainGroup struct {
	VsphereDeploymentZone *vspherev1.VSphereDeploymentZone
	VsphereFailureDomain  *vspherev1.VSphereFailureDomain
}

const (
	// VsphereDataCenterConfigNameLabel is label for VsphereDataCenter name in Cluster.Spec.VsphereDataCenter.Name.
	VsphereDataCenterConfigNameLabel = "infrastructure.cluster.x-k8s.io/vsphere-datacenter-config-name"
	// ClusterNameLabel is label for cluster name.
	ClusterNameLabel = "infrastructure.cluster.x-k8s.io/cluster-name"
)

// Objects returns a list of API objects for a collection of failure domain groups.
func (f *FailureDomains) Objects() []client.Object { _ = "STUB: not implemented"; return nil }

func templateNamesForFailureDomains(spec *cluster.Spec) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// FailureDomainTemplateName generates the template name for failure domain.
func FailureDomainTemplateName(spec *cluster.Spec, failureDomainName string) string {
	_ = "STUB: not implemented"
	return ""
}

// FailureDomainsSpec generates a vSphere Failure domains spec for the cluster.
func FailureDomainsSpec(logger logr.Logger, spec *cluster.Spec) (*FailureDomains, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
