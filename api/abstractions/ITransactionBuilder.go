package abstractions

import (
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/host"
	"github.com/globalpayments/go-sdk/api/entities/enums/hosterror"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	networkentities "github.com/globalpayments/go-sdk/api/network/entities"
	"github.com/globalpayments/go-sdk/api/network/enums/cardissuerentrytag"
	paymentmethods "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
)

type ITransactionBuilder interface {
	IBaseBuilder
	WithServiceCode(serviceCode string)
	WithEcommerceAuthIndicator(ecommerceAuthIndicator string)
	WithEcommerceData1(ecommerceData1 string)
	WithEcommerceData2(ecommerceData2 string)
	WithZipCode(zipCode string)
	GetTransactionType() transactiontype.TransactionType
	SetTransactionType(transactionType transactiontype.TransactionType)
	GetTransactionModifier() transactionmodifier.TransactionModifier
	SetTransactionModifier(transactionModifier transactionmodifier.TransactionModifier)
	GetPaymentMethod() paymentmethods.IPaymentMethod
	SetPaymentMethod(paymentMethod paymentmethods.IPaymentMethod)
	GetSimulatedHostErrors() map[host.Host][]hosterror.HostError
	GetBatchNumber() *int
	GetCompanyId() string
	GetIssuerData() map[cardissuerentrytag.CardIssuerEntryTag]string
	GetFollowOnStan() *int
	GetPriorMessageInformation() *networkentities.PriorMessageInformation
	SetPriorMessageInformation(priorMessageInformation *networkentities.PriorMessageInformation)
	GetSequenceNumber() *int
	GetSystemTraceAuditNumber() *int
	GetUniqueDeviceId() string
	GetTransactionMatchingData() *networkentities.TransactionMatchingData
	IsTerminalError() bool
	WithDescription(description string)
	WithIssuerData(issuerData map[cardissuerentrytag.CardIssuerEntryTag]string)
	WithFollowOnStan(followOnStan *int)
	WithSequenceNumber(sequenceNumber *int)
	WithSystemTraceAuditNumber(systemTraceAuditNumber *int)
	WithUniqueDeviceId(uniqueDeviceId string)
	WithTransactionMatchingData(transactionMatchingData *networkentities.TransactionMatchingData)
	SetTerminalError(terminalError bool)
	SetTaxAmount(taxAmount *decimal.Decimal)
	SetTipAmount(tipAmount *decimal.Decimal)
	SetSurchargeAmount(surchargeAmount *decimal.Decimal)
	WithCashBackAmount(cashBackAmount *decimal.Decimal)
	SetInvoiceNumber(invoiceNumber string)
	SetPosSequenceNumber(sn string)
	SetCvn(cvn string)
	SetAmount(amount *decimal.Decimal)
	SetTagData(tagData string)
	SetEmvMaxPinEntry(emvMaxPinEntry string)
	SetCardSequenceNumber(cardSequenceNumber string)
	SetEcommerceAuthIndicator(ecommerceAuthIndicator string)
	SetEcommerceData1(ecommerceData1 string)
	SetEcommerceData2(ecommerceData2 string)
	SetTransactionDate(transactionDate string)
	SetTransactionTime(transactionTime string)
	SetZipCode(zipCode string)
}
