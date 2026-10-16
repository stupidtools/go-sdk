package builders

import (
	"context"
	"errors"
	abstractions2 "github.com/globalpayments/go-sdk/api/abstractions"
	validations2 "github.com/globalpayments/go-sdk/api/builders/validations"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/billing"
	"github.com/globalpayments/go-sdk/api/entities/enums/accounttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/addresstype"
	"github.com/globalpayments/go-sdk/api/entities/enums/aliasaction"
	"github.com/globalpayments/go-sdk/api/entities/enums/bnplshippingmethod"
	"github.com/globalpayments/go-sdk/api/entities/enums/emvchipcondition"
	"github.com/globalpayments/go-sdk/api/entities/enums/emvfallbackcondition"
	"github.com/globalpayments/go-sdk/api/entities/enums/emvlastchipread"
	"github.com/globalpayments/go-sdk/api/entities/enums/fraudfiltermode"
	"github.com/globalpayments/go-sdk/api/entities/enums/host"
	"github.com/globalpayments/go-sdk/api/entities/enums/hosterror"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodusagemode"
	"github.com/globalpayments/go-sdk/api/entities/enums/phonenumbertype"
	"github.com/globalpayments/go-sdk/api/entities/enums/recurringsequence"
	"github.com/globalpayments/go-sdk/api/entities/enums/recurringtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/reversalreasoncode"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/recurring"
	"github.com/globalpayments/go-sdk/api/entities/remittancereferencetype"
	networkentities "github.com/globalpayments/go-sdk/api/network/entities"
	"github.com/globalpayments/go-sdk/api/network/enums/feetype"
	pabstractions "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type AuthorizationBuilder struct {
	*TransactionBuilder
	AccountType               accounttype.AccountType
	Alias                     string
	AliasAction               aliasaction.AliasAction
	AllowDuplicates           bool
	GenerateReceipt           bool
	IsAvs                     bool
	AllowPartialAuth          bool
	AmountEstimated           bool
	AuthAmount                *decimal.Decimal
	BalanceInquiryType        inquirytype.InquiryType
	BillingAddress            *base.Address
	CardBrandTransactionId    string
	CardHolderLanguage        string
	ClerkId                   string
	ClientTransactionId       string
	ConvenienceAmount         *decimal.Decimal
	Currency                  string
	CustomerId                string
	CustomerData              *recurring.Customer
	CustomData                [][]string
	CustomerIpAddress         string
	DecisionManager           *entities.DecisionManager
	DynamicDescriptor         string
	EcommerceInfo             *entities.EcommerceInfo
	EmvFallbackCondition      *emvfallbackcondition.EmvFallbackCondition
	EmvChipCondition          *emvchipcondition.EmvChipCondition
	EmvLastChipRead           *emvlastchipread.EmvLastChipRead
	FraudFilterMode           fraudfiltermode.FraudFilterMode
	FraudRules                *entities.FraudRuleCollection
	Gratuity                  *decimal.Decimal
	HostedPaymentData         *entities.HostedPaymentData
	IdempotencyKey            string
	Level2Request             bool
	MessageAuthenticationCode string
	MultiCapture              bool
	OfflineAuthCode           string
	OneTimePayment            bool
	OrderId                   string
	PaymentApplicationVersion string
	PaymentMethodUsageMode    paymentmethodusagemode.PaymentMethodUsageMode
	ReversalReasonCode        reversalreasoncode.ReversalReasonCode
	ReplacementCardValue      string
	ReplacementCardPin        string
	ReplacementCardType       string
	HomePhone                 base.PhoneNumber
	WorkPhone                 base.PhoneNumber
	ShippingPhone             base.PhoneNumber
	MobilePhone               base.PhoneNumber
	RemittanceReferenceType   remittancereferencetype.RemittanceReferenceType
	RemittanceReferenceValue  string
	ProductId                 string
	MiscProductData           []base.Product
	RecurringSequence         recurringsequence.RecurringSequence
	RecurringType             recurringtype.RecurringType
	RequestMultiUseToken      bool
	RequestUniqueToken        bool
	ScheduleId                string
	ShippingAddress           *base.Address
	ShippingAmount            *decimal.Decimal
	ShippingDiscount          *decimal.Decimal
	OrderDetails              *base.OrderDetails
	StoredCredential          *entities.StoredCredential
	SupplementaryData         map[string][][]string
	MaskedDataResponse        *bool
	hasEmvFallbackData        bool
	Timestamp                 string
	TransactionInitiator      storedcredentialinitiator.StoredCredentialInitiator
	BNPLShippingMethod        bnplshippingmethod.BNPLShippingMethod
	Bills                     []billing.Bill
	FeeAmount                 *decimal.Decimal
	FeeType                   feetype.FeeType
	FollowOnTimestamp         string
	ShiftNumber               string
	TransportData             string
	GoodsSold                 string
	EWICIssuingEntity         string
	CheckCustomerId           string
	RawMICRData               string
	Country                   string
	PaymentPurposeCode        string
}

