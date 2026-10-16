package exceptions

type BuilderException struct {
	message        string
	innerException error
}

func NewBuilderException(message string, innerExceptions ...error) *BuilderException {
	var innerException error
	if len(innerExceptions) > 0 {
		innerException = innerExceptions[0]
	}
	return &BuilderException{message: message, innerException: innerException}
}

func (e *BuilderException) Error() string {
	return e.message
}

func (e *BuilderException) InnerException() error {
	return e.innerException
}
