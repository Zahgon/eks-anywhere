package vsphere

import (
	"context"
	_ "embed"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/collection"
	"github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/govmomi"
)

const (
	vsphereRootPath = "/"
)

type PrivAssociation struct {
	objectType   string
	privsContent string
	path         string
}

type missingPriv struct {
	Username    string   `yaml:"username"`
	ObjectType  string   `yaml:"objectType"`
	Path        string   `yaml:"path"`
	Permissions []string `yaml:"permissions"`
}

type VSphereClientBuilder interface {
	Build(ctx context.Context, host string, username string, password string, insecure bool, datacenter string) (govmomi.VSphereClient, error)
}

// ResourcePaths defines an interface for objects that contain vSphere resource path information.
type ResourcePaths interface {
	ResourcePaths() map[string]string
}

type Validator struct {
	govc                 ProviderGovcClient
	vSphereClientBuilder VSphereClientBuilder
}

// NewValidator initializes the client for VSphere provider validations.
func NewValidator(govc ProviderGovcClient, vscb VSphereClientBuilder) *Validator {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateVCenterAccess(ctx context.Context, server string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) ValidateVCenterConfig(ctx context.Context, datacenterConfig *anywherev1.VSphereDatacenterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateFailureDomains validates the provided list of failure domains.
func (v *Validator) ValidateFailureDomains(ctx context.Context, vsphereClusterSpec *Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Error message here if Failure Domain not being used by workernodegroups?
// Skipping further validation currently
// return fmt.Errorf("failure domain defined, but no worker node group references")

func (v *Validator) validateWorkerNodeGroupDomains(vsphereClusterSpec *Spec, providedFailureDomains collection.Set[string]) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (v *Validator) validateFailureDomainResources(ctx context.Context, vsphereClusterSpec *Spec, failureDomains []anywherev1.FailureDomain) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateMachineConfigTagsExist(ctx context.Context, machineConfigs []*anywherev1.VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateClusterMachineConfigs validates all the attributes of etcd, control plane, and worker node VSphereMachineConfigs.
func (v *Validator) ValidateClusterMachineConfigs(ctx context.Context, vsphereClusterSpec *Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: move this to api Cluster validations

// Temporary until we remove the need to pass a bool pointer
// TODO: remove side effects from this implementation or directly move it to set defaults (pointer to bool is not needed)

func (v *Validator) validateControlPlaneIp(ip string) error {
	_ = "STUB: not implemented"
	// check if controlPlaneEndpointIp is valid
	return nil
}

func (v *Validator) validateTemplates(ctx context.Context, spec *Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) getTemplatePath(ctx context.Context, datacenter, templatePath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v *Validator) validateTemplateTags(ctx context.Context, templatePath string, requiredTags []string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: maybe add help text about to how to tag a template?

func (v *Validator) validateBRHardDiskSize(ctx context.Context, spec *Spec, machineConfigSpec *anywherev1.VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// 2GB in KB to avoid roundoff errors

// 20GB in KB to avoid roundoff errors

// 22GB in KB to avoid roundoff errors

func (v *Validator) validateThumbprint(ctx context.Context, datacenterConfig *anywherev1.VSphereDatacenterConfig) error {
	_ = "STUB: not implemented"
	// No need to validate thumbprint in insecure mode
	return nil
}

// If cert is not self signed, thumbprint is ignored

func (v *Validator) validateDatacenter(ctx context.Context, datacenter string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateNetwork(ctx context.Context, network string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateNetworksFieldUsage(ctx context.Context, vsphereClusterSpec *Spec) error {
	_ = "STUB: not implemented"
	// Check control plane - should NOT have networks field
	return nil
}

// Check etcd - should NOT have networks field

// Validate worker node networks

func (v *Validator) validateWorkerMachineConfigNetworks(ctx context.Context, machineConfig *anywherev1.VSphereMachineConfig, workerGroupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) collectResourcePathConfig(_ context.Context, spec *Spec) ([]ResourcePaths, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) validateVsphereUserPrivs(ctx context.Context, vSphereClusterSpec *Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func markPrivsValidationPass(passed bool, username string) { _ = "STUB: not implemented"; return }

func (v *Validator) validateUserPrivs(ctx context.Context, spec *Spec, vuc *config.VSphereUserConfig) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// validate global root priv settings are correct

// ToDo: add more sophisticated validation around a scenario where someone has uploaded templates
// on their own and does not want to allow EKSA user write access to templates
// Verify privs on the template

func (v *Validator) validateCPUserPrivs(ctx context.Context, spec *Spec, vuc *config.VSphereUserConfig) (bool, error) {
	_ = "STUB: not implemented"
	// CP role just needs read only
	return false, nil
}

func (v *Validator) validatePrivs(ctx context.Context, privObjs []PrivAssociation, vsc govmomi.VSphereClient) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkRequiredPrivs(requiredPrivs []string, hasPrivs []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) getMissingPrivs(ctx context.Context, vsc govmomi.VSphereClient, path string, objType string, requiredPrivsContent string, username string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) sameOSFamily(configs map[string]*anywherev1.VSphereMachineConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *Validator) sameTemplate(configs map[string]*anywherev1.VSphereMachineConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func getRandomMachineConfig(configs map[string]*anywherev1.VSphereMachineConfig) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func sliceIfNotNil(machines ...*anywherev1.VSphereMachineConfig) []*anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}
