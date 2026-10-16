package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"testing"
	"time"
)

func TestPorticoTokenManagementTests(t *testing.T) {
	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MTyMAQBiHVEAewvIzXVFcmUd2UcyBge_eCpaASUp0A"
	config.ServiceUrl = "https://cert.api2.heartlandportico.com"
	config.EnableLogging = true
	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	updateToken(t)
	deleteToken(t)
	uniqueTokenRequest(t)
}

func updateToken(t *testing.T) {
	ctx := context.Background()
	card := *paymentmethods.NewCreditCardData()
	card.SetNumber("4111111111111111")
	exMonth := 12
	exYear := time.Now().Year() + 1
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")

	transaction := card.Verify()
	transaction.WithRequestMultiUseToken(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Token update failed with error: %s", err.Error())
	}
	if response == nil || response.GetToken() == "" {
		t.Errorf("Token not returned")
	}

	cardAsToken := *paymentmethods.NewCreditCardData()
	cardAsToken.SetToken(response.GetToken())
	newExpYear := 2026
	newExpMonth := 1
	cardAsToken.SetExpYear(&newExpYear)
	cardAsToken.SetExpMonth(&newExpMonth)
	gateway, err := api.LoadGateway()
	if err != nil {
		t.Errorf("Load gateway failed with error: %s", err.Error())
	}
	response2, err := cardAsToken.UpdateToken(ctx, gateway)
	if err != nil {
		t.Errorf("Token update failed with error: %s", err.Error())
	}
	if !response2 {
		t.Errorf("Token expiry update failed")
	}
}

func deleteToken(t *testing.T) {
	ctx := context.Background()
	card := *paymentmethods.NewCreditCardData()
	card.SetNumber("5454545454545454")
	exMonth := 12
	exYear := time.Now().Year() + 1
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")

	transaction := card.Verify()
	transaction.WithRequestMultiUseToken(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Token delete failed with error: %s", err.Error())
	}
	if response == nil || response.GetToken() == "" {
		t.Errorf("Token not returned")
	}

	cardAsToken := *paymentmethods.NewCreditCardData()
	cardAsToken.SetToken(response.GetToken())
	gateway, err := api.LoadGateway()
	ok, err := cardAsToken.DeleteToken(ctx, gateway)
	if !ok {
		t.Errorf("Token deletion failed")
	}
}

func uniqueTokenRequest(t *testing.T) {
	ctx := context.Background()

	// Get the test card equivalent to TestCards.VisaManual()
	card := *paymentmethods.NewCreditCardData()
	card.SetNumber("5454545454545454")
	exMonth := 12
	exYear := time.Now().Year() + 1
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")

	// First verification with non-unique token request
	transaction1 := card.Verify()
	transaction1.WithRequestMultiUseToken(true)
	response1, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction1)
	if err != nil {
		t.Errorf("First verification failed with error: %s", err.Error())
	}
	if response1 == nil || response1.GetToken() == "" {
		t.Errorf("Token not returned in first response")
	}

	// Second verification with non-unique token request
	transaction2 := card.Verify()
	transaction2.WithRequestMultiUseToken(true)
	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction2)
	if err != nil {
		t.Errorf("Second verification failed with error: %s", err.Error())
	}
	if response2 == nil || response2.GetToken() == "" {
		t.Errorf("Token not returned in second response")
	}

	// Check tokens are equal
	if response1.GetToken() != response2.GetToken() {
		t.Errorf("Expected equal tokens, got %s and %s", response1.GetToken(), response2.GetToken())
	}

	// Third verification with unique token request
	transaction3 := card.Verify()
	transaction3.WithRequestMultiUseToken(true)
	transaction3.WithRequestUniqueToken(true)
	response3, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction3)
	if err != nil {
		t.Errorf("Third verification failed with error: %s", err.Error())
	}
	if response3 == nil || response3.GetToken() == "" {
		t.Errorf("Token not returned in third response")
	}

	// Check tokens are different
	if response1.GetToken() == response3.GetToken() {
		t.Errorf("Expected different tokens, got same token: %s", response1.GetToken())
	}
}
