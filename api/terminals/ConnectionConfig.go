package terminals

import (
	"github.com/globalpayments/go-sdk/api/configuredservices"
	"github.com/globalpayments/go-sdk/api/entities/enums/baudrate"
	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/databits"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/entities/enums/parity"
	"github.com/globalpayments/go-sdk/api/entities/enums/stopbits"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
)

type ConnectionConfig struct {
	*serviceconfigs.Configuration
	ConnectionMode    connectionmodes.ConnectionModes
	BaudRate          baudrate.BaudRate
	Parity            parity.Parity
	StopBits          stopbits.StopBits
	DataBits          databits.DataBits
	IpAddress         string
	Port              int
	DeviceType        devicetype.DeviceType
	RequestIdProvider abstractions.IRequestIdProvider
}

func NewConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{Configuration: &serviceconfigs.Configuration{Timeout: 30000}}
}

func (c *ConnectionConfig) SetConnectionMode(connectionMode connectionmodes.ConnectionModes) {
	c.ConnectionMode = connectionMode
}

func (c *ConnectionConfig) SetBaudRate(baudRate baudrate.BaudRate) {
	c.BaudRate = baudRate
}

func (c *ConnectionConfig) SetParity(parity parity.Parity) {
	c.Parity = parity
}

func (c *ConnectionConfig) SetStopBits(stopBits stopbits.StopBits) {
	c.StopBits = stopBits
}

func (c *ConnectionConfig) SetDataBits(dataBits databits.DataBits) {
	c.DataBits = dataBits
}

func (c *ConnectionConfig) SetIpAddress(ipAddress string) {
	c.IpAddress = ipAddress
}

func (c *ConnectionConfig) SetPort(port int) {
	c.Port = port
}

func (c *ConnectionConfig) GetIpAddress() string {
	return c.IpAddress
}

func (c *ConnectionConfig) GetPort() int {
	return c.Port
}

func (c *ConnectionConfig) SetDeviceType(deviceType devicetype.DeviceType) {
	c.DeviceType = deviceType
}

func (c *ConnectionConfig) SetRequestIDProvider(requestIDProvider abstractions.IRequestIdProvider) {
	c.RequestIdProvider = requestIDProvider
}

func (c *ConnectionConfig) ConfigureContainer(services *configuredservices.ConfiguredServices) error {
	switch c.DeviceType {
	case devicetype.UPA_DEVICE:
		ctrl, err := NewUpaController(c)
		if err != nil {
			return err
		}
		services.SetDeviceController(ctrl)
	default:
	}
	return nil
}

func (c ConnectionConfig) IsValidated() bool {
	return c.Validated
}

func (c *ConnectionConfig) Validate() error {
	if c.ConnectionMode == connectionmodes.TCP_IP || c.ConnectionMode == connectionmodes.HTTP {
		if stringutils.IsNullOrEmpty(c.IpAddress) {
			return exceptions.NewConfigurationException("IpAddress is required for TCP or HTTP communication modes.")
		}
		if c.Port == 0 {
			return exceptions.NewConfigurationException("Port is required for TCP or HTTP communication modes.")
		}
	}
	c.Validated = true
	return nil
}
