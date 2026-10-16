package upa

import (
	"context"
	"github.com/globalpayments/go-sdk/api/services/deviceservice"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
	"strings"
	"testing"
	"time"

	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities/enums/connectionmodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
	"github.com/globalpayments/go-sdk/api/terminals"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"github.com/globalpayments/go-sdk/tests/integration/gateways/terminals/hpa"
)

func TestUpaCreditTests(t *testing.T) {
	var device abstractions.IDeviceInterface

	config := terminals.NewConnectionConfig()
	config.Port = 8081
	config.IpAddress = "192.168.12.217"
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

	creditSale(t, device)
	creditSaleWithContextCancellation(t, device)
	creditVoidTerminalTrans(t, device)
	creditCancelledTrans(t, device)
	creditBlindRefund(t, device)
	creditDeletePreAuths(t, device)
	incrementalAuths(t, device)
	tipAdjust(t, device)
}

func creditSale(t *testing.T, device abstractions.IDeviceInterface) {
	val, _ := stringutils.ToDecimalAmount("12.01")
	terminal, err := device.CreditSale(val)
	if err != nil {
		t.Errorf("Credit setup failed with error: %s", err.Error())
	}
	response, err := api.ExecuteTerminal(terminal)

	if err != nil {
		t.Errorf("Credit sale swipe failed with error: %s", err.Error())
	}

	runBasicTests(t, response)
	target, _ := stringutils.ToDecimalAmount("12.01")
	actual := response.GetTransactionAmount()
	if actual == nil {
		t.Errorf("Transaction did not complete")
	} else if target.Cmp(*actual) != 0 {
		t.Errorf("Transaction amount expected %v but got %v", target, actual)
	}
}

func creditSaleWithContextCancellation(t *testing.T, device abstractions.IDeviceInterface) {
	val, err := stringutils.ToDecimalAmount("12.01")
	if err != nil {
		t.Fatalf("Failed to convert amount: %v", err)
	}

	terminal, err := device.CreditSale(val)
	if err != nil {
		t.Fatalf("Credit setup failed with error: %s", err.Error())
	}

	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a channel to receive the result of ExecuteTerminalWithContext
	resultChan := make(chan struct {
		response terminalresponse.ITerminalResponse
		err      error
	})

	// Start ExecuteTerminalWithContext in a goroutine
	go func() {
		response, err := api.ExecuteTerminalWithContext(ctx, terminal)
		resultChan <- struct {
			response terminalresponse.ITerminalResponse
			err      error
		}{response, err}
	}()

	// Wait for 5 seconds, then cancel the context
	time.Sleep(5 * time.Second)
	cancel()

	// Wait for the result
	select {
	case result := <-resultChan:
		if result.err == nil {
			t.Errorf("Expected an error due to cancellation, but got nil")
		} else if !strings.Contains(result.err.Error(), "context canceled") {
			t.Errorf("Expected a context cancellation error, but got: %v", result.err)
		}
	case <-time.After(15 * time.Second):
		t.Errorf("Test timed out waiting for ExecuteTerminalWithContext to return")
	}

}

func creditBlindRefund(t *testing.T, device abstractions.IDeviceInterface) {
	val, _ := stringutils.ToDecimalAmount("12.01")
	terminal, err := device.CreditRefund(val)
	if err != nil {
		t.Errorf("Refund setup failed with error: %s", err.Error())
	}
	response, err := api.ExecuteTerminal(terminal)
	if err != nil {
		t.Errorf("Credit refund failed with error: %s", err.Error())
	}

	runBasicTests(t, response)
	target, _ := stringutils.ToDecimalAmount("12.01")
	actual := response.GetTransactionAmount()
	if actual == nil {
		t.Errorf("Transaction did not complete")
	} else if target.Cmp(*actual) != 0 {
		t.Errorf("Transaction amount expected %v but got %v", target, actual)
	}
}

func creditVoidTerminalTrans(t *testing.T, device abstractions.IDeviceInterface) {
	val, _ := stringutils.ToDecimalAmount("12.34")
	terminal, err := device.CreditSale(val)
	if err != nil {
		t.Errorf("Credit setup failed with error: %s", err.Error())
	}
	response1, err := api.ExecuteTerminal(terminal)

	if err != nil {
		t.Errorf("Void terminal trans failed with error: %s", err.Error())
	}

	runBasicTests(t, response1)
	terminal2, err := device.CreditVoid()
	if err != nil {
		t.Errorf("Void setup failed with error: %s", err.Error())
	}
	response2, err := api.ExecuteTerminal(terminal2.
		WithTerminalRefNumber(response1.GetTerminalRefNumber()))

	if err != nil {
		t.Errorf("Void terminal trans failed with error: %s", err.Error())
	}

	runBasicTests(t, response2)
}

