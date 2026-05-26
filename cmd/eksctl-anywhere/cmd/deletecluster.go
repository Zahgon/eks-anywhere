package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
)

type deleteClusterOptions struct {
	clusterOptions
	wConfig               string
	forceCleanup          bool
	hardwareFileName      string
	tinkerbellBootstrapIP string
	providerOptions       *dependencies.ProviderOptions
}

var dc = &deleteClusterOptions{
	providerOptions: &dependencies.ProviderOptions{
		Tinkerbell: &dependencies.TinkerbellOptions{
			BMCOptions: &hardware.BMCOptions{
				RPC: &hardware.RPCOpts{},
			},
		},
	},
}

var deleteClusterCmd = &cobra.Command{
	Use:          "cluster (<cluster-name>|-f <config-file>)",
	Short:        "Workload cluster",
	Long:         "This command is used to delete workload clusters created by eksctl anywhere",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := dc.validate(cmd.Context(), args); err != nil {
			return err
		}
		if err := dc.deleteCluster(cmd.Context()); err != nil {
			return fmt.Errorf("failed to delete cluster: %v", err)
		}
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteClusterCmd)
	deleteClusterCmd.Flags().StringVarP(&dc.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration, required if <cluster-name> is not provided")
	deleteClusterCmd.Flags().StringVarP(&dc.wConfig, "w-config", "w", "", "Kubeconfig file to use when deleting a workload cluster")
	deleteClusterCmd.Flags().BoolVar(&dc.forceCleanup, "force-cleanup", false, "Force deletion of previously created bootstrap cluster")
	hideForceCleanup(deleteClusterCmd.Flags())
	deleteClusterCmd.Flags().StringVar(&dc.managementKubeconfig, "kubeconfig", "", "kubeconfig file pointing to a management cluster")
	deleteClusterCmd.Flags().StringVar(&dc.bundlesOverride, "bundles-override", "", "Override default Bundles manifest (not recommended)")
	tinkerbellFlags(deleteClusterCmd.Flags(), dc.providerOptions.Tinkerbell.BMCOptions.RPC)
}

func (dc *deleteClusterOptions) validate(ctx context.Context, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (dc *deleteClusterOptions) deleteCluster(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
