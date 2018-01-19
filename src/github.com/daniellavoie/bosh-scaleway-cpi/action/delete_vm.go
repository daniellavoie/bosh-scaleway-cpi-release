package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscvm "github.com/daniellavoie/bosh-scaleway-cpi/vm"
)

type DeleteVMMethod struct {
	vmFinder bscvm.Finder
}

func NewDeleteVMMethod(vmFinder bscvm.Finder) DeleteVMMethod {
	return DeleteVMMethod{vmFinder: vmFinder}
}

func (a DeleteVMMethod) DeleteVM(cid apiv1.VMCID) error {
	vm, _, err := a.vmFinder.Find(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding vm '%s'", cid)
	}

	err = vm.Delete()
	if err != nil {
		return bosherr.WrapErrorf(err, "Deleting vm '%s'", cid)
	}

	return nil
}
