package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bscstem "github.com/daniellavoie/bosh-scaleway-cpi/stemcell"
	"errors"
	"github.com/scaleway/scaleway-cli/pkg/api"
	"fmt"
	"github.com/docker/docker/cli/compose/loader"
)

type SWCreator struct {
	uuidGen  boshuuid.Generator
	scaleway *api.ScalewayAPI

	agentOptions apiv1.AgentOptions
	logger       boshlog.Logger
}

func NewSWCreator(
	uuidGen boshuuid.Generator,
	scaleway *api.ScalewayAPI,
	agentOptions apiv1.AgentOptions,
	logger boshlog.Logger,
) SWCreator {
	return SWCreator{
		uuidGen:      uuidGen,
		scaleway:     scaleway,
		agentOptions: agentOptions,
		logger:       logger,
	}
}

func (c SWCreator) Create(
	agentID apiv1.AgentID, stemcell bscstem.Stemcell, props VMProps,
	networks apiv1.Networks, env apiv1.VMEnv) (VM, error) {

	image, err := c.getUbuntuImageId()
	if err != nil {
		return nil, err
	}

	serverId, err := c.createServer(image)
	if err != nil {
		return nil, err
	}



	return nil, errors.New("not mplemented yet")
}

func (c SWCreator) createServer(name string, imageGuid string) (string, error) {
	definition := api.ScalewayServerDefinition{Name: name, Image: imageGuid, CommercialType: }

	c.scaleway.PostServer(definition)
}

func (c SWCreator) getUbuntuImageId() (string, error) {
	os := "Ubuntu Trusty"
	arch := "x86_64"

	images, err := c.scaleway.GetImages()
	if err != nil {
		return "", err
	}

	for _, image := range *images {
		if image.Name == os {
			for _, version := range image.Versions {
				for _, localImage := range version.LocalImages {
					if localImage.Arch == arch {
						return localImage.ID, nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("could not found an image for %s:%s", os, arch)
}

func (c SWCreator) resolveNetworkIPCIDR(networks apiv1.Networks) (string, error) {
	return "", errors.New("not implemented yet")
}
