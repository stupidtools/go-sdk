package transactions

import (
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders"
	entities2 "github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodusagemode"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/globalpayments/go-sdk/api/network/entities"
	"github.com/globalpayments/go-sdk/api/network/enums/cardissuerentrytag"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	AdditionalResponseCode      string
	AuthorizedAmount            *decimal.Decimal
	AutoReversed                bool
	AutoSettleFlag              string
	AvailableBalance            *decimal.Decimal
	AvsResponseCode             string
	AvsResponseMessage          string
	AvsAddressResponse          string
	BalanceAmount               *decimal.Decimal
	BatchSummary                *entities2.BatchSummary
	DebitMac                    *entities2.DebitMac
	CardBrandTransactionId      string
	CardType                    string
	CardLast4                   string
	FingerPrint                 string
	FingerPrintIndicator        string
	CardNumber                  string
	CardExpMonth                int
	CardExpYear                 int
	CavvResponseCode            string
	CommercialIndicator         string
	ConvenienceFee              *decimal.Decimal
	CvnResponseCode             string
	CvnResponseMessage          string
	EmvIssuerResponse           string
	GiftCard                    *paymentmethods.GiftCard
	HostResponseDate            string
	MultiCapture                bool
	MultiCapturePaymentCount    int
	MultiCaptureSequence        int
	NonApprovedDataCollectToken []string
	FormatErrorDataCollectToken []string
	IssuerData                  map[cardissuerentrytag.CardIssuerEntryTag]string
	MessageInformation          *entities.PriorMessageInformation
	PointsBalanceAmount         *decimal.Decimal
	RecurringDataCode           string
	ReferenceNumber             string
	ResponseCode                string
	ResponseDate                string
	ResponseMessage             string
	ResponseValues              map[string]string
	SchemeId                    string
	Timestamp                   string
	TransactionDescriptor       string
	TransactionToken            string
	Token                       string
	TokenUsageMode              paymentmethodusagemode.PaymentMethodUsageMode
	TransactionReference        *references.TransactionReference
	TransactionDate             string
	TransactionTime             string
	TransactionCode             string
	TransactionSummary          transactionsummary.TransactionSummary
	CustomerReceipt             string
	MerchantReceipt             string
	CheckNumber                 string
	RoutingNumber               string
	BankNumber                  string
	BranchTransitNumber         string
	BsbNumber                   string
	FinancialInstitutionNumber  string
	CustomerFeeAmount           *decimal.Decimal
	ReceiptText                 string
	AdditionalDuplicateData     *entities2.AdditionalDuplicateData
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) Increment(amt *decimal.Decimal) abstractions.IManagementBuilder {
	builder := builders.NewManagementBuilder(transactiontype.Auth)
	builder.WithPaymentMethod(t.TransactionReference)
	builder.WithAmount(amt)
	builder.SetTransactionType(transactiontype.Increment)
	return builder
}

func (t *Transaction) Reverse() abstractions.IManagementBuilder {
	var amt *decimal.Decimal
	amt = nil
	if t.TransactionReference != nil {
		amt = t.TransactionReference.OriginalAmount
	}
	return t.ReverseWithAmount(amt)
}

func (t *Transaction) ReverseWithAmount(amt *decimal.Decimal) abstractions.IManagementBuilder {
	builder := builders.NewManagementBuilder(transactiontype.Reversal)
	builder.WithPaymentMethod(t.TransactionReference)
	builder.WithAmount(amt)
	return builder
}

func (t *Transaction) Refund() abstractions.IManagementBuilder {
	var amt *decimal.Decimal
	amt = nil
	if t.TransactionReference != nil {
		amt = t.TransactionReference.OriginalAmount
	}
	return t.RefundWithAmount(amt)
}

func (t *Transaction) RefundWithAmount(amt *decimal.Decimal) abstractions.IManagementBuilder {
	builder := builders.NewManagementBuilder(transactiontype.Refund)
	builder.WithPaymentMethod(t.TransactionReference)
	builder.WithAmount(amt)
	return builder
}

func (t *Transaction) Capture() abstractions.IManagementBuilder {
	var amt *decimal.Decimal
	return t.CaptureWithAmount(amt)
}

func (t *Transaction) CaptureWithAmount(amt *decimal.Decimal) abstractions.IManagementBuilder {
	builder := builders.NewManagementBuilder(transactiontype.Capture)
	builder.WithPaymentMethod(t.TransactionReference)
	builder.WithAmount(amt)

	if t.MultiCapture {
		builder.WithMultiCaptureAndPaymentCount(&t.MultiCaptureSequence, &t.MultiCapturePaymentCount)
	}

	return builder
}

