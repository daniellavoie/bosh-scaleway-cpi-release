package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscdisk "github.com/daniellavoie/bosh-scaleway-cpi/disk"
)

type DeleteDiskMethod struct {
	diskFinder bscdisk.Finder
}

func NewDeleteDiskMethod(diskFinder bscdisk.Finder) DeleteDiskMethod {
	return DeleteDiskMethod{diskFinder: diskFinder}
}

func (a DeleteDiskMethod) DeleteDisk(cid apiv1.DiskCID) error {
	disk, err := a.diskFinder.Find(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding disk '%s'", cid)
	}

	err = disk.Delete()
	if err != nil {
		return bosherr.WrapErrorf(err, "Deleting disk '%s'", cid)
	}

	return nil
}
