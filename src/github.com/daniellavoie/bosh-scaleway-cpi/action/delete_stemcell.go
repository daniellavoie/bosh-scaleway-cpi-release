package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscstem "github.com/daniellavoie/bosh-scaleway-cpi/stemcell"
)

type DeleteStemcellMethod struct {
	stemcellFinder bscstem.Finder
}

func NewDeleteStemcellMethod(stemcellFinder bscstem.Finder) DeleteStemcellMethod {
	return DeleteStemcellMethod{stemcellFinder: stemcellFinder}
}

func (a DeleteStemcellMethod) DeleteStemcell(cid apiv1.StemcellCID) error {
	stemcell, found, err := a.stemcellFinder.Find(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding stemcell '%s'", cid)
	}

	if found {
		err := stemcell.Delete()
		if err != nil {
			return bosherr.WrapErrorf(err, "Deleting stemcell '%s'", cid)
		}
	}

	return nil
}
