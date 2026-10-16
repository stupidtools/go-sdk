package abstractions

type ISignatureResponse interface {
	IDeviceResponse
	GetSignatureData() []byte
}
