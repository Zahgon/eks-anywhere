package docker

import (
	"context"

	"github.com/go-logr/logr"
	dockerv1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	capiyaml "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

type (
	// Workers represents the docker specific CAPI spec for worker nodes.
	Workers        = clusterapi.Workers[*dockerv1.DockerMachineTemplate]
	workersBuilder = capiyaml.WorkersBuilder[*dockerv1.DockerMachineTemplate]
)

// WorkersSpec generates a Docker specific CAPI spec for an eks-a cluster worker nodes.
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

func machineTemplateMapping() yamlutil.Mapping[*dockerv1.DockerMachineTemplate] {
	_ = "STUB: not implemented"
	return nil
}
