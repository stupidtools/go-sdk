package certifications

import (
	"context"
	"fmt"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/services/batchservice"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"github.com/globalpayments/go-sdk/tests/integration/gateways/terminals/hpa"
	"github.com/globalpayments/go-sdk/tests/testdata"
	"testing"
	"time"
)

func TestPorticoEcommerceCert(t *testing.T) {

	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MTeSAQAfG1UA9qQDrzl-kz4toXvARyieptFwSKP24w"
	config.VersionNumber = "3026"
	config.EnableLogging = true
	config.SafDataSupported = false
	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	verifyVisa001(t)
	verifyMCFive002(t)
	verifyDiscover003(t)
	verifyAmex004(t)
	creditSaleWithReversalVisa005and006(t)
	creditSaleMCFive007(t)
	creditSaleDiscover008(t)
	creditSaleAmex009(t)
	creditSaleJCB010(t)
	creditSaleMCTwo011(t)
	creditAuthAndCaptureVisa012Aand012B(t)
	creditAuthAndCaptureMCFive013Aand013B(t)
	creditAuthDiscover014A(t)
	creditAuthMCTwo015A(t)
	creditSaleWithGratuityEditVisa016(t)
	creditSaleWithGratuityEditMCfive017(t)
	creditSaleWithGratuityCaptureVisa018(t)
	creditSaleWithGratuityCaptureMCfive019(t)
	creditSaleLVL2Visa020(t)
	creditSaleLVL2Visa021(t)
	creditSaleLVL2Visa022(t)
	creditSaleLVL2Visa023(t)
	creditSaleLVL2MCFive024(t)
	creditSaleLVL2MCFive025(t)
	creditSaleLVL2MCFive026(t)
	creditSaleLVL2MCFive027(t)
	creditSaleLVL2Amex028(t)
	creditSaleLVL2Amex029(t)
	creditSaleLVL2Amex030(t)
	creditSaleLVL2Amex031(t)
	creditSaleLVL3Visa032(t)
	creditSaleLVL3Visa033(t)
	creditSaleLVL3Visa034(t)
	creditSaleLVL3MCFive035(t)
	creditSaleLVL3MCFive036(t)
	creditIncrementalAuthVisa037(t)
	creditIncrementalAuthMCFive038(t)
	creditIncrementalAuthVisa039(t)
	creditTimeoutReversalVisa040(t)
	creditCloseBatch(t)
}

func verifyVisa001(t *testing.T) {
	ctx := context.Background()
	card := testdata.VisaManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	transaction := card.Verify()
	transaction.WithAllowDuplicates(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit verify failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("verifyVisa001 completed successfully:  TRID - " + response.GetTransactionId())
		}
	}

}

func verifyMCFive002(t *testing.T) {
	ctx := context.Background()
	card := testdata.MasterCard5Manual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	transaction := card.Verify()
	transaction.WithAllowDuplicates(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit verify failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("verifyMCFive002 completed successfully:  TRID - " + response.GetTransactionId())
		}
	}
}

func verifyDiscover003(t *testing.T) {
	ctx := context.Background()
	card := testdata.DiscoverManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	transaction := card.Verify()
	transaction.WithAllowDuplicates(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit verify failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("verifyDiscover003 completed successfully:  TRID - " + response.GetTransactionId())
		}
	}
}

func verifyAmex004(t *testing.T) {
	ctx := context.Background()
	card := testdata.AmexManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	transaction := card.Verify()
	transaction.WithAllowDuplicates(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit verify failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("verifyAmex004 completed successfully:  TRID - " + response.GetTransactionId())
		}
	}
}

