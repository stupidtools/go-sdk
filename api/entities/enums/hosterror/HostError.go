package hosterror

type HostError string

const (
	Connection  HostError = "Connection"
	SendFailure HostError = "SendFailure"
	Timeout     HostError = "Timeout"
)
