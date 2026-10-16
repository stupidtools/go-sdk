package abstractions

type IEODResponse interface {
	IDeviceResponse
	GetBatchId() *int
}