func creditSaleWithReversalVisa005and006(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.01")
	card := testdata.VisaManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860 Dallas Pkwy")
	transaction := card.ChargeWithAmount(val)
	transaction.WithAllowDuplicates(true)
	transaction.WithRequestMultiUseToken(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	cliTrId := hpa.NewRandomIdProvider().GetRequestIdString()
	transaction.WithClientTransactionId(cliTrId)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("creditSaleVisa005 completed successfully:  TRID - " + response.GetTransactionId())
		}

	}
	transactionId := response.GetTransactionId()
	revTrans := rebuilders.FromId(transactionId, card.GetPaymentMethodType()).ReverseWithAmount(val)
	revTrans.WithClientTransactionId(cliTrId)
	reverseResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, revTrans)
	if err != nil || reverseResponse == nil {
		t.Errorf("Reverse with transaction ID failed with error: %s", err.Error())
		return
	}
	if reverseResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected reverse response code with transaction ID: %s", reverseResponse.GetResponseCode())
	} else {
		fmt.Println("creditReversalVisa006 completed successfully:  TRID - " + reverseResponse.GetTransactionId())
	}
}

func creditSaleMCFive007(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.02")
	card := testdata.MasterCard5Manual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	transaction := card.ChargeWithAmount(val)
	transaction.WithAllowDuplicates(true)
	transaction.WithRequestMultiUseToken(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("creditSaleMCFive007 completed successfully:  TRID - " + response.GetTransactionId())
		}

	}
}

func creditSaleDiscover008(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.03")
	card := testdata.DiscoverManual(false, false)
	address := base.NewAddressWithStreet("750241234", "6860")
	transaction := card.ChargeWithAmount(val)
	transaction.WithAllowDuplicates(true)
	transaction.WithRequestMultiUseToken(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("creditSaleDiscover008 completed successfully:  TRID - " + response.GetTransactionId())
		}

	}
}

func creditSaleAmex009(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.04")
	card := testdata.AmexManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860 Dallas Pkwy")
	transaction := card.ChargeWithAmount(val)
	transaction.WithAllowDuplicates(true)
	transaction.WithRequestMultiUseToken(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("creditSaleAmex009 completed successfully:  TRID - " + response.GetTransactionId())
		}

	}
}

func creditSaleJCB010(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.05")
	card := testdata.JcbManual(false, false)
	address := base.NewAddressWithStreet("750241234", "6860 Dallas Pkwy")
	transaction := card.ChargeWithAmount(val)
	transaction.WithAllowDuplicates(true)
	transaction.WithRequestMultiUseToken(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("creditSaleJCB010 completed successfully:  TRID - " + response.GetTransactionId())
		}

	}
}

func creditSaleMCTwo011(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.30")
	card := testdata.MasterCard5Manual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	transaction := card.ChargeWithAmount(val)
	transaction.WithAllowDuplicates(true)
	transaction.WithRequestMultiUseToken(true)
	transaction.WithAddress(address)
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		} else {
			fmt.Println("creditSaleMCTwo011 completed successfully:  TRID - " + response.GetTransactionId())
		}

	}
}

func creditAuthAndCaptureVisa012Aand012B(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.06")
	card := testdata.VisaManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860 Dallas Pkwy")
	authorization := card.AuthorizeWithAmount(val, false)
	authorization.WithRequestMultiUseToken(true)
	authorization.WithCurrency("USD")
	authorization.WithAllowDuplicates(true)
	authorization.WithAddress(address)
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditAuthVisa012A completed successfully:  TRID - " + response.GetTransactionId())
	}

	capture := response.CaptureWithAmount(val)
	capture.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	captureResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, capture)
	if err != nil {
		t.Errorf("Capture failed with error: %s", err.Error())
		return
	}

	if captureResponse == nil {
		t.Errorf("Capture did not complete")
		return
	}
	if captureResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected capture response code: %s", captureResponse.GetResponseMessage())
	} else {
		fmt.Println("creditCaptureVisa012B completed successfully:  TRID - " + captureResponse.GetTransactionId())
	}
}

