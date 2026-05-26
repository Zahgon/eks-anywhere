package tags

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/executables"
)

type Factory struct {
	client GovcClient
}

type GovcClient interface {
	ListTags(ctx context.Context) ([]executables.Tag, error)
	CreateTag(ctx context.Context, tag, category string) error
	AddTag(ctx context.Context, path, tag string) error
	ListCategories(ctx context.Context) ([]string, error)
	CreateCategoryForVM(ctx context.Context, name string) error
}

func NewFactory(client GovcClient) *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) TagTemplate(ctx context.Context, templatePath string, tagsByCategory map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}
