package stemcell

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type SWImporter struct {
	dirPath string

	logTag string
	logger boshlog.Logger
}

func NewSWImporter(
	dirPath string,
	logger boshlog.Logger,
) SWImporter {
	return SWImporter{
		dirPath: dirPath,

		logTag: "SWImporter",
		logger: logger,
	}
}

func (i SWImporter) ImportFromPath(imagePath string) (Stemcell, error) {
	// TODO
	return nil, nil
}
