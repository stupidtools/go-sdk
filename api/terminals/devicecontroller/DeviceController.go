package devicecontroller

import (
	"context"
	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
)

type IDeviceController interface {
	GetConnectionModes() connectionmodes.ConnectionModes
	GetDeviceType() devicetype.DeviceType
	Send(message abstractions.IDeviceMessage) ([]byte, error)
	ConfigureInterface() abstractions.IDeviceInterface
	ProcessTransaction(builder *builders.TerminalAuthBuilder) (terminalresponse.ITerminalResponse, error)
	ManageTransaction(builder *builders.TerminalManageBuilder) (terminalresponse.ITerminalResponse, error)
	ProcessTransactionWithContext(ctx context.Context, builder *builders.TerminalAuthBuilder) (terminalresponse.ITerminalResponse, error)
	ManageTransactionWithContext(ctx context.Context, builder *builders.TerminalManageBuilder) (terminalresponse.ITerminalResponse, error)
	GetSettings() abstractions.ITerminalConfiguration
	SetSettings(settings abstractions.ITerminalConfiguration)
	GetDeviceCommInterface() abstractions.IDeviceCommInterface
	SetDeviceCommInterface(deviceInterface abstractions.IDeviceCommInterface)
	GetRequestIdProvider() abstractions.IRequestIdProvider
	SetRequestIdProvider(provider abstractions.IRequestIdProvider)
	Dispose()
}

//this type does not fully implement the interface.  it needs to be extended to do so
type DeviceController struct {
	connectionModes     connectionmodes.ConnectionModes
	deviceType          devicetype.DeviceType
	deviceCommInterface abstractions.IDeviceCommInterface
	settings            abstractions.ITerminalConfiguration
	requestIdProvider   abstractions.IRequestIdProvider
}

func NewDeviceController() *DeviceController {
	return &DeviceController{}
}

func (dc *DeviceController) GetConnectionModes() connectionmodes.ConnectionModes {
	return dc.connectionModes
}

func (dc *DeviceController) GetDeviceType() devicetype.DeviceType {
	return dc.deviceType
}

func (dc *DeviceController) SetDeviceType(d devicetype.DeviceType) {
	dc.deviceType = d
}

func (dc *DeviceController) GetDeviceCommInterface() abstractions.IDeviceCommInterface {
	return dc.deviceCommInterface
}

func (dc *DeviceController) SetDeviceCommInterface(deviceInterface abstractions.IDeviceCommInterface) {
	dc.deviceCommInterface = deviceInterface
}

func (dc *DeviceController) GetSettings() abstractions.ITerminalConfiguration {
	return dc.settings
}

func (dc *DeviceController) SetSettings(settings abstractions.ITerminalConfiguration) {
	dc.settings = settings
}

func (dc *DeviceController) GetRequestIdProvider() abstractions.IRequestIdProvider {
	return dc.requestIdProvider
}

func (dc *DeviceController) SetRequestIdProvider(provider abstractions.IRequestIdProvider) {
	dc.requestIdProvider = provider
}

func (dc *DeviceController) Dispose() {
	// implementation here
}