func (t *Transaction) Edit() abstractions.IManagementBuilder {
	builder := builders.NewManagementBuilder(transactiontype.Edit)
	builder.WithPaymentMethod(t.TransactionReference)

	if t.CommercialIndicator != "" {
		builder.SetTransactionModifier(transactionmodifier.LevelII)
	}

	if t.CardType != "" {
		builder.WithCardType(t.CardType)
	}

	return builder
}

func (t *Transaction) Void(amt *decimal.Decimal, force bool) abstractions.IManagementBuilder {
	builder := builders.NewManagementBuilder(transactiontype.Void)
	builder.WithAmount(amt)
	builder.WithPaymentMethod(t.TransactionReference)
	builder.WithForceToHost(force)
	return builder
}

func (t *Transaction) GetGiftCard() *paymentmethods.GiftCard { return t.GiftCard }

func (t *Transaction) SetGiftCard(g *paymentmethods.GiftCard) {
	t.GiftCard = g
}

func (t *Transaction) GetAutoReversed() bool { return t.AutoReversed }

func (t *Transaction) SetAutoReversed(r bool) {
	t.AutoReversed = r
}

func (t *Transaction) GetAdditionalResponseCode() string {
	return t.AdditionalResponseCode
}

func (t *Transaction) SetAdditionalResponseCode(value string) {
	t.AdditionalResponseCode = value
}

func (t *Transaction) GetAuthorizedAmount() *decimal.Decimal {
	return t.AuthorizedAmount
}

func (t *Transaction) SetAuthorizedAmount(value *decimal.Decimal) {
	t.AuthorizedAmount = value
}

func (t *Transaction) GetAutoSettleFlag() string {
	return t.AutoSettleFlag
}

func (t *Transaction) SetAutoSettleFlag(value string) {
	t.AutoSettleFlag = value
}

func (t *Transaction) GetAvailableBalance() *decimal.Decimal {
	return t.AvailableBalance
}

func (t *Transaction) SetAvailableBalance(value *decimal.Decimal) {
	t.AvailableBalance = value
}

func (t *Transaction) GetAvsResponseCode() string {
	return t.AvsResponseCode
}

func (t *Transaction) SetAvsResponseCode(value string) {
	t.AvsResponseCode = value
}

func (t *Transaction) GetAvsResponseMessage() string {
	return t.AvsResponseMessage
}

func (t *Transaction) SetAvsResponseMessage(value string) {
	t.AvsResponseMessage = value
}

func (t *Transaction) GetAvsAddressResponse() string {
	return t.AvsAddressResponse
}

func (t *Transaction) SetAvsAddressResponse(value string) {
	t.AvsAddressResponse = value
}

func (t *Transaction) GetBalanceAmount() *decimal.Decimal {
	return t.BalanceAmount
}

func (t *Transaction) SetBalanceAmount(value *decimal.Decimal) {
	t.BalanceAmount = value
}

func (t *Transaction) GetBatchSummary() *entities2.BatchSummary {
	return t.BatchSummary
}

func (t *Transaction) SetBatchSummary(value *entities2.BatchSummary) {
	t.BatchSummary = value
}

func (t *Transaction) GetDebitMac() *entities2.DebitMac {
	return t.DebitMac
}

func (t *Transaction) SetDebitMac(value *entities2.DebitMac) {
	t.DebitMac = value
}

func (t *Transaction) GetCardBrandTransactionId() string {
	return t.CardBrandTransactionId
}

func (t *Transaction) SetCardBrandTransactionId(value string) {
	t.CardBrandTransactionId = value
}

func (t *Transaction) GetCardType() string {
	return t.CardType
}

func (t *Transaction) SetCardType(value string) {
	t.CardType = value
}

func (t *Transaction) GetCardLast4() string {
	return t.CardLast4
}

func (t *Transaction) SetCardLast4(value string) {
	t.CardLast4 = value
}

func (t *Transaction) GetFingerPrint() string {
	return t.FingerPrint
}

func (t *Transaction) SetFingerPrint(value string) {
	t.FingerPrint = value
}

func (t *Transaction) GetFingerPrintIndicator() string {
	return t.FingerPrintIndicator
}

func (t *Transaction) SetFingerPrintIndicator(value string) {
	t.FingerPrintIndicator = value
}

func (t *Transaction) GetCardNumber() string {
	return t.CardNumber
}

func (t *Transaction) SetCardNumber(value string) {
	t.CardNumber = value
}

func (t *Transaction) GetCardExpMonth() int {
	return t.CardExpMonth
}

func (t *Transaction) SetCardExpMonth(value int) {
	t.CardExpMonth = value
}

func (t *Transaction) GetCardExpYear() int {
	return t.CardExpYear
}

func (t *Transaction) SetCardExpYear(value int) {
	t.CardExpYear = value
}

