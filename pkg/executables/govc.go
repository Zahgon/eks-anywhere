package executables

import (
	"bytes"
	"context"
	_ "embed"
	"time"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	govcPath             = "govc"
	govcUsernameKey      = "GOVC_USERNAME"
	govcPasswordKey      = "GOVC_PASSWORD"
	govcURLKey           = "GOVC_URL"
	govcInsecure         = "GOVC_INSECURE"
	govcDatacenterKey    = "GOVC_DATACENTER"
	govcTlsHostsFile     = "govc_known_hosts"
	govcTlsKnownHostsKey = "GOVC_TLS_KNOWN_HOSTS"
	vSphereServerKey     = "VSPHERE_SERVER"
	byteToGiB            = 1073741824.0
	DeployOptsFile       = "deploy-opts.json"
	disk1                = "Hard disk 1"
	disk2                = "Hard disk 2"
	MemoryAvailable      = "Memory_Available"
)

var requiredEnvs = []string{govcUsernameKey, govcPasswordKey, govcURLKey, govcInsecure, govcDatacenterKey}

type networkMapping struct {
	Name    string `json:"Name,omitempty"`
	Network string `json:"Network,omitempty"`
}

type deployOption struct {
	DiskProvisioning string           `json:"DiskProvisioning,omitempty"`
	NetworkMapping   []networkMapping `json:"NetworkMapping,omitempty"`
}

type FolderType string

const (
	datastore     FolderType = "datastore"
	vm            FolderType = "vm"
	maxRetries               = 5
	backOffPeriod            = 5 * time.Second
)

type Govc struct {
	writer filewriter.FileWriter
	Executable
	*retrier.Retrier
	requiredEnvs *syncSlice
	envMap       map[string]string
}

type GovcOpt func(*Govc)

func NewGovc(executable Executable, writer filewriter.FileWriter, opts ...GovcOpt) *Govc {
	_ = "STUB: not implemented"
	return nil
}

func WithGovcEnvMap(envMap map[string]string) GovcOpt {
	_ = "STUB: not implemented"
	return *new(GovcOpt)
}

