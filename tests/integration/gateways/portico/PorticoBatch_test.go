package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/services/batchservice"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"strings"
	"testing"
	"time"
)

const BatchNotOpen = "Transaction was rejected because it requires a batch to be open."
const BatchEmpty = "Batch close was rejected because no transactions are associated with the currently open batch"

func TestPorticoBatchTests(t *testing.T) {
	var card paymentmethods.CreditCardData
	var commercialData entities.CommercialData
	var additionalTaxDetails entities.AdditionalTaxDetails
	var commercialLineItem entities.CommercialLineItem
	var discountDetails entities.DiscountDetails

	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MaePAQBr-1QAqjfckFC8FTbRTT120bVQUlfVOjgCBw"
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

	batchClose(t, card)
}

func batchClose(t *testing.T, card paymentmethods.CreditCardData) {
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
	builder := batchservice.CloseBatch()
	batchResponse, batchErr := api.ExecuteGateway[transactions.Transaction](ctx, builder)
	if batchErr != nil {
		errString := batchErr.Error()
		if !strings.Contains(errString, BatchNotOpen) && !strings.Contains(errString, BatchEmpty) {
			t.Errorf("CloseBatch failed with error: %s", batchErr.Error())
			return
		}
	}
	if batchResponse.GetResponseCode() != "00" {
		t.Errorf("Unexpected batch response code: %s", batchResponse.GetResponseMessage())
		return
	}
	batchSummary := batchResponse.GetBatchSummary()
	if batchSummary.BatchId == nil {
		t.Errorf("No Batch ID returned")
		return
	}
}
