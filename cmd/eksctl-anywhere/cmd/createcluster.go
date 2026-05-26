package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/aws/eks-anywhere/cmd/eksctl-anywhere/cmd/aflag"
	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
	"github.com/aws/eks-anywhere/pkg/validations/createvalidations"
)

type createClusterOptions struct {
	clusterOptions
	timeoutOptions
	forceClean            bool
	skipIpCheck           bool
	hardwareCSVPath       string
	tinkerbellBootstrapIP string
	installPackages       string
	skipValidations       []string
	providerOptions       *dependencies.ProviderOptions
}

var cc = &createClusterOptions{
	providerOptions: &dependencies.ProviderOptions{
		Tinkerbell: &dependencies.TinkerbellOptions{
			BMCOptions: &hardware.BMCOptions{
				RPC: &hardware.RPCOpts{},
			},
		},
	},
}

var createClusterCmd = &cobra.Command{
	Use:          "cluster -f <cluster-config-file> [flags]",
	Short:        "Create workload cluster",
	Long:         "This command is used to create workload clusters",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	RunE:         cc.createCluster,
}

func init() {
	createCmd.AddCommand(createClusterCmd)
	applyClusterOptionFlags(createClusterCmd.Flags(), &cc.clusterOptions)
	applyTimeoutFlags(createClusterCmd.Flags(), &cc.timeoutOptions)
	applyTinkerbellHardwareFlag(createClusterCmd.Flags(), &cc.hardwareCSVPath)
	aflag.String(aflag.TinkerbellBootstrapIP, &cc.tinkerbellBootstrapIP, createClusterCmd.Flags())
	createClusterCmd.Flags().BoolVar(&cc.forceClean, "force-cleanup", false, "Force deletion of previously created bootstrap cluster")
	hideForceCleanup(createClusterCmd.Flags())
	createClusterCmd.Flags().BoolVar(&cc.skipIpCheck, "skip-ip-check", false, "Skip check for whether cluster control plane ip is in use")
	createClusterCmd.Flags().StringVar(&cc.installPackages, "install-packages", "", "Location of curated packages configuration files to install to the cluster")
	createClusterCmd.Flags().StringArrayVar(&cc.skipValidations, "skip-validations", []string{}, fmt.Sprintf("Bypass create validations by name. Valid arguments you can pass are --skip-validations=%s", strings.Join(createvalidations.SkippableValidations[:], ",")))
	tinkerbellFlags(createClusterCmd.Flags(), cc.providerOptions.Tinkerbell.BMCOptions.RPC)

	aflag.MarkRequired(createClusterCmd.Flags(), aflag.ClusterConfig.Name)
}

func tinkerbellFlags(fs *pflag.FlagSet, r *hardware.RPCOpts) { _ = "STUB: not implemented"; return }

func (cc *createClusterOptions) createCluster(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
