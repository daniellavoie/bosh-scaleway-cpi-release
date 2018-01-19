package action

import (
	"strings"

	bscvm "github.com/daniellavoie/bosh-scaleway-cpi/vm"
)

type VMCloudProperties struct {
	instanceType string `json:"instance_type"`
}

type alwaysString string

func (s *alwaysString) UnmarshalJSON(data []byte) error {
	*s = alwaysString(strings.TrimPrefix(strings.TrimSuffix(string(data), `"`), `"`))
	return nil
}

func (cp VMCloudProperties) AsVMProps() (bscvm.VMProps, error) {
	return bscvm.VMProps{}, nil
}
