package interfaces

import (
	"github.com/globalpayments/go-sdk/api/configuredservices"
	"github.com/globalpayments/go-sdk/api/entities/enums/baudrate"
	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/databits"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/entities/enums/parity"
	"github.com/globalpayments/go-sdk/api/entities/enums/stopbits"
	"github.com/globalpayments/go-sdk/api/logging"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
)

type IConnectionConfig interface {
	GetIpAddress() string
	GetPort() int
	GetRequestLogger() logging.IRequestLogger
	SetRequestLogger(logging.IRequestLogger)
	ConfigureContainer(*configuredservices.ConfiguredServices) error
	GetTimeout() int
	IsValidated() bool
	SetBaudRate(baudrate.BaudRate)
	SetConnectionMode(connectionmodes.ConnectionModes)
	SetDataBits(databits.DataBits)
	SetDeviceType(devicetype.DeviceType)
	SetIpAddress(string)
	SetParity(parity.Parity)
	SetPort(int)
	SetRequestIDProvider(abstractions.IRequestIdProvider)
	SetStopBits(stopbits.StopBits)
	SetTimeout(int)
	Validate() error
}
