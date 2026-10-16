package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/securethreedversion"
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestPorticoCreditTests(t *testing.T) {
	var card paymentmethods.CreditCardData
	var clientTxnID string
	var commercialData entities.CommercialData
	var additionalTaxDetails entities.AdditionalTaxDetails
	var commercialLineItem entities.CommercialLineItem
	var discountDetails entities.DiscountDetails

	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MTeSAQAfG1UA9qQDrzl-kz4toXvARyieptFwSKP24w"
	config.VersionNumber = "3026"
	config.EnableLogging = true
	config.SafDataSupported = false
	err := api.ConfigureService(config, "default")
	if err != nil {
		t.Errorf("Failed to configure service with error: %s", err.Error())
	}

	card = *paymentmethods.NewCreditCardData()
	card.SetNumber("4111111111111111")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")

	clientTxnID = generateRandomClientTransactionID()

	commercialData = *entities.NewCommercialDataWithLevel(taxtype.SalesTax, transactionmodifier.LevelIII)
	commercialData.PONumber = "9876543210"
	val, _ := stringutils.ToDecimalAmount("10")
	commercialData.TaxAmount = val
	commercialData.DestinationPostalCode = "85212"
	commercialData.DestinationCountryCode = "USA"
	commercialData.OriginPostalCode = "22193"
	commercialData.SummaryCommodityCode = "SSC"
	commercialData.CustomerReferenceId = "UVATREF162"
	commercialData.OrderDate = stringutils.ToStandardDateString(time.Now())
	commercialData.FreightAmount, _ = stringutils.ToDecimalAmount("10")
	commercialData.DutyAmount, _ = stringutils.ToDecimalAmount("10")

	additionalTaxDetails = *entities.NewAdditionalTaxDetails()
	additionalTaxDetails.TaxAmount, _ = stringutils.ToDecimalAmount("10")
	additionalTaxDetails.TaxRate, _ = stringutils.ToDecimalAmount("10")
	commercialData.AdditionalTaxDetails = &additionalTaxDetails

	commercialLineItem = *entities.NewCommercialLineItem()
	commercialLineItem.Description = "PRODUCT 1 NOTES"
	commercialLineItem.ProductCode = "PRDCD1"
	commercialLineItem.UnitCost, _ = stringutils.ToDecimalAmount(".01")
	commercialLineItem.Quantity, _ = stringutils.ToDecimalAmount("1")
	commercialLineItem.UnitOfMeasure = "METER"
	commercialLineItem.TotalAmount, _ = stringutils.ToDecimalAmount("10")

	discountDetails = *entities.NewDiscountDetails()
	discountDetails.DiscountAmount, _ = stringutils.ToDecimalAmount("1")
	commercialLineItem.DiscountDetails = &discountDetails

	commercialData.AddLineItems(&commercialLineItem)

	creditSaleWithCard(t, card, clientTxnID, commercialData)
	creditVerify(t, card)
	creditAuthorization(t, card, clientTxnID)
	creditReverseWithCard(t, card, commercialData)
	creditVoidFromTransactionId(t, card, clientTxnID)
	creditAuthReversalOnGatewayTimeout(t, card)
	creditLevelIITransaction(t, card)
	creditLevelIIITransaction(t, card)
	test3DSecrueV1(t, card)
	test3DSecrueV2(t, card)
}

func generateRandomClientTransactionID() string {
	randomID := rand.Intn(999999-10000) + 10000
	return strconv.Itoa(randomID)
}

func creditSaleWithCard(t *testing.T, card paymentmethods.CreditCardData, clientTxnID string, commercialData entities.CommercialData) {
	val, _ := stringutils.ToDecimalAmount("15")
	ctx := context.Background()
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithClientTransactionId(clientTxnID)
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		}
		tr := response.GetTransactionReference()
		if tr == nil {
			t.Errorf("Client reference not returned")
		}
		if tr.ClientTransactionId != clientTxnID {
			t.Errorf("Client transaction ID mismatch. Expected: %s, Got: %s", clientTxnID, tr.ClientTransactionId)
		}
	}
}

func creditVerify(t *testing.T, card paymentmethods.CreditCardData) {
	transaction := card.Verify()
	ctx := context.Background()
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit verify failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		}
	}
}

