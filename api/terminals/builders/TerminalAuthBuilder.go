package builders

import (
	"context"
	"github.com/globalpayments/go-sdk/api/builders/validations"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/currencytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
	"github.com/shopspring/decimal"
)

type TerminalAuthBuilder struct {
	*TerminalBuilder
	Address                *base.Address
	AllowDuplicates        bool
	Amount                 *decimal.Decimal
	AuthCode               string
	CardBrandStorage       storedcredentialinitiator.StoredCredentialInitiator
	CardBrandTransactionId string
	CashBackAmount         *decimal.Decimal
	CommercialRequest      bool
	Currency               currencytype.CurrencyType
	CustomerCode           string
	Gratuity               *decimal.Decimal
	InvoiceNumber          string
	PoNumber               string
	RequestMultiUseToken   bool
	SignatureCapture       bool
	TaxAmount              *decimal.Decimal
	TaxExempt              string
	TaxExemptId            string
	TokenRequest           *int
	TokenValue             string
	TransactionId          string
	GiftTransactionType    transactiontype.TransactionType
}

func (tb *TerminalAuthBuilder) GetAddress() *base.Address {
	return tb.Address
}

func (tb *TerminalAuthBuilder) GetAllowDuplicates() bool {
	return tb.AllowDuplicates
}

func (tb *TerminalAuthBuilder) GetAmount() *decimal.Decimal {
	return tb.Amount
}

func (tb *TerminalAuthBuilder) GetCashBackAmount() *decimal.Decimal {
	return tb.CashBackAmount
}

func (tb *TerminalAuthBuilder) GetCardBrandStorage() storedcredentialinitiator.StoredCredentialInitiator {
	return tb.CardBrandStorage
}

func (tb *TerminalAuthBuilder) GetCardBrandTransactionId() string {
	return tb.CardBrandTransactionId
}

func (tb *TerminalAuthBuilder) GetCommercialRequest() bool {
	return tb.CommercialRequest
}

func (tb *TerminalAuthBuilder) GetCurrency() currencytype.CurrencyType {
	return tb.Currency
}

func (tb *TerminalAuthBuilder) GetCustomerCode() string {
	return tb.CustomerCode
}

func (tb *TerminalAuthBuilder) GetGratuity() *decimal.Decimal {
	return tb.Gratuity
}

func (tb *TerminalAuthBuilder) GetInvoiceNumber() string {
	return tb.InvoiceNumber
}

func (tb *TerminalAuthBuilder) IsRequestMultiUseToken() bool {
	return tb.RequestMultiUseToken
}

func (tb *TerminalAuthBuilder) IsSignatureCapture() bool {
	return tb.SignatureCapture
}

func (tb *TerminalAuthBuilder) GetPoNumber() string {
	return tb.PoNumber
}

func (tb *TerminalAuthBuilder) GetTaxAmount() *decimal.Decimal {
	return tb.TaxAmount
}

func (tb *TerminalAuthBuilder) GetTaxExempt() string {
	return tb.TaxExempt
}

func (tb *TerminalAuthBuilder) GetTransactionId() string {
	return tb.TransactionId
}

func (tb *TerminalAuthBuilder) GetTaxExemptId() string {
	return tb.TaxExemptId
}

func (tb *TerminalAuthBuilder) GetTokenRequest() *int {
	return tb.TokenRequest
}

func (tb *TerminalAuthBuilder) GetGiftTransactionType() transactiontype.TransactionType {
	return tb.GiftTransactionType
}

func (tb *TerminalAuthBuilder) WithTokenRequest(tokenRequest *int) *TerminalAuthBuilder {
	tb.TokenRequest = tokenRequest
	return tb
}

func (tb *TerminalAuthBuilder) WithTokenValue(tokenValue string) *TerminalAuthBuilder {
	tb.TokenValue = tokenValue
	return tb
}

func (tb *TerminalAuthBuilder) WithAddress(address *base.Address) *TerminalAuthBuilder {
	tb.Address = address
	return tb
}

func (tb *TerminalAuthBuilder) WithAllowDuplicates(allowDuplicates bool) *TerminalAuthBuilder {
	tb.AllowDuplicates = allowDuplicates
	return tb
}

func (tb *TerminalAuthBuilder) WithAmount(amount *decimal.Decimal) *TerminalAuthBuilder {
	tb.Amount = amount
	return tb
}

func (tb *TerminalAuthBuilder) WithCashBack(value *decimal.Decimal) *TerminalAuthBuilder {
	tb.CashBackAmount = value
	return tb
}

func (tb *TerminalAuthBuilder) WithCardBrandStorage(value storedcredentialinitiator.StoredCredentialInitiator) *TerminalAuthBuilder {
	tb.CardBrandStorage = value
	return tb
}

func (tb *TerminalAuthBuilder) WithCardBrandStorageAndTransactionId(initiatorValue storedcredentialinitiator.StoredCredentialInitiator, cardBrandTransId string) *TerminalAuthBuilder {
	tb.CardBrandStorage = initiatorValue
	tb.CardBrandTransactionId = cardBrandTransId
	return tb
}

func (tb *TerminalAuthBuilder) WithCommercialRequest(value bool) *TerminalAuthBuilder {
	tb.CommercialRequest = value
	return tb
}

