package abstractions

import "github.com/globalpayments/go-sdk/api/configuredservices"

type IConfiguration interface {
	GetTimeout() int
	SetTimeout(timeout int)
	Validate() error
	IsValidated() bool
	ConfigureContainer(config *configuredservices.ConfiguredServices) error
}
