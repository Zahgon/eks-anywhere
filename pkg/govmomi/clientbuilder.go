package govmomi

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
)

type VSphereClient interface {
	Username() string
	GetPrivsOnEntity(ctx context.Context, path string, objType string, username string) ([]string, error)
}

type VMOMIFinderBuilder interface {
	Build(arg0 *vim25.Client, arg1 ...bool) VMOMIFinder
}

type VMOMISessionBuilder interface {
	Build(ctx context.Context, u *url.URL, insecure bool) (*govmomi.Client, error)
}

type VMOMIAuthorizationManagerBuilder interface {
	Build(c *vim25.Client) *object.AuthorizationManager
}

type vMOMIClientBuilder struct {
	vfb VMOMIFinderBuilder
	gcb VMOMISessionBuilder
	amb VMOMIAuthorizationManagerBuilder
}

func NewVMOMIClientBuilder() *vMOMIClientBuilder { _ = "STUB: not implemented"; return nil }

func NewVMOMIClientBuilderOverride(vfb VMOMIFinderBuilder, gcb VMOMISessionBuilder, amb VMOMIAuthorizationManagerBuilder) *vMOMIClientBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (vcb *vMOMIClientBuilder) Build(ctx context.Context, host string, username string, password string, insecure bool, datacenter string) (VSphereClient, error) {
	_ = "STUB: not implemented"
	return *new(VSphereClient), nil
}

// start gvmc
