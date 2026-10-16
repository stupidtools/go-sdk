package abstractions

type IInitializeResponse interface {
	IDeviceResponse
	GetSerialNumber() string
}