func (a *AuthorizationBuilder) GetAlias() string {
	return a.Alias
}

func (a *AuthorizationBuilder) GetAliasAction() aliasaction.AliasAction {
	return a.AliasAction
}

func (a *AuthorizationBuilder) IsAllowDuplicates() bool {
	return a.AllowDuplicates
}

func (a *AuthorizationBuilder) IsAllowPartialAuth() bool {
	return a.AllowPartialAuth
}

func (a *AuthorizationBuilder) GetAmount() *decimal.Decimal {
	return a.amount
}

func (a *AuthorizationBuilder) IsAmountEstimated() bool {
	return a.AmountEstimated
}

func (a *AuthorizationBuilder) GetAuthAmount() *decimal.Decimal {
	return a.AuthAmount
}

func (a *AuthorizationBuilder) GetBalanceInquiryType() inquirytype.InquiryType {
	return a.BalanceInquiryType
}

func (a *AuthorizationBuilder) GetBillingAddress() *base.Address {
	return a.BillingAddress
}

func (a *AuthorizationBuilder) GetBills() []billing.Bill {
	return a.Bills
}

func (a *AuthorizationBuilder) GetCardBrandTransactionId() string {
	return a.CardBrandTransactionId
}

func (a *AuthorizationBuilder) GetCashBackAmount() *decimal.Decimal {
	return a.cashBackAmount
}

func (a *AuthorizationBuilder) GetClerkId() string {
	return a.ClerkId
}

func (a *AuthorizationBuilder) GetClientTransactionId() string {
	return a.ClientTransactionId
}

func (a *AuthorizationBuilder) GetCurrency() string {
	return a.Currency
}

func (a *AuthorizationBuilder) GetCustomer() *recurring.Customer {
	return a.GetCustomerData()
}

func (a *AuthorizationBuilder) GetCustomerId() string {
	return a.CustomerId
}

func (a *AuthorizationBuilder) GetCustomerIpAddress() string {
	return a.CustomerIpAddress
}

func (a *AuthorizationBuilder) GetCustomerData() *recurring.Customer {
	return a.CustomerData
}

func (a *AuthorizationBuilder) GetCustomData() [][]string {
	return a.CustomData
}

func (a *AuthorizationBuilder) GetCvn() string {
	return a.cvn
}

func (a *AuthorizationBuilder) GetDescription() string {
	return a.description
}

func (a *AuthorizationBuilder) GetDecisionManager() *entities.DecisionManager {
	return a.DecisionManager
}

func (a *AuthorizationBuilder) GetDynamicDescriptor() string {
	return a.DynamicDescriptor
}

func (a *AuthorizationBuilder) GetEcommerceInfo() *entities.EcommerceInfo {
	return a.EcommerceInfo
}

func (a *AuthorizationBuilder) GetFraudFilterMode() fraudfiltermode.FraudFilterMode {
	return a.FraudFilterMode
}

func (a *AuthorizationBuilder) GetGratuity() *decimal.Decimal {
	return a.Gratuity
}

func (a *AuthorizationBuilder) GetHostedPaymentData() *entities.HostedPaymentData {
	return a.HostedPaymentData
}

func (a *AuthorizationBuilder) GetInvoiceNumber() string {
	return a.invoiceNumber
}

func (a *AuthorizationBuilder) IsLevel2Request() bool {
	return a.Level2Request
}

func (a *AuthorizationBuilder) GetOfflineAuthCode() string {
	return a.OfflineAuthCode
}

