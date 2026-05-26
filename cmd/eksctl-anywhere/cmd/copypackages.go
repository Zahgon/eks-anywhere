package cmd

import (
	"context"
	"fmt"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"

	packagesv1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/logger"
	"github.com/aws/eks-anywhere/pkg/registry"
)

// copyPackagesCmd is the context for the copy packages command.
var copyPackagesCmd = &cobra.Command{
	Use:          "packages <destination-registry>",
	Short:        "Copy curated package images and charts from source registries to a destination registry",
	Long:         `Copy all the EKS Anywhere curated package images and helm charts from source registries to a destination registry. Registry credentials are fetched from docker config.`,
	SilenceUsage: true,
	RunE:         runCopyPackages,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("A destination registry must be specified as an argument")
		}
		return nil
	},
}

var cs = registry.NewCredentialStore()

func init() {
	copyCmd.AddCommand(copyPackagesCmd)

	copyPackagesCmd.Flags().StringVar(&cpc.srcImageRegistry, "src-image-registry", "", "The source registry that stores container images")
	if err := copyPackagesCmd.MarkFlagRequired("src-image-registry"); err != nil {
		logger.Fatal(err, "Cannot mark flag as required")
	}
	copyPackagesCmd.Flags().StringVar(&cpc.kubeVersion, "kube-version", "", "The kubernetes version of the package bundle to copy")
	if err := copyPackagesCmd.MarkFlagRequired("kube-version"); err != nil {
		logger.Fatal(err, "Cannot mark flag as required")
	}
	copyPackagesCmd.Flags().StringVar(&cpc.srcChartRegistry, "src-chart-registry", "", "The source registry that stores helm charts (default src-image-registry)")
	copyPackagesCmd.Flags().BoolVar(&cpc.dstPlainHTTP, "dst-plain-http", false, "Whether or not to use plain http for destination registry")
	copyPackagesCmd.Flags().BoolVar(&cpc.dstInsecure, "dst-insecure", false, "Skip TLS verification against the destination registry")
	copyPackagesCmd.Flags().BoolVar(&cpc.dryRun, "dry-run", false, "Dry run will show what artifacts would be copied, but not actually copy them")

	// making oras client to use dockerconfig
	if err := cs.Init(); err != nil {
		panic(err)
	}
	auth.DefaultClient.Credential = func(ctx context.Context, registry string) (auth.Credential, error) {
		return cs.Credential(registry)
	}
}

var cpc = copyPackagesConfig{}

var publicPackages = []string{"ecr-token-refresher", "eks-anywhere-packages", "credential-provider-package"}

// copyPackagesConfig copies packages specified in a bundle to a destination.
type copyPackagesConfig struct {
	destRegistry     string
	srcImageRegistry string
	srcChartRegistry string
	kubeVersion      string
	dstPlainHTTP     bool
	dstInsecure      bool
	dryRun           bool
}

func runCopyPackages(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// copy package bundle yaml after charts and images

func getTagsFromChartValues(chartValues map[string]any, res map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func getPackageBundleTag(kubeVersion string) string { _ = "STUB: not implemented"; return "" }

func getPackageBundle(ctx context.Context, registry, kubeVersion string) (*packagesv1.PackageBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyArtifacts(ctx context.Context, bundle *packagesv1.PackageBundle) error {
	_ = "STUB: not implemented"
	return nil
}

func copyImages(ctx context.Context, images []packagesv1.VersionImages, tags map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func getChartValues(chartURL string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func orasCopy(ctx context.Context, repo, srcRegistry, srcRef, dstRegistry, dstRef string) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

func setUpDstRepo(dst *remote.Repository, c *copyPackagesConfig) { _ = "STUB: not implemented"; return }
