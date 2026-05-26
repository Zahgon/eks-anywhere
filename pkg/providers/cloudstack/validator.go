package cloudstack

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/networkutils"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
	"github.com/aws/eks-anywhere/pkg/types"
)

type Validator struct {
	cmk         ProviderCmkClient
	netClient   networkutils.NetClient
	skipIpCheck bool
}

func NewValidator(cmk ProviderCmkClient, netClient networkutils.NetClient, skipIpCheck bool) *Validator {
	_ = "STUB: not implemented"
	return nil
}

type localAvailabilityZone struct {
	*anywherev1.CloudStackAvailabilityZone
	ZoneId   string
	DomainId string
}

// ProviderCmkClient defines the methods used by Cmk as a separate interface to be mockable when injected into other objects.
type ProviderCmkClient interface {
	GetManagementApiEndpoint(profile string) (string, error)
	ValidateServiceOfferingPresent(ctx context.Context, profile string, zoneId string, serviceOffering anywherev1.CloudStackResourceIdentifier) error
	ValidateDiskOfferingPresent(ctx context.Context, profile string, zoneId string, diskOffering anywherev1.CloudStackResourceDiskOffering) error
	ValidateTemplatePresent(ctx context.Context, profile string, domainId string, zoneId string, account string, template anywherev1.CloudStackResourceIdentifier) error
	ValidateAffinityGroupsPresent(ctx context.Context, profile string, domainId string, account string, affinityGroupIds []string) error
	ValidateZoneAndGetId(ctx context.Context, profile string, zone anywherev1.CloudStackZone) (string, error)
	ValidateNetworkPresent(ctx context.Context, profile string, domainId string, network anywherev1.CloudStackResourceIdentifier, zoneId string, account string) error
	ValidateDomainAndGetId(ctx context.Context, profile string, domain string) (string, error)
	ValidateAccountPresent(ctx context.Context, profile string, account string, domainId string) error
}

func (v *Validator) ValidateCloudStackDatacenterConfig(ctx context.Context, datacenterConfig *anywherev1.CloudStackDatacenterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func generateLocalAvailabilityZones(ctx context.Context, datacenterConfig *anywherev1.CloudStackDatacenterConfig) ([]localAvailabilityZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: dry out machine configs validations.
// Cyclomatic complexity is high. The exception below can probably be removed once the above todo is done.
// nolint:gocyclo
func (v *Validator) ValidateClusterMachineConfigs(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// validate template field name contains cluster kubernetes version for the control plane machine.

// validate template field name contains cluster kubernetes version for the external etcd machine.

// validate template field of worker group spec with the kubernetes version of each workerNodeGroup - in case of modular upgrade.

func (v *Validator) ValidateControlPlaneEndpointUniqueness(endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateMachineConfig(ctx context.Context, datacenterConfig *anywherev1.CloudStackDatacenterConfig, machineConfig *anywherev1.CloudStackMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateTemplateMatchesKubernetesVersion(ctx context.Context, templateName string, kubernetesVersionName string) error {
	_ = "STUB: not implemented"
	// Replace 1.23, 1-23, 1_23 to 123 in the template name string.
	return nil
}

// Replace 1-23 to 123 in the kubernetesversion string.

// This will return an error if the template name does not contain specified kubernetes version.
// For ex if the kubernetes version is 1.23,
// the template name should include 1.23 or 1-23, 1_23 or 123 i.e. kubernetes-1-23-eks in the string.

// ValidateSecretsUnchanged checks the secret to see if it has not been changed.
func (v *Validator) ValidateSecretsUnchanged(ctx context.Context, cluster *types.Cluster, execConfig *decoder.CloudStackExecConfig, client ProviderKubectlClient) error {
	_ = "STUB: not implemented"
	return nil
}

// When the secret is not found we allow for new secrets

func secretDifferentFromProfile(secret *corev1.Secret, profile decoder.CloudStackProfileConfig) bool {
	_ = "STUB: not implemented"
	return false
}
