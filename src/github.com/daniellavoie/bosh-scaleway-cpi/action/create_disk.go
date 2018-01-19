package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscdisk "github.com/daniellavoie/bosh-scaleway-cpi/disk"
)

type CreateDiskMethod struct {
	diskCreator bscdisk.Creator
}

func NewCreateDiskMethod(diskCreator bscdisk.Creator) CreateDiskMethod {
	return CreateDiskMethod{diskCreator: diskCreator}
}

func (a CreateDiskMethod) CreateDisk(size int, _ apiv1.DiskCloudProps, _ *apiv1.VMCID) (apiv1.DiskCID, error) {
	disk, err := a.diskCreator.Create(size)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Creating disk of size '%d'", size)
	}

	return disk.ID(), nil
}
