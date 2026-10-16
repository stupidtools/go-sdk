package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"testing"
)

func TestPorticoDebitTests(t *testing.T) {
	var track *paymentmethods.DebitTrackData

	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MaePAQBr-1QAqjfckFC8FTbRTT120bVQUlfVOjgCBw"
	config.SafDataSupported = false
	config.VersionNumber = "3026"
	config.EnableLogging = true
	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	track = paymentmethods.NewDebitTrackData()
	track.SetValue("E1050711%B4012001000000016^VI TEST CREDIT^251200000000000000000000?|LO04K0WFOmdkDz0um+GwUkILL8ZZOP6Zc4rCpZ9+kg2T3JBT4AEOilWTI|+++++++Dbbn04ekG|11;4012001000000016=25120000000000000000?|1u2F/aEhbdoPixyAPGyIDv3gBfF|+++++++Dbbn04ekG|00|||/wECAQECAoFGAgEH2wYcShV78RZwb3NAc2VjdXJlZXhjaGFuZ2UubmV0PX50qfj4dt0lu9oFBESQQNkpoxEVpCW3ZKmoIV3T93zphPS3XKP4+DiVlM8VIOOmAuRrpzxNi0TN/DWXWSjUC8m/PI2dACGdl/hVJ/imfqIs68wYDnp8j0ZfgvM26MlnDbTVRrSx68Nzj2QAgpBCHcaBb/FZm9T7pfMr2Mlh2YcAt6gGG1i2bJgiEJn8IiSDX5M2ybzqRT86PCbKle/XCTwFFe1X|")
	track.SetPinBlock("32539F50C245A6A93D123412324000AA")
	// Assume that SetEncryptionData is correctly implemented in the Go SDK
	track.SetEncryptionData(base.EncryptionDataVersion1())
	debitSaleWithTrack(t, track)
	debitRefund(t, track)
	debitReverse(t, track)

}

func debitSaleWithTrack(t *testing.T, track *paymentmethods.DebitTrackData) {
	val, _ := stringutils.ToDecimalAmount("14.01")
	ctx := context.Background()
	transaction := track.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Debit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		}
		// Reversal logic
		reversalTransaction := rebuilders.NewTransactionRebuilder(track.GetPaymentMethodType())
		reversalTransaction.WithAmount(val)
		reversalTransaction.WithPaymentMethod(track)
		toReverse := reversalTransaction.Build()
		reversal, err := api.ExecuteGateway[transactions.Transaction](ctx, toReverse.Reverse())
		if err != nil {
			t.Errorf("Reversal failed with error: %s", err.Error())
		}
		if reversal == nil {
			t.Errorf("Reversal did not complete")
		} else {
			if reversal.GetResponseCode() != "00" {
				t.Errorf("Unexpected reversal response code: %s", reversal.GetResponseCode())
			}
		}
	}
}

func debitRefund(t *testing.T, track *paymentmethods.DebitTrackData) {
	ctx := context.Background()
	refundAmount, _ := stringutils.ToDecimalAmount("16.01")
	refundTransaction := track.RefundWithAmount(refundAmount)
	refundTransaction.WithCurrency("USD")
	refundTransaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, refundTransaction)
	if err != nil {
		t.Errorf("Debit refund failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Refund transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected refund response code: %s", response.GetResponseCode())
		}
	}
}

func debitReverse(t *testing.T, track *paymentmethods.DebitTrackData) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.01")
	transaction := track.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Debit reverse charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Charge transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected charge response code: %s", response.GetResponseCode())
		}

		reverseTransaction := track.ReverseWithAmount(val)
		reverseTransaction.WithCurrency("USD")
		reverseTransaction.WithAllowDuplicates(true)
		response, err = api.ExecuteGateway[transactions.Transaction](ctx, reverseTransaction)
		if err != nil {
			t.Errorf("Debit reverse failed with error: %s", err.Error())
		}
		if response == nil {
			t.Errorf("Reverse transaction did not complete")
		} else {
			if response.GetResponseCode() != "00" {
				t.Errorf("Unexpected reverse response code: %s", response.GetResponseCode())
			}
		}
	}
}
