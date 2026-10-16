package builders

import (
	"context"
	abstractions2 "github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders/validations"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/billing"
	"github.com/globalpayments/go-sdk/api/entities/enums/accounttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/alternativepaymenttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/emvchipcondition"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodusagemode"
	"github.com/globalpayments/go-sdk/api/entities/enums/reversalreasoncode"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/recurring"
	"github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/shopspring/decimal"
)

type ManagementBuilder struct {
	*TransactionBuilder
	AccountType                 accounttype.AccountType
	AlternativePaymentType      alternativepaymenttype.AlternativePaymentType
	Amount                      *decimal.Decimal
	AuthAmount                  *decimal.Decimal
	CardType                    string
	CardBrandTransactionId      string
	CashBackAmount              *decimal.Decimal
	ClerkId                     string
	ClientTransactionId         string
	ConvenienceAmount           *decimal.Decimal
	Currency                    string
	CustomerId                  string
	CustomerIpAddress           string
	CustomerInitiated           bool
	MultiCapture                bool
	Description                 string
	DisputeId                   string
	DynamicDescriptor           string
	ForceToHost                 bool
	Gratuity                    *decimal.Decimal
	IdempotencyKey              string
	InvoiceNumber               string
	MultiCapturePaymentCount    *int
	MultiCaptureSequence        *int
	OrderId                     string
	PayerAuthenticationResponse string
	PoNumber                    string
	ProductId                   string
	ReferenceNumber             string
	ShiftNumber                 string
	SupplementaryData           map[string][][2]string
	SurchargeAmount             *decimal.Decimal
	TagData                     string
	PaymentMethodUsageMode      paymentmethodusagemode.PaymentMethodUsageMode
	TaxAmount                   *decimal.Decimal
	TaxType                     taxtype.TaxType
	Timestamp                   string
	TransactionCount            *int
	TransactionInitiator        storedcredentialinitiator.StoredCredentialInitiator
	TransportData               string
	TotalCredits                *decimal.Decimal
	TotalDebits                 *decimal.Decimal
	TotalAmount                 *decimal.Decimal
	BatchReference              string
	Bills                       []billing.Bill
	EcommerceInfo               entities.EcommerceInfo
	StoredCredential            entities.StoredCredential
	PaymentPurposeCode          string
	DataCollectResponseCode     int
	Approvalcode                string
	SettlementAmount            *decimal.Decimal
	TotalSales                  *decimal.Decimal
	TotalReturns                *decimal.Decimal
	AllowDuplicates             bool
	Customer                    recurring.Customer
	Country                     string
	GenerateReceipt             bool
	UsageMode                   paymentmethodusagemode.PaymentMethodUsageMode
	UsageLimit                  *int
	MiscProductData             []base.Product
	Reference                   string
	CommercialData              *entities.CommercialData
	EmvChipCondition            *emvchipcondition.EmvChipCondition
	ReversalReasonCode          reversalreasoncode.ReversalReasonCode
}

func (a *ManagementBuilder) Execute(ctx context.Context, gateway abstractions2.IPaymentGateway) (abstractions2.ITransaction, error) {
	return gateway.ManageTransaction(ctx, a)
}

func (a *ManagementBuilder) GetAccountType() accounttype.AccountType {
	return a.AccountType
}

func (a *ManagementBuilder) GetEmvChipCondition() *emvchipcondition.EmvChipCondition {
	return a.EmvChipCondition
}

func (ab *ManagementBuilder) WithChipCondition(value emvchipcondition.EmvChipCondition) {
	ab.EmvChipCondition = &value
}

func (a *ManagementBuilder) GetReversalReasonCode() reversalreasoncode.ReversalReasonCode {
	return a.ReversalReasonCode
}

func (a *ManagementBuilder) WithReversalReasonCode(code reversalreasoncode.ReversalReasonCode) {
	a.ReversalReasonCode = code
}

func (mb *ManagementBuilder) WithMiscProductData(values []base.Product) {
	mb.MiscProductData = values
}

// GetAlternativePaymentType gets the alternativePaymentType field.
func (mb *ManagementBuilder) GetAlternativePaymentType() alternativepaymenttype.AlternativePaymentType {
	return mb.AlternativePaymentType
}

// GetAmount gets the amount field.
func (mb *ManagementBuilder) GetAmount() *decimal.Decimal {
	return mb.Amount
}

// GetAuthAmount gets the authAmount field.
func (mb *ManagementBuilder) GetAuthAmount() *decimal.Decimal {
	return mb.AuthAmount
}

// GetBills gets the bills field.
func (mb *ManagementBuilder) GetBills() []billing.Bill {
	return mb.Bills
}

// GetCashBackAmount gets the cashBackAmount field.
func (mb *ManagementBuilder) GetCashBackAmount() *decimal.Decimal {
	return mb.CashBackAmount
}

