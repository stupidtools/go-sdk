package transactionsummary

import (
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/securethreedversion"
	"github.com/globalpayments/go-sdk/api/gateways/helpers"
	"github.com/globalpayments/go-sdk/api/utils"
	"github.com/shopspring/decimal"
	"strconv"
)

type TransactionSummary struct {
	AccountDataSource       string
	AdjustmentAmount        *decimal.Decimal
	AdjustmentCurrency      string
	AdjustmentReason        string
	Amount                  *decimal.Decimal
	AttachmentInfo          string
	AcquirerReferenceNumber string
	AuthorizedAmount        *decimal.Decimal
	AuthCode                string
	AvsResponseCode         string
	BaseAmount              *decimal.Decimal
	BatchCloseDate          string
	BatchId                 string
	BatchSequenceNumber     string
	BillingAddress          *base.Address
	BrandReference          string
	CaptureAmount           *decimal.Decimal
	CardHolderFirstName     string
	CardHolderLastName      string
	CardHolderName          string
	CardSwiped              string
	CardType                string
	CavvResponseCode        string
	Channel                 string
	ClerkId                 string
	ClientTransactionId     string
	CompanyName             string
	ConvenienceAmount       *decimal.Decimal
	Currency                string
	CustomerFirstName       string
	CustomerId              string
	CustomerLastName        string
	CvnResponseCode         string
	DebtRepaymentIndicator  bool
	DepositAmount           *decimal.Decimal
	DepositCurrency         string
	DepositDate             string
	DepositReference        string
	DepositStatus           string
	Description             string
	DeviceId                *int
	EciIndicator            string
	EmvChipCondition        string
	EmvIssuerResponse       string
	EntryMode               string
	FraudRuleInfo           string
	FullyCaptured           bool
	CashBackAmount          *decimal.Decimal
	GratuityAmount          *decimal.Decimal
	HasEcomPaymentData      bool
	HasEmvTags              bool
	HostTimeOut             *bool
	InvoiceNumber           string
	IssuerResponseCode      string
	IssuerResponseMessage   string
	IssuerTransactionId     string
	GatewayResponseCode     string
	GatewayResponseMessage  string
	GiftCurrency            string
	MaskedAlias             string
	MaskedCardNumber        string
	MerchantCategory        string
	MerchantDbaName         string
	MerchantHierarchy       string
	MerchantId              string
	MerchantName            string
	MerchantNumber          string
	OneTimePayment          bool
	OrderId                 string
	OriginalTransactionId   string
	PaymentMethodKey        string
	PaymentType             string
	PoNumber                string
	RecurringDataCode       string
	ReferenceNumber         string
	RepeatCount             *int
	ResponseDate            string
	SafTotal                *decimal.Decimal
	SafReferenceNumber      string
	ScheduleId              string
	SchemeReferenceData     string
	ServiceName             string
	SettlementAmount        *decimal.Decimal
	ShippingAmount          *decimal.Decimal
	SiteTrace               string
	Status                  string
	SurchargeAmount         *decimal.Decimal
	TaxAmount               *decimal.Decimal
	TaxType                 string
	ThreeDSecure            *entities.ThreeDSecure
	TerminalId              string
	TokenPanLastFour        string
	TransactionDate         string
	TransactionLocalDate    string
	TransactionDescriptor   string
	TransactionStatus       string
	TransactionId           string
	UniqueDeviceId          string
	Username                string
	XId                     string
	AccountNumberLast4      string
	AccountType             string
	TransactionType         string
	CardEntryMethod         string
	AmountDue               *decimal.Decimal
	Country                 string
	Language                string
	PaymentPurposeCode      string
	VerificationCode        string
	BatchAmount             *decimal.Decimal
}

func (ts *TransactionSummary) GetAccountDataSource() string {
	return ts.AccountDataSource
}

func (ts *TransactionSummary) GetAdjustmentAmount() *decimal.Decimal {
	return ts.AdjustmentAmount
}

func (ts *TransactionSummary) GetAdjustmentCurrency() string {
	return ts.AdjustmentCurrency
}

func (ts *TransactionSummary) GetAdjustmentReason() string {
	return ts.AdjustmentReason
}

