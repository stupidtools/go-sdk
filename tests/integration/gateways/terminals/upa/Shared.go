package upa

import (
	"fmt"
	"testing"

	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
)

type ResponseHandler struct{}

func (r *ResponseHandler) MessageSent(msg string) {
	fmt.Println(msg)
}

func runBasicTests(t *testing.T, response abstractions.IDeviceResponse) {
	if response == nil {
		t.Error("response is nil")
	}
	if response.GetDeviceResponseCode() != "00" {
		t.Errorf("Device response code expected 00 but got %s", response.GetDeviceResponseCode())
	}
	if response.GetStatus() != "Success" {
		t.Errorf("Transaction status expected Success but got %s", response.GetStatus())
	}
}