func (a *AuthorizationBuilder) IsOneTimePayment() bool {
	return a.OneTimePayment
}

func (a *AuthorizationBuilder) GetOrderId() string {
	return a.OrderId
}

func (a *AuthorizationBuilder) GetProductId() string {
	return a.ProductId
}

func (a *AuthorizationBuilder) IsRequestMultiUseToken() bool {
	return a.RequestMultiUseToken
}

func (a *AuthorizationBuilder) IsRequestUniqueToken() bool {
	return a.RequestUniqueToken
}

func (a *AuthorizationBuilder) GetRecurringSequence() recurringsequence.RecurringSequence {
	return a.RecurringSequence
}

func (a *AuthorizationBuilder) GetRecurringType() recurringtype.RecurringType {
	return a.RecurringType
}

func (a *AuthorizationBuilder) GetScheduleId() string {
	return a.ScheduleId
}

func (a *AuthorizationBuilder) GetShippingAddress() *base.Address {
	return a.ShippingAddress
}

func (a *AuthorizationBuilder) GetStoredCredential() *entities.StoredCredential {
	return a.StoredCredential
}

func (a *AuthorizationBuilder) GetSurchargeAmount() *decimal.Decimal {
	return a.surchargeAmount
}

func (a *AuthorizationBuilder) GetTimestamp() string {
	return a.Timestamp
}

func (a *AuthorizationBuilder) HasEmvFallbackData() bool {
	return a.EmvFallbackCondition != nil || a.EmvChipCondition != nil || a.PaymentApplicationVersion != ""
}

func (a *AuthorizationBuilder) GetConvenienceAmount() *decimal.Decimal {
	return a.ConvenienceAmount
}

func (a *AuthorizationBuilder) GetShippingAmount() *decimal.Decimal {
	return a.ShippingAmount
}

func (a *AuthorizationBuilder) GetSupplementaryData() map[string][][]string {
	return a.SupplementaryData
}

func (a *AuthorizationBuilder) GetAccountType() accounttype.AccountType {
	return a.AccountType
}

func (a *AuthorizationBuilder) GetEmvChipCondition() *emvchipcondition.EmvChipCondition {
	return a.EmvChipCondition
}

func (a *AuthorizationBuilder) GetEmvLastChipRead() *emvlastchipread.EmvLastChipRead {
	return a.EmvLastChipRead
}

func (a *AuthorizationBuilder) GetMessageAuthenticationCode() string {
	return a.MessageAuthenticationCode
}

func (a *AuthorizationBuilder) IsMultiCapture() bool {
	return a.MultiCapture
}

func (a *AuthorizationBuilder) GetMiscProductData() []base.Product {
	return a.MiscProductData
}

func (a *AuthorizationBuilder) GetTagData() string {
	return a.tagData
}

func (a *AuthorizationBuilder) GetTransactionInitiator() storedcredentialinitiator.StoredCredentialInitiator {
	return a.TransactionInitiator
}

func (a *AuthorizationBuilder) GetReversalReasonCode() reversalreasoncode.ReversalReasonCode {
	return a.ReversalReasonCode
}

func (a *AuthorizationBuilder) WithReversalReasonCode(code reversalreasoncode.ReversalReasonCode) {
	a.ReversalReasonCode = code
}

func (a *AuthorizationBuilder) WithReplacementCard(value string, pin string, cardType string) {
	a.ReplacementCardValue = value
	a.ReplacementCardPin = pin
	a.ReplacementCardType = cardType
}

func (a *AuthorizationBuilder) GetReplacementCardValue() string {
	return a.ReplacementCardValue
}

func (a *AuthorizationBuilder) GetReplacementCardPin() string {
	return a.ReplacementCardPin
}

func (a *AuthorizationBuilder) GetReplacementCardType() string {
	return a.ReplacementCardType
}

func (a *AuthorizationBuilder) GetFeeAmount() *decimal.Decimal {
	return a.FeeAmount
}

func (a *AuthorizationBuilder) GetFeeType() feetype.FeeType {
	return a.FeeType
}

func (a *AuthorizationBuilder) GetFollowOnTimestamp() string {
	return a.FollowOnTimestamp
}

func (a *AuthorizationBuilder) GetShiftNumber() string {
	return a.ShiftNumber
}

