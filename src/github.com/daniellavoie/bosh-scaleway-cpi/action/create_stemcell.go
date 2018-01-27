package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscstem "github.com/daniellavoie/bosh-scaleway-cpi/stemcell"
)

type CreateStemcellMethod struct {
	stemcellImporter bscstem.Importer
}

func NewCreateStemcellMethod(stemcellImporter bscstem.Importer) CreateStemcellMethod {
	return CreateStemcellMethod{stemcellImporter: stemcellImporter}
}

func (a CreateStemcellMethod) CreateStemcell(imagePath string, _ apiv1.StemcellCloudProps) (apiv1.StemcellCID, error) {
	stemcell, err := a.stemcellImporter.ImportFromPath(imagePath)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "Importing stemcell from '%s'", imagePath)
	}

	return stemcell.ID(), nil
}
