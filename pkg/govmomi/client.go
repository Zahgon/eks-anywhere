package govmomi

import (
	"context"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	VSphereTypeFolder         = "Folder"
	VSphereTypeNetwork        = "Network"
	VSphereTypeResourcePool   = "ResourcePool"
	VSphereTypeDatastore      = "Datastore"
	VSphereTypeVirtualMachine = "VirtualMachine"
	VSphereTypeComputeCluster = "ComputeCluster"
)

type VMOMIAuthorizationManager interface {
	FetchUserPrivilegeOnEntities(ctx context.Context, entities []types.ManagedObjectReference, userName string) ([]types.UserPrivilegeResult, error)
}

type VMOMIFinder interface {
	Datastore(ctx context.Context, path string) (*object.Datastore, error)
	Folder(ctx context.Context, path string) (*object.Folder, error)
	Network(ctx context.Context, path string) (object.NetworkReference, error)
	ResourcePool(ctx context.Context, path string) (*object.ResourcePool, error)
	ClusterComputeResource(ctx context.Context, path string) (*object.ClusterComputeResource, error)
	VirtualMachine(ctx context.Context, path string) (*object.VirtualMachine, error)
	Datacenter(ctx context.Context, path string) (*object.Datacenter, error)
	SetDatacenter(dc *object.Datacenter) *find.Finder
}

type VMOMIClient struct {
	Gcvm                 *govmomi.Client
	Finder               VMOMIFinder
	username             string
	AuthorizationManager VMOMIAuthorizationManager
}

func NewVMOMIClientCustom(gcvm *govmomi.Client, f VMOMIFinder, username string, am VMOMIAuthorizationManager) *VMOMIClient {
	_ = "STUB: not implemented"
	return nil
}

func (vsc *VMOMIClient) Username() string { _ = "STUB: not implemented"; return "" }

func (vsc *VMOMIClient) GetPrivsOnEntity(ctx context.Context, path string, objType string, username string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vsc *VMOMIClient) getFolder(ctx context.Context, path string) (types.ManagedObjectReference, error) {
	_ = "STUB: not implemented"
	return *new(types.ManagedObjectReference), nil
}

func (vsc *VMOMIClient) getNetwork(ctx context.Context, path string) (types.ManagedObjectReference, error) {
	_ = "STUB: not implemented"
	return *new(types.ManagedObjectReference), nil
}

func (vsc *VMOMIClient) getDatastore(ctx context.Context, path string) (types.ManagedObjectReference, error) {
	_ = "STUB: not implemented"
	return *new(types.ManagedObjectReference), nil
}

func (vsc *VMOMIClient) getResourcePool(ctx context.Context, path string) (types.ManagedObjectReference, error) {
	_ = "STUB: not implemented"
	return *new(types.ManagedObjectReference), nil
}

func (vsc *VMOMIClient) getComputeCluster(ctx context.Context, path string) (types.ManagedObjectReference, error) {
	_ = "STUB: not implemented"
	return *new(types.ManagedObjectReference), nil
}

func (vsc *VMOMIClient) getVirtualMachine(ctx context.Context, path string) (types.ManagedObjectReference, error) {
	_ = "STUB: not implemented"
	return *new(types.ManagedObjectReference), nil
}
