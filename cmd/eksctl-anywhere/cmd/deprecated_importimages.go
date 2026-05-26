package cmd

//////////////////////////////////////////////////////
//
// WARNING: The command defined in this file is DEPRECATED.
//
// See ./import_images.go for the newer command.
//
//////////////////////////////////////////////////////

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/helm"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type importImagesOptions struct {
	fileName string
}

var opts = &importImagesOptions{}

const ociPrefix = "oci://"

func init() {
	rootCmd.AddCommand(importImagesCmdDeprecated)
	importImagesCmdDeprecated.Flags().StringVarP(&opts.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	err := importImagesCmdDeprecated.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking filename flag as required: %v", err)
	}
}

var importImagesCmdDeprecated = &cobra.Command{
	Use:          "import-images",
	Short:        "Push EKS Anywhere images to a private registry (Deprecated)",
	Long:         "This command is used to import images from an EKS Anywhere release bundle into a private registry",
	PreRunE:      preRunImportImagesCmd,
	SilenceUsage: true,
	Deprecated:   "use `eksctl anywhere import images` instead",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := importImages(cmd.Context(), opts.fileName); err != nil {
			return err
		}
		return nil
	},
}

//gocyclo:ignore
func importImages(ctx context.Context, clusterSpecPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func importImage(ctx context.Context, docker *executables.Docker, image string, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func importCharts(ctx context.Context, helm helm.Client, charts map[string]*v1alpha1.Image, endpoint, username, password string) error {
	_ = "STUB: not implemented"
	return nil
}

func importChart(ctx context.Context, helm helm.Client, chart v1alpha1.Image, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func preRunImportImagesCmd(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func getChartUriAndVersion(chart v1alpha1.Image) (uri, version string) {
	_ = "STUB: not implemented"
	return "", ""
}

func pushChartURI(chart v1alpha1.Image, registryEndpoint string) string {
	_ = "STUB: not implemented"
	return ""
}
