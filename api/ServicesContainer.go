package api

import (
	"fmt"
	abstractions2 "github.com/globalpayments/go-sdk/api/abstractions"
	apiabstractions "github.com/globalpayments/go-sdk/api/serviceconfigs/abstractions"
	"sync"

	"github.com/globalpayments/go-sdk/api/configuredservices"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/devicecontroller"
)

type ServicesContainer struct {
	configurations map[string]*configuredservices.ConfiguredServices
}

var servicesContainerInstance *ServicesContainer
var servicesContainerOnce sync.Once

func GetServiceContainerInstance() *ServicesContainer {
	servicesContainerOnce.Do(func() {
		servicesContainerInstance = &ServicesContainer{
			configurations: make(map[string]*configuredservices.ConfiguredServices),
		}
	})
	return servicesContainerInstance
}

func (sc *ServicesContainer) GetDeviceInterface(configName string) (abstractions.IDeviceInterface, error) {
	if config, ok := sc.configurations[configName]; ok {
		return config.GetDeviceInterface(), nil
	}
	return nil, exceptions.NewApiException("The specified configuration has not been configured for terminal interaction.")
}

func (sc *ServicesContainer) GetDeviceController(configName string) (devicecontroller.IDeviceController, error) {
	if config, ok := sc.configurations[configName]; ok {
		return config.GetDeviceController(), nil
	}
	return nil, exceptions.NewApiException("The specified configuration has not been configured for terminal interaction.")
}

func (sc *ServicesContainer) GetGateway(configName string) (abstractions2.IPaymentGateway, error) {
	if config, ok := sc.configurations[configName]; ok {
		return config.GetGateway(), nil
	}
	return nil, exceptions.NewApiException("The specified configuration has not been configured for gateway interaction.")
}

func Configure(config *ServicesConfig) error {
	return ConfigureWithConfigName(config, "default")
}

func ConfigureWithConfigName(config *ServicesConfig, configName string) error {
	err := config.Validate()
	if err != nil {
		return err
	}

	err = ConfigureService(config.GetDeviceConnectionConfig(), configName)
	if err != nil {
		return err
	}

	if config.GetDeviceConnectionConfig() != nil {

	}

	return nil
}

func ConfigureService(config apiabstractions.IConfiguration, configName string) error {
	if config == nil {
		return GetServiceContainerInstance().removeConfiguration(configName)
	}

	if !config.IsValidated() {
		err := config.Validate()
		if err != nil {
			return err
		}
	}

	cs := GetServiceContainerInstance().getConfiguration(configName)
	err := config.ConfigureContainer(cs)
	if err != nil {
		return err
	}

	return GetServiceContainerInstance().addConfiguration(configName, cs)
}

func (sc *ServicesContainer) getConfiguration(configName string) *configuredservices.ConfiguredServices {
	if config, ok := sc.configurations[configName]; ok {
		return config
	}
	return &configuredservices.ConfiguredServices{}
}

func (sc *ServicesContainer) addConfiguration(configName string, cs *configuredservices.ConfiguredServices) error {
	fmt.Printf("[Add Configuration] - %s\n", configName)
	if _, ok := sc.configurations[configName]; ok {
		delete(sc.configurations, configName)
	}
	sc.configurations[configName] = cs
	return nil
}

func (sc *ServicesContainer) removeConfiguration(configName string) error {
	fmt.Printf("[Remove Configuration] - %s\n", configName)
	if _, ok := sc.configurations[configName]; ok {
		delete(sc.configurations, configName)
	}
	return nil
}

func (sc *ServicesContainer) Dispose() {
	for _, cs := range sc.configurations {
		cs.Dispose()
	}
}