func (ts *TransactionSummary) GetAmount() *decimal.Decimal {
	return ts.Amount
}

func (ts *TransactionSummary) GetAttachmentInfo() string {
	return ts.AttachmentInfo
}

func (ts *TransactionSummary) GetAcquirerReferenceNumber() string {
	return ts.AcquirerReferenceNumber
}

func (ts *TransactionSummary) GetAuthorizedAmount() *decimal.Decimal {
	return ts.AuthorizedAmount
}

func (ts *TransactionSummary) GetAuthCode() string {
	return ts.AuthCode
}

func (ts *TransactionSummary) GetAvsResponseCode() string {
	return ts.AvsResponseCode
}

func (ts *TransactionSummary) GetBaseAmount() *decimal.Decimal {
	return ts.BaseAmount
}

func (ts *TransactionSummary) GetBatchCloseDate() string {
	return ts.BatchCloseDate
}

func (ts *TransactionSummary) GetBatchId() string {
	return ts.BatchId
}

func (ts *TransactionSummary) GetBatchSequenceNumber() string {
	return ts.BatchSequenceNumber
}

func (ts *TransactionSummary) GetBillingAddress() *base.Address {
	return ts.BillingAddress
}

func (ts *TransactionSummary) GetBrandReference() string {
	return ts.BrandReference
}

func (ts *TransactionSummary) GetCaptureAmount() *decimal.Decimal {
	return ts.CaptureAmount
}

func (ts *TransactionSummary) GetCardHolderFirstName() string {
	return ts.CardHolderFirstName
}

func (ts *TransactionSummary) GetCardHolderLastName() string {
	return ts.CardHolderLastName
}

func (ts *TransactionSummary) GetCardHolderName() string {
	return ts.CardHolderName
}

func (ts *TransactionSummary) GetCardSwiped() string {
	return ts.CardSwiped
}

func (ts *TransactionSummary) GetCardType() string {
	return ts.CardType
}

func (ts *TransactionSummary) GetCavvResponseCode() string {
	return ts.CavvResponseCode
}

func (ts *TransactionSummary) GetChannel() string {
	return ts.Channel
}

func (ts *TransactionSummary) GetClerkId() string {
	return ts.ClerkId
}

func (ts *TransactionSummary) GetClientTransactionId() string {
	return ts.ClientTransactionId
}

func (ts *TransactionSummary) GetCompanyName() string {
	return ts.CompanyName
}

func (ts *TransactionSummary) GetConvenienceAmount() *decimal.Decimal {
	return ts.ConvenienceAmount
}

func (ts *TransactionSummary) GetCurrency() string {
	return ts.Currency
}

func (ts *TransactionSummary) GetCustomerFirstName() string {
	return ts.CustomerFirstName
}

func (ts *TransactionSummary) GetCustomerId() string {
	return ts.CustomerId
}

func (ts *TransactionSummary) GetCustomerLastName() string {
	return ts.CustomerLastName
}

func (ts *TransactionSummary) GetCvnResponseCode() string {
	return ts.CvnResponseCode
}

func (ts *TransactionSummary) GetDebtRepaymentIndicator() bool {
	return ts.DebtRepaymentIndicator
}

func (ts *TransactionSummary) GetDepositAmount() *decimal.Decimal {
	return ts.DepositAmount
}

func (ts *TransactionSummary) GetDepositCurrency() string {
	return ts.DepositCurrency
}

func (ts *TransactionSummary) GetDepositDate() string {
	return ts.DepositDate
}

func (ts *TransactionSummary) GetDepositReference() string {
	return ts.DepositReference
}

func (ts *TransactionSummary) GetDepositStatus() string {
	return ts.DepositStatus
}

func (ts *TransactionSummary) GetDescription() string {
	return ts.Description
}

func (ts *TransactionSummary) GetDeviceId() *int {
	return ts.DeviceId
}

func (ts *TransactionSummary) GetEciIndicator() string {
	return ts.EciIndicator
}

func (ts *TransactionSummary) GetEmvChipCondition() string {
	return ts.EmvChipCondition
}

func (ts *TransactionSummary) GetEmvIssuerResponse() string {
	return ts.EmvIssuerResponse
}

