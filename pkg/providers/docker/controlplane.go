package docker

import (
	"context"

	"github.com/go-logr/logr"
	dockerv1beta2 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	yamlcapi "github.com/aws/eks-anywhere/pkg/clusterapi/yaml"
	"github.com/aws/eks-anywhere/pkg/yamlutil"
)

// ControlPlane represents a CAPI Docker control plane.
type ControlPlane = clusterapi.ControlPlane[*dockerv1beta2.DockerCluster, *dockerv1beta2.DockerMachineTemplate]

type controlPlaneBuilder = yamlcapi.ControlPlaneBuilder[*dockerv1beta2.DockerCluster, *dockerv1beta2.DockerMachineTemplate]

// ControlPlaneSpec builds a docker ControlPlane definition based on an eks-a cluster spec.
func ControlPlaneSpec(ctx context.Context, logger logr.Logger, client kubernetes.Client, spec *cluster.Spec) (*ControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newControlPlaneParser(logger logr.Logger) (*yamlutil.Parser, *controlPlaneBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
