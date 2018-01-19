package stemcell

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type SWStemcell struct {
	id      apiv1.StemcellCID
	dirPath string

	logger boshlog.Logger
}

func NewSWStemcell(
	id apiv1.StemcellCID,
	dirPath string,
	logger boshlog.Logger,
) SWStemcell {
	return SWStemcell{id: id, dirPath: dirPath, logger: logger}
}

func (s SWStemcell) ID() apiv1.StemcellCID { return s.id }

func (s SWStemcell) DirPath() string { return s.dirPath }

func (s SWStemcell) Delete() error {
	s.logger.Debug("SWStemcell", "Deleting stemcell '%s'", s.id)

	// TODO

	return nil
}
