package exceptions

type MessageException struct {
	message        string
	innerException error
}

func NewMessageException(message string, innerExceptions ...error) *MessageException {
	var innerException error
	if len(innerExceptions) > 0 {
		innerException = innerExceptions[0]
	}
	return &MessageException{message: message, innerException: innerException}
}

func (e *MessageException) Error() string {
	return e.message
}

func (e *MessageException) InnerException() error {
	return e.innerException
}
