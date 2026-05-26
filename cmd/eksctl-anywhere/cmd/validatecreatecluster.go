package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
)

type validateOptions struct {
	clusterOptions
	hardwareCSVPath       string
	tinkerbellBootstrapIP string
	providerOptions       *dependencies.ProviderOptions
}

var valOpt = &validateOptions{
	providerOptions: &dependencies.ProviderOptions{
		Tinkerbell: &dependencies.TinkerbellOptions{
			BMCOptions: &hardware.BMCOptions{
				RPC: &hardware.RPCOpts{},
			},
		},
	},
}

var validateCreateClusterCmd = &cobra.Command{
	Use:          "cluster -f <cluster-config-file> [flags]",
	Short:        "Validate create cluster",
	Long:         "Use eksctl anywhere validate create cluster to validate the create cluster action",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	RunE:         valOpt.validateCreateCluster,
}

func init() {
	validateCreateCmd.AddCommand(validateCreateClusterCmd)
	applyTinkerbellHardwareFlag(validateCreateClusterCmd.Flags(), &valOpt.hardwareCSVPath)
	validateCreateClusterCmd.Flags().StringVarP(&valOpt.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	validateCreateClusterCmd.Flags().StringVar(&valOpt.tinkerbellBootstrapIP, "tinkerbell-bootstrap-ip", "", "Override the local tinkerbell IP in the bootstrap cluster")

	if err := validateCreateClusterCmd.MarkFlagRequired("filename"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	tinkerbellFlags(validateCreateClusterCmd.Flags(), valOpt.providerOptions.Tinkerbell.BMCOptions.RPC)
}

func (valOpt *validateOptions) validateCreateCluster(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