func (a *AuthorizationBuilder) GetTransportData() string {
	return a.TransportData
}

func (a *AuthorizationBuilder) GetCardHolderLanguage() string {
	return a.CardHolderLanguage
}

func (a *AuthorizationBuilder) WithAccountType(value accounttype.AccountType) {
	a.AccountType = value
}

func (a *AuthorizationBuilder) WithAddress(value *base.Address) {
	a.WithAddressWithType(value, addresstype.Billing)
}

func (a *AuthorizationBuilder) WithAddressWithType(value *base.Address, addrType addresstype.AddressType) {
	if value != nil {
		value.Type = addrType
		if addrType == addresstype.Billing {
			a.BillingAddress = value
		} else {
			a.ShippingAddress = value
		}
	}
}

func (a *AuthorizationBuilder) WithAlias(action aliasaction.AliasAction, value string) {
	a.Alias = value
	a.AliasAction = action
}

func (a *AuthorizationBuilder) WithAllowDuplicates(value bool) {
	a.AllowDuplicates = value
}

func (a *AuthorizationBuilder) WithGenerateReceipt(value bool) {
	a.GenerateReceipt = value
}

func (a *AuthorizationBuilder) WithAvs(value bool) {
	a.IsAvs = value
}

func (a *AuthorizationBuilder) WithAllowPartialAuth(value bool) {
	a.AllowPartialAuth = value
}

func (a *AuthorizationBuilder) WithAmount(value *decimal.Decimal) {
	a.amount = value
}

func (a *AuthorizationBuilder) WithAmountEstimated(value bool) {
	a.AmountEstimated = value
}

func (a *AuthorizationBuilder) WithAuthAmount(value *decimal.Decimal) {
	a.AuthAmount = value
}

func (a *AuthorizationBuilder) WithBalanceInquiryType(value inquirytype.InquiryType) {
	a.BalanceInquiryType = value
}

func (a *AuthorizationBuilder) WithCardHolderLanguage(value string) {
	a.CardHolderLanguage = value
}

func (a *AuthorizationBuilder) WithBills(bills ...billing.Bill) {
	a.Bills = bills
}

func (ab *AuthorizationBuilder) WithChipCondition(value emvchipcondition.EmvChipCondition) {
	ab.EmvChipCondition = &value
}

func (ab *AuthorizationBuilder) WithFallbackCondition(value emvfallbackcondition.EmvFallbackCondition) {
	ab.EmvFallbackCondition = &value
}

func (ab *AuthorizationBuilder) WithLastChipRead(value emvlastchipread.EmvLastChipRead) {
	ab.EmvLastChipRead = &value
}

func (ab *AuthorizationBuilder) WithClerkId(value string) {
	ab.ClerkId = value
}

func (ab *AuthorizationBuilder) WithClientTransactionId(value string) {
	if ab.transactionType == transactiontype.Reversal {
		if ref, ok := ab.paymentMethod.(*references.TransactionReference); ok {
			ref.SetClientTransactionId(value)
		} else {
			ref := references.NewTransactionReference()
			ref.SetClientTransactionId(value)
			if ab.paymentMethod != nil {
				ref.SetPaymentMethodType(ab.paymentMethod.GetPaymentMethodType())
			}
			ab.paymentMethod = ref
		}
	} else {
		ab.ClientTransactionId = value
	}
}

func (ab *AuthorizationBuilder) WithCommercialRequest(value bool) {
	ab.Level2Request = value
}

func (ab *AuthorizationBuilder) WithConvenienceAmt(value *decimal.Decimal) {
	ab.ConvenienceAmount = value
}

func (ab *AuthorizationBuilder) WithCurrency(value string) {
	ab.Currency = value
}

func (b *AuthorizationBuilder) WithCustomer(value *recurring.Customer) {
	b.WithCustomerData(value)
}

func (b *AuthorizationBuilder) WithCustomerId(value string) {
	b.CustomerId = value
}

func (b *AuthorizationBuilder) WithCustomerIpAddress(value string) {
	b.CustomerIpAddress = value
}

func (b *AuthorizationBuilder) WithCustomerData(value *recurring.Customer) {
	b.CustomerData = value
}

