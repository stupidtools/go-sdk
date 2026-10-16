package serviceconfigs

import (
	"github.com/globalpayments/go-sdk/api/configuredservices"
	"github.com/globalpayments/go-sdk/api/entities/enums/environment"
	"github.com/globalpayments/go-sdk/api/entities/enums/serviceendpoints"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/gateways"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"strings"
)

type PorticoConfig struct {
	GatewayConfig
	SiteId                      string
	LicenseId                   string
	DeviceId                    string
	Username                    string
	Password                    string
	DeveloperId                 string
	VersionNumber               string
	SecretApiKey                string
	UniqueDeviceId              string
	SdkNameVersion              string
	CertificationStr            string
	TerminalID                  string
	X509CertificatePath         string
	X509CertificateBase64String string
	ProPayUS                    bool
	SafDataSupported            bool
}

func NewPorticoConfig() *PorticoConfig {
	return &PorticoConfig{GatewayConfig: *NewGatewayConfig()}
}

func (gc *PorticoConfig) ConfigureContainer(services *configuredservices.ConfiguredServices) error {

	if gc.ServiceUrl == "" {
		if gc.Environment == environment.PRODUCTION {
			gc.ServiceUrl = serviceendpoints.PORTICO_PRODUCTION.GetValue()
		} else {
			gc.ServiceUrl = serviceendpoints.PORTICO_TEST.GetValue()
		}
	}

	gateway := gateways.NewPorticoConnector()
	gateway.SiteId = gc.SiteId
	gateway.LicenseId = gc.LicenseId
	gateway.DeviceId = gc.DeviceId
	gateway.Username = gc.Username
	gateway.Password = gc.Password
	gateway.SecretApiKey = gc.SecretApiKey
	gateway.DeveloperId = gc.DeveloperId
	gateway.VersionNumber = gc.VersionNumber
	gateway.IsSAFDataSupported = gc.SafDataSupported
	gateway.SdkNameVersion = gc.SdkNameVersion
	gateway.Timeout = gc.Timeout
	gateway.ServiceUrl = gc.ServiceUrl + "/Hps.Exchange.PosGateway/PosGatewayService.asmx"
	gateway.EnableLogging = gc.EnableLogging
	gateway.RequestLogger = gc.RequestLogger
	gateway.WebProxy = gc.WebProxy

	services.SetGateway(gateway)

	//if stringutils.IsNullOrEmpty(gc.DataClientId) {
	//	services.SetReportingService(gateway)
	//}

	//payplan := &PayPlanConnector{}
	//payplan.SetSecretApiKey(gc.SecretApiKey).
	//	SetTimeout(gc.Timeout).
	//	SetServiceUrl(gc.ServiceUrl + gc.GetPayPlanEndpoint()).
	//	SetEnableLogging(gc.EnableLogging).
	//	SetRequestLogger(gc.RequestLogger).
	//	SetWebProxy(gc.WebProxy)
	//
	//services.SetRecurringConnector(payplan)

	// TODO: Implement ProPayConnector

	return nil
}

func (pc *PorticoConfig) getPayPlanEndpoint() string {
	if (!stringutils.IsNullOrEmpty(pc.SecretApiKey) && strings.Contains(strings.ToLower(pc.SecretApiKey), "cert")) ||
		(stringutils.IsNullOrEmpty(pc.SecretApiKey) && pc.Environment == "TEST") {
		return "/Portico.PayPlan.v2/"
	}
	return "/PayPlan.v2/"
}

func (pc *PorticoConfig) Validate() error {
	err := pc.GatewayConfig.Validate()
	if err != nil {
		return err
	}

	// Portico api key validation
	if pc.SecretApiKey != "" {
		if pc.SiteId != "" || pc.LicenseId != "" || pc.DeviceId != "" || pc.Username != "" || pc.Password != "" {
			return exceptions.NewConfigurationException("Configuration contains both secret api key and legacy credentials. These are mutually exclusive.")
		}
	}

	// Portico legacy validation
	if pc.SiteId != "" || pc.LicenseId != "" || pc.DeviceId != "" || pc.Username != "" || pc.Password != "" {
		if pc.SiteId == "" || pc.LicenseId == "" || pc.DeviceId == "" || pc.Username == "" || pc.Password == "" {
			return exceptions.NewConfigurationException("Site, License, Device, Username and Password should all have values for this configuration.")
		}
	}

	return nil
}