func creditAuthReversalOnGatewayTimeout(t *testing.T, card paymentmethods.CreditCardData) {
	val, _ := stringutils.ToDecimalAmount("911")
	ctx := context.Background()
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithClientTransactionId("9873212654")
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteTimeoutReversibleGateway(ctx, transaction)
	if err != nil {
		t.Errorf("Credit autoreverse failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if !response.GetAutoReversed() {
			t.Errorf("Transacction was unable to autoreverse")
		}
	}
}

func creditAuthorization(t *testing.T, card paymentmethods.CreditCardData, clientTxnID string) {
	val, _ := stringutils.ToDecimalAmount("14")
	ctx := context.Background()
	authorization := card.AuthorizeWithAmount(val, false)
	authorization.WithCurrency("USD")
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
	}

	captureAmount, _ := stringutils.ToDecimalAmount("16")
	gratuity, _ := stringutils.ToDecimalAmount("2")
	capture := response.CaptureWithAmount(captureAmount)
	capture.WithGratuity(gratuity)
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
	}
}

func creditReverseWithCard(t *testing.T, card paymentmethods.CreditCardData, commercialData entities.CommercialData) {
	val, _ := stringutils.ToDecimalAmount("18")
	ctx := context.Background()
	transaction := card.ChargeWithAmount(val)
	transaction.WithCurrency("USD")
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Credit sale failed with error: %s", err.Error())
	}
	if response == nil {
		t.Errorf("Transaction did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code: %s", response.GetResponseMessage())
		}
	}

	reversalVal, _ := stringutils.ToDecimalAmount("18")
	reversalTransaction := card.ReverseWithAmount(reversalVal)
	reversalTransaction.WithCurrency("USD")
	reversalResponse, revErr := api.ExecuteGateway[transactions.Transaction](ctx, reversalTransaction)
	if revErr != nil {
		t.Errorf("Credit reversal failed with error: %s", revErr.Error())
	}
	if reversalResponse == nil {
		t.Errorf("Reversal did not complete")
	} else {
		if reversalResponse.GetResponseCode() != "00" {
			t.Errorf("Unexpected reversal response code: %s", reversalResponse.GetResponseMessage())
		}
	}

}

func creditVoidFromTransactionId(t *testing.T, card paymentmethods.CreditCardData, clientTxnID string) {
	val, _ := stringutils.ToDecimalAmount("10")
	ctx := context.Background()
	transaction := card.AuthorizeWithAmount(val, false)
	transaction.WithCurrency("USD")
	transaction.WithAllowDuplicates(true)
	response, err := api.ExecuteGateway[transactions.Transaction](ctx, transaction)
	if err != nil {
		t.Errorf("Authorization failed with error: %s", err.Error())
	}

	if response == nil {
		t.Errorf("Authorization did not complete")
	} else {
		if response.GetResponseCode() != "00" {
			t.Errorf("Unexpected response code during authorization: %s", response.GetResponseMessage())
		}

		transactionId := response.GetTransactionId()
		voidTransaction := rebuilders.FromId(transactionId, transaction.GetPaymentMethod().GetPaymentMethodType()).Void(nil, false)
		voidResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, voidTransaction)
		if err != nil {
			t.Errorf("Void transaction failed with error: %s", err.Error())
		}

		if voidResponse == nil {
			t.Errorf("Void transaction did not complete")
		} else {
			if voidResponse.GetResponseCode() != "00" {
				t.Errorf("Unexpected response code during void: %s", voidResponse.GetResponseMessage())
			}
		}
	}
}

func creditLevelIITransaction(t *testing.T, card paymentmethods.CreditCardData) {
	address := base.NewAddressWithStreet("750241234", "6860 Dallas Pkwy")

	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("112.34")
	charge := card.ChargeWithAmount(amount)
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

	if response.GetCommercialIndicator() != "B" {
		t.Errorf("Unexpected commercial indicator: %s", response.GetCommercialIndicator())
	}

	edit := response.Edit()
	edit.WithPoNumber("9876543210")
	edit.WithTaxType(taxtype.NotUsed)

	cpcResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, edit)
	if err != nil {
		t.Fatalf("Edit failed with error: %s", err.Error())
	}

	if cpcResponse == nil {
		t.Fatal("Edit did not complete")
	}

	if cpcResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", cpcResponse.GetResponseMessage())
	}
}

