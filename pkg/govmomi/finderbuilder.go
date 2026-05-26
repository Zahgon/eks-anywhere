package govmomi

import (
	"github.com/vmware/govmomi/vim25"
)

type vMOMIFinderBuilder struct{}

func NewVMOMIFinderBuilder() *vMOMIFinderBuilder { _ = "STUB: not implemented"; return nil }

func (*vMOMIFinderBuilder) Build(client *vim25.Client, all ...bool) VMOMIFinder {
	_ = "STUB: not implemented"
	return *new(VMOMIFinder)
}
