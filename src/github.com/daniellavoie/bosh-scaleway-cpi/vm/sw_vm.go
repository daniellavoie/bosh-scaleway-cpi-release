package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscdisk "github.com/daniellavoie/bosh-scaleway-cpi/disk"
	"errors"
)

type ScalewayVM struct {
	id     apiv1.VMCID
	logger boshlog.Logger
}

func NewScalewayVM(
	id apiv1.VMCID,
	logger boshlog.Logger,
) ScalewayVM {
	return ScalewayVM{
		id:     id,
		logger: logger,
	}
}

func (vm ScalewayVM) ID() apiv1.VMCID { return vm.id }

func (vm ScalewayVM) Delete() error {
	return errors.New("not implemented yet")
}

func (vm ScalewayVM) AttachDisk(disk bscdisk.Disk) error {
	return errors.New("not implemented yet")
}

func (vm ScalewayVM) DetachDisk(disk bscdisk.Disk) error {
	return errors.New("not implemented yet")
}
