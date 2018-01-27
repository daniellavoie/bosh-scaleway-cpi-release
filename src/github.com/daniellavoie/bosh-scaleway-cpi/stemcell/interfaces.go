package stemcell

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type Importer interface {
	ImportFromPath(imagePath string) (Stemcell, error)
}

type Finder interface {
	Find(apiv1.StemcellCID) (Stemcell, bool, error)
}

type Stemcell interface {
	ID() apiv1.StemcellCID
	ImageName() string

	Delete() error
}