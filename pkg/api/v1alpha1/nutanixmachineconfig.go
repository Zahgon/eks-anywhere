package v1alpha1

// NutanixIdentifierType is an enumeration of different resource identifier types.
type NutanixIdentifierType string

// NutanixGPUIdentifierType is an enumeration of different GPU identifier types.
type NutanixGPUIdentifierType string

// NutanixBootType is an enumeration of different boot types.
type NutanixBootType string

func (c NutanixIdentifierType) String() string { _ = "STUB: not implemented"; return "" }

func (c NutanixGPUIdentifierType) String() string { _ = "STUB: not implemented"; return "" }

const (
	// NutanixMachineConfigKind is the kind for a NutanixMachineConfig.
	NutanixMachineConfigKind = "NutanixMachineConfig"

	// NutanixIdentifierUUID is a resource identifier identifying the object by UUID.
	NutanixIdentifierUUID NutanixIdentifierType = "uuid"
	// NutanixIdentifierName is a resource identifier identifying the object by Name.
	NutanixIdentifierName NutanixIdentifierType = "name"

	// NutanixGPUIdentifierDeviceID is a GPU identifier identifying the object by DeviceID.
	NutanixGPUIdentifierDeviceID NutanixGPUIdentifierType = "deviceID"
	// NutanixGPUIdentifierName is a GPU identifier identifying the object by Name.
	NutanixGPUIdentifierName NutanixGPUIdentifierType = "name"

	// NutanixBootTypeLegacy is a resource identifier identifying the legacy boot type for virtual machines.
	NutanixBootTypeLegacy NutanixBootType = "legacy"

	// NutanixBootTypeUEFI is a resource identifier identifying the UEFI boot type for virtual machines.
	NutanixBootTypeUEFI NutanixBootType = "uefi"

	defaultNutanixOSFamily         = Ubuntu
	defaultNutanixSystemDiskSizeGi = "40Gi"
	defaultNutanixMemorySizeGi     = "4Gi"
	defaultNutanixVCPUsPerSocket   = 1
	defaultNutanixVCPUSockets      = 2

	// DefaultNutanixMachineConfigUser is the default username we set in machine config.
	DefaultNutanixMachineConfigUser string = "eksa"
)

// NutanixResourceIdentifier holds the identity of a Nutanix Prism resource (cluster, image, subnet, etc.)
//
// +union.
type NutanixResourceIdentifier struct {
	// Type is the identifier type to use for this resource.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=uuid;name
	Type NutanixIdentifierType `json:"type"`

	// uuid is the UUID of the resource in the PC.
	// +optional
	UUID *string `json:"uuid,omitempty"`

	// name is the resource name in the PC
	// +optional
	Name *string `json:"name,omitempty"`
}

// NutanixCategoryIdentifier holds the identity of a Nutanix Prism Central category.
type NutanixCategoryIdentifier struct {
	// key is the Key of the category in the Prism Central.
	// +kubebuilder:validation:Required
	Key string `json:"key,omitempty"`

	// value is the category value linked to the key in the Prism Central.
	// +kubebuilder:validation:Required
	Value string `json:"value,omitempty"`
}

// NutanixGPUIdentifier holds VM GPU device configuration.
type NutanixGPUIdentifier struct {
	// deviceID is the device ID of the GPU device.
	// +optional
	DeviceID *int64 `json:"deviceID,omitempty"`

	// vendorID is the vendor ID of the GPU device.
	// +optional
	Name string `json:"name,omitempty"`

	// type is the type of the GPU device.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=deviceID;name
	Type NutanixGPUIdentifierType `json:"type"`
}

// NutanixMachineConfigGenerateOpt is a functional option that can be passed to NewNutanixMachineConfigGenerate to
// customize the generated machine config
//
// +kubebuilder:object:generate=false
type NutanixMachineConfigGenerateOpt func(config *NutanixMachineConfigGenerate)

// NewNutanixMachineConfigGenerate returns a new instance of NutanixMachineConfigGenerate
// used for generating yaml for generate clusterconfig command.
func NewNutanixMachineConfigGenerate(name string, opts ...NutanixMachineConfigGenerateOpt) *NutanixMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *NutanixMachineConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *NutanixMachineConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *NutanixMachineConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func setNutanixMachineConfigDefaults(machineConfig *NutanixMachineConfig) {
	_ = "STUB: not implemented"
	return
}

func validateNutanixMachineConfig(c *NutanixMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMinimumNutanixMachineSpecs(c *NutanixMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNutanixReferences(c *NutanixMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNutanixResourceReference(i *NutanixResourceIdentifier, resource string, mcName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNutanixCategorySlice(i []NutanixCategoryIdentifier, mcName string) error {
	_ = "STUB: not implemented"
	return nil
}
