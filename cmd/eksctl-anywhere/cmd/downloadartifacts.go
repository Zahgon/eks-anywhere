package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/files"
)

type downloadArtifactsOptions struct {
	downloadDir     string
	fileName        string
	bundlesOverride string
	dryRun          bool
	retainDir       bool
}

var downloadArtifactsopts = &downloadArtifactsOptions{}

func init() {
	downloadCmd.AddCommand(downloadArtifactsCmd)
	downloadArtifactsCmd.Flags().StringVarP(&downloadArtifactsopts.bundlesOverride, "bundles-override", "", "", "Override default Bundles manifest (not recommended)")
	downloadArtifactsCmd.Flags().StringVarP(&downloadArtifactsopts.fileName, "filename", "f", "", "[Deprecated] Filename that contains EKS-A cluster configuration")
	downloadArtifactsCmd.Flags().StringVarP(&downloadArtifactsopts.downloadDir, "download-dir", "d", "eks-anywhere-downloads", "Directory to download the artifacts to")
	downloadArtifactsCmd.Flags().BoolVarP(&downloadArtifactsopts.dryRun, "dry-run", "", false, "Print the manifest URIs without downloading them")
	downloadArtifactsCmd.Flags().BoolVarP(&downloadArtifactsopts.retainDir, "retain-dir", "r", false, "Do not delete the download folder after creating a tarball")
}

var downloadArtifactsCmd = &cobra.Command{
	Use:          "artifacts",
	Short:        "Download EKS Anywhere artifacts/manifests to a tarball on disk",
	Long:         "This command is used to download the S3 artifacts from an EKS Anywhere bundle manifest and package them into a tarball",
	PreRunE:      preRunDownloadArtifactsCmd,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := downloadArtifacts(cmd.Context(), downloadArtifactsopts); err != nil {
			return err
		}
		return nil
	},
}

func downloadArtifacts(context context.Context, opts *downloadArtifactsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// download the eks-a-release.yaml

// This can happen if the provider is not GA and not added to the bundle-release corresponding to an EKS-A release

func preRunDownloadArtifactsCmd(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadArtifact(filePath, artifactUri string, reader *files.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func createTarball(downloadDir string) error { _ = "STUB: not implemented"; return nil }