func (ts *TransactionSummary) GetEntryMode() string {
	return ts.EntryMode
}

func (ts *TransactionSummary) GetFraudRuleInfo() string {
	return ts.FraudRuleInfo
}

func (ts *TransactionSummary) GetFullyCaptured() bool {
	return ts.FullyCaptured
}

func (ts *TransactionSummary) GetCashBackAmount() *decimal.Decimal {
	return ts.CashBackAmount
}

func (ts *TransactionSummary) GetGratuityAmount() *decimal.Decimal {
	return ts.GratuityAmount
}

func (ts *TransactionSummary) GetHasEcomPaymentData() bool {
	return ts.HasEcomPaymentData
}

func (ts *TransactionSummary) GetHasEmvTags() bool {
	return ts.HasEmvTags
}

func (ts *TransactionSummary) GetHostTimeOut() *bool {
	return ts.HostTimeOut
}

func (ts *TransactionSummary) GetInvoiceNumber() string {
	return ts.InvoiceNumber
}

func (ts *TransactionSummary) GetIssuerResponseCode() string {
	return ts.IssuerResponseCode
}

func (ts *TransactionSummary) GetIssuerResponseMessage() string {
	return ts.IssuerResponseMessage
}

func (ts *TransactionSummary) GetIssuerTransactionId() string {
	return ts.IssuerTransactionId
}

func (ts *TransactionSummary) GetGatewayResponseCode() string {
	return ts.GatewayResponseCode
}

func (ts *TransactionSummary) GetGatewayResponseMessage() string {
	return ts.GatewayResponseMessage
}

func (ts *TransactionSummary) GetGiftCurrency() string {
	return ts.GiftCurrency
}

func (ts *TransactionSummary) GetMaskedAlias() string {
	return ts.MaskedAlias
}

func (ts *TransactionSummary) GetMaskedCardNumber() string {
	return ts.MaskedCardNumber
}

func (ts *TransactionSummary) GetMerchantCategory() string {
	return ts.MerchantCategory
}

func (ts *TransactionSummary) GetMerchantDbaName() string {
	return ts.MerchantDbaName
}

func (ts *TransactionSummary) GetMerchantHierarchy() string {
	return ts.MerchantHierarchy
}

func (ts *TransactionSummary) GetMerchantId() string {
	return ts.MerchantId
}

func (ts *TransactionSummary) GetMerchantName() string {
	return ts.MerchantName
}

func (ts *TransactionSummary) GetMerchantNumber() string {
	return ts.MerchantNumber
}

func (ts *TransactionSummary) GetOneTimePayment() bool {
	return ts.OneTimePayment
}

func (ts *TransactionSummary) GetOrderId() string {
	return ts.OrderId
}

func (ts *TransactionSummary) GetOriginalTransactionId() string {
	return ts.OriginalTransactionId
}

func (ts *TransactionSummary) GetPaymentMethodKey() string {
	return ts.PaymentMethodKey
}

func (ts *TransactionSummary) GetPaymentType() string {
	return ts.PaymentType
}

func (ts *TransactionSummary) GetPoNumber() string {
	return ts.PoNumber
}

func (ts *TransactionSummary) GetRecurringDataCode() string {
	return ts.RecurringDataCode
}

func (ts *TransactionSummary) GetReferenceNumber() string {
	return ts.ReferenceNumber
}

func (ts *TransactionSummary) GetRepeatCount() *int {
	return ts.RepeatCount
}

func (ts *TransactionSummary) GetResponseDate() string {
	return ts.ResponseDate
}

func (ts *TransactionSummary) GetSafReferenceNumber() string {
	return ts.SafReferenceNumber
}

func (ts *TransactionSummary) GetScheduleId() string {
	return ts.ScheduleId
}

func (ts *TransactionSummary) GetSchemeReferenceData() string {
	return ts.SchemeReferenceData
}

func (ts *TransactionSummary) GetServiceName() string {
	return ts.ServiceName
}

func (ts *TransactionSummary) GetSettlementAmount() *decimal.Decimal {
	return ts.SettlementAmount
}

func (ts *TransactionSummary) GetShippingAmount() *decimal.Decimal {
	return ts.ShippingAmount
}

