package certifications

import (
	"context"
	"fmt"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"github.com/globalpayments/go-sdk/tests/integration/gateways/terminals/hpa"
	"github.com/globalpayments/go-sdk/tests/testdata"
	"testing"
)

func TestPorticoEcommerceCofCert(t *testing.T) {

	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MXDMBQDwa3IAA4GV7NGMqQA_wFR3_TNeamFWoNUu_Q"
	config.VersionNumber = "3026"
	config.EnableLogging = true
	config.SafDataSupported = false
	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	cofMC(t)
	cofVisa(t)
	cofVisaPreAuth(t)
	cofDiscover(t)
	cofAmex(t)
	cofVisaPurchase2(t)

}

func cofMC(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("5.06")
	card := testdata.CofM2Card(false, false)
	address := base.NewAddressWithStreet("", "3032920104 CORPORATE SQ")
	builder := card.ChargeWithAmount(val)
	builder.WithAllowDuplicates(true)
	builder.WithRequestMultiUseToken(true)
	builder.WithCommercialRequest(true)
	builder.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder.WithAddress(address)
	builder.WithCardBrandStorage(storedcredentialinitiator.CardHolder)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	} else if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		return

	} else {
		fmt.Println("COF test 1 part 1 complete:  TRID - " + response.GetTransactionId())
	}
	tokenCard := paymentmethods.NewCreditCardData()
	tokenCard.SetToken(response.Token)
	val2, _ := stringutils.ToDecimalAmount("6.25")
	builder2 := tokenCard.ChargeWithAmount(val2)
	builder2.WithAllowDuplicates(true)
	builder2.WithCommercialRequest(true)
	builder2.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder2.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.Merchant, response.GetCardBrandTransactionId())
	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, builder2)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response2 == nil {
		t.Errorf("COF trsansaction did not complete")
		return
	} else if response2.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response2.GetResponseCode())
		return

	} else {
		fmt.Println("COF test 1 completed successful:  TRID - " + response2.GetTransactionId())
	}
}

func cofVisa(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("5.01")
	card := testdata.CofV2Card(false, false)
	address := base.NewAddressWithStreet("", "3032920104 CORPORATE SQ")

	builder := card.ChargeWithAmount(val)
	builder.WithAllowDuplicates(true)
	builder.WithRequestMultiUseToken(true)
	builder.WithAddress(address)
	builder.WithCommercialRequest(true)
	builder.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder.WithCardBrandStorage(storedcredentialinitiator.CardHolder)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	} else if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		return
	} else {
		fmt.Println("COF test 2 part 1 complete:  TRID - " + response.GetTransactionId())
	}

	tokenCard := paymentmethods.NewCreditCardData()
	tokenCard.SetToken(response.Token)
	val2, _ := stringutils.ToDecimalAmount("5.01")
	builder2 := tokenCard.ChargeWithAmount(val2)
	builder2.WithAllowDuplicates(true)
	builder2.WithCommercialRequest(true)
	builder2.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder2.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.Merchant, response.GetCardBrandTransactionId())

	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, builder2)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response2 == nil {
		t.Errorf("COF transaction did not complete")
		return
	} else if response2.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response2.GetResponseCode())
		return
	} else {
		fmt.Println("COF test 2 completed successful:  TRID - " + response2.GetTransactionId())
	}
}

func cofVisaPreAuth(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("5.51")
	card := testdata.CofV2Card(false, false)
	address := base.NewAddressWithStreet("", "3032920104 CORPORATE SQ")

	// Initial pre-auth
	builder := card.AuthorizeWithAmount(val, false)
	builder.WithAllowDuplicates(true)
	builder.WithRequestMultiUseToken(true)
	builder.WithAddress(address)
	builder.WithCommercialRequest(true)
	builder.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder.WithCardBrandStorage(storedcredentialinitiator.CardHolder)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Pre-auth failed with error: %s", err.Error())
		return
	}

	if response == nil || response.GetResponseCode() != "00" {
		t.Errorf("Pre-auth failed: %s", response.GetResponseCode())
		return
	} else {
		fmt.Println("COF test 3 part 1 complete:  TRID - " + response.GetTransactionId())
	}

	// Initial capture
	builder2 := response.CaptureWithAmount(val)
	builder2.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, builder2)
	if err != nil || response2 == nil || response2.GetResponseCode() != "00" {
		t.Errorf("Capture failed: %s", err.Error())
		return
	} else {
		fmt.Println("COF test 3 part 2 complete:  TRID - " + response2.GetTransactionId())
	}

	// Subsequent pre-auth using stored credential
	tokenCard := paymentmethods.NewCreditCardData()
	tokenCard.SetToken(response.Token)
	builder3 := tokenCard.AuthorizeWithAmount(val, false)
	builder3.WithAllowDuplicates(true)
	builder3.WithCommercialRequest(true)
	builder3.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder3.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.Merchant, response.GetCardBrandTransactionId())

	response3, err := api.ExecuteGateway[transactions.Transaction](ctx, builder3)
	if err != nil || response3 == nil || response3.GetResponseCode() != "00" {
		t.Errorf("Subsequent pre-auth failed: %s", err.Error())
		return
	} else {
		fmt.Println("COF test 3 part 3 complete:  TRID - " + response3.GetTransactionId())
	}

	// Subsequent capture
	builder4 := response3.CaptureWithAmount(val)
	builder4.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	response4, err := api.ExecuteGateway[transactions.Transaction](ctx, builder4)
	if err != nil || response4 == nil || response4.GetResponseCode() != "00" {
		t.Errorf("Subsequent capture failed: %s", err.Error())
		return
	}

	fmt.Println("COF Visa pre-auth test 3 completed successfully:  TRID - " + response4.GetTransactionId())
}