func (g *Govc) exec(ctx context.Context, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (g *Govc) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (g *Govc) Logout(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Commands that skip cert verification will have a different session.
// So we try to destroy it as well here to avoid leaving it orphaned

// SearchTemplate looks for a vm template with the same base name as the provided template path.
// If found, it returns the full qualified path to the template.
// If multiple matching templates are found, it returns an error.
func (g *Govc) SearchTemplate(ctx context.Context, datacenter, template string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *Govc) LibraryElementExists(ctx context.Context, library string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type libElement struct {
	ContentVersion string `json:"content_version"`
}

func (g *Govc) GetLibraryElementContentVersion(ctx context.Context, element string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *Govc) DeleteLibraryElement(ctx context.Context, element string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) ResizeDisk(ctx context.Context, datacenter, template, diskName string, diskSizeInGB int) error {
	_ = "STUB: not implemented"
	return nil
}

type deviceInfoResponse struct {
	Devices []VirtualDevice
}

// VirtualDevice describes a virtual device for a VM.
type VirtualDevice struct {
	Name         string
	DeviceInfo   deviceInfo
	CapacityInKB float64
}

type deviceInfo struct {
	Label string
}

// DevicesInfo returns the device info for te provided virtual machine.
func (g *Govc) DevicesInfo(ctx context.Context, datacenter, template string, args ...string) ([]VirtualDevice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVMDiskSizeInGB returns the size of the first disk on the VM in GB.
func (g *Govc) GetVMDiskSizeInGB(ctx context.Context, vm, datacenter string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetHardDiskSize returns the size of all the hard disks for given VM.
func (g *Govc) GetHardDiskSize(ctx context.Context, vm, datacenter string) (map[string]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Govc) TemplateHasSnapshot(ctx context.Context, template string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type datastoreResponse struct {
	Datastores []types.Datastores `json:"Datastores"`
}

func (g *Govc) GetWorkloadAvailableSpace(ctx context.Context, datastore string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (g *Govc) CreateLibrary(ctx context.Context, datastore, library string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) DeployTemplateFromLibrary(ctx context.Context, templateDir, templateName, library, datacenter, datastore, network, resourcePool string, resizeBRDisk bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get devices information template to identify second disk properly

// For 1.22 we switched to using one disk for BR, so it for now as long as the boolean is set, and we only see
// one disk, we can assume this is for 1.22. This loop would need to change if that assumption changes
// in the future, but 1.20 and 1.21 are still using dual disks which is why we need to check for the second
// disk first. Since this loop will get all kinds of devices and not just hard disks, we need to do these
// checks based on the label.

// Get the name of the hard disk and resize the disk to 20G

func (g *Govc) ImportTemplate(ctx context.Context, library, ovaURL, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) DeployTemplate(ctx context.Context, library, templateName, vmName, deployFolder, datacenter, datastore, network, resourcePool string, deployOptionsOverride []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) DeleteTemplate(ctx context.Context, resourcePool, templatePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) markAsVM(ctx context.Context, resourcePool, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) removeSnapshotsFromVM(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) deleteVM(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) createVMSnapshot(ctx context.Context, datacenter, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) markVMAsTemplate(ctx context.Context, datacenter, vmName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) getEnvMap() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *Govc) validateAndSetupCreds() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Govc) CleanupVms(ctx context.Context, clusterName string, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) ValidateVCenterConnection(ctx context.Context, server string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) ValidateVCenterAuthentication(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) IsCertSelfSigned(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (g *Govc) GetCertThumbprint(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *Govc) ConfigureCertThumbprint(ctx context.Context, server, thumbprint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) DatacenterExists(ctx context.Context, datacenter string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (g *Govc) NetworkExists(ctx context.Context, network string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetDatastorePath validates and returns the full path to a datastore in the specified datacenter.
// Returns an error if the datastore doesn't exist or if the path is invalid.
func (g *Govc) GetDatastorePath(ctx context.Context, datacenter string, datastorePath string, envMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetFolderPath validates or creates a folder in the specified datacenter.
// Returns the full path to the folder or an error if creation fails.
func (g *Govc) GetFolderPath(ctx context.Context, datacenter string, folder string, envMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetResourcePoolPath finds and validates a resource pool in the specified datacenter.
// Returns an error if the pool doesn't exist or if multiple matching pools are found.
func (g *Govc) GetResourcePoolPath(ctx context.Context, datacenter string, resourcePool string, envMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetComputeClusterPath finds and validates a compute cluster in the specified datacenter.
// Returns an error if the compute cluster doesn't exist or if multiple matching compute clusters are found.
func (g *Govc) GetComputeClusterPath(ctx context.Context, datacenter string, computeCluster string, envMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ValidateVCenterSetupMachineConfig validates that all resources specified in a
// VSphereMachineConfig exist and are accessible.
func (g *Govc) ValidateVCenterSetupMachineConfig(ctx context.Context, datacenterConfig *v1alpha1.VSphereDatacenterConfig, machineConfig *v1alpha1.VSphereMachineConfig, _ *bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateFailureDomainConfig validates that all resources specified in a VSphere
// failure domain exist and are accessible.
func (g *Govc) ValidateFailureDomainConfig(ctx context.Context, datacenterConfig *v1alpha1.VSphereDatacenterConfig, failureDomain *v1alpha1.FailureDomain) error {
	_ = "STUB: not implemented"
	return nil
}

func prependPath(folderType FolderType, folderPath string, datacenter string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *Govc) createFolder(ctx context.Context, envMap map[string]string, folderPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) isValidPath(ctx context.Context, envMap map[string]string, path string) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Govc) GetTags(ctx context.Context, path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tag struct to represent a vSphere Tag.
type Tag struct {
	Id         string
	Name       string
	CategoryId string `json:"category_id,omitempty"`
}

// ListTags list all vSphere tags in vCenter.
func (g *Govc) ListTags(ctx context.Context) ([]Tag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Govc) AddTag(ctx context.Context, path, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) CreateTag(ctx context.Context, tag, category string) error {
	_ = "STUB: not implemented"
	return nil
}

type category struct {
	Id              string
	Name            string
	Cardinality     string
	AssociableTypes []string `json:"associable_types,omitempty"`
}

func (g *Govc) ListCategories(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type objectType string

const virtualMachine objectType = "VirtualMachine"

func (g *Govc) CreateCategoryForVM(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Govc) createCategory(ctx context.Context, name string, objectTypes []objectType) error {
	_ = "STUB: not implemented"
	return nil
}

func getDeployOptions(network string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// needed for Ubuntu

// needed for Bottlerocket

// CreateUser creates a user.
func (g *Govc) CreateUser(ctx context.Context, username string, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// UserExists checks if a user exists.
func (g *Govc) UserExists(ctx context.Context, username string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CreateGroup creates a group.
func (g *Govc) CreateGroup(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupExists checks if a group exists.
func (g *Govc) GroupExists(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// govc returns empty response when group exists

// AddUserToGroup adds a user to a group.
func (g *Govc) AddUserToGroup(ctx context.Context, name string, username string) error {
	_ = "STUB: not implemented"
	return nil
}

// RoleExists checks if a role exists.
func (g *Govc) RoleExists(ctx context.Context, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CreateRole creates a role with specified privileges.
func (g *Govc) CreateRole(ctx context.Context, name string, privileges []string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetGroupRoleOnObject sets a role for a given group on target object.
func (g *Govc) SetGroupRoleOnObject(ctx context.Context, principal string, role string, object string, domain string) error {
	_ = "STUB: not implemented"
	return nil
}

type resourcePoolInfo struct {
	ResourcePoolIdentifier *resourcePool
}

type resourcePool struct {
	memoryUsage string
	memoryLimit string
}

// GetResourcePoolInfo returns the pool info for the provided resource pool.
func (g *Govc) GetResourcePoolInfo(ctx context.Context, datacenter, resourcepool string, args ...string) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getPoolInfo parses resource pool response and returns memory requirements.
func getPoolInfo(rp *resourcePool) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getValueFromString cleans the input string and returns the extracted numerical value.
func getValueFromString(str string) (int, error) { _ = "STUB: not implemented"; return 0, nil }