func (ts *TransactionSummary) GetSiteTrace() string {
	return ts.SiteTrace
}

func (ts *TransactionSummary) GetStatus() string {
	return ts.Status
}

func (ts *TransactionSummary) GetSurchargeAmount() *decimal.Decimal {
	return ts.SurchargeAmount
}

func (ts *TransactionSummary) GetTaxAmount() *decimal.Decimal {
	return ts.TaxAmount
}

func (ts *TransactionSummary) GetTaxType() string {
	return ts.TaxType
}

func (ts *TransactionSummary) GetTerminalId() string {
	return ts.TerminalId
}

func (ts *TransactionSummary) GetThreeDSecure() *entities.ThreeDSecure {
	return ts.ThreeDSecure
}

func (ts *TransactionSummary) GetTokenPanLastFour() string {
	return ts.TokenPanLastFour
}

func (ts *TransactionSummary) GetTransactionDate() string {
	return ts.TransactionDate
}

func (ts *TransactionSummary) GetTransactionLocalDate() string {
	return ts.TransactionLocalDate
}

func (ts *TransactionSummary) GetTransactionDescriptor() string {
	return ts.TransactionDescriptor
}

func (ts *TransactionSummary) GetTransactionStatus() string {
	return ts.TransactionStatus
}

func (ts *TransactionSummary) GetTransactionId() string {
	return ts.TransactionId
}

func (ts *TransactionSummary) GetUniqueDeviceId() string {
	return ts.UniqueDeviceId
}

func (ts *TransactionSummary) GetUsername() string {
	return ts.Username
}

func (ts *TransactionSummary) GetXId() string {
	return ts.XId
}

func (ts *TransactionSummary) GetAccountNumberLast4() string {
	return ts.AccountNumberLast4
}

func (ts *TransactionSummary) GetAccountType() string {
	return ts.AccountType
}

func (ts *TransactionSummary) GetTransactionType() string {
	return ts.TransactionType
}

func (ts *TransactionSummary) GetCardEntryMethod() string {
	return ts.CardEntryMethod
}

func (ts *TransactionSummary) GetAmountDue() *decimal.Decimal {
	return ts.AmountDue
}

func (ts *TransactionSummary) GetCountry() string {
	return ts.Country
}

func (ts *TransactionSummary) GetLanguage() string {
	return ts.Language
}

func (ts *TransactionSummary) GetPaymentPurposeCode() string {
	return ts.PaymentPurposeCode
}

func (ts *TransactionSummary) GetVerificationCode() string {
	return ts.VerificationCode
}

func (ts *TransactionSummary) GetBatchAmount() *decimal.Decimal {
	return ts.BatchAmount
}

func (ts *TransactionSummary) SetAccountDataSource(val string) {
	ts.AccountDataSource = val
}

func (ts *TransactionSummary) SetAdjustmentAmount(val *decimal.Decimal) {
	ts.AdjustmentAmount = val
}

func (ts *TransactionSummary) SetAdjustmentCurrency(val string) {
	ts.AdjustmentCurrency = val
}

func (ts *TransactionSummary) SetAdjustmentReason(val string) {
	ts.AdjustmentReason = val
}

func (ts *TransactionSummary) SetAmount(val *decimal.Decimal) {
	ts.Amount = val
}

func (ts *TransactionSummary) SetAttachmentInfo(val string) {
	ts.AttachmentInfo = val
}

func (ts *TransactionSummary) SetAcquirerReferenceNumber(val string) {
	ts.AcquirerReferenceNumber = val
}

func (ts *TransactionSummary) SetAuthorizedAmount(val *decimal.Decimal) {
	ts.AuthorizedAmount = val
}

func (ts *TransactionSummary) SetAuthCode(val string) {
	ts.AuthCode = val
}

func (ts *TransactionSummary) SetAvsResponseCode(val string) {
	ts.AvsResponseCode = val
}

func (ts *TransactionSummary) SetBaseAmount(val *decimal.Decimal) {
	ts.BaseAmount = val
}

func (ts *TransactionSummary) SetBatchCloseDate(val string) {
	ts.BatchCloseDate = val
}

