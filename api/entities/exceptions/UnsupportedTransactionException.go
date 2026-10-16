package exceptions

type UnsupportedTransactionException struct {
	message        string
	innerException error
}

func NewUnsupportedTransactionException(message string, innerExceptions ...error) *UnsupportedTransactionException {
	var innerException error
	if len(innerExceptions) > 0 {
		innerException = innerExceptions[0]
	}
	return &UnsupportedTransactionException{message: message, innerException: innerException}
}

func (e *UnsupportedTransactionException) Error() string {
	return e.message
}

func (e *UnsupportedTransactionException) InnerException() error {
	return e.innerException
}