func creditCancelledTrans(t *testing.T, device abstractions.IDeviceInterface) {
	val, _ := stringutils.ToDecimalAmount("12.34")
	terminal, err := device.CreditSale(val)
	if err != nil {
		t.Errorf("Credit setup failed with error: %s", err.Error())
	}
	response, err := api.ExecuteTerminal(terminal)
	if err != nil {
		t.Errorf("Cancelled trans failed with error: %s", err.Error())
	}

	if response == nil {
		t.Error("response is nil")
	}

	if response.GetStatus() != "Failed" {
		t.Errorf("Transaction status expected Failed but got %s", response.GetStatus())
	}

	if response.GetDeviceResponseCode() != "APP001" {
		t.Errorf("Device response code expected APP001 but got %s", response.GetDeviceResponseCode())
	}

	if response.GetDeviceResponseText() != "TRANSACTION CANCELLED BY USER" {
		t.Errorf("Device response text expected TRANSACTION CANCELLED BY USER but got %s", response.GetDeviceResponseText())
	}
}

func creditDeletePreAuths(t *testing.T, device abstractions.IDeviceInterface) {
	// Initial authorization amount
	amount1, _ := stringutils.ToDecimalAmount("10.00")
	terminal1, err := device.CreditAuth(amount1)
	if err != nil {
		t.Errorf("Initial credit authorization setup failed with error: %s", err.Error())
		return
	}
	response1, err := api.ExecuteTerminal(terminal1)
	if err != nil {
		t.Errorf("Initial credit authorization failed with error: %s", err.Error())
		return
	}

	// Ensure that the first transaction response is not nil
	if response1 == nil {
		t.Errorf("First transaction did not complete")
		return
	}

	response2, err := device.PreAuthDelete(response1.GetTransactionId(), amount1)
	if err != nil {
		t.Errorf("Authorization delete failed with error: %s", err.Error())
		return
	}

	// Ensure that the first transaction response is not nil
	if response2 == nil {
		t.Errorf("Second transaction (auth delete) did not complete")
		return
	}
	if response2.GetDeviceResponseCode() != "00" {
		t.Errorf("Device response code expected 00 but got %s", response2.GetDeviceResponseCode())
	}
}

func incrementalAuths(t *testing.T, device abstractions.IDeviceInterface) {
	// Initial authorization amount
	amount1, _ := stringutils.ToDecimalAmount("10.00")
	terminal1, err := device.CreditAuth(amount1)
	if err != nil {
		t.Errorf("Initial credit authorization setup failed with error: %s", err.Error())
		return
	}
	response1, err := api.ExecuteTerminal(terminal1)
	if err != nil {
		t.Errorf("Initial credit authorization failed with error: %s", err.Error())
		return
	}

	// Ensure that the first transaction response is not nil
	if response1 == nil {
		t.Errorf("First transaction did not complete")
		return
	}

	// Incremental authorization amount
	amount2, _ := stringutils.ToDecimalAmount("5.00")
	terminal2, err := device.CreditAuth(amount2)
	if err != nil {
		t.Errorf("Incremental credit authorization setup failed with error: %s", err.Error())
		return
	}
	terminal2.WithTransactionId(response1.GetTransactionId())
	response2, err := api.ExecuteTerminal(terminal2)
	if err != nil {
		t.Errorf("Incremental credit authorization failed with error: %s", err.Error())
		return
	}

	// Ensure that the second transaction response is not nil
	if response2 == nil {
		t.Errorf("Second transaction did not complete")
		return
	}

	// Check if the response code is "00"
	if response2.GetDeviceResponseCode() != "00" {
		t.Errorf("Device response code expected 00 but got %s", response2.GetDeviceResponseCode())
	}
}

func tipAdjust(t *testing.T, device abstractions.IDeviceInterface) {
	// Convert string amount to decimal
	val, _ := stringutils.ToDecimalAmount("12.34")
	grat, _ := stringutils.ToDecimalAmount("0.00")
	clerkId := 420
	// Setup and execute the credit sale
	terminal, err := device.CreditSale(val)
	if err != nil {
		t.Errorf("Credit setup failed with error: %s", err.Error())
		return
	}
	terminal.WithGratuity(grat)
	terminal.WithClerkId(&clerkId)
	response1, err := api.ExecuteTerminal(terminal)
	if err != nil {
		t.Errorf("Tip adjust failed with error: %s", err.Error())
		return
	}

	runBasicTests(t, response1)

	// Convert tip amount to decimal
	tipVal, _ := stringutils.ToDecimalAmount("1.50")

	// Setup and execute the tip adjust
	terminal2, err := device.TipAdjust(tipVal)
	if err != nil {
		t.Errorf("Tip adjust setup failed with error: %s", err.Error())
		return
	}
	tm := terminal2.WithTerminalRefNumber(response1.GetTerminalRefNumber()).WithTransactionId(response1.GetTransactionId())
	tm.WithClerkId(&clerkId)
	response2, err := api.ExecuteTerminal(tm)

	if err != nil {
		t.Errorf("Tip adjust transaction failed with error: %s", err.Error())
		return
	}

	runBasicTests(t, response2)

	// Asserting the results
	if !response2.GetTipAmount().Equals(*tipVal) {
		t.Errorf("Expected tip amount: %s, but got: %s", tipVal.String(), response2.GetTipAmount().String())
	}

	expectedTransactionAmount, _ := stringutils.ToDecimalAmount("13.84")
	if !response2.GetTransactionAmount().Equals(*expectedTransactionAmount) {
		t.Errorf("Expected transaction amount: %s, but got: %s", expectedTransactionAmount.String(), response2.GetTransactionAmount().String())
	}
}