func (ts *TransactionSummary) SetBatchId(val string) {
	ts.BatchId = val
}

func (ts *TransactionSummary) SetBatchSequenceNumber(val string) {
	ts.BatchSequenceNumber = val
}

func (ts *TransactionSummary) SetBillingAddress(val *base.Address) {
	ts.BillingAddress = val
}

func (ts *TransactionSummary) SetBrandReference(val string) {
	ts.BrandReference = val
}

func (ts *TransactionSummary) SetCaptureAmount(val *decimal.Decimal) {
	ts.CaptureAmount = val
}

func (ts *TransactionSummary) SetCardHolderFirstName(val string) {
	ts.CardHolderFirstName = val
}

func (ts *TransactionSummary) SetCardHolderLastName(val string) {
	ts.CardHolderLastName = val
}

func (ts *TransactionSummary) SetCardHolderName(val string) {
	ts.CardHolderName = val
}

func (ts *TransactionSummary) SetCardSwiped(val string) {
	ts.CardSwiped = val
}

func (ts *TransactionSummary) SetCardType(val string) {
	ts.CardType = val
}

func (ts *TransactionSummary) SetCavvResponseCode(val string) {
	ts.CavvResponseCode = val
}

func (ts *TransactionSummary) SetChannel(val string) {
	ts.Channel = val
}

func (ts *TransactionSummary) SetClerkId(val string) {
	ts.ClerkId = val
}

func (ts *TransactionSummary) SetClientTransactionId(val string) {
	ts.ClientTransactionId = val
}

func (ts *TransactionSummary) SetCompanyName(val string) {
	ts.CompanyName = val
}

func (ts *TransactionSummary) SetConvenienceAmount(val *decimal.Decimal) {
	ts.ConvenienceAmount = val
}

func (ts *TransactionSummary) SetCurrency(val string) {
	ts.Currency = val
}

func (ts *TransactionSummary) SetCustomerFirstName(val string) {
	ts.CustomerFirstName = val
}

func (ts *TransactionSummary) SetCustomerId(val string) {
	ts.CustomerId = val
}

func (ts *TransactionSummary) SetCustomerLastName(val string) {
	ts.CustomerLastName = val
}

func (ts *TransactionSummary) SetCvnResponseCode(val string) {
	ts.CvnResponseCode = val
}

func (ts *TransactionSummary) SetDebtRepaymentIndicator(val bool) {
	ts.DebtRepaymentIndicator = val
}

func (ts *TransactionSummary) SetDepositAmount(val *decimal.Decimal) {
	ts.DepositAmount = val
}

func (ts *TransactionSummary) SetDepositCurrency(val string) {
	ts.DepositCurrency = val
}

func (ts *TransactionSummary) SetDepositDate(val string) {
	ts.DepositDate = val
}

func (ts *TransactionSummary) SetDepositReference(val string) {
	ts.DepositReference = val
}

func (ts *TransactionSummary) SetDepositStatus(val string) {
	ts.DepositStatus = val
}

func (ts *TransactionSummary) SetDescription(val string) {
	ts.Description = val
}

func (ts *TransactionSummary) SetDeviceId(val *int) {
	ts.DeviceId = val
}

func (ts *TransactionSummary) SetEciIndicator(val string) {
	ts.EciIndicator = val
}

func (ts *TransactionSummary) SetEmvChipCondition(val string) {
	ts.EmvChipCondition = val
}

func (ts *TransactionSummary) SetEmvIssuerResponse(val string) {
	ts.EmvIssuerResponse = val
}

func (ts *TransactionSummary) SetEntryMode(val string) {
	ts.EntryMode = val
}

func (ts *TransactionSummary) SetFraudRuleInfo(val string) {
	ts.FraudRuleInfo = val
}

func (ts *TransactionSummary) SetFullyCaptured(val bool) {
	ts.FullyCaptured = val
}

func (ts *TransactionSummary) SetCashBackAmount(val *decimal.Decimal) {
	ts.CashBackAmount = val
}

func (ts *TransactionSummary) SetGratuityAmount(val *decimal.Decimal) {
	ts.GratuityAmount = val
}

func (ts *TransactionSummary) SetHasEcomPaymentData(val bool) {
	ts.HasEcomPaymentData = val
}

