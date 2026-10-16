package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"testing"
)

func TestPorticoEbtTests(t *testing.T) {
	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MaePAQBr-1QAqjfckFC8FTbRTT120bVQUlfVOjgCBw"
	config.ServiceUrl = "https://cert.api2.heartlandportico.com"

	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	card := paymentmethods.NewEBTCardData()
	card.SetNumber("4012002000060016")
	card.SetExpMonth(stringutils.StringToIntPointer("12"))
	card.SetExpYear(stringutils.StringToIntPointer("2025"))
	card.SetCvn("123")
	card.SetPinBlock("32539F50C245A6A93D123412324000AA")

	track := paymentmethods.NewEBTTrackData()
	track.SetValue("%B4012002000060016^VI TEST CREDIT^251210118039000000000396?;4012002000060016=25121011803939600000?")
	track.SetPinBlock("32539F50C245A6A93D123412324000AA")
	// Assume that SetEncryptionData is correctly implemented in the Go SDK
	track.SetEncryptionData(base.EncryptionDataVersion1())

	ebtSale(t, card)
	ebtReversal(t, card)
	ebtTrackBalanceInquiry(t, card)
	ebtTrackRefund(t, card)
	ebtTrackSale(t, card)
	ebtRefundFromTransactionId(t)
}

func ebtSale(t *testing.T, card *paymentmethods.EBTCardData) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("EBT sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		}
	}
}

func ebtReversal(t *testing.T, card *paymentmethods.EBTCardData) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("9")
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("EBT charge for reversal failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Charge transaction for reversal did not complete")
		return
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected charge response code for reversal: %s", response.GetResponseCode())
		return
	}

	reversalTransaction := rebuilders.NewTransactionRebuilder(paymentmethodtype.EBT).WithTransactionId(response.GetTransactionId())
	reversalTransaction.WithAmount(val)
	targetTransaction := reversalTransaction.Build()
	_, err = api.ExecuteGateway[transactions.Transaction](ctx, targetTransaction.Reverse())
	if err != nil {
		t.Errorf("EBT reversal failed with error: %s", err.Error())
	}
}

func ebtTrackBalanceInquiry(t *testing.T, card *paymentmethods.EBTCardData) {
	ctx := context.Background()
	transaction := card.BalanceInquiryWithType(inquirytype.Foodstamp)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("EBT balance inquiry failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Balance inquiry transaction did not complete")
		return
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code for balance inquiry: %s", response.GetResponseCode())
	}
}

func ebtTrackSale(t *testing.T, card *paymentmethods.EBTCardData) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("11")
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("EBT sale failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Sale transaction did not complete")
		return
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected sale response code: %s", response.GetResponseCode())
	}
}

func ebtTrackRefund(t *testing.T, card *paymentmethods.EBTCardData) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("11")
	refundTransaction := card.RefundWithAmount(val)
	refundTransaction.WithCurrency("USD")
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, refundTransaction)
	if err != nil {
		t.Errorf("EBT refund failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Refund transaction did not complete")
		return
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected refund response code: %s", response.GetResponseCode())
	}
}

func ebtRefundFromTransactionId(t *testing.T) {
	ctx := context.Background()
	reversalTransaction := rebuilders.NewTransactionRebuilder(paymentmethodtype.EBT).WithTransactionId("1234567890")
	targetTransaction := reversalTransaction.Build()
	_, err := api.ExecuteGateway[transactions.Transaction](ctx, targetTransaction.Refund())
	if err == nil {
		t.Errorf("Expected UnsupportedTransactionException for refund from transaction ID")
	}
}
