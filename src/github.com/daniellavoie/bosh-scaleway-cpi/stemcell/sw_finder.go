package stemcell

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"errors"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

type SWFinder struct {
	imageName string

	scaleway *api.ScalewayAPI

	logger boshlog.Logger
}

func NewSWFinder(imageName string, scaleway *api.ScalewayAPI, logger boshlog.Logger) SWFinder {
	return SWFinder{imageName: imageName, scaleway: scaleway, logger: logger}
}

func (f SWFinder) Find(id apiv1.StemcellCID) (Stemcell, bool, error) {
	return nil, false, errors.New("not implemented yet")
}
