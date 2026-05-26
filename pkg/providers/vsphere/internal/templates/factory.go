package templates

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/providers/vsphere/internal/tags"
)

const (
	libraryContentCorrupted    = "1"
	libraryContentDoesNotExist = "-1"
)

type Factory struct {
	client          GovcClient
	datacenter      string
	datastore       string
	network         string
	resourcePool    string
	templateLibrary string
	tagsFactory     *tags.Factory
}

type GovcClient interface {
	CreateLibrary(ctx context.Context, datastore, library string) error
	DeployTemplateFromLibrary(ctx context.Context, templateDir, templateName, library, datacenter, datastore, network, resourcePool string, resizeBRDisk bool) error
	SearchTemplate(ctx context.Context, datacenter, template string) (string, error)
	ImportTemplate(ctx context.Context, library, ovaURL, name string) error
	LibraryElementExists(ctx context.Context, library string) (bool, error)
	GetLibraryElementContentVersion(ctx context.Context, element string) (string, error)
	DeleteLibraryElement(ctx context.Context, element string) error
	ListTags(ctx context.Context) ([]executables.Tag, error)
	CreateTag(ctx context.Context, tag, category string) error
	AddTag(ctx context.Context, path, tag string) error
	ListCategories(ctx context.Context) ([]string, error)
	CreateCategoryForVM(ctx context.Context, name string) error
	CreateUser(ctx context.Context, username string, password string) error
	UserExists(ctx context.Context, username string) (bool, error)
	CreateGroup(ctx context.Context, name string) error
	GroupExists(ctx context.Context, name string) (bool, error)
	AddUserToGroup(ctx context.Context, name string, username string) error
	RoleExists(ctx context.Context, name string) (bool, error)
	CreateRole(ctx context.Context, name string, privileges []string) error
	SetGroupRoleOnObject(ctx context.Context, principal string, role string, object string, domain string) error
}

func NewFactory(client GovcClient, datacenter, datastore, network, resourcePool, templateLibrary string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) CreateIfMissing(ctx context.Context, datacenter string, machineConfig *v1alpha1.VSphereMachineConfig, ovaURL string, tagsByCategory map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: move this out of the factory into the defaulter, it's a side effect

func (f *Factory) createTemplate(ctx context.Context, templatePath, ovaURL, osFamily string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add rough estimate timing?

func (f *Factory) createLibraryIfMissing(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) importOVAIfMissing(ctx context.Context, templateName, ovaURL string) error {
	_ = "STUB: not implemented"
	return nil
}