func creditAuthAndCaptureMCFive013Aand013B(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.07")
	card := testdata.MasterCard5Manual(false, false)
	address := base.NewAddressWithStreet("750241234", "6860 Dallas Pkwy")
	authorization := card.AuthorizeWithAmount(val, false)
	authorization.WithRequestMultiUseToken(true)
	authorization.WithCurrency("USD")
	authorization.WithAllowDuplicates(true)
	authorization.WithAddress(address)
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditAuthMCFive013A completed successfully:  TRID - " + response.GetTransactionId())
	}

	capture := response.CaptureWithAmount(val)
	capture.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	captureResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, capture)
	if err != nil {
		t.Errorf("Capture failed with error: %s", err.Error())
		return
	}

	if captureResponse == nil {
		t.Errorf("Capture did not complete")
		return
	}
	if captureResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected capture response code: %s", captureResponse.GetResponseMessage())
	} else {
		fmt.Println("creditCaptureMCFive013B completed successfully:  TRID - " + captureResponse.GetTransactionId())
	}
}

func creditAuthDiscover014A(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.08")
	card := testdata.DiscoverManual(false, false)
	address := base.NewAddressWithStreet("75024", "6860")
	authorization := card.AuthorizeWithAmount(val, false)
	authorization.WithRequestMultiUseToken(true)
	authorization.WithCurrency("USD")
	authorization.WithAllowDuplicates(true)
	authorization.WithAddress(address)
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println(" creditAuthDiscover014A completed successfully:  TRID - " + response.GetTransactionId())
	}

}

func creditAuthMCTwo015A(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("17.09")
	card := testdata.MasterCard2Manual(false, false)
	address := base.NewAddressWithStreet("750241234", "6860 Dallas Pkwy")
	authorization := card.AuthorizeWithAmount(val, false)
	authorization.WithRequestMultiUseToken(true)
	authorization.WithCurrency("USD")
	authorization.WithAllowDuplicates(true)
	authorization.WithAddress(address)
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println(" creditAuthMCTwo015A completed successfully:  TRID - " + response.GetTransactionId())
	}

}

func creditSaleWithGratuityEditVisa016(t *testing.T) {
	card := testdata.VisaManual(false, false)
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("15.12")
	gratuity, _ := stringutils.ToDecimalAmount("3.00")
	charge := card.ChargeWithAmount(amount)
	charge.WithRequestMultiUseToken(true)
	charge.WithCurrency("USD")
	charge.WithAllowDuplicates(true)
	charge.WithCommercialRequest(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	edit := response.Edit()
	edit.WithGratuity(gratuity)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	edit.SetTransactionModifier(transactionmodifier.None)

	editResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if editResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if editResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", editResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleWithGratuityEditVisa016 completed successfully:  TRID - " + editResponse.GetTransactionId())
	}
}

func creditSaleWithGratuityEditMCfive017(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "")
	card := testdata.MasterCard5Manual(false, false)
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("15.13")
	gratuity, _ := stringutils.ToDecimalAmount("3.00")
	charge := card.ChargeWithAmount(amount)
	charge.WithRequestMultiUseToken(true)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithAllowDuplicates(true)
	charge.WithCommercialRequest(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	edit := response.Edit()
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	edit.WithGratuity(gratuity)
	edit.SetTransactionModifier(transactionmodifier.None)

	editResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if editResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if editResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", editResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleWithGratuityEditMCfive017 completed successfully:  TRID - " + editResponse.GetTransactionId())
	}
}

func creditSaleWithGratuityCaptureVisa018(t *testing.T) {
	val, _ := stringutils.ToDecimalAmount("15.11")
	gratuity, _ := stringutils.ToDecimalAmount("3.50")
	address := base.NewAddressWithStreet("75024", "")
	ctx := context.Background()
	card := testdata.VisaManual(false, false)
	authorization := card.ChargeWithAmount(val)
	authorization.WithCurrency("USD")
	authorization.WithRequestMultiUseToken(true)
	authorization.WithGratuity(gratuity)
	authorization.WithAddress(address)
	authorization.WithAllowDuplicates(true)
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditSaleWithGratuityCaptureVisa018 completed successfully:  TRID - " + response.GetTransactionId())
	}
}

