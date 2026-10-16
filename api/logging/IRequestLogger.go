package logging

type IRequestLogger interface {
	RequestSent(request string) error
	ResponseReceived(response string) error
}
