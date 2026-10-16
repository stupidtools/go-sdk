package upa

import (
	"github.com/globalpayments/go-sdk/api/services/deviceservice"
	"strings"
	"testing"

	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/terminals"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/tests/integration/gateways/terminals/hpa"
)

func TestUpaAdminTests(t *testing.T) {
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

	ping(t, device)
	cancel(t, device)
	reboot(t, device)
	lineItems(t, device)
	lineItemsBulk(t, device)
	promptForSignature(t, device)
	getSafReport(t, device)
	sendSAF(t, device)
	deleteSAF(t, device)
}

func ping(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.Ping()
	if err != nil {
		t.Fatal(err)
	}
	runBasicTests(t, response)
}

func cancel(t *testing.T, device abstractions.IDeviceInterface) {
	err := device.Cancel()
	if err != nil {
		t.Fatal(err)
	}
}

func lineItemsBulk(t *testing.T, device abstractions.IDeviceInterface) {
	lineItems := make([][2]string, 0)
	for i := 0; i < 10; i++ {
		lineItems = append(lineItems, [2]string{"Line Item 2", "2.00"})
	}
	_, err := device.AddLineItemBulk(lineItems)
	if err != nil {
		t.Fatal(err)
	}
	err = device.Cancel()
	if err != nil {
		t.Fatal(err)
	}
}

func lineItems(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.AddLineItem("Line Item 2", "2.00")
	if err != nil {
		t.Fatal(err)
	}
	runBasicTests(t, response)
	err = device.Cancel()
	if err != nil {
		t.Fatal(err)
	}
}

func promptForSignature(t *testing.T, device abstractions.IDeviceInterface) {

	response, err := device.PromptForSignature()
	if err != nil {
		t.Fatal(err)
	}
	runBasicTests(t, response)
	if response.GetSignatureData() == nil {
		t.Fatal("Signature data is nil")
	}

}

func reboot(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.Reboot()
	if err != nil {
		t.Fatal(err)
	}
	runBasicTests(t, response)
}

func deleteSAF(t *testing.T, device abstractions.IDeviceInterface) {
	// this requires an actual transaction number and reference number to work.  The sample data here will fail.
	response, err := device.SafDelete("12345", "0001")
	if err != nil {
		t.Errorf("SafDelete failed with error: %s", err.Error())
	}

	runBasicTests(t, response)

}

func sendSAF(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.SendStoreAndForward()
	if err != nil {
		t.Errorf("SendStoreAndForward failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Response is null")
	}

	if response.GetDeviceResponseCode() != "00" {
		t.Errorf("Device response code not as expected: expected '00', got '%s'", response.GetDeviceResponseCode())
	}

	status := response.GetStatus()
	if strings.ToLower(status) != "success" {
		t.Errorf("Status not as expected: expected 'success', got '%s'", status)
	}
}

func getSafReport(t *testing.T, device abstractions.IDeviceInterface) {
	response, err := device.SafSummaryReport()
	if err != nil {
		t.Errorf("SafSummaryReport failed with error: %s", err.Error())
	}
	if response == nil {
		t.Errorf("Response is nil")
	}

	if response.GetDeviceResponseCode() != "00" {
		t.Errorf("Device response code not correct, got: %s, want: %s", response.GetDeviceResponseCode(), "00")
	}

	if strings.ToLower(response.GetStatus()) != "success" {
		t.Errorf("Response status not correct, got: %s, want: %s", response.GetStatus(), "success")
	}

}
