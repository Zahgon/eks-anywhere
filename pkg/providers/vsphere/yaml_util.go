package vsphere

import (
	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

const (
	// VSphereFailureDomainKind is kind for capv failure domain.
	VSphereFailureDomainKind = "VSphereFailureDomain"
	// VSphereDeploymentZoneKind is kind for capv vsphere deployment zone.
	VSphereDeploymentZoneKind = "VSphereDeploymentZone"
)

// FailureDomainsYamlProcessor handles parsing and transformation of failure domains YAML into FailureDomains Objects.
type FailureDomainsYamlProcessor struct {
	parser                *yamlutil.Parser
	failureDomainsBuilder *FailureDomainsBuilder
}

// FailureDomainsBuilder implements yamlutil.Builder
// It's a wrapper around FailureDomains to provide yaml parsing functionality.
type FailureDomainsBuilder struct {
	FailureDomains *FailureDomains
}

// NewFailureDomainsYamlProcessor initializes and returns a new FailureDomainsYamlProcessor.
func NewFailureDomainsYamlProcessor(logger logr.Logger) *FailureDomainsYamlProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessYAML processes the YAML content and returns the failure domains.
func (fbp *FailureDomainsYamlProcessor) ProcessYAML(failureDomainYaml []byte) (*FailureDomains, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fbp *FailureDomainsYamlProcessor) registerFailureDomainMappings() error {
	_ = "STUB: not implemented"
	return nil
}

// BuildFromParsed reads parsed objects in ObjectLookup and sets them in the FailureDomains.
func (fb *FailureDomainsBuilder) BuildFromParsed(lookup yamlutil.ObjectLookup) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessFailureDomainObjects finds all necessary objects in the parsed objects and sets them in FailureDomains.
func (fb *FailureDomainsBuilder) ProcessFailureDomainObjects(f *FailureDomains, lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}

// ProcessFailureDomainGroupObjects looks in the parsed objects for the VsphereFailureDomain by the kind, apiversion
// and VsphereFailureDomain name. Once it is found, it sets in the FailureDomainGroup.
// VsphereDeploymentZone needs to be already set in the FailureDomainGroup.
func (fb *FailureDomainsBuilder) ProcessFailureDomainGroupObjects(g *FailureDomainGroup, lookup yamlutil.ObjectLookup) {
	_ = "STUB: not implemented"
	return
}
