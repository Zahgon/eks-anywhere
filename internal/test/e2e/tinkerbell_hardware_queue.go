package e2e

import (
	"sync"
	"time"

	"github.com/aws/eks-anywhere/internal/pkg/api"
)

// Default timeout for Tink tests to poll for hardware.
const hwPollingTimeout = 120 * time.Minute

// hardwareCatalogue has a thread safe FIFO queue implementation to facilitate hardware reservation.
type hardwareCatalogue struct {
	hws []*api.Hardware
	mu  sync.Mutex
}

func (hwQu *hardwareCatalogue) reserveHardware(count int) ([]*api.Hardware, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hwQu *hardwareCatalogue) releaseHardware(hws []*api.Hardware) {
	_ = "STUB: not implemented"
	return
}

func (hwQu *hardwareCatalogue) shuffleHardware() { _ = "STUB: not implemented"; return }

func newHardwareCatalogue(hws []*api.Hardware) *hardwareCatalogue {
	_ = "STUB: not implemented"
	return nil
}