func (tb *TerminalAuthBuilder) WithCurrency(value currencytype.CurrencyType) *TerminalAuthBuilder {
	tb.Currency = value
	return tb
}

func (tb *TerminalAuthBuilder) WithCustomerCode(value string) *TerminalAuthBuilder {
	tb.CustomerCode = value
	return tb
}

func (tb *TerminalAuthBuilder) WithGratuity(gratuity *decimal.Decimal) *TerminalAuthBuilder {
	tb.Gratuity = gratuity
	return tb
}

func (tb *TerminalAuthBuilder) WithInvoiceNumber(invoiceNumber string) *TerminalAuthBuilder {
	tb.InvoiceNumber = invoiceNumber
	return tb
}

func (tb *TerminalAuthBuilder) WithPoNumber(value string) *TerminalAuthBuilder {
	tb.PoNumber = value
	return tb
}

func (tb *TerminalAuthBuilder) WithRequestMultiUseToken(requestMultiUseToken bool) *TerminalAuthBuilder {
	tb.RequestMultiUseToken = requestMultiUseToken
	return tb
}

func (tb *TerminalAuthBuilder) WithSignatureCapture(signatureCapture bool) *TerminalAuthBuilder {
	tb.SignatureCapture = signatureCapture
	return tb
}

func (tb *TerminalAuthBuilder) WithTaxAmount(value *decimal.Decimal) *TerminalAuthBuilder {
	tb.TaxAmount = value
	return tb
}

func (tb *TerminalAuthBuilder) WithTaxType(value taxtype.TaxType) *TerminalAuthBuilder {
	return tb.WithTaxTypeAndExemptId(value, "")
}

func (tb *TerminalAuthBuilder) WithTaxTypeAndExemptId(value taxtype.TaxType, taxExemptId string) *TerminalAuthBuilder {
	if value == taxtype.TaxExempt {
		tb.TaxExempt = "1"
	} else {
		tb.TaxExempt = "0"
	}
	tb.TaxExemptId = taxExemptId
	return tb
}

func (tb *TerminalAuthBuilder) WithGiftTransactionType(value transactiontype.TransactionType) *TerminalAuthBuilder {
	tb.GiftTransactionType = value
	return tb
}

func (tb *TerminalAuthBuilder) ExecuteWithName(ctx context.Context, configName string, device ITerminalBuilderDevice) (terminalresponse.ITerminalResponse, error) {
	err := tb.TerminalBuilder.Execute(configName)
	if err != nil {
		return nil, err
	}
	return device.ProcessTransactionWithContext(ctx, tb)
}

func (tb *TerminalAuthBuilder) SetupValidations() {
	tb.Validations = *validations.NewValidations()
	tb.Validations.Of(transactiontype.Sale.LongValue() | transactiontype.Auth.LongValue()).Check("amount").IsNotNull("Amount cannot be empty")
	tb.Validations.Of(transactiontype.Refund.LongValue()).Check("amount").IsNotNull("Amount cannot be empty")
	//tb.Validations.Of(transactiontype.Auth.LongValue()).
	//	With(paymentmethodtype.Credit).
	//	When("transactionId").IsNotNull("").
	//	Check("authCode").IsNotNull("AuthCode cannot be empty")

	tb.Validations.Of(transactiontype.Refund.LongValue()).
		With(paymentmethodtype.Credit).
		When("transactionId").IsNotNull("").
		Check("authCode").IsNotNull("Auth code is required")

	tb.Validations.Of(paymentmethodtype.Gift.LongValue()).Check("currency").IsNotNull("Currency cannot be empty")
	tb.Validations.Of(transactiontype.AddValue.LongValue()).Check("amount").IsNotNull("Amount cannot be empty")

	tb.Validations.Of(paymentmethodtype.EBT.LongValue()).
		With(transactiontype.Balance).
		When("currency").IsNotNull("").
		Check("currency").IsNotEqual(currencytype.Voucher, "Currency cannot be Voucher")

	tb.Validations.Of(transactiontype.BenefitWithdrawal.LongValue()).
		When("currency").IsNotNull("").
		Check("currency").IsEqual(currencytype.CashBenefits, "Currency must be CashBenefits")

	tb.Validations.Of(paymentmethodtype.EBT.LongValue()).
		With(transactiontype.Refund).
		Check("allowDuplicates").IsEqual(false, "Duplicates are not allowed")

	tb.Validations.Of(paymentmethodtype.EBT.LongValue()).
		With(transactiontype.BenefitWithdrawal).
		Check("allowDuplicates").IsEqual(false, "Duplicates are not allowed")
}

func (tb *TerminalAuthBuilder) WithTransactionId(value string) *TerminalAuthBuilder {
	paymentMethod := tb.GetPaymentMethod()
	if !references.IsItATransactionReference(paymentMethod) {
		paymentMethod = references.NewTransactionReference()
	}
	tr := paymentMethod.(*references.TransactionReference)
	tr.SetTransactionId(value)
	tb.SetPaymentMethod(tr)
	tb.TransactionId = value
	return tb
}

func NewTerminalAuthBuilder(transactionType transactiontype.TransactionType, paymentMethod paymentmethodtype.PaymentMethodType) *TerminalAuthBuilder {
	return &TerminalAuthBuilder{
		TerminalBuilder: NewTerminalBuilder(transactionType, paymentMethod),
	}
}
