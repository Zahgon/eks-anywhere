package cleanup

import (
	"context"
	"time"

	"github.com/bmc-toolbox/bmclib/v2"
	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
)

const (
	cleanupRetries       = 5
	retryBackoff         = 10 * time.Second
	cloudstackNetworkVar = "T_CLOUDSTACK_NETWORK"
)

func CleanUpAwsTestResources(storageBucket string, maxAge string, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

// one week

func CleanUpVsphereTestResources(ctx context.Context, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func VsphereRmVms(ctx context.Context, clusterName string, opts ...executables.GovcOpt) error {
	_ = "STUB: not implemented"
	return nil
}

// CloudstackTestResources cleans up resources on the CloudStack environment.
// This can include VMs as well as duplicate networks.
func CloudstackTestResources(ctx context.Context, clusterName string, dryRun bool, deleteDuplicateNetworks bool) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanupCloudstackDuplicateNetworks(ctx context.Context, cmk *executables.Cmk, execConfig *decoder.CloudStackExecConfig, deleteDuplicateNetworks bool) error {
	_ = "STUB: not implemented"
	return nil
}

// NutanixTestResources cleans up any leftover VMs in Nutanix after a test run.
func NutanixTestResources(clusterName, endpoint, port string, insecure, ignoreErrors bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TinkerbellTestResources cleans up machines by powering them down.
func TinkerbellTestResources(inventoryCSVFilePath string, ignoreErrors bool) error {
	_ = "STUB: not implemented"
	return nil
}

func powerOffHardwarePool(hardware map[string]*hardware.Machine, ignoreErrors bool) error {
	_ = "STUB: not implemented"
	return nil
}

func powerOffHardware(h *hardware.Machine, ignoreErrors bool) (reterror error) {
	_ = "STUB: not implemented"
	return nil
}

func handlePowerOffHardwareError(err error, ignoreErrors bool) error {
	_ = "STUB: not implemented"
	return nil
}

// newBmclibClient creates a new BMClib client.
func newBmclibClient(log logr.Logger, hostIP, username, password string) *bmclib.Client {
	_ = "STUB: not implemented"
	return nil
}
