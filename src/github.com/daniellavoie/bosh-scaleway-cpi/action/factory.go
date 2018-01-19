package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscdisk "github.com/daniellavoie/bosh-scaleway-cpi/disk"
	bscstem "github.com/daniellavoie/bosh-scaleway-cpi/stemcell"
	bscvm "github.com/daniellavoie/bosh-scaleway-cpi/vm"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

type IaasFactory struct {
	stemcellImporter bscstem.Importer
	stemcellFinder   bscstem.Finder

	vmCreator bscvm.Creator
	vmFinder  bscvm.Finder

	diskCreator bscdisk.Creator
	diskFinder  bscdisk.Finder
}

type CPI struct {
	InfoMethod

	CreateStemcellMethod
	DeleteStemcellMethod

	CreateVMMethod
	DeleteVMMethod
	CalculateVMCloudPropertiesMethod
	HasVMMethod
	RebootVMMethod
	SetVMMetadataMethod
	GetDisksMethod

	CreateDiskMethod
	DeleteDiskMethod
	AttachDiskMethod
	DetachDiskMethod
	HasDiskMethod
}

func NewFactory(
	cmdRunner boshsys.CmdRunner,
	uuidGen boshuuid.Generator,
	scaleway *api.ScalewayAPI,
	opts FactoryOpts,
	logger boshlog.Logger,
) IaasFactory {
	stemcellImporter := bscstem.NewSWImporter(
		opts.StemcellsDir, logger)

	stemcellFinder := bscstem.NewSWFinder(opts.StemcellsDir, logger)
	vmCreator := bscvm.NewSWCreator(
		uuidGen, scaleway, opts.Agent, logger)

	vmFinder := bscvm.NewSWFinder(scaleway, logger)

	diskFactory := bscdisk.NewSWFactory(opts.DisksDir, scaleway, uuidGen, cmdRunner, logger)

	return IaasFactory{
		stemcellImporter,
		stemcellFinder,
		vmCreator,
		vmFinder,
		diskFactory,
		diskFactory,
	}
}

func (f IaasFactory) New(_ apiv1.CallContext) (apiv1.CPI, error) {
	return CPI{
		NewInfoMethod(),

		NewCreateStemcellMethod(f.stemcellImporter),
		NewDeleteStemcellMethod(f.stemcellFinder),

		NewCreateVMMethod(f.stemcellFinder, f.vmCreator),
		NewDeleteVMMethod(f.vmFinder),
		NewCalculateVMCloudPropertiesMethod(),
		NewHasVMMethod(f.vmFinder),
		NewRebootVMMethod(),
		NewSetVMMetadataMethod(),
		NewGetDisksMethod(f.vmFinder),

		NewCreateDiskMethod(f.diskCreator),
		NewDeleteDiskMethod(f.diskFinder),
		NewAttachDiskMethod(f.vmFinder, f.diskFinder),
		NewDetachDiskMethod(f.vmFinder, f.diskFinder),
		NewHasDiskMethod(f.diskFinder),
	}, nil
}
