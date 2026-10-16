package configuredservices

import (
	abstractions2 "github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/devicecontroller"
)

type ConfiguredServices struct {
	deviceInterface  abstractions.IDeviceInterface
	deviceController devicecontroller.IDeviceController
	gatewayConnector abstractions2.IPaymentGateway
}

func (cs *ConfiguredServices) GetGateway() abstractions2.IPaymentGateway {
	return cs.gatewayConnector
}

func (cs *ConfiguredServices) SetGateway(gw abstractions2.IPaymentGateway) {
	cs.gatewayConnector = gw
}

func (cs *ConfiguredServices) GetDeviceInterface() abstractions.IDeviceInterface {
	return cs.deviceInterface
}

func (cs *ConfiguredServices) GetDeviceController() devicecontroller.IDeviceController {
	return cs.deviceController
}

func (cs *ConfiguredServices) SetDeviceController(deviceController devicecontroller.IDeviceController) error {
	cs.deviceController = deviceController
	deviceInterface := (deviceController).ConfigureInterface()
	cs.deviceInterface = deviceInterface
	return nil
}

func (cs *ConfiguredServices) Dispose() {
	(cs.deviceController).Dispose()
}