func (b *AuthorizationBuilder) WithCustomData(value ...string) {
	if b.CustomData == nil {
		b.CustomData = make([][]string, 0)
	}
	b.CustomData = append(b.CustomData, value)
}

func (b *AuthorizationBuilder) WithCvn(value string) {
	b.cvn = value
}

func (b *AuthorizationBuilder) WithCardSequenceNumber(value string) {
	b.cardSequenceNumber = value
}

func (b *AuthorizationBuilder) WithDescription(value string) {
	b.description = value
}

func (b *AuthorizationBuilder) WithDecisionManager(value *entities.DecisionManager) {
	b.DecisionManager = value
}

func (b *AuthorizationBuilder) WithDynamicDescriptor(value string) {
	b.DynamicDescriptor = value
}

func (b *AuthorizationBuilder) WithEcommerceInfo(value *entities.EcommerceInfo) {
	b.EcommerceInfo = value
}

func (b *AuthorizationBuilder) WithFraudFilter(fraudFilterMode fraudfiltermode.FraudFilterMode, fraudRules ...*entities.FraudRuleCollection) {
	b.FraudFilterMode = fraudFilterMode
	if len(fraudRules) > 0 {
		b.FraudRules = fraudRules[0]
	}
}

func (b *AuthorizationBuilder) WithGratuity(value *decimal.Decimal) {
	b.Gratuity = value
}

func (b *AuthorizationBuilder) WithHostedPaymentData(value *entities.HostedPaymentData) {
	b.HostedPaymentData = value
}

func (b *AuthorizationBuilder) WithIdempotencyKey(value string) {
	b.IdempotencyKey = value
}

func (b *AuthorizationBuilder) WithInvoiceNumber(value string) {
	b.invoiceNumber = value
}

func (b *AuthorizationBuilder) WithMessageAuthenticationCode(value string) {
	b.MessageAuthenticationCode = value
}

func (b *AuthorizationBuilder) WithMultiCapture(value bool) {
	b.MultiCapture = value
}

func (b *AuthorizationBuilder) WithOfflineAuthCode(value string) {
	b.OfflineAuthCode = value
	b.transactionModifier = transactionmodifier.Offline
}

func (b *AuthorizationBuilder) WithOneTimePayment(value bool) {
	b.OneTimePayment = value
	b.transactionModifier = transactionmodifier.Recurring
}

func (b *AuthorizationBuilder) WithOrderId(value string) {
	b.OrderId = value
}

func (b *AuthorizationBuilder) WithPosSequenceNumber(value string) {
	b.SetPosSequenceNumber(value)
}

func (b *AuthorizationBuilder) WithMiscProductData(values []base.Product) {
	b.MiscProductData = values
}

func (b *AuthorizationBuilder) WithProductId(value string) {
	b.ProductId = value
}

func (b *AuthorizationBuilder) WithPaymentApplicationVersion(value string) {
	b.PaymentApplicationVersion = value
}

func (b *AuthorizationBuilder) WithPaymentMethodUsageMode(value paymentmethodusagemode.PaymentMethodUsageMode) {
	b.PaymentMethodUsageMode = value
}

func (b *AuthorizationBuilder) WithPhoneNumber(phoneCountryCode, number string, phoneType phonenumbertype.PhoneNumberType) {
	phoneNumber := base.PhoneNumber{
		CountryCode: phoneCountryCode,
		Number:      number,
	}

	switch phoneType {
	case phonenumbertype.Home:
		b.HomePhone = phoneNumber
	case phonenumbertype.Work:
		b.WorkPhone = phoneNumber
	case phonenumbertype.Shipping:
		b.ShippingPhone = phoneNumber
	case phonenumbertype.Mobile:
		b.MobilePhone = phoneNumber
	}
}

func (b *AuthorizationBuilder) WithPaymentMethod(value pabstractions.IPaymentMethod) {
	b.paymentMethod = value
}

func (b *AuthorizationBuilder) WithPriorMessageInformation(value *networkentities.PriorMessageInformation) {
	b.priorMessageInformation = value
}

func (b *AuthorizationBuilder) WithRecurringInfo(typeValue recurringtype.RecurringType, sequence recurringsequence.RecurringSequence) {
	b.RecurringType = typeValue
	b.RecurringSequence = sequence
}

