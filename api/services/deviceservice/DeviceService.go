package deviceservice

import (
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/terminals"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
)

func DeviceServiceCreate(config *terminals.ConnectionConfig) (abstractions.IDeviceInterface, error) {
	return DeviceServiceCreateWithName(config, "default")
}

func DeviceServiceCreateWithName(config *terminals.ConnectionConfig, configName string) (abstractions.IDeviceInterface, error) {
	api.ConfigureService(config, configName)
	inst, err := api.GetServiceContainerInstance().GetDeviceInterface(configName)
	return inst, err
}
