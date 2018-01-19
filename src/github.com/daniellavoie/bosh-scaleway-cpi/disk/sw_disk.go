package disk

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

type SWDisk struct {
	id   apiv1.DiskCID
	path string

	scaleway *api.ScalewayAPI

	logger boshlog.Logger
}

func NewFSDisk(
	id apiv1.DiskCID,
	path string,

	scaleway *api.ScalewayAPI,

	logger boshlog.Logger,
) SWDisk {
	return SWDisk{id: id, path: path, scaleway: scaleway}
}

func (s SWDisk) ID() apiv1.DiskCID { return s.id }

func (s SWDisk) Path() string { return s.path }

func (s SWDisk) Exists() (bool, error) {
	// TODO
	return false, nil
}

func (s SWDisk) Delete() error {
	// TODO
	return nil
}
