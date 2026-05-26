package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/cmd/eksctl-anywhere/cmd/aflag"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/logger"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
	"github.com/aws/eks-anywhere/pkg/validations/upgradevalidations"
)

type upgradeClusterOptions struct {
	clusterOptions
	timeoutOptions
	wConfig               string
	forceClean            bool
	hardwareCSVPath       string
	tinkerbellBootstrapIP string
	skipValidations       []string
	providerOptions       *dependencies.ProviderOptions
}

var uc = &upgradeClusterOptions{
	providerOptions: &dependencies.ProviderOptions{
		Tinkerbell: &dependencies.TinkerbellOptions{
			BMCOptions: &hardware.BMCOptions{
				RPC: &hardware.RPCOpts{},
			},
		},
	},
}

var upgradeClusterCmd = &cobra.Command{
	Use:          "cluster",
	Short:        "Upgrade workload cluster",
	Long:         "This command is used to upgrade workload clusters",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if uc.forceClean {
			logger.MarkFail(forceCleanupDeprecationMessageForUpgrade)
			return errors.New("please remove the --force-cleanup flag")
		}

		if err := uc.upgradeCluster(cmd, args); err != nil {
			return fmt.Errorf("failed to upgrade cluster: %v", err)
		}
		return nil
	},
}

func init() {
	upgradeCmd.AddCommand(upgradeClusterCmd)
	applyClusterOptionFlags(upgradeClusterCmd.Flags(), &uc.clusterOptions)
	applyTimeoutFlags(upgradeClusterCmd.Flags(), &uc.timeoutOptions)
	applyTinkerbellHardwareFlag(upgradeClusterCmd.Flags(), &uc.hardwareCSVPath)
	upgradeClusterCmd.Flags().StringVarP(&uc.wConfig, "w-config", "w", "", "Kubeconfig file to use when upgrading a workload cluster")
	upgradeClusterCmd.Flags().BoolVar(&uc.forceClean, "force-cleanup", false, "Force deletion of previously created bootstrap cluster")
	hideForceCleanup(upgradeClusterCmd.Flags())
	upgradeClusterCmd.Flags().StringArrayVar(&uc.skipValidations, "skip-validations", []string{}, fmt.Sprintf("Bypass upgrade validations by name. Valid arguments you can pass are --skip-validations=%s", strings.Join(upgradevalidations.SkippableValidations[:], ",")))
	aflag.MarkRequired(createClusterCmd.Flags(), aflag.ClusterConfig.Name)
	tinkerbellFlags(upgradeClusterCmd.Flags(), uc.providerOptions.Tinkerbell.BMCOptions.RPC)
}

// nolint:gocyclo
func (uc *upgradeClusterOptions) upgradeCluster(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (uc *upgradeClusterOptions) commonValidations(ctx context.Context) (cluster *v1alpha1.Cluster, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
