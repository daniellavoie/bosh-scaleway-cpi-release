package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

type SWFinder struct {

	scaleway *api.ScalewayAPI
	logTag string
	logger boshlog.Logger
}

func NewSWFinder(
	scaleway *api.ScalewayAPI,
	logger boshlog.Logger,
) SWFinder {
	return SWFinder{
		scaleway: scaleway,
		logTag: "vm.VMFinder",
		logger: logger,
	}
}

func (f SWFinder) Find(id apiv1.VMCID) (VM, bool, error) {
	// TODO - Implement
	return nil, false, nil
}
