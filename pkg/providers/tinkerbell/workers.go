package tinkerbell

import (
	"context"

	"github.com/go-logr/logr"

	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	capiyaml "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

type (
	// Workers represents the Tinkerbell specific CAPI spec for worker nodes.
	Workers        = clusterapi.Workers[*tinkerbellv1.TinkerbellMachineTemplate]
	workersBuilder = capiyaml.WorkersBuilder[*tinkerbellv1.TinkerbellMachineTemplate]
)

// WorkersSpec generates a Tinkerbell specific CAPI spec for an eks-a cluster worker nodes.
// It talks to the cluster with a client to detect changes in immutable objects and generates new
// names for them.
func WorkersSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, spec *cluster.Spec) (*Workers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWorkersParserAndBuilder(logger logr.Logger) (*yamlutil.Parser, *workersBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func machineTemplateMapping() yamlutil.Mapping[*tinkerbellv1.TinkerbellMachineTemplate] {
	_ = "STUB: not implemented"
	return nil
}