func creditLevelIIITransaction(t *testing.T, card paymentmethods.CreditCardData) {
	address := base.NewAddress("75024")
	address.StreetAddr1 = "6860"

	commercialData := entities.NewCommercialDataWithLevel(taxtype.SalesTax, transactionmodifier.LevelIII)
	commercialData.PONumber = "9876543210"
	commercialData.TaxAmount, _ = stringutils.ToDecimalAmount("10")
	commercialData.DestinationPostalCode = "85212"
	commercialData.DestinationCountryCode = "USA"
	commercialData.OriginPostalCode = "22193"
	commercialData.SummaryCommodityCode = "SSC"
	commercialData.CustomerReferenceId = "UVATREF162"
	commercialData.OrderDate = stringutils.ToStandardDateString(time.Now())
	commercialData.FreightAmount, _ = stringutils.ToDecimalAmount("10")
	commercialData.DutyAmount, _ = stringutils.ToDecimalAmount("10")

	additionalTaxDetails := entities.NewAdditionalTaxDetails()
	additionalTaxDetails.TaxAmount, _ = stringutils.ToDecimalAmount("10")
	additionalTaxDetails.TaxRate, _ = stringutils.ToDecimalAmount("10")
	commercialData.AdditionalTaxDetails = additionalTaxDetails

	commercialLineItem := entities.NewCommercialLineItem()
	commercialLineItem.Description = "PRODUCT 1 NOTES"
	commercialLineItem.ProductCode = "PRDCD1"
	commercialLineItem.Quantity, _ = stringutils.ToDecimalAmount("10") // Changed to 1 from 1000
	commercialLineItem.UnitOfMeasure = "METER"

	discountDetails := entities.NewDiscountDetails()
	discountDetails.DiscountAmount, _ = stringutils.ToDecimalAmount("10") // Changed to 10 from 1
	commercialLineItem.DiscountDetails = discountDetails

	commercialData.AddLineItems(commercialLineItem)

	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("114.12") // Changed to match the working XML
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithAddress(address)
	charge.WithCommercialRequest(true)
	charge.WithAllowDuplicates(true)

	chargeResponse, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if chargeResponse == nil {
		t.Fatal("Transaction did not complete")
	}

	if chargeResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected response code: %s", chargeResponse.GetResponseMessage())
	}

	edit := chargeResponse.Edit()
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
	}
}

func test3DSecrueV1(t *testing.T, card paymentmethods.CreditCardData) {
	ecom := &entities.ThreeDSecure{
		Cavv:    "XXXXf98AAajXbDRg3HSUMAACAAA=",
		Xid:     "0l35fwh1sys3ojzyxelu4ddhmnu5zfke5vst",
		Eci:     "5",
		Version: securethreedversion.ONE,
	}
	card.ThreeDSecure = ecom
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("10")
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithInvoiceNumber("1234567890")
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", response.GetResponseMessage())
	}
}

func test3DSecrueV2(t *testing.T, card paymentmethods.CreditCardData) {
	ecom := &entities.ThreeDSecure{
		Cavv:    "XXXXf98AAajXbDRg3HSUMAACAAA=",
		Xid:     "0l35fwh1sys3ojzyxelu4ddhmnu5zfke5vst",
		Eci:     "5",
		Version: securethreedversion.TWO,
	}
	card.ThreeDSecure = ecom
	ctx := context.Background()
	amount, _ := stringutils.ToDecimalAmount("10")
	charge := card.ChargeWithAmount(amount)
	charge.WithCurrency("USD")
	charge.WithInvoiceNumber("1234567890")
	charge.WithAllowDuplicates(true)

	response, err := api.ExecuteGateway[transactions.Transaction](ctx, charge)
	if err != nil {
		t.Fatalf("Charge failed with error: %s", err.Error())
	}

	if response.GetResponseCode() != "00" {
		t.Errorf("Unexpected edit response code: %s", response.GetResponseMessage())
	}
}
