package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type createSupportBundleOptions struct {
	fileName              string
	wConfig               string
	since                 string
	sinceTime             string
	bundleConfig          string
	hardwareFileName      string
	tinkerbellBootstrapIP string
	bundlesManifest       string
	auditLogs             bool
}

var csbo = &createSupportBundleOptions{}

var supportbundleCmd = &cobra.Command{
	Use:          "support-bundle -f my-cluster.yaml",
	Short:        "Generate a support bundle",
	Long:         "This command is used to create a support bundle to troubleshoot a cluster",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := csbo.validate(cmd.Context()); err != nil {
			return err
		}
		if err := csbo.createBundle(cmd.Context(), csbo.since, csbo.sinceTime, csbo.bundleConfig, csbo.bundlesManifest, csbo.auditLogs); err != nil {
			return fmt.Errorf("failed to create support bundle: %v", err)
		}
		return nil
	},
}

func init() {
	generateCmd.AddCommand(supportbundleCmd)
	supportbundleCmd.Flags().StringVarP(&csbo.sinceTime, "since-time", "", "", "Collect pod logs after a specific datetime(RFC3339) like 2021-06-28T15:04:05Z")
	supportbundleCmd.Flags().StringVarP(&csbo.since, "since", "", "", "Collect pod logs in the latest duration like 5s, 2m, or 3h.")
	supportbundleCmd.Flags().StringVarP(&csbo.bundleConfig, "bundle-config", "", "", "Bundle Config file to use when generating support bundle")
	supportbundleCmd.Flags().StringVarP(&csbo.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	supportbundleCmd.Flags().StringVarP(&csbo.wConfig, "w-config", "w", "", "Kubeconfig file to use when creating support bundle for a workload cluster")
	supportbundleCmd.Flags().StringVarP(&csbo.bundlesManifest, "bundles-manifest", "", "", "Bundles manifest to use when generating support bundle (required for generating support bundle in airgap environment)")
	supportbundleCmd.Flags().BoolVarP(&csbo.auditLogs, "audit-logs", "", false, "Include the latest api server audit log file in the support bundle")
	err := supportbundleCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func (csbo *createSupportBundleOptions) validate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (csbo *createSupportBundleOptions) createBundle(ctx context.Context, since, sinceTime, bundleConfig string, bundlesManifest string, auditLogs bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Also collect management cluster bundle for workload cluster debugging
// This provides CAPI resources and controller logs which are essential for debugging

// This is a self-managed cluster, skip management cluster bundle collection

// Check if management cluster kubeconfig exists
