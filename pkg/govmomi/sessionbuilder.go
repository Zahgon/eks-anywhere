package govmomi

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
)

type vMOMISessionBuilder struct{}

func NewvMOMISessionBuilder() *vMOMIClientBuilder { _ = "STUB: not implemented"; return nil }

func (*vMOMISessionBuilder) Build(ctx context.Context, u *url.URL, insecure bool) (*govmomi.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
