package exceptions

type ApiException struct {
	message        string
	innerException error
}

func NewApiException(message string, innerExceptions ...error) *ApiException {
	var innerException error
	if len(innerExceptions) > 0 {
		innerException = innerExceptions[0]
	}
	return &ApiException{message: message, innerException: innerException}
}

func (e *ApiException) Error() string {
	return e.message
}

func (e *ApiException) InnerException() error {
	return e.innerException
}
