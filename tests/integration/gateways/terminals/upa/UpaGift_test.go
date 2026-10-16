package upa

import (
	"github.com/globalpayments/go-sdk/api/services/deviceservice"
	"testing"

	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/terminals"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"github.com/globalpayments/go-sdk/tests/integration/gateways/terminals/hpa"
)

func TestUpaGiftTests(t *testing.T) {
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

	giftAddValue(t, device)

}

func giftAddValue(t *testing.T, device abstractions.IDeviceInterface) {
	val, _ := stringutils.ToDecimalAmount("12.01")
	terminal, err := device.GiftAddValue(val)
	if err != nil {
		t.Errorf("Gift setup failed with error: %s", err.Error())
	}
	response, err := api.ExecuteTerminal(terminal.WithGiftTransactionType(transactiontype.Sale))

	if err != nil {
		t.Errorf("Gift sale swipe failed with error: %s", err.Error())
	}

	runBasicTests(t, response)
	unmasked := response.GetUnmaskedCardNumber()
	target := "5022440000000000098"
	if unmasked != target {
		t.Errorf("Gift card number incorrect. Wanted: %s Actual: %s", target, unmasked)
	}
}