// GetClerkId gets the clerkId field.
func (mb *ManagementBuilder) GetClerkId() string {
	return mb.ClerkId
}

// GetClientTransactionId gets the clientTransactionId field.
func (mb *ManagementBuilder) GetClientTransactionId() string {
	return mb.ClientTransactionId
}

func (mb *ManagementBuilder) GetTransactionId() string {
	if t, ok := mb.paymentMethod.(*references.TransactionReference); ok {
		return t.TransactionId
	}
	return ""
}

func (mb *ManagementBuilder) GetCardBrandTransactionId() string {
	return mb.CardBrandTransactionId
}

func (mb *ManagementBuilder) GetTransactionInitiator() storedcredentialinitiator.StoredCredentialInitiator {
	return mb.TransactionInitiator
}

// GetConvenienceAmount gets the convenienceAmount field.
func (mb *ManagementBuilder) GetConvenienceAmount() *decimal.Decimal {
	return mb.ConvenienceAmount
}

// GetCurrency gets the currency field.
func (mb *ManagementBuilder) GetCurrency() string {
	return mb.Currency
}

// GetCustomerId gets the customerId field.
func (mb *ManagementBuilder) GetCustomerId() string {
	return mb.CustomerId
}

// GetCustomerIpAddress gets the customerIpAddress field.
func (mb *ManagementBuilder) GetCustomerIpAddress() string {
	return mb.CustomerIpAddress
}

// IsCustomerInitiated gets the customerInitiated field.
func (mb *ManagementBuilder) IsCustomerInitiated() bool {
	return mb.CustomerInitiated
}

// GetDescription gets the description field.
func (mb *ManagementBuilder) GetDescription() string {
	return mb.Description
}

// IsForceToHost gets the forceToHost field.
func (mb *ManagementBuilder) IsForceToHost() bool {
	return mb.ForceToHost
}

// GetGratuity gets the gratuity field.
func (mb *ManagementBuilder) GetGratuity() *decimal.Decimal {
	return mb.Gratuity
}

// GetInvoiceNumber gets the invoiceNumber field.
func (mb *ManagementBuilder) GetInvoiceNumber() string {
	return mb.InvoiceNumber
}

// GetOrderId gets the orderId field.
func (mb *ManagementBuilder) GetOrderId() string {
	return mb.OrderId
}

func (mb *ManagementBuilder) GetCommercialData() *entities.CommercialData {
	return mb.CommercialData
}

// GetPayerAuthenticationResponse gets the payerAuthenticationResponse field.
func (mb *ManagementBuilder) GetPayerAuthenticationResponse() string {
	return mb.PayerAuthenticationResponse
}

// GetPoNumber gets the poNumber field.
func (mb *ManagementBuilder) GetPoNumber() string {
	return mb.PoNumber
}

// GetProductId gets the productId field.
func (mb *ManagementBuilder) GetProductId() string {
	return mb.ProductId
}

// GetReferenceNumber gets the referenceNumber field.
func (mb *ManagementBuilder) GetReferenceNumber() string {
	return mb.ReferenceNumber
}

func (m *ManagementBuilder) GetShiftNumber() string {
	return m.ShiftNumber
}

func (m *ManagementBuilder) GetSupplementaryData() map[string][][2]string {
	return m.SupplementaryData
}

func (m *ManagementBuilder) GetSurchargeAmount() *decimal.Decimal {
	return m.SurchargeAmount
}

func (m *ManagementBuilder) GetTagData() string {
	return m.TagData
}

func (m *ManagementBuilder) GetTaxAmount() *decimal.Decimal {
	return m.TaxAmount
}

func (m *ManagementBuilder) GetTaxType() taxtype.TaxType {
	return m.TaxType
}

func (m *ManagementBuilder) GetTimestamp() string {
	return m.Timestamp
}

func (m *ManagementBuilder) GetTransportData() string {
	return m.TransportData
}

func (m *ManagementBuilder) GetTransactionCount() *int {
	return m.TransactionCount
}

func (m *ManagementBuilder) GetTotalCredits() *decimal.Decimal {
	return m.TotalCredits
}

func (m *ManagementBuilder) GetTotalDebits() *decimal.Decimal {
	return m.TotalDebits
}

func (m *ManagementBuilder) GetCardType() string {
	return m.CardType
}

func (m *ManagementBuilder) WithCommercialData(cd *entities.CommercialData) {
	m.CommercialData = cd
	if cd.CommercialIndicator == transactionmodifier.LevelIII {
		m.SetTransactionModifier(transactionmodifier.LevelIII)
	} else {
		m.SetTransactionModifier(transactionmodifier.LevelII)
	}
}
func (m *ManagementBuilder) WithCardType(ct string) {
	m.CardType = ct
}

