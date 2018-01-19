package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscdisk "github.com/daniellavoie/bosh-scaleway-cpi/disk"
)

type HasDiskMethod struct {
	diskFinder bscdisk.Finder
}

func NewHasDiskMethod(diskFinder bscdisk.Finder) HasDiskMethod {
	return HasDiskMethod{diskFinder: diskFinder}
}

func (a HasDiskMethod) HasDisk(cid apiv1.DiskCID) (bool, error) {
	disk, err := a.diskFinder.Find(cid)
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding disk '%s'", cid)
	}

	return disk.Exists()
}
