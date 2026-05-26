package clustermarshaller

import (
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
)

func MarshalClusterSpec(clusterSpec *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If a GitOpsConfig is present, marshal the GitOpsConfig to file; otherwise, use the FluxConfig
// Allows us to use the FluxConfig internally but preserve the provided spec while GitOpsConfig is being deprecated

func WriteClusterConfig(clusterSpec *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig, writer filewriter.FileWriter) error {
	_ = "STUB: not implemented"
	return nil
}