func (b *AuthorizationBuilder) WithRequestMultiUseToken(value bool) {
	b.RequestMultiUseToken = value
}

func (b *AuthorizationBuilder) WithRequestUniqueToken(value bool) {
	b.RequestUniqueToken = value
}

func (ab *AuthorizationBuilder) WithTransactionId(value string) {
	if tr, ok := ab.paymentMethod.(*references.TransactionReference); ok {
		tr.SetTransactionId(value)
	} else {
		ref := references.NewTransactionReference()
		ref.SetTransactionId(value)
		if ab.paymentMethod != nil {
			ref.SetPaymentMethodType(ab.paymentMethod.GetPaymentMethodType())
		}

		ab.paymentMethod = ref
	}
}

func (ab *AuthorizationBuilder) WithModifier(value transactionmodifier.TransactionModifier) {
	ab.transactionModifier = value
}

func (ab *AuthorizationBuilder) WithScheduleId(value string) {
	ab.ScheduleId = value
}

func (ab *AuthorizationBuilder) WithShippingAmt(value *decimal.Decimal) {
	ab.ShippingAmount = value
}

func (ab *AuthorizationBuilder) WithShippingDiscount(value *decimal.Decimal) {
	ab.ShippingDiscount = value
}

func (ab *AuthorizationBuilder) WithOrderDetails(value *base.OrderDetails) {
	ab.OrderDetails = value
}

func (ab *AuthorizationBuilder) WithSimulatedHostErrors(h host.Host, errors ...hosterror.HostError) {
	if ab.simulatedHostErrors == nil {
		ab.simulatedHostErrors = make(map[host.Host][]hosterror.HostError)
	}

	if _, exists := ab.simulatedHostErrors[h]; !exists {
		ab.simulatedHostErrors[h] = []hosterror.HostError{}
	}
	for _, error := range errors {
		ab.simulatedHostErrors[h] = append(ab.simulatedHostErrors[h], error)
	}
}

func (ab *AuthorizationBuilder) WithStoredCredential(value *entities.StoredCredential) {
	ab.StoredCredential = value
}

func (ab *AuthorizationBuilder) WithSupplementaryData(t string, values ...string) {
	if ab.SupplementaryData == nil {
		ab.SupplementaryData = make(map[string][][]string)
	}

	if _, exists := ab.SupplementaryData[t]; !exists {
		ab.SupplementaryData[t] = [][]string{}
	}

	ab.SupplementaryData[t] = append(ab.SupplementaryData[t], values)
}

func (ab *AuthorizationBuilder) WithSurchargeAmount(value *decimal.Decimal) {
	ab.surchargeAmount = value
}

func (ab *AuthorizationBuilder) WithTimestamp(value string) {
	ab.WithTimestampAndFollowOn(value, "")
}

func (ab *AuthorizationBuilder) WithTimestampAndFollowOn(value, followOn string) {
	ab.Timestamp = value
	ab.FollowOnTimestamp = followOn
}

func (ab *AuthorizationBuilder) WithTagData(value string) {
	ab.tagData = value
}

func (ab *AuthorizationBuilder) WithBatchNumber(batchNumber *int) {
	val := 0
	ab.WithBatchAndSequenceNumber(batchNumber, &val)
}

func (ab *AuthorizationBuilder) WithBatchAndSequenceNumber(batchNumber, sequenceNumber *int) {
	ab.batchNumber = batchNumber
	ab.sequenceNumber = sequenceNumber
}

func (ab *AuthorizationBuilder) WithCardBrandStorage(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator) {
	ab.WithCardBrandStorageAndTransactionId(transactionInitiator, "")
}

func (ab *AuthorizationBuilder) WithCardBrandStorageAndTransactionId(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator, value string) {
	ab.TransactionInitiator = transactionInitiator
	ab.CardBrandTransactionId = value
}

func (ab *AuthorizationBuilder) WithCompanyId(companyId string) {
	ab.companyId = companyId
}

func (ab *AuthorizationBuilder) WithFee(feeType feetype.FeeType, feeAmount *decimal.Decimal) {
	ab.FeeType = feeType
	ab.FeeAmount = feeAmount
}

func (ab *AuthorizationBuilder) WithShiftNumber(value string) {
	ab.ShiftNumber = value
}