func (t *Transaction) GetCavvResponseCode() string {
	return t.CavvResponseCode
}

func (t *Transaction) SetCavvResponseCode(value string) {
	t.CavvResponseCode = value
}

func (t *Transaction) GetCommercialIndicator() string {
	return t.CommercialIndicator
}

func (t *Transaction) SetCommercialIndicator(value string) {
	t.CommercialIndicator = value
}

func (t *Transaction) GetConvenienceFee() *decimal.Decimal {
	return t.ConvenienceFee
}

func (t *Transaction) SetConvenienceFee(value *decimal.Decimal) {
	t.ConvenienceFee = value
}

func (t *Transaction) GetCvnResponseCode() string {
	return t.CvnResponseCode
}

func (t *Transaction) SetCvnResponseCode(value string) {
	t.CvnResponseCode = value
}

func (t *Transaction) GetCvnResponseMessage() string {
	return t.CvnResponseMessage
}

func (t *Transaction) SetCvnResponseMessage(value string) {
	t.CvnResponseMessage = value
}

func (t *Transaction) GetEmvIssuerResponse() string {
	return t.EmvIssuerResponse
}

func (t *Transaction) SetEmvIssuerResponse(value string) {
	t.EmvIssuerResponse = value
}

func (t *Transaction) GetHostResponseDate() string {
	return t.HostResponseDate
}

func (t *Transaction) SetHostResponseDate(value string) {
	t.HostResponseDate = value
}

func (t *Transaction) GetMultiCapture() bool {
	return t.MultiCapture
}

func (t *Transaction) SetMultiCapture(value bool) {
	t.MultiCapture = value
}

func (t *Transaction) GetMultiCapturePaymentCount() int {
	return t.MultiCapturePaymentCount
}

func (t *Transaction) SetMultiCapturePaymentCount(value int) {
	t.MultiCapturePaymentCount = value
}

func (t *Transaction) GetMultiCaptureSequence() int {
	return t.MultiCaptureSequence
}

func (t *Transaction) SetMultiCaptureSequence(value int) {
	t.MultiCaptureSequence = value
}

func (t *Transaction) GetNonApprovedDataCollectToken() []string {
	return t.NonApprovedDataCollectToken
}

func (t *Transaction) SetNonApprovedDataCollectToken(value []string) {
	t.NonApprovedDataCollectToken = value
}

func (t *Transaction) GetFormatErrorDataCollectToken() []string {
	return t.FormatErrorDataCollectToken
}

func (t *Transaction) SetFormatErrorDataCollectToken(value []string) {
	t.FormatErrorDataCollectToken = value
}

func (t *Transaction) GetIssuerData() map[cardissuerentrytag.CardIssuerEntryTag]string {
	return t.IssuerData
}

func (t *Transaction) SetIssuerData(value map[cardissuerentrytag.CardIssuerEntryTag]string) {
	t.IssuerData = value
}

func (t *Transaction) GetMessageInformation() *entities.PriorMessageInformation {
	return t.MessageInformation
}

func (t *Transaction) SetMessageInformation(value *entities.PriorMessageInformation) {
	t.MessageInformation = value
}

func (t *Transaction) GetPointsBalanceAmount() *decimal.Decimal {
	return t.PointsBalanceAmount
}

func (t *Transaction) SetPointsBalanceAmount(value *decimal.Decimal) {
	t.PointsBalanceAmount = value
}

func (t *Transaction) GetRecurringDataCode() string {
	return t.RecurringDataCode
}

func (t *Transaction) SetRecurringDataCode(value string) {
	t.RecurringDataCode = value
}

func (t *Transaction) GetReferenceNumber() string {
	return t.ReferenceNumber
}

func (t *Transaction) SetReferenceNumber(value string) {
	t.ReferenceNumber = value
}

func (t *Transaction) GetResponseCode() string {
	return t.ResponseCode
}

func (t *Transaction) SetResponseCode(value string) {
	t.ResponseCode = value
}

func (t *Transaction) GetResponseDate() string {
	return t.ResponseDate
}

func (t *Transaction) SetResponseDate(value string) {
	t.ResponseDate = value
}

func (t *Transaction) GetResponseMessage() string {
	return t.ResponseMessage
}

func (t *Transaction) SetResponseMessage(value string) {
	t.ResponseMessage = value
}

func (t *Transaction) GetResponseValues() map[string]string {
	return t.ResponseValues
}

func (t *Transaction) SetResponseValues(value map[string]string) {
	t.ResponseValues = value
}

func (t *Transaction) GetSchemeId() string {
	return t.SchemeId
}

func (t *Transaction) SetSchemeId(value string) {
	t.SchemeId = value
}

func (t *Transaction) GetTimestamp() string {
	return t.Timestamp
}

func (t *Transaction) SetTimestamp(value string) {
	t.Timestamp = value
}

