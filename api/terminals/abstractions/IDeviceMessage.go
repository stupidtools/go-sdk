package abstractions

type IDeviceMessage interface {
	IsKeepAlive() bool
	SetKeepAlive(keepAlive bool)
	IsAwaitResponse() bool
	SetAwaitResponse(awaitResponse bool)
	GetSendBuffer() []byte
}
