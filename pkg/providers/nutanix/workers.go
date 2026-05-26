package nutanix

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/nutanix-cloud-native/cluster-api-provider-nutanix/api/v1beta1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	capiyaml "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

type (
	// Workers represents the nutanix specific CAPI spec for worker nodes.
	Workers        = clusterapi.Workers[*v1beta1.NutanixMachineTemplate]
	workersBuilder = capiyaml.WorkersBuilder[*v1beta1.NutanixMachineTemplate]
)

// WorkersSpec generates a nutanix specific CAPI spec for an eks-a cluster worker nodes.
// It talks to the cluster with a client to detect changes in immutable objects and generates new
// names for them.
func WorkersSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, spec *cluster.Spec) (*Workers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTemplateNames(spec *cluster.Spec, templateBuilder *TemplateBuilder, machineConfigs map[string]*v1alpha1.NutanixMachineConfig) (map[string]string, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseWorkersYaml(logger logr.Logger, workersYaml []byte) (*Workers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWorkersParserAndBuilder(logger logr.Logger) (*yamlutil.Parser, *workersBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func machineTemplateMapping() yamlutil.Mapping[*v1beta1.NutanixMachineTemplate] {
	_ = "STUB: not implemented"
	return nil
}

func getMachineTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*v1beta1.NutanixMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func machineTemplateEquals(new, old *v1beta1.NutanixMachineTemplate) bool {
	_ = "STUB: not implemented"
	return false
}
