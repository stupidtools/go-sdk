package api

import (
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/terminals"
)

type ServicesConfig struct {
	GatewayConfig          *serviceconfigs.GatewayConfig
	DeviceConnectionConfig *terminals.ConnectionConfig
}

func (c *ServicesConfig) GetGatewayConfig() *serviceconfigs.GatewayConfig {
	return c.GatewayConfig
}

func (c *ServicesConfig) SetGatewayConfig(gatewayConfig *serviceconfigs.GatewayConfig) {
	c.GatewayConfig = gatewayConfig
}

func (c *ServicesConfig) GetDeviceConnectionConfig() *terminals.ConnectionConfig {
	return c.DeviceConnectionConfig
}

func (c *ServicesConfig) SetDeviceConnectionConfig(deviceConnectionConfig *terminals.ConnectionConfig) {
	c.DeviceConnectionConfig = deviceConnectionConfig
}

func (c *ServicesConfig) SetTimeout(timeout int) {
	if c.GatewayConfig != nil {
		c.GatewayConfig.SetTimeout(timeout)
	}
	if c.DeviceConnectionConfig != nil {
		c.DeviceConnectionConfig.SetTimeout(timeout)
	}
}

func (c *ServicesConfig) Validate() error {
	var configErr error
	if c.GatewayConfig != nil {
		configErr = c.GatewayConfig.Validate()
	}
	if configErr == nil && c.DeviceConnectionConfig != nil {
		configErr = c.DeviceConnectionConfig.Validate()
	}
	if configErr != nil {
		return exceptions.NewConfigurationException(configErr.Error())
	}
	return nil
}
