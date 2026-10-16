package upa

import (
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/services/deviceservice"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"testing"

	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/terminals"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/tests/integration/gateways/terminals/hpa"
)

func TestUpaBatchTests(t *testing.T) {
	var device abstractions.IDeviceInterface

	config := terminals.NewConnectionConfig()
	config.Port = 8081
	config.IpAddress = "192.168.12.197"
	config.Timeout = 45000
	config.RequestIdProvider = hpa.NewRandomIdProvider()
	config.DeviceType = devicetype.UPA_DEVICE
	config.ConnectionMode = connectionmodes.TCP_IP

	device, err := deviceservice.DeviceServiceCreate(config)
	if err != nil {
		t.Errorf("Failed to create device with error: %s", err.Error())
	}

	if device == nil {
		t.Error("device is nil")
	}

	device.SetOnMessageSent(&ResponseHandler{})
	incrementalAuths(t, device)
	openTabDetailsReport(t, device)
	endOfDay(t, device)

}

func openTabDetailsReport(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.GetOpenTabDetails()
	if err != nil {
		t.Errorf("OpenTabDetails retrieval failed with error: %s", err.Error())
		return
	}

	runBasicTests(t, response)

	transactions := response.GetTransactionSummaries()

	for _, transaction := range transactions {
		gratuity, _ := stringutils.ToDecimalAmount("0.00")
		terminal, err := device.CreditCapture(transaction.GetAuthorizedAmount())
		if err != nil {
			t.Errorf("Credit Capture setup failed for transaction ID %s with error: %s", transaction.GetTransactionId(), err.Error())
			continue
		}

		terminal.WithGratuity(gratuity)
		terminal.WithTransactionId(transaction.GetTransactionId())

		response, err := api.ExecuteTerminal(terminal)

		if err != nil {
			t.Errorf("Credit Capture execution failed for transaction ID %s with error: %s", transaction.GetTransactionId(), err.Error())
		} else {
			runBasicTests(t, response)
		}
	}
}

func endOfDay(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.EndOfDay()
	if err != nil {
		t.Errorf("EndOfDay failed with error: %s", err.Error())
	}

	runBasicTests(t, response)
	if *response.GetBatchId() == 0 {
		t.Errorf("Batch ID not returned")
	}
}
