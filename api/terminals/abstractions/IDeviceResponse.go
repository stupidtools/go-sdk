package abstractions

type IDeviceResponse interface {
	GetStatus() string
	SetStatus(status string)
	GetCommand() string
	SetCommand(command string)
	GetVersion() string
	SetVersion(version string)
	GetDeviceResponseCode() string
	SetDeviceResponseCode(deviceResponseCode string)
	GetDeviceResponseText() string
	SetDeviceResponseText(deviceResponseMessage string)
	ToString() string
}