func (ab *AuthorizationBuilder) WithSystemTraceAuditNumber(original *int) {
	ab.WithSystemTraceAuditNumberAndFollowOn(original, nil)
}

func (ab *AuthorizationBuilder) WithSystemTraceAuditNumberAndFollowOn(original *int, followOn *int) {
	ab.systemTraceAuditNumber = original
	ab.followOnStan = followOn
}

func (ab *AuthorizationBuilder) WithTerminalError(value bool) {
	ab.terminalError = value
}

func (ab *AuthorizationBuilder) WithTransportData(value string) {
	ab.TransportData = value
}

func (ab *AuthorizationBuilder) WithTransactionMatchingData(value *networkentities.TransactionMatchingData) {
	ab.transactionMatchingData = value
}

func (ab *AuthorizationBuilder) WithUniqueDeviceId(value string) {
	ab.uniqueDeviceId = value
}

func (ab *AuthorizationBuilder) WithTaxAmount(taxAmount *decimal.Decimal) {
	ab.taxAmount = taxAmount
}

func (ab *AuthorizationBuilder) WithTipAmount(tipAmount *decimal.Decimal) {
	ab.tipAmount = tipAmount
}

func (ab *AuthorizationBuilder) WithRemittanceReference(remittanceReferenceType remittancereferencetype.RemittanceReferenceType, remittanceReferenceValue string) {
	ab.RemittanceReferenceType = remittanceReferenceType
	ab.RemittanceReferenceValue = remittanceReferenceValue
}

func (ab *AuthorizationBuilder) WithBNPLShippingMethod(value bnplshippingmethod.BNPLShippingMethod) error {
	if ab.paymentMethod.GetPaymentMethodType() != paymentmethodtype.BNPL {
		return errors.New("The selected payment method doesn't support this property!")
	}

	ab.BNPLShippingMethod = value
	return nil
}

func (ab *AuthorizationBuilder) WithTransactiontype(t transactiontype.TransactionType) {
	ab.transactionType = t
}

func (ab *AuthorizationBuilder) WithPaymentPurposeCode(paymentPurposeCode string) {
	ab.PaymentPurposeCode = paymentPurposeCode
}

func (b *AuthorizationBuilder) WithEmvMaxPinEntry(emvMaxPinEntry string) {
	b.emvMaxPinEntry = emvMaxPinEntry
}

func (b *AuthorizationBuilder) WithServiceCode(serviceCode string) {
	b.serviceCode = serviceCode
}

func (b *AuthorizationBuilder) WithGoodsSold(goodsSold string) {
	b.GoodsSold = goodsSold
}

func (b *AuthorizationBuilder) WithCheckCustomerId(checkCustomerId string) {
	b.CheckCustomerId = checkCustomerId
}

func (b *AuthorizationBuilder) WithRawMICRData(rawMICRData string) {
	b.RawMICRData = rawMICRData
}

func (b *AuthorizationBuilder) WithEWICIssuingEntity(eWICIssuingEntity string) {
	b.EWICIssuingEntity = eWICIssuingEntity
}

func (b *AuthorizationBuilder) WithCountry(country string) {
	b.Country = country
}

func (b *AuthorizationBuilder) WithMaskedDataResponse(value *bool) {
	b.MaskedDataResponse = value
}

func (a *AuthorizationBuilder) Serialize(ctx context.Context, gateway abstractions2.IPaymentGateway) (string, error) {
	return a.SerializeWithConfig(ctx, "default", gateway)
}

func (a *AuthorizationBuilder) SerializeWithConfig(ctx context.Context, configName string, gateway abstractions2.IPaymentGateway) (string, error) {
	a.transactionModifier = transactionmodifier.HostedRequest
	_, err := gateway.ProcessAuthorization(ctx, a)

	if err != nil {
		return "", err
	}
	if gateway.SupportsHostedPayments() {
		return gateway.SerializeRequest(ctx, a)
	}
	return "", errors.New("Your current gateway does not support hosted payments.")
}

func (a *AuthorizationBuilder) Execute(ctx context.Context, gateway abstractions2.IPaymentGateway) (abstractions2.ITransaction, error) {
	return gateway.ProcessAuthorization(ctx, a)
}