func creditSaleWithGratuityCaptureMCfive019(t *testing.T) {
	val, _ := stringutils.ToDecimalAmount("15.12")
	gratuity, _ := stringutils.ToDecimalAmount("3.50")
	ctx := context.Background()
	card := testdata.MasterCard5Manual(false, false)
	authorization := card.ChargeWithAmount(val)
	authorization.WithCurrency("USD")
	authorization.WithGratuity(gratuity)
	authorization.WithRequestMultiUseToken(true)
	authorization.WithAllowDuplicates(true)
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditSaleWithGratuityCaptureMCfive0019 completed successfully:  TRID - " + response.GetTransactionId())
	}
}

func creditSaleLVL2Visa020(t *testing.T) {
	address := base.NewAddressWithStreet("750241234", "6860 Dallas Pkwy")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("112.34")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "B" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.NotUsed)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("ccreditSaleLVL2Visa020 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Visa021(t *testing.T) {
	address := base.NewAddressWithStreet("750241234", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("112.34")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "B" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithTaxAmount(tax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Visa021 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Visa022(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("123.45")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "R" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.TaxExempt)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Visa022 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Visa023(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("134.56")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithTaxAmount(tax)
	edit.WithPoNumber("9876543210")
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Visa023 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2MCFive024(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.06")
	card := testdata.MasterCard5Manual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.NotUsed)
	edit.WithPoNumber("9876543210")
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2MCFive024 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2MCFive025(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.06")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.MasterCard5Manual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithTaxAmount(tax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2MCFive025 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2MCFive026(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.09")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.MasterCard5Manual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithTaxAmount(tax)
	edit.WithPoNumber("9876543210")
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2MCFive026 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2MCFive027(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.09")
	card := testdata.MasterCard5Manual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.TaxExempt)
	edit.WithPoNumber("9876543210")
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2MCFive027 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Amex028(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("110.10")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.AmexManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "0" {
		t.Errorf("Unexpected non-empty commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithTaxAmount(tax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Amex028 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Amex029(t *testing.T) {
	address := base.NewAddressWithStreet("750241234", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("110.11")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.AmexManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "0" {
		t.Errorf("Unexpected non-empty commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithTaxAmount(tax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Amex029 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Amex030(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.12")
	card := testdata.AmexManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "0" {
		t.Errorf("Unexpected non-empty commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.NotUsed)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Amex030 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL2Amex031(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.13")
	card := testdata.AmexManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "0" {
		t.Errorf("Unexpected non-empty commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithTaxType(taxtype.TaxExempt)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL2Amex031 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL3Visa032(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "6860")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.34")
	tax, _ := stringutils.ToDecimalAmount("1")
	discount, _ := stringutils.ToDecimalAmount("1")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "0" {
		t.Errorf("Unexpected non-empty commercial indicator: %s", response.GetCommercialIndicator())
	}
	commercialData := entities.NewCommercialDataWithLevel(taxtype.SalesTax, transactionmodifier.LevelIII)
	commercialData.OrderDate = stringutils.ToStandardDateString(time.Now())
	commercialData.DiscountAmount = discount
	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	edit.WithTaxAmount(tax)
	edit.WithCommercialData(commercialData)
	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL3Visa032 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL3Visa033(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("123.45")
	duty, _ := stringutils.ToDecimalAmount("1")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "R" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	commercialData := entities.NewCommercialDataWithLevel(taxtype.TaxExempt, transactionmodifier.LevelIII)
	commercialData.DutyAmount = duty
	commercialData.DestinationCountryCode = "USA"

	edit := response.Edit()
	edit.WithTaxType(taxtype.TaxExempt)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	edit.WithCommercialData(commercialData)

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL3Visa033 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL3Visa034(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("134.56")
	tax, _ := stringutils.ToDecimalAmount("1")
	discount, _ := stringutils.ToDecimalAmount("1")
	card := testdata.VisaManual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	commercialData := entities.NewCommercialDataWithLevel(taxtype.SalesTax, transactionmodifier.LevelIII)
	commercialData.PONumber = "9876543210"
	commercialData.OrderDate = stringutils.ToStandardDateString(time.Now())
	commercialData.DiscountAmount = discount

	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	edit.WithTaxAmount(tax)
	edit.WithCommercialData(commercialData)

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL3Visa034 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL3MCFive035(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.08")
	tax, _ := stringutils.ToDecimalAmount("1")
	card := testdata.MasterCard5Manual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "S" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	commercialData := entities.NewCommercialDataWithLevel(taxtype.SalesTax, transactionmodifier.LevelIII)
	commercialLineItem := entities.NewCommercialLineItem()
	commercialLineItem.Description = "PRODUCT 1 NOTES"
	commercialLineItem.Quantity, _ = stringutils.ToDecimalAmount("10")
	commercialLineItem.TotalAmount, _ = stringutils.ToDecimalAmount("100")
	commercialLineItem.UnitOfMeasure = "METER"
	commercialData.AddLineItems(commercialLineItem)
	edit := response.Edit()
	edit.WithTaxType(taxtype.SalesTax)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	edit.WithTaxAmount(tax)
	edit.WithCommercialData(commercialData)

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL3MCFive035 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditSaleLVL3MCFive036(t *testing.T) {
	address := base.NewAddressWithStreet("75024", "")
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("111.14")
	card := testdata.MasterCard5Manual(false, false)
	charge := card.ChargeWithAmount(amount)
	charge.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response == nil {
		t.Fatal("Transaction did not complete")
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
	}

	if response.GetCommercialIndicator() != "R" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	commercialData := entities.NewCommercialDataWithLevel(taxtype.TaxExempt, transactionmodifier.LevelIII)
	commercialLineItem := entities.NewCommercialLineItem()
	commercialLineItem.Description = "PRODUCT 1 NOTES"
	commercialLineItem.Quantity, _ = stringutils.ToDecimalAmount("10")
	commercialLineItem.TotalAmount, _ = stringutils.ToDecimalAmount("100")
	commercialLineItem.UnitOfMeasure = "METER"
	commercialData.AddLineItems(commercialLineItem)

	edit := response.Edit()
	edit.WithTaxType(taxtype.TaxExempt)
	edit.WithCommercialData(commercialData)
	edit.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	} else {
		fmt.Println("creditSaleLVL3MCFive036 completed successfully:  TRID - " + cpcResponse.GetTransactionId())
	}
}

func creditIncrementalAuthVisa037(t *testing.T) {

	card := testdata.VisaManual(false, false)
	authAmt, _ := stringutils.ToDecimalAmount("115")
	incrAmt, _ := stringutils.ToDecimalAmount("138")
	ctx := context.Background()
	authorization := card.AuthorizeWithAmount(authAmt, false)
	authorization.WithCurrency("USD")
	authorization.WithCommercialRequest(true)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	authorization.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditIncrementalAuthVisa037-a completed successfully:  TRID - " + response.GetTransactionId())
	}

	incrAuth := rebuilders.FromId(response.GetTransactionId(), authorization.GetPaymentMethod().GetPaymentMethodType()).Increment(incrAmt)
	incrAuth.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	incrResp, err := api.ExecuteGateway[transactions.Transaction](ctx, incrAuth)

	if incrResp == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if incrResp.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditIncrementalAuthVisa037-b completed successfully:  TRID - " + incrResp.GetTransactionId())
	}

	capture := response.CaptureWithAmount(incrAmt)
	capture.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	captureResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, capture)
	if err != nil {
		t.Errorf("Capture failed with error: %s", err.Error())
		return
	}

	if captureResponse == nil {
		t.Errorf("Capture did not complete")
		return
	}
	if captureResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected capture response code: %s", captureResponse.GetResponseMessage())
	} else {
		fmt.Println("creditIncrementalAuthVisa037-c completed successfully:  TRID - " + captureResponse.GetTransactionId())
	}
}

func creditIncrementalAuthMCFive038(t *testing.T) {
	card := testdata.MasterCard5Manual(false, false)
	authAmt, _ := stringutils.ToDecimalAmount("116")
	incrAmt, _ := stringutils.ToDecimalAmount("140")
	ctx := context.Background()
	authorization := card.AuthorizeWithAmount(authAmt, false)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	authorization.WithCurrency("USD")
	authorization.WithCommercialRequest(true)
	authorization.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditIncrementalAuthMCFive038-a completed successfully:  TRID - " + response.GetTransactionId())
	}

	incrAuth := rebuilders.FromId(response.GetTransactionId(), authorization.GetPaymentMethod().GetPaymentMethodType()).Increment(incrAmt)
	incrAuth.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	incrAuth.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.CardHolder, response.GetCardBrandTransactionId())
	incrResp, err := api.ExecuteGateway[transactions.Transaction](ctx, incrAuth)

	if incrResp == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if incrResp.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditIncrementalAuthMCFive038-b completed successfully:  TRID - " + incrResp.GetTransactionId())
	}

	capture := response.CaptureWithAmount(incrAmt)
	capture.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	captureResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, capture)
	if err != nil {
		t.Errorf("Capture failed with error: %s", err.Error())
		return
	}

	if captureResponse == nil {
		t.Errorf("Capture did not complete")
		return
	}
	if captureResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected capture response code: %s", captureResponse.GetResponseMessage())
	} else {
		fmt.Println("creditIncrementalAuthMCFive038-c completed successfully:  TRID - " + captureResponse.GetTransactionId())
	}
}

func creditIncrementalAuthVisa039(t *testing.T) {
	card := testdata.VisaManual(false, false)
	authAmt, _ := stringutils.ToDecimalAmount("117")
	incrAmt, _ := stringutils.ToDecimalAmount("142")
	ctx := context.Background()
	authorization := card.AuthorizeWithAmount(authAmt, false)
	authorization.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	authorization.WithCurrency("USD")
	authorization.WithCommercialRequest(true)
	authorization.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, authorization)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
		return
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		return
	} else {
		fmt.Println("creditIncrementalAuthVisa039-a completed successfully:  TRID - " + response.GetTransactionId())
	}

	incrAuth := rebuilders.FromId(response.GetTransactionId(), authorization.GetPaymentMethod().GetPaymentMethodType()).Increment(incrAmt)
	incrAuth.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	incrResp, err := api.ExecuteGateway[transactions.Transaction](ctx, incrAuth)

	if incrResp == nil {
		t.Errorf("Transaction did not complete")
		return
	}
	if incrResp.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", incrResp.GetResponseMessage())
		return
	} else {
		fmt.Println("creditIncrementalAuthVisa039-b completed successfully:  TRID - " + incrResp.GetTransactionId())
	}

	capture := response.CaptureWithAmount(incrAmt)
	capture.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	captureResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, capture)
	if err != nil {
		t.Errorf("Capture failed with error: %s", err.Error())
		return
	}

	if captureResponse == nil {
		t.Errorf("Capture did not complete")
		return
	}
	if captureResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected capture response code: %s", captureResponse.GetResponseMessage())
	} else {
		fmt.Println("creditIncrementalAuthVisa039-c completed successfully:  TRID - " + captureResponse.GetTransactionId())
	}
}

func creditTimeoutReversalVisa040(t *testing.T) {
	card := testdata.VisaManual(false, false)
	val, _ := stringutils.ToDecimalAmount("10.33")
	ctx := context.Background()
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithCommercialRequest(true)
	transaction.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteTimeoutReversibleGateway(ctx, transaction)
	if err != nil {
		t.Errorf("Credit autoreverse failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else if !response.GetAutoReversed() {
		t.Errorf("Transacction was unable to autoreverse")

	} else {
		fmt.Println("creditTimeoutReversalVisa040 completed successfully:  TRID - " + response.GetTransactionId())
	}
}

func creditCloseBatch(t *testing.T) {
	ctx := context.Background()
	builder := batchservice.CloseBatch()
	_, batchErr := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if batchErr != nil {
		t.Errorf("Unable to close batch")
	} else {
		fmt.Println("Batch successfully closed")
	}
}
