package executables

import (
	"bytes"
	"context"
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
)

//go:embed config/cmk.ini
var cmkConfigTemplate string

const (
	cmkPath                           = "cmk"
	cmkConfigFileNameTemplate         = "cmk_%s.ini"
	defaultCloudStackPreflightTimeout = "30"
	rootDomain                        = "ROOT"
	domainDelimiter                   = "/"
)

// Cmk this struct wraps around the CloudMonkey executable CLI to perform operations against a CloudStack endpoint.
type Cmk struct {
	writer     filewriter.FileWriter
	executable Executable
	configMap  map[string]decoder.CloudStackProfileConfig
}

type listTemplatesResponse struct {
	CmkTemplates []cmkTemplate `json:"template"`
}

func (c *Cmk) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cmk) ValidateTemplatePresent(ctx context.Context, profile string, domainId string, zoneId string, account string, template v1alpha1.CloudStackResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

// SearchTemplate looks for a template by name or by id and returns template name if found.
func (c *Cmk) SearchTemplate(ctx context.Context, profile string, template v1alpha1.CloudStackResourceIdentifier) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Cmk) ValidateServiceOfferingPresent(ctx context.Context, profile string, zoneId string, serviceOffering v1alpha1.CloudStackResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmk) ValidateDiskOfferingPresent(ctx context.Context, profile string, zoneId string, diskOffering v1alpha1.CloudStackResourceDiskOffering) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmk) ValidateAffinityGroupsPresent(ctx context.Context, profile string, domainId string, account string, affinityGroupIds []string) error {
	_ = "STUB: not implemented"
	return nil
}

// account must be specified with a domainId
// domainId can be specified without account

func (c *Cmk) ValidateZoneAndGetId(ctx context.Context, profile string, zone v1alpha1.CloudStackZone) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Cmk) ValidateDomainAndGetId(ctx context.Context, profile string, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// "list domains" API does not support querying by domain path, so here we extract the domain name which is the last part of the input domain

// EnsureNoDuplicateNetwork ensures that there are no duplicate networks with the name networkName.
// If it finds duplicates that are not shared networks, it deletes them.
func (c *Cmk) EnsureNoDuplicateNetwork(ctx context.Context, profile string, networkName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmk) ValidateNetworkPresent(ctx context.Context, profile string, domainId string, network v1alpha1.CloudStackResourceIdentifier, zoneId string, account string) error {
	_ = "STUB: not implemented"
	return nil
}

// account must be specified within a domainId
// domainId can be specified without account

// filter by network name -- cmk does not support name= filter
// if network id and name are both provided, the following code is to confirm name matches return value retrieved by id.
// if only name is provided, the following code is to only get networks with specified name.

func (c *Cmk) ValidateAccountPresent(ctx context.Context, profile string, account string, domainId string) error {
	_ = "STUB: not implemented"
	// If account is not specified then no need to check its presence
	return nil
}

// NewCmk initializes CloudMonkey executable to query CloudStack via CLI.
func NewCmk(executable Executable, writer filewriter.FileWriter, config *decoder.CloudStackExecConfig) (*Cmk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cmk) GetManagementApiEndpoint(profile string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Cmk) CleanupVms(ctx context.Context, profile string, clusterName string, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmk) exec(ctx context.Context, profile string, args ...string) (stdout bytes.Buffer, err error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (c *Cmk) buildCmkConfigFile(profile string) (configFile string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

type cmkTemplate struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Zonename string `json:"zonename"`
}

type cmkServiceOffering struct {
	CpuNumber int    `json:"cpunumber"`
	CpuSpeed  int    `json:"cpuspeed"`
	Memory    int    `json:"memory"`
	Id        string `json:"id"`
	Name      string `json:"name"`
}

type cmkNetwork struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type cmkResourceIdentifier struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type cmkDiskOffering struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Customized bool   `json:"iscustomized"`
}

type cmkAffinityGroup struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type cmkDomain struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type cmkAccount struct {
	RoleType string `json:"roletype"`
	Domain   string `json:"domain"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}
