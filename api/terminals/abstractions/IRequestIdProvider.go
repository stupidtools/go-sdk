package abstractions

type IRequestIdProvider interface {
	GetRequestId() int
}
