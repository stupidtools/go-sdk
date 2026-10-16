package exceptions

type ConfigurationException struct {
	message        string
	innerException error
}

func NewConfigurationException(message string, innerExceptions ...error) *ConfigurationException {
	var innerException error
	if len(innerExceptions) > 0 {
		innerException = innerExceptions[0]
	}
	return &ConfigurationException{message: message, innerException: innerException}
}

func (e *ConfigurationException) Error() string {
	return e.message
}

func (e *ConfigurationException) InnerException() error {
	return e.innerException
}