func NewAuthorizationBuilder(transactionType transactiontype.TransactionType) *AuthorizationBuilder {
	return NewAuthorizationBuilderWithPaymentMethod(transactionType, nil)
}

func NewAuthorizationBuilderWithPaymentMethod(transactionType transactiontype.TransactionType, paymentMethod pabstractions.IPaymentMethod) *AuthorizationBuilder {
	tb := NewTransactionBuilder(transactionType, paymentMethod)
	ab := &AuthorizationBuilder{
		TransactionBuilder: tb,
	}
	ab.WithPaymentMethod(paymentMethod)
	return ab
}

func (ab *AuthorizationBuilder) SetupValidations() {
	validations := validations2.NewValidations()

	// Using LongValue() method on enum constants to get their int64 representation.
	authSaleRefundAddValue := []int64{transactiontype.Auth.LongValue(), transactiontype.Sale.LongValue(), transactiontype.Refund.LongValue(), transactiontype.AddValue.LongValue()}
	authSale := []int64{transactiontype.Auth.LongValue(), transactiontype.Sale.LongValue()}

	validations.OfGeneral(enumToString(authSaleRefundAddValue), transactionmodifier.None.LongValue()).
		Check("AuthAmount").IsNotNull("").
		Check("Currency").IsNotNull("").
		Check("PaymentMethod").IsNotNull("")

	validations.OfGeneral(enumToString(authSale), transactionmodifier.HostedRequest.LongValue()).
		Check("AuthAmount").IsNotNull("").
		Check("Currency").IsNotNull("")

	validations.Of(transactiontype.Verify.LongValue()).
		Check("Currency").IsNotNull("")

	validations.Of(transactionmodifier.HostedRequest.LongValue()).
		Check("Currency").IsNotNull("")

	validations.OfGeneral(enumToString(authSale), transactionmodifier.Offline.LongValue()).
		Check("AuthAmount").IsNotNull("").
		Check("Currency").IsNotNull("").
		Check("OfflineAuthCode").IsNotNull("")

	validations.Of(transactiontype.BenefitWithdrawal.LongValue()).
		Check("AuthAmount").IsNotNull("").
		Check("Currency").IsNotNull("").
		Check("PaymentMethod").IsNotNull("")

	validations.Of(transactionmodifier.CashBack.LongValue()).
		Check("AuthAmount").IsNotNull("").
		Check("Currency").IsNotNull("").
		Check("PaymentMethod").IsNotNull("")

	validations.Of(transactiontype.Balance.LongValue()).
		Check("PaymentMethod").IsNotNull("")

	validations.Of(transactiontype.Alias.LongValue()).
		Check("AliasAction").IsNotNull("").
		Check("Alias").IsNotNull("")

	validations.Of(transactiontype.Replace.LongValue()).
		Check("ReplacementCard").IsNotNull("")

	validations.Of(paymentmethodtype.ACH.LongValue()).
		Check("BillingAddress").IsNotNull("")

	validations.Of(paymentmethodtype.Debit.LongValue()).
		When("ReversalReasonCode").IsNotNull("").
		Check("TransactionType").IsEqual(transactiontype.Reversal.LongValue(), "")

	debitCredit := []int64{paymentmethodtype.Debit.LongValue(), paymentmethodtype.Credit.LongValue()}
	validations.OfGeneral(enumToString(debitCredit), 0).
		When("EmvChipCondition").IsNotNull("").
		Check("TagData").IsNull("")

	validations.OfGeneral(enumToString(debitCredit), 0).
		When("TagData").IsNotNull("").
		Check("EmvChipCondition").IsNull("")

	validations.OfGeneral(enumToString(authSale), transactionmodifier.EncryptedMobile.LongValue()).
		Check("PaymentMethod").IsNotNull("")

	validations.Of(transactiontype.DccRateLookup.LongValue()).
		Check("DccRateData").IsNotNull("")

	validations.Of(transactiontype.Tokenize.LongValue()).
		Check("PaymentMethod").IsNotNull("")

}

// enumToString converts a slice of int64 enums to a string representation.
func enumToString(enums []int64) string {
	var strEnums []string
	for _, e := range enums {
		strEnums = append(strEnums, strconv.FormatInt(e, 10))
	}
	return strings.Join(strEnums, ",")
}