func (ts *TransactionSummary) SetHasEmvTags(val bool) {
	ts.HasEmvTags = val
}

func (ts *TransactionSummary) SetHostTimeOut(val *bool) {
	ts.HostTimeOut = val
}

func (ts *TransactionSummary) SetInvoiceNumber(val string) {
	ts.InvoiceNumber = val
}

func (ts *TransactionSummary) SetIssuerResponseCode(val string) {
	ts.IssuerResponseCode = val
}

func (ts *TransactionSummary) SetIssuerResponseMessage(val string) {
	ts.IssuerResponseMessage = val
}

func (ts *TransactionSummary) SetIssuerTransactionId(val string) {
	ts.IssuerTransactionId = val
}

func (ts *TransactionSummary) SetGatewayResponseCode(val string) {
	ts.GatewayResponseCode = val
}

func (ts *TransactionSummary) SetGatewayResponseMessage(val string) {
	ts.GatewayResponseMessage = val
}

func (ts *TransactionSummary) SetGiftCurrency(val string) {
	ts.GiftCurrency = val
}

func (ts *TransactionSummary) SetMaskedAlias(val string) {
	ts.MaskedAlias = val
}

func (ts *TransactionSummary) SetMaskedCardNumber(val string) {
	ts.MaskedCardNumber = val
}

func (ts *TransactionSummary) SetMerchantCategory(val string) {
	ts.MerchantCategory = val
}

func (ts *TransactionSummary) SetMerchantDbaName(val string) {
	ts.MerchantDbaName = val
}

func (ts *TransactionSummary) SetMerchantHierarchy(val string) {
	ts.MerchantHierarchy = val
}

func (ts *TransactionSummary) SetMerchantId(val string) {
	ts.MerchantId = val
}

func (ts *TransactionSummary) SetMerchantName(val string) {
	ts.MerchantName = val
}

func (ts *TransactionSummary) SetMerchantNumber(val string) {
	ts.MerchantNumber = val
}

func (ts *TransactionSummary) SetOneTimePayment(val bool) {
	ts.OneTimePayment = val
}

func (ts *TransactionSummary) SetPaymentMethodKey(val string) {
	ts.PaymentMethodKey = val
}

func (ts *TransactionSummary) SetPaymentType(val string) {
	ts.PaymentType = val
}

func (ts *TransactionSummary) SetPoNumber(val string) {
	ts.PoNumber = val
}

func (ts *TransactionSummary) SetRecurringDataCode(val string) {
	ts.RecurringDataCode = val
}

func (ts *TransactionSummary) SetReferenceNumber(val string) {
	ts.ReferenceNumber = val
}

func (ts *TransactionSummary) SetRepeatCount(val *int) {
	ts.RepeatCount = val
}

func (ts *TransactionSummary) SetResponseDate(val string) {
	ts.ResponseDate = val
}

func (ts *TransactionSummary) SetSafReferenceNumber(val string) {
	ts.SafReferenceNumber = val
}

func (ts *TransactionSummary) SetScheduleId(val string) {
	ts.ScheduleId = val
}

func (ts *TransactionSummary) SetSchemeReferenceData(val string) {
	ts.SchemeReferenceData = val
}

func (ts *TransactionSummary) SetServiceName(val string) {
	ts.ServiceName = val
}

func (ts *TransactionSummary) SetSettlementAmount(val *decimal.Decimal) {
	ts.SettlementAmount = val
}

func (ts *TransactionSummary) SetShippingAmount(val *decimal.Decimal) {
	ts.ShippingAmount = val
}

func (ts *TransactionSummary) SetSiteTrace(val string) {
	ts.SiteTrace = val
}

func (ts *TransactionSummary) SetStatus(val string) {
	ts.Status = val
}

func (ts *TransactionSummary) SetSurchargeAmount(val *decimal.Decimal) {
	ts.SurchargeAmount = val
}

func (ts *TransactionSummary) SetTaxAmount(val *decimal.Decimal) {
	ts.TaxAmount = val
}

func (ts *TransactionSummary) SetTaxType(val string) {
	ts.TaxType = val
}

