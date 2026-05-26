package features

// These are environment variables used as flags to enable/disable features.
const (
	CloudStackKubeVipDisabledEnvVar = "CLOUDSTACK_KUBE_VIP_DISABLED"
	CheckpointEnabledEnvVar         = "CHECKPOINT_ENABLED"
	UseControllerForCli             = "USE_CONTROLLER_FOR_CLI"
	VSphereInPlaceEnvVar            = "VSPHERE_IN_PLACE_UPGRADE"
	APIServerExtraArgsEnabledEnvVar = "API_SERVER_EXTRA_ARGS_ENABLED"
)

func FeedGates(featureGates []string) { _ = "STUB: not implemented"; return }

type Feature struct {
	Name     string
	IsActive func() bool
}

func IsActive(feature Feature) bool { _ = "STUB: not implemented"; return false }

// ClearCache is mainly used for unit tests as of now.
func ClearCache() { _ = "STUB: not implemented"; return }

func CloudStackKubeVipDisabled() Feature { _ = "STUB: not implemented"; return *new(Feature) }

func CheckpointEnabled() Feature { _ = "STUB: not implemented"; return *new(Feature) }

// VSphereInPlaceUpgradeEnabled is the feature flag for performing in-place upgrades with the vSphere provider.
func VSphereInPlaceUpgradeEnabled() Feature { _ = "STUB: not implemented"; return *new(Feature) }

// APIServerExtraArgsEnabled is the feature flag for configuring api server extra args.
func APIServerExtraArgsEnabled() Feature { _ = "STUB: not implemented"; return *new(Feature) }