func (m *ManagementBuilder) WithSettlementAmount(value *decimal.Decimal) {
	m.SettlementAmount = value

}

func (m *ManagementBuilder) WithPaymentPurposeCode(paymentPurposeCode string) {
	m.PaymentPurposeCode = paymentPurposeCode

}

func (b *ManagementBuilder) WithApprovalCode(value string) {
	b.Approvalcode = value

}

func (b *ManagementBuilder) WithDataCollectResponseCode(value int) {
	b.DataCollectResponseCode = value

}

func (b *ManagementBuilder) WithAlternativePaymentType(value alternativepaymenttype.AlternativePaymentType) {
	b.AlternativePaymentType = value

}

func (b *ManagementBuilder) WithAmount(value *decimal.Decimal) {
	b.Amount = value

}

func (b *ManagementBuilder) WithAuthAmount(value *decimal.Decimal) {
	b.AuthAmount = value

}

func (b *ManagementBuilder) WithBatchTotals(transactionCount int, totalDebits, totalCredits *decimal.Decimal) {
	b.TransactionCount = &transactionCount
	b.TotalDebits = totalDebits
	b.TotalCredits = totalCredits

}

func (b *ManagementBuilder) WithBatchTotalsAmount(totalAmount, totalDebits, totalCredits *decimal.Decimal) {
	b.TotalAmount = totalAmount
	b.TotalDebits = totalDebits
	b.TotalCredits = totalCredits

}

func (b *ManagementBuilder) WithBatchTotalTransaction(transactionCount int, totalSales, totalReturns *decimal.Decimal) {
	b.TransactionCount = &transactionCount
	b.TotalSales = totalSales
	b.TotalReturns = totalReturns

}

func (b *ManagementBuilder) WithBatchReference(value string) {
	b.BatchReference = value

}

func (b *ManagementBuilder) WithBills(bills ...billing.Bill) {
	b.Bills = bills

}

func (b *ManagementBuilder) WithCardBrandStorage(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator) {
	b.WithCardBrandStorageAndTransactionId(transactionInitiator, "")
}

func (b *ManagementBuilder) WithCardBrandStorageAndTransactionId(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator, value string) {
	b.TransactionInitiator = transactionInitiator
	b.CardBrandTransactionId = value
}

func (b *ManagementBuilder) WithMultiCapture(sequence *int) {
	b.WithMultiCaptureAndPaymentCount(sequence, nil)
}
func (b *ManagementBuilder) WithMultiCaptureAndPaymentCount(sequence, paymentCount *int) {
	b.MultiCapture = true
	defaultVal := 1
	if sequence != nil {
		b.MultiCaptureSequence = sequence
	} else {
		b.MultiCaptureSequence = &defaultVal
	}
	if paymentCount != nil {
		b.MultiCapturePaymentCount = paymentCount
	} else {
		b.MultiCapturePaymentCount = &defaultVal
	}

}

func (b *ManagementBuilder) WithCashBackAmount(value *decimal.Decimal) {
	b.CashBackAmount = value

}

func (b *ManagementBuilder) WithClerkId(value string) {
	b.ClerkId = value

}

func (b *ManagementBuilder) WithClientTransactionId(value string) {
	b.ClientTransactionId = value

}

func (b *ManagementBuilder) WithConvenienceAmt(value *decimal.Decimal) {
	b.ConvenienceAmount = value

}

func (b *ManagementBuilder) WithCurrency(value string) {
	b.Currency = value

}

func (b *ManagementBuilder) WithCustomerId(value string) {
	b.CustomerId = value

}

func (b *ManagementBuilder) WithCustomerIpAddress(value string) {
	b.CustomerIpAddress = value

}

func (b *ManagementBuilder) WithCustomerInitiated(value bool) {
	b.CustomerInitiated = value

}

func (m *ManagementBuilder) WithDescription(value string) {
	m.Description = value

}

func (m *ManagementBuilder) WithDisputeId(value string) {
	m.DisputeId = value

}

func (m *ManagementBuilder) WithDynamicDescriptor(value string) {
	m.DynamicDescriptor = value

}

func (m *ManagementBuilder) WithForceToHost(value bool) {
	m.ForceToHost = value

}

func (m *ManagementBuilder) WithGratuity(value *decimal.Decimal) {
	m.Gratuity = value

}

func (m *ManagementBuilder) WithInvoiceNumber(value string) {
	m.InvoiceNumber = value

}

func (m *ManagementBuilder) WithIdempotencyKey(value string) {
	m.IdempotencyKey = value

}

func (m *ManagementBuilder) WithPaymentMethod(value abstractions.IPaymentMethod) {
	m.paymentMethod = value

}

