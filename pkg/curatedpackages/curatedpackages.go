package curatedpackages

import (
	"context"
	"strings"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere-packages/pkg/bundle"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	license = `The Amazon EKS Anywhere Curated Packages are only available to customers with the
Amazon EKS Anywhere Enterprise Subscription`
	width = 86
)

var userMsgSeparator = strings.Repeat("-", width)

// CreateBundleManager builds a new bundle Manager.
func CreateBundleManager(log logr.Logger) bundle.RegistryClient {
	_ = "STUB: not implemented"
	return *new(bundle.RegistryClient)
}

func parseKubeVersion(kubeVersion string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func GetVersionBundle(reader Reader, eksaVersion string, spec *v1alpha1.Cluster) (*releasev1.VersionsBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PrintLicense() {
	_ = "STUB: not implemented"
	// Currently, use the width of the longest line to repeat the dashes
	// Sample Output
	// -------------------------------------------------------------------------------------
	// The Amazon EKS Anywhere Curated Packages are only available to customers with the
	// Amazon EKS Anywhere Enterprise Subscription
	// -------------------------------------------------------------------------------------
	return
}

// PullLatestBundle reads the contents of the artifact using the latest bundle.
func PullLatestBundle(ctx context.Context, log logr.Logger, artifact string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PushBundle(ctx context.Context, ref, fileName string, fileContent []byte) error {
	_ = "STUB: not implemented"
	// Parse the reference to get repository
	return nil
}

// Configure repository with insecure option if needed

// Create memory store

// Create a descriptor with proper digest calculation

// Calculate the digest

// Create descriptor using OCI spec

// Store the content in memory

// Push the content to the registry

func GetRegistry(uri string) string { _ = "STUB: not implemented"; return "" }
