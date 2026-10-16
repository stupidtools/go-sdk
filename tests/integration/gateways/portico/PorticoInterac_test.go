package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities/enums/accounttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/emvchipcondition"
	"github.com/globalpayments/go-sdk/api/entities/enums/entrymethod"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"testing"
)

func TestPorticoInteracTests(t *testing.T) {
	var track *paymentmethods.DebitTrackData
	var tagData string

	config := serviceconfigs.NewPorticoConfig()
	config.LicenseId = "374209"
	config.SiteId = "374391"
	config.DeviceId = "5246"
	config.Username = "gateway1082907"
	config.Password = "$Test1234"
	config.ServiceUrl = "https://cert.api2.heartlandportico.com"
	config.EnableLogging = true

	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	track = paymentmethods.NewDebitTrackData()
	track.SetValue(";0012030000000003=2812220016290740?")
	track.SetEntryMethod(entrymethod.Swipe)
	// Assume SetEntryMethod is correctly implemented

	tagData = "82021C008407A0000002771010950580000000009A031709289C01005F280201245F2A0201245F3401019F02060000000010009F03060000000000009F080200019F090200019F100706010A03A420009F1A0201249F26089CC473F4A4CE18D39F2701809F3303E0F8C89F34030100029F3501229F360200639F370435EFED379F410400000019"

	debitAuth(t, track, tagData)
	debitAddToBatch(t, track, tagData)
}

func debitAuth(t *testing.T, track *paymentmethods.DebitTrackData, tagData string) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("10")
	transaction := track.AuthorizeWithAmount(val, false)
	transaction.WithCurrency("USD")
	transaction.WithAllowDuplicates(true)
	transaction.WithCardHolderLanguage("ENGLISH")
	transaction.WithPosSequenceNumber("000010010180")
	transaction.WithTagData(tagData)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Debit auth failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		}
	}
}

func debitAddToBatch(t *testing.T, track *paymentmethods.DebitTrackData, tagData string) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("14.01")
	authorizationTransaction := track.AuthorizeWithAmount(val, false)
	authorizationTransaction.WithCurrency("USD")
	authorizationTransaction.WithAllowDuplicates(true)
	authorizationTransaction.WithCardHolderLanguage("ENGLISH")
	authorizationTransaction.WithPosSequenceNumber("000010010180")
	authorizationTransaction.WithTagData(tagData)
	authorizationTransaction.WithAccountType(accounttype.Savings)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorizationTransaction)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
	}
	if response == nil {
		t.Errorf("Authorization transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected authorization response code: %s", response.GetResponseCode())
		}
		// Capture logic
		captureVal, _ := stringutils.ToDecimalAmount("16")
		captureTransaction := rebuilders.NewTransactionRebuilder(paymentmethodtype.Debit)
		captureTransaction.WithTransactionId(response.GetTransactionId())
		captureTransaction.WithPaymentMethod(track)
		builder := captureTransaction.Build().CaptureWithAmount(captureVal)
		builder.WithCurrency("USD")
		builder.WithPosSequenceNumber("000010010180")
		builder.WithTagData(tagData)
		builder.WithAccountType(accounttype.Savings)
		builder.WithChipCondition(emvchipcondition.ChipFailPreviousFail)

		capture, err := api.ExecuteGateway[transactions.Transaction](ctx, captureTransaction.Build().Capture())
		if err != nil {
			t.Errorf("Capture failed with error: %s", err.Error())
		}
		if capture == nil {
			t.Errorf("Capture transaction did not complete")
		} else {
			if capture.GetResponseCode() != "00" {
				t.Errorf("Unexpected capture response code: %s", capture.GetResponseCode())
			}
		}
	}
}
