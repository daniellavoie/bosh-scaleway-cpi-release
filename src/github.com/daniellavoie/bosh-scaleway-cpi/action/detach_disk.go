package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscdisk "github.com/daniellavoie/bosh-scaleway-cpi/disk"
	bscvm "github.com/daniellavoie/bosh-scaleway-cpi/vm"
)

type DetachDiskMethod struct {
	vmFinder   bscvm.Finder
	diskFinder bscdisk.Finder
}

func NewDetachDiskMethod(vmFinder bscvm.Finder, diskFinder bscdisk.Finder) DetachDiskMethod {
	return DetachDiskMethod{vmFinder, diskFinder}
}

func (a DetachDiskMethod) DetachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	vm, found, err := a.vmFinder.Find(vmCID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding VM '%s'", vmCID)
	}

	if !found {
		return bosherr.Errorf("Expected to find VM '%s'", vmCID)
	}

	disk, err := a.diskFinder.Find(diskCID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding disk '%s'", diskCID)
	}

	err = vm.DetachDisk(disk)
	if err != nil {
		return bosherr.WrapErrorf(err, "Detaching disk '%s' to VM '%s'", diskCID, vmCID)
	}

	return nil
}