func cofDiscover(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("5.00")
	card := testdata.CofD1Card(false, false)
	address := base.NewAddressWithStreet("", "60015 2500 Lake Cook Road")

	builder := card.ChargeWithAmount(val)
	builder.WithAllowDuplicates(true)
	builder.WithRequestMultiUseToken(true)
	builder.WithCommercialRequest(true)
	builder.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder.WithAddress(address)
	builder.WithCardBrandStorage(storedcredentialinitiator.CardHolder)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	} else if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		return
	} else {
		fmt.Println("COF test 4 part 1 complete:  TRID - " + response.GetTransactionId())
	}

	tokenCard := paymentmethods.NewCreditCardData()
	tokenCard.SetToken(response.Token)
	val2, _ := stringutils.ToDecimalAmount("8.00")
	builder2 := tokenCard.ChargeWithAmount(val2)
	builder2.WithAllowDuplicates(true)
	builder2.WithCommercialRequest(true)
	builder2.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder2.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.Merchant, response.GetCardBrandTransactionId())

	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, builder2)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response2 == nil {
		t.Errorf("COF transaction did not complete")
		return
	} else if response2.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response2.GetResponseCode())
		return
	} else {
		fmt.Println("COF Discover test completed successfully:  TRID -" + response2.GetTransactionId())
	}
}

func cofAmex(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("4.44")
	card := testdata.CofA1Card(false, false)
	address := base.NewAddressWithStreet("", "3032920104 CORPORATE SQ")

	builder := card.ChargeWithAmount(val)
	builder.WithAllowDuplicates(true)
	builder.WithRequestMultiUseToken(true)
	builder.WithAddress(address)
	builder.WithCommercialRequest(true)
	builder.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder.WithCardBrandStorage(storedcredentialinitiator.CardHolder)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	} else if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		return
	} else {
		fmt.Println("COF test 5 part 1 complete:  TRID - " + response.GetTransactionId())
	}

	tokenCard := paymentmethods.NewCreditCardData()
	tokenCard.SetToken(response.Token)
	builder2 := tokenCard.ChargeWithAmount(val)
	builder2.WithAllowDuplicates(true)
	builder2.WithCommercialRequest(true)
	builder2.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder2.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.Merchant, response.GetCardBrandTransactionId())

	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, builder2)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response2 == nil {
		t.Errorf("COF transaction did not complete")
		return
	} else if response2.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response2.GetResponseCode())
		return
	} else {
		fmt.Println("COF Amex test completed successfully:  TRID -" + response2.GetTransactionId())
	}
}

func cofVisaPurchase2(t *testing.T) {
	ctx := context.Background()
	val, _ := stringutils.ToDecimalAmount("5.25")
	card := testdata.CofV2Card(false, false)
	address := base.NewAddressWithStreet("", "3032920104 CORPORATE SQ")

	builder := card.ChargeWithAmount(val)
	builder.WithAllowDuplicates(true)
	builder.WithRequestMultiUseToken(true)
	builder.WithAddress(address)
	builder.WithCommercialRequest(true)
	builder.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder.WithCardBrandStorage(storedcredentialinitiator.CardHolder)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
		return
	} else if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response.GetResponseCode())
		return
	} else {
		fmt.Println("COF test visa 2 part 1 complete:  TRID - " + response.GetTransactionId())
	}

	tokenCard := paymentmethods.NewCreditCardData()
	tokenCard.SetToken(response.Token)
	val2, _ := stringutils.ToDecimalAmount("7.25")
	builder2 := tokenCard.ChargeWithAmount(val2)
	builder2.WithAllowDuplicates(true)
	builder2.WithCommercialRequest(true)
	builder2.WithClientTransactionId(hpa.NewRandomIdProvider().GetRequestIdString())
	builder2.WithCardBrandStorageAndTransactionId(storedcredentialinitiator.Merchant, response.GetCardBrandTransactionId())

	response2, err := api.ExecuteGateway[transactions.Transaction](ctx, builder2)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response2 == nil {
		t.Errorf("COF transaction did not complete")
		return
	} else if response2.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", response2.GetResponseCode())
		return
	} else {
		fmt.Println("COF test visa 2 completed successful:  TRID - " + response2.GetTransactionId())
	}
}
