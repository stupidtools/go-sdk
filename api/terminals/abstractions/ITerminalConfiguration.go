package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/baudrate"
	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/databits"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/entities/enums/parity"
	"github.com/globalpayments/go-sdk/api/entities/enums/stopbits"
)

type ITerminalConfiguration interface {
	GetConnectionMode() connectionmodes.ConnectionModes
	SetConnectionMode(connectionMode connectionmodes.ConnectionModes)
	GetIpAddress() string
	SetIpAddress(ipAddress string)
	GetPort() int
	SetPort(port int)
	GetBaudRate() baudrate.BaudRate
	SetBaudRate(baudRate baudrate.BaudRate)
	GetParity() parity.Parity
	SetParity(parity parity.Parity)
	GetStopBits() stopbits.StopBits
	SetStopBits(stopBits stopbits.StopBits)
	GetDataBits() databits.DataBits
	SetDataBits(dataBits databits.DataBits)
	GetTimeout() int
	SetTimeout(timeout int)
	Validate() error
	GetDeviceType() devicetype.DeviceType
	SetDeviceType(deviceType devicetype.DeviceType)
	GetRequestIdProvider() IRequestIdProvider
	SetRequestIdProvider(requestIdProvider IRequestIdProvider)
}
