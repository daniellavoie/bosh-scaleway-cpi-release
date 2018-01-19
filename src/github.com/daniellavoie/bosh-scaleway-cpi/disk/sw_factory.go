package disk

import (
	"path/filepath"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

type SWFactory struct {
	dirPath string

	scaleway *api.ScalewayAPI

	logTag string
	logger boshlog.Logger
}

func NewSWFactory(
	dirPath string,
	scaleway *api.ScalewayAPI,
	uuidGen boshuuid.Generator,
	cmdRunner boshsys.CmdRunner,
	logger boshlog.Logger,
) SWFactory {
	return SWFactory{
		dirPath: dirPath,

		scaleway: scaleway,

		logTag: "disk.FSFactory",
		logger: logger,
	}
}

func (f SWFactory) Create(size int) (Disk, error) {
	// TODO
	return nil, nil
}

func (f SWFactory) Find(id apiv1.DiskCID) (Disk, error) {
	return NewFSDisk(id, filepath.Join(f.dirPath, id.AsString()), f.scaleway, f.logger), nil
}
