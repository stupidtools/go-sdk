package serviceconfigs

import (
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
)

type GatewayConfig struct {
	Configuration
	UseDataReportingService bool
	DataClientId            string
	DataClientSecret        string
	DataClientUserId        string
	DataClientServiceUrl    string
}

func NewGatewayConfig() *GatewayConfig {
	return &GatewayConfig{}
}

func (gc *GatewayConfig) ConfigureContainer() error {
	return exceptions.NewConfigurationException("Not Implemented")

}

func (gc *GatewayConfig) Validate() error {
	return gc.Configuration.Validate()
}
