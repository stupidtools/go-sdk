package portico

import (
	"context"
	"testing"

	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"github.com/globalpayments/go-sdk/tests/citesting"
	"github.com/shopspring/decimal"
)

func amt(s string) *decimal.Decimal {
	v, _ := stringutils.ToDecimalAmount(s)
	return v
}

func newCard() paymentmethods.CreditCardData {
	card := *paymentmethods.NewCreditCardData()
	card.SetNumber("4111111111111111")
	month, year := 12, 2025
	card.SetExpMonth(&month)
	card.SetExpYear(&year)
	card.SetCvn("123")
	return card
}

func configurePortico(h *citesting.CiTestingHarness) {
	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MTeSAQAfG1UA9qQDrzl-kz4toXvARyieptFwSKP24w"
	config.DeveloperId = "002914"
	config.VersionNumber = "3026"
	config.EnableLogging = true
	config.ServiceUrl = h.GetTestingURL()
	if err := api.ConfigureService(config, "default"); err != nil {
		panic(err)
	}
}

func TestPorticoCreditCITests(t *testing.T) {
	h := citesting.New(
		"https://cert.api2.heartlandportico.com",
		citesting.Locked,
		"PorticoCreditCITests",
	)
	ctx := context.Background()

	t.Run("CreditSale", func(t *testing.T) {
		h.SetFunction("Portico|Credit Transactions|CreditSale")
		configurePortico(h)
		clientTxnID := h.GenerateRandomID("creditSale")
		card := newCard()

		charge := card.ChargeWithAmount(amt("15.5"))
		charge.WithCurrency("USD")
		charge.WithClientTransactionId(clientTxnID)
		charge.WithUniqueDeviceId("5678")
		charge.WithAllowDuplicates(true)

		response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
		if err != nil {
			t.Fatalf("Credit sale failed: %s", err.Error())
		}
		if response.GetResponseCode() != "00" {
			t.Fatalf("Unexpected response code: %s", response.GetResponseMessage())
		}
		if ref := response.GetTransactionReference(); ref == nil || ref.ClientTransactionId != clientTxnID {
			t.Fatalf("Client transaction ID mismatch for %s", clientTxnID)
		}
	})

	t.Run("CreditTxnEdit", func(t *testing.T) {
		h.SetFunction("Portico|Credit Transactions|CreditTxnEdit - aka Gratuity")
		configurePortico(h)
		clientTxnID := h.GenerateRandomID("creditTxnEdit_charge")
		card := newCard()

		charge := card.ChargeWithAmount(amt("15"))
		charge.WithCurrency("USD")
		charge.WithClientTransactionId(clientTxnID)
		charge.WithAllowDuplicates(true)
		chargeResp, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
		if err != nil {
			t.Fatalf("Charge failed: %s", err.Error())
		}
		if chargeResp.GetResponseCode() != "00" {
			t.Fatalf("Unexpected charge response code: %s", chargeResp.GetResponseMessage())
		}

		editID := h.GenerateRandomID("creditTxnEdit_edit")
		edit := chargeResp.Edit()
		edit.WithAmount(amt("17"))
		edit.WithCurrency("USD")
		edit.WithGratuity(amt("2"))
		edit.WithClientTransactionId(editID)
		editResp, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
		if err != nil {
			t.Fatalf("Edit failed: %s", err.Error())
		}
		if editResp.GetResponseCode() != "00" {
			t.Fatalf("Unexpected edit response code: %s", editResp.GetResponseMessage())
		}
		if ref := editResp.GetTransactionReference(); ref == nil || ref.ClientTransactionId != editID {
			t.Fatalf("Edit client transaction ID mismatch for %s", editID)
		}
	})

	t.Run("CreditAdditionalAuth", func(t *testing.T) {
		h.SetFunction("Portico|Credit Transactions|CreditAdditionalAuth")
		configurePortico(h)
		clientTxnID := h.GenerateRandomID("creditAdditionalAuth_auth")
		card := newCard()

		auth := card.AuthorizeWithAmount(amt("10"), false)
		auth.WithCurrency("USD")
		auth.WithClientTransactionId(clientTxnID)
		auth.WithAllowDuplicates(true)
		authResp, err := api.ExecuteGateway[transactions.Transaction](ctx, auth)
		if err != nil {
			t.Fatalf("Authorize failed: %s", err.Error())
		}
		if authResp.GetResponseCode() != "00" {
			t.Fatalf("Unexpected auth response code: %s", authResp.GetResponseMessage())
		}
		if ref := authResp.GetTransactionReference(); ref == nil || ref.ClientTransactionId != clientTxnID {
			t.Fatalf("Client transaction ID mismatch for %s", clientTxnID)
		}

		additionalID := h.GenerateRandomID("creditAdditionalAuth_additional")
		add := authResp.Increment(amt("10"))
		add.WithCurrency("USD")
		add.WithClientTransactionId(additionalID)
		addResp, err := api.ExecuteGateway[transactions.Transaction](ctx, add)
		if err != nil {
			t.Fatalf("Additional auth failed: %s", err.Error())
		}
		if addResp.GetResponseCode() != "00" {
			t.Fatalf("Unexpected additional-auth response code: %s", addResp.GetResponseMessage())
		}
		if ref := addResp.GetTransactionReference(); ref == nil || ref.ClientTransactionId != additionalID {
			t.Fatalf("Additional-auth client transaction ID mismatch for %s", additionalID)
		}
	})
}