func (ts *TransactionSummary) SetTerminalId(val string) {
	ts.TerminalId = val
}
func (ts *TransactionSummary) SetThreeDSecure(s *entities.ThreeDSecure) {
	ts.ThreeDSecure = s
}

func (ts *TransactionSummary) SetTokenPanLastFour(val string) {
	ts.TokenPanLastFour = val
}

func (ts *TransactionSummary) SetTransactionDate(val string) {
	ts.TransactionDate = val
}

func (ts *TransactionSummary) SetTransactionLocalDate(val string) {
	ts.TransactionLocalDate = val
}

func (ts *TransactionSummary) SetTransactionDescriptor(val string) {
	ts.TransactionDescriptor = val
}

func (ts *TransactionSummary) SetTransactionStatus(val string) {
	ts.TransactionStatus = val
}

func (ts *TransactionSummary) SetTransactionId(val string) {
	ts.TransactionId = val
}

func (ts *TransactionSummary) SetUniqueDeviceId(val string) {
	ts.UniqueDeviceId = val
}

func (ts *TransactionSummary) SetUsername(val string) {
	ts.Username = val
}

func (ts *TransactionSummary) SetXId(val string) {
	ts.XId = val
}

func (ts *TransactionSummary) SetAccountNumberLast4(val string) {
	ts.AccountNumberLast4 = val
}

func (ts *TransactionSummary) SetAccountType(val string) {
	ts.AccountType = val
}

func (ts *TransactionSummary) SetTransactionType(val string) {
	ts.TransactionType = val
}

func (ts *TransactionSummary) SetCardEntryMethod(val string) {
	ts.CardEntryMethod = val
}

func (ts *TransactionSummary) SetAmountDue(val *decimal.Decimal) {
	ts.AmountDue = val
}

func (ts *TransactionSummary) SetCountry(val string) {
	ts.Country = val
}

func (ts *TransactionSummary) SetLanguage(val string) {
	ts.Language = val
}

func (ts *TransactionSummary) SetPaymentPurposeCode(val string) {
	ts.PaymentPurposeCode = val
}

func (ts *TransactionSummary) SetVerificationCode(val string) {
	ts.VerificationCode = val
}

func (ts *TransactionSummary) SetBatchAmount(val *decimal.Decimal) {
	ts.BatchAmount = val
}

