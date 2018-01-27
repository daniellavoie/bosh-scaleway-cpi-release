package vm

import (
	"encoding/json"
	"fmt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type metadataService struct {
	agentEnvService  string
	registryOptions  RegistryOptions
	userDataFilePath string
	metadataFilePath string

	logTag string
	logger boshlog.Logger
}

func NewMetadataService(
	logger boshlog.Logger,
) MetadataService {
	return &metadataService{
		userDataFilePath: "/var/vcap/bosh/scaleway-cpi-user-data.json",
		metadataFilePath: "/var/vcap/bosh/scaleway-cpi-metadata.json",

		logTag: "vm.metadataService",
		logger: logger,
	}
}

type RegistryType struct {
	Endpoint string
}

type UserDataContentsType struct {
	Registry RegistryType
}

type MetadataContentsType struct {
	InstanceID string `json:"instance-id"`
}

func (ms *metadataService) Save(wardenFileService WardenFileService, instanceID apiv1.VMCID) error {
	var endpoint string

	if ms.agentEnvService == "registry" {
		endpoint = fmt.Sprintf(
			"http://%s:%s@%s:%d",
			ms.registryOptions.Username,
			ms.registryOptions.Password,
			ms.registryOptions.Host,
			ms.registryOptions.Port,
		)
	} else {
		endpoint = "/var/vcap/bosh/scaleway-cpi-agent-env.json"
	}

	userDataContents := UserDataContentsType{
		Registry: RegistryType{
			Endpoint: endpoint,
		},
	}

	jsonBytes, err := json.Marshal(userDataContents)
	if err != nil {
		return bosherr.WrapError(err, "Marshalling user data")
	}

	ms.logger.Debug(ms.logTag, "Saving user data to %s", ms.userDataFilePath)

	err = wardenFileService.Upload(ms.userDataFilePath, jsonBytes)
	if err != nil {
		return bosherr.WrapError(err, "Saving user data")
	}

	metadataContents := MetadataContentsType{
		InstanceID: instanceID.AsString(),
	}

	jsonBytes, err = json.Marshal(metadataContents)
	if err != nil {
		return bosherr.WrapError(err, "Marshalling metadata")
	}

	err = wardenFileService.Upload(ms.metadataFilePath, jsonBytes)
	if err != nil {
		return bosherr.WrapError(err, "Saving metadata")
	}

	return nil
}
