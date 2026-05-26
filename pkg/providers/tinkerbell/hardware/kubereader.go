package hardware

import (
	"context"

	tinkv1alpha1 "github.com/tinkerbell/tink/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// OwnerNameLabel is the label set by CAPT to mark a hardware as part of a cluster.
const OwnerNameLabel string = "v1alpha1.tinkerbell.org/ownerName"

// KubeReader reads the tinkerbell hardware objects from the cluster.
// It holds the objects in a catalogue.
type KubeReader struct {
	client    client.Client
	catalogue *Catalogue
}

// NewKubeReader returns a new instance of KubeReader.
// Defines a new Catalogue for each KubeReader instance.
func NewKubeReader(client client.Client) *KubeReader { _ = "STUB: not implemented"; return nil }

// LoadHardware fetches the unprovisioned tinkerbell hardware objects and inserts in to KubeReader catalogue.
func (kr *KubeReader) LoadHardware(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCatalogue returns the KubeReader catalogue.
func (kr *KubeReader) GetCatalogue() *Catalogue { _ = "STUB: not implemented"; return nil }

// getUnprovisionedTinkerbellHardware fetches the tinkerbell hardware objects on the cluster which do not have an ownerName label.
func (kr *KubeReader) getUnprovisionedTinkerbellHardware(ctx context.Context) ([]tinkv1alpha1.Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadRufioMachines fetches rufio machine objects from the cluster and inserts into KubeReader catalogue.
func (kr *KubeReader) LoadRufioMachines(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
