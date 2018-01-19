package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bscvm "github.com/daniellavoie/bosh-scaleway-cpi/vm"
)

type GetDisksMethod struct {
	vmFinder bscvm.Finder
}

func NewGetDisksMethod(vmFinder bscvm.Finder) GetDisksMethod {
	return GetDisksMethod{vmFinder}
}

func (a GetDisksMethod) GetDisks(cid apiv1.VMCID) ([]apiv1.DiskCID, error) {
	// todo implement
	return nil, nil
}
