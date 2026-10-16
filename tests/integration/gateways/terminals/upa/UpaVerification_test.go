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

func TestUpaVerificationTests(t *testing.T) {
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

	test010bAuthCapture(t, device)

}

func test010bAuthCapture(t *testing.T, device abstractions.IDeviceInterface) {
	val, _ := stringutils.ToDecimalAmount("15.12")
	terminal, err := device.CreditAuth(val)
	if err != nil {
		t.Errorf("Credit Auth setup failed with error: %s", err.Error())
		return
	}

	response, err := api.ExecuteTerminal(terminal)
	if err != nil {
		t.Errorf("Credit Auth failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Response is nil")
		return
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Expected response code '00' but got %s", response.GetResponseCode())
	}

	captureVal, _ := stringutils.ToDecimalAmount("18.12")
	tipVal, _ := stringutils.ToDecimalAmount("0.00")
	captureTerminal, err := device.CreditCapture(captureVal)
	captureTerminal.WithTransactionId(response.GetTransactionId()).
		WithGratuity(tipVal)

	if err != nil {
		t.Errorf("Credit Capture setup failed with error: %s", err.Error())
		return
	}

	captureResponse, err := api.ExecuteTerminal(captureTerminal)
	if err != nil {
		t.Errorf("Credit Capture failed with error: %s", err.Error())
		return
	}

	// EMV receipt requirements
	target, _ := stringutils.ToDecimalAmount("18.12")
	actual := captureResponse.GetTransactionAmount()
	if actual == nil {
		t.Errorf("Capture transaction did not complete")
	} else if target.Cmp(*actual) != 0 {
		t.Errorf("Capture transaction amount expected %v but got %v", target, actual)
	}

	if captureResponse.GetMaskedCardNumber() == "" {
		t.Errorf("Masked card number is empty")
	}

	if captureResponse.GetApplicationPreferredName()+captureResponse.GetApplicationLabel() == "" {
		t.Errorf("Application preferred name and label are empty")
	}

	if captureResponse.GetApplicationId() == "" {
		t.Errorf("Application ID is empty")
	}

	if captureResponse.GetApplicationCryptogramType() == "" {
		t.Errorf("Application Cryptogram Type is empty")
	}

	if captureResponse.GetApplicationCryptogram() == "" {
		t.Errorf("Application Cryptogram is empty")
	}

	if captureResponse.GetEntryMethod() == "" {
		t.Errorf("Entry method is empty")
	}

	if captureResponse.GetCardHolderName() == "" {
		t.Errorf("Cardholder name is empty")
	}
}