func (m *ManagementBuilder) WithPayerAuthenticationResponse(value string) {
	m.PayerAuthenticationResponse = value

}

func (m *ManagementBuilder) WithPoNumber(value string) {
	m.transactionModifier = transactionmodifier.LevelII
	m.PoNumber = value

}

func (m *ManagementBuilder) WithPosSequenceNumber(value string) {
	m.posSequenceNumber = value

}

func (mb *ManagementBuilder) WithProductId(value string) {
	mb.ProductId = value
}

func (mb *ManagementBuilder) WithReferenceNumber(value string) {
	mb.ReferenceNumber = value
}

func (mb *ManagementBuilder) WithShiftNumber(value string) {
	mb.ShiftNumber = value
}

func (mb *ManagementBuilder) WithEcommerceInfo(value entities.EcommerceInfo) {
	mb.EcommerceInfo = value
}

func (mb *ManagementBuilder) WithStoredCredential(value entities.StoredCredential) {
	mb.StoredCredential = value
}

func (mb *ManagementBuilder) WithSupplementaryData(typeKey string, values ...string) {
	if mb.SupplementaryData == nil {
		mb.SupplementaryData = make(map[string][][2]string)
	}

	if _, ok := mb.SupplementaryData[typeKey]; !ok {
		mb.SupplementaryData[typeKey] = [][2]string{}
	}

	mb.SupplementaryData[typeKey] = append(mb.SupplementaryData[typeKey], [2]string{values[0], values[1]})
}

func (mb *ManagementBuilder) WithSurchargeAmount(value *decimal.Decimal) {
	mb.SurchargeAmount = value
}

func (mb *ManagementBuilder) WithTagData(value string) {
	mb.TagData = value
}

func (mb *ManagementBuilder) WithPaymentMethodUsageMode(value paymentmethodusagemode.PaymentMethodUsageMode) {
	mb.PaymentMethodUsageMode = value
}

func (mb *ManagementBuilder) WithTaxAmount(value *decimal.Decimal) {
	mb.TaxAmount = value
}

func (mb *ManagementBuilder) WithTaxType(value taxtype.TaxType) {
	mb.TaxType = value
}

func (mb *ManagementBuilder) WithTimestamp(value string) {
	mb.Timestamp = value
}

func (mb *ManagementBuilder) WithTransportData(value string) {
	mb.TransportData = value
}

func (mb *ManagementBuilder) WithReference(value string) {
	mb.Reference = value
}

func (a *ManagementBuilder) WithAccountType(value accounttype.AccountType) {
	a.AccountType = value
}

func NewManagementBuilder(transactionType transactiontype.TransactionType) *ManagementBuilder {
	return &ManagementBuilder{
		TransactionBuilder: NewTransactionBuilder(transactionType, nil),
	}
}

func NewManagementBuilderWithPaymentMethod(transactionType transactiontype.TransactionType, paymentMethod abstractions.IPaymentMethod) *ManagementBuilder {
	return &ManagementBuilder{
		TransactionBuilder: NewTransactionBuilder(transactionType, paymentMethod),
	}
}

func (mb *ManagementBuilder) setupValidations(validations *validations.Validations) {
	transactionTypes := []int64{
		transactiontype.Capture.LongValue(),
		transactiontype.Edit.LongValue(),
		transactiontype.Hold.LongValue(),
		transactiontype.Release.LongValue(),
	}
	for _, tType := range transactionTypes {
		validations.Of(tType).Check("PaymentMethod").IsNotNull("")
	}

	validations.Of(transactiontype.Refund.LongValue()).
		When("Amount").IsNotNull("").
		Check("Currency").IsNotNull("")

	validations.Of(transactiontype.VerifySignature.LongValue()).
		Check("PayerAuthenticationResponse").IsNotNull("").
		Check("Amount").IsNotNull("").
		Check("Currency").IsNotNull("").
		Check("OrderId").IsNotNull("")

	tokenTypes := []int64{
		transactiontype.TokenDelete.LongValue(),
		transactiontype.TokenUpdate.LongValue(),
	}
	for _, tType := range tokenTypes {
		validations.Of(tType).Check("PaymentMethod").IsNotNull("")

	}

	validations.Of(transactiontype.PayLinkUpdate.LongValue()).
		Check("Amount").IsNotNull("").
		Check("PayByLinkData").IsNotNull("")
	// Assuming propertyOf and isNotNull() methods for nested property checks
	validations.Of(transactiontype.PayLinkUpdate.LongValue()).
		Check("PayByLinkData.UsageMode").IsNotNull("").
		Check("PayByLinkData.UsageLimit").IsNotNull("").
		Check("PayByLinkData.Type").IsNotNull("")

	validations.Of(transactiontype.Reversal.LongValue()).
		Check("PaymentMethod").IsNotNull("")

}
