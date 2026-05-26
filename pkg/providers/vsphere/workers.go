package vsphere

import (
	"context"

	"github.com/go-logr/logr"
	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	capiyaml "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

type (
	// Workers represents the vSphere specific CAPI spec for worker nodes.
	Workers        = clusterapi.Workers[*vspherev1.VSphereMachineTemplate]
	workersBuilder = capiyaml.WorkersBuilder[*vspherev1.VSphereMachineTemplate]
)

// WorkersSpec generates a vSphere specific CAPI spec for an eks-a cluster worker nodes.
// It talks to the cluster with a client to detect changes in immutable objects and generates new
// names for them.
func WorkersSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, spec *cluster.Spec) (*Workers, error) {
	_ = "STUB: not implemented"
	// TODO(g-gaston): refactor template builder so it doesn't behave differently for controller and CLI
	// TODO(g-gaston): do we need time.Now if the names are not dependent on a timestamp anymore?
	return nil, nil
}

func newWorkersParserAndBuilder(logger logr.Logger) (*yamlutil.Parser, *workersBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func machineTemplateMapping() yamlutil.Mapping[*vspherev1.VSphereMachineTemplate] {
	_ = "STUB: not implemented"
	return nil
}