func (t *Transaction) GetTransactionDescriptor() string {
	return t.TransactionDescriptor
}

func (t *Transaction) SetTransactionDescriptor(value string) {
	t.TransactionDescriptor = value
}

func (t *Transaction) GetTransactionToken() string {
	return t.TransactionToken
}

func (t *Transaction) SetTransactionToken(value string) {
	t.TransactionToken = value
}

func (t *Transaction) GetToken() string {
	return t.Token
}

func (t *Transaction) SetToken(value string) {
	t.Token = value
}

func (t *Transaction) GetTokenUsageMode() paymentmethodusagemode.PaymentMethodUsageMode {
	return t.TokenUsageMode
}

func (t *Transaction) SetTokenUsageMode(value paymentmethodusagemode.PaymentMethodUsageMode) {
	t.TokenUsageMode = value
}

func (t *Transaction) GetTransactionReference() *references.TransactionReference {
	return t.TransactionReference
}

func (t *Transaction) SetTransactionReference(value *references.TransactionReference) {
	t.TransactionReference = value
}

func (t *Transaction) GetTransactionId() string {
	ref := t.GetTransactionReference()
	if ref != nil {
		return ref.GetTransactionId()
	}
	return ""
}

func (t *Transaction) GetTransactionDate() string {
	return t.TransactionDate
}

func (t *Transaction) SetTransactionDate(transactionDate string) {
	t.TransactionDate = transactionDate
}

func (t *Transaction) GetTransactionTime() string {
	return t.TransactionTime
}

func (t *Transaction) SetTransactionTime(transactionTime string) {
	t.TransactionTime = transactionTime
}

func (t *Transaction) GetTransactionCode() string {
	return t.TransactionCode
}

func (t *Transaction) SetTransactionCode(transactionCode string) {
	t.TransactionCode = transactionCode
}

func (t *Transaction) GetTransactionSummary() transactionsummary.TransactionSummary {
	return t.TransactionSummary
}

func (t *Transaction) SetTransactionSummary(transactionSummary transactionsummary.TransactionSummary) {
	t.TransactionSummary = transactionSummary
}

func (t *Transaction) GetCustomerReceipt() string {
	return t.CustomerReceipt
}

func (t *Transaction) SetCustomerReceipt(customerReceipt string) {
	t.CustomerReceipt = customerReceipt
}

func (t *Transaction) GetMerchantReceipt() string {
	return t.MerchantReceipt
}

func (t *Transaction) SetMerchantReceipt(merchantReceipt string) {
	t.MerchantReceipt = merchantReceipt
}

func (t *Transaction) GetCheckNumber() string {
	return t.CheckNumber
}

func (t *Transaction) SetCheckNumber(checkNumber string) {
	t.CheckNumber = checkNumber
}

func (t *Transaction) GetRoutingNumber() string {
	return t.RoutingNumber
}

func (t *Transaction) SetRoutingNumber(routingNumber string) {
	t.RoutingNumber = routingNumber
}

func (t *Transaction) GetBankNumber() string {
	return t.BankNumber
}

func (t *Transaction) SetBankNumber(bankNumber string) {
	t.BankNumber = bankNumber
}

func (t *Transaction) GetBranchTransitNumber() string {
	return t.BranchTransitNumber
}

func (t *Transaction) SetBranchTransitNumber(branchTransitNumber string) {
	t.BranchTransitNumber = branchTransitNumber
}

func (t *Transaction) GetBsbNumber() string {
	return t.BsbNumber
}

func (t *Transaction) SetBsbNumber(bsbNumber string) {
	t.BsbNumber = bsbNumber
}

func (t *Transaction) GetFinancialInstitutionNumber() string {
	return t.FinancialInstitutionNumber
}

func (t *Transaction) SetFinancialInstitutionNumber(financialInstitutionNumber string) {
	t.FinancialInstitutionNumber = financialInstitutionNumber
}

func (t *Transaction) GetCustomerFeeAmount() *decimal.Decimal {
	return t.CustomerFeeAmount
}

func (t *Transaction) SetCustomerFeeAmount(customerFeeAmount *decimal.Decimal) {
	t.CustomerFeeAmount = customerFeeAmount
}

func (t *Transaction) GetReceiptText() string {
	return t.ReceiptText
}

func (t *Transaction) SetReceiptText(receiptText string) {
	t.ReceiptText = receiptText
}

func (t *Transaction) GetAdditionalDuplicateData() *entities2.AdditionalDuplicateData {
	return t.AdditionalDuplicateData
}

func (t *Transaction) SetAdditionalDuplicateData(additionalDuplicateData *entities2.AdditionalDuplicateData) {
	t.AdditionalDuplicateData = additionalDuplicateData
}