func HydrateTransactionSummary(root *utils.Element) *TransactionSummary {
	summary := &TransactionSummary{}

	summary.AccountDataSource = root.GetString("AcctDataSrc")
	summary.Amount = root.GetDecimal("Amt")

	summary.AuthCode = root.GetString("AuthCode")

	summary.AuthorizedAmount = root.GetDecimal("AuthAmt")

	// Assuming `getDateTime` returns a string in a suitable format. Otherwise, conversion might be necessary.
	summary.BatchCloseDate = root.GetString("BatchCloseDT")
	summary.BatchSequenceNumber = root.GetString("BatchSeqNbr")

	summary.CaptureAmount = root.GetDecimal("CaptureAmtInfo")

	summary.CardHolderFirstName = root.GetString("CardHolderFirstName")
	summary.CardHolderLastName = root.GetString("CardHolderLastName")

	address := base.NewAddressWithStreet(root.GetString("CardHolderZip"), root.GetString("CardHolderAddr"))

	address.City = root.GetString("CardHolderCity")
	address.Province = root.GetString("CardHolderState")
	summary.BillingAddress = address

	summary.CardSwiped = root.GetString("CardSwiped")
	summary.CardType = root.GetString("CardType")
	summary.CavvResponseCode = root.GetString("CAVVResultCode")
	summary.ClerkId = root.GetString("ClerkID")
	summary.ClientTransactionId = root.GetString("ClientTxnId")
	summary.CompanyName = root.GetString("Company")

	summary.ConvenienceAmount = root.GetDecimal("ConvenienceAmtInfo")

	summary.CustomerFirstName = root.GetString("CustomerFirstname")
	summary.CustomerId = root.GetString("CustomerID")
	summary.CustomerLastName = root.GetString("CustomerLastname")
	summary.DebtRepaymentIndicator = root.GetString("DebtRepaymentIndicator") == "1"
	summary.Description = root.GetString("Description")
	summary.EmvChipCondition = root.GetString("EMVChipCondition")
	summary.EmvIssuerResponse = root.GetString("EMVIssuerResp")

	summary.FraudRuleInfo = root.GetString("FraudInfoRule")
	summary.FullyCaptured = root.GetString("FullyCapturedInd") == "1"
	summary.GatewayResponseCode = helpers.NormalizeResponse(root.GetString("GatewayRspCode"))
	summary.GatewayResponseMessage = root.GetString("GatewayRspMsg")
	summary.GiftCurrency = root.GetString("GiftCurrency")

	summary.GratuityAmount = root.GetDecimal("GratuityAmtInfo")

	summary.HasEmvTags = root.GetString("HasEMVTag") == "1"
	summary.HasEcomPaymentData = root.GetString("HasEComPaymentData") == "1"
	summary.InvoiceNumber = root.GetString("InvoiceNbr")
	summary.IssuerResponseCode = root.GetString("IssuerRspCode")
	summary.IssuerResponseMessage = root.GetString("IssuerRspText")
	summary.IssuerTransactionId = root.GetString("IssTxnId")
	summary.MaskedAlias = root.GetString("GiftMaskedAlias")
	summary.MaskedCardNumber = root.GetString("MaskedCardNbr")
	summary.OneTimePayment = root.GetString("OneTime") == "1"
	summary.OriginalTransactionId = root.GetString("OriginalGatewayTxnId")
	summary.PaymentMethodKey = root.GetString("PaymentMethodKey")
	summary.PaymentType = root.GetString("PaymentType")
	summary.PoNumber = root.GetString("CardHolderPONbr")
	summary.RecurringDataCode = root.GetString("RecurringDataCode")
	summary.ReferenceNumber = root.GetString("RefNbr")

	// Convert to appropriate time format if necessary
	summary.ResponseDate = root.GetString("RspDT")
	summary.ScheduleId = root.GetString("ScheduleID")
	summary.ServiceName = root.GetString("ServiceName")

	summary.SettlementAmount = root.GetDecimal("SettlementAmt")

	summary.ShippingAmount = root.GetDecimal("ShippingAmtInfo")

	summary.Status = root.GetString("Status")
	summary.SurchargeAmount = root.GetDecimal("SurchargeAmtInfo")

	summary.TaxType = root.GetString("TaxType")
	summary.TokenPanLastFour = root.GetString("TokenPANLast4")

	// Assuming `getDateTime` for transaction date conversion if necessary
	summary.TransactionDate = root.GetString("TxnUtcDT")
	summary.TransactionDescriptor = root.GetString("TxnDescriptor")
	summary.TransactionId = root.GetString("GatewayTxnId")
	summary.TransactionStatus = root.GetString("TxnStatus")
	summary.UniqueDeviceId = root.GetString("UniqueDeviceId")
	summary.Username = root.GetString("UserName")

	// 3DSecure
	if root.Has("Secure3D") {
		secure3D := &entities.ThreeDSecure{}
		secure3D.AuthenticationValue = root.GetString("AuthenticationValue")
		secure3D.DirectoryServerTransactionId = root.GetString("DirectoryServerTxnId")
		secure3D.Eci = root.GetString("ECI")

		versionString := root.GetString("Version")
		if versionString == "" {
			// Default to version One
			secure3D.Version = securethreedversion.ONE
		} else {
			versionNumber, err := strconv.Atoi(versionString)
			if err != nil {
				// Handle error, possibly log it
				secure3D.Version = securethreedversion.ONE // Default to ONE in case of error
			} else {
				secure3D.Version = getSecure3dVersion(versionNumber)
			}
		}

		summary.ThreeDSecure = secure3D
	}

	return summary
}

// Assuming this function exists to convert int to Secure3dVersion
func getSecure3dVersion(versionNumber int) securethreedversion.Secure3dVersion {
	// Implementation depends on how Secure3dVersion is defined in your codebase
	// This is a placeholder implementation
	switch versionNumber {
	case 1:
		return securethreedversion.ONE
	case 2:
		return securethreedversion.TWO
	default:
		return securethreedversion.ONE
	}
}
