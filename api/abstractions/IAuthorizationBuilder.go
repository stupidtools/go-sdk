package abstractions

import (
	"context"
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
	"github.com/shopspring/decimal"
)

type IAuthorizationBuilder interface {
	ITransactionBuilder
	GetAlias() string
	GetAliasAction() aliasaction.AliasAction
	IsAllowDuplicates() bool
	IsAllowPartialAuth() bool
	GetAmount() *decimal.Decimal
	IsAmountEstimated() bool
	GetAuthAmount() *decimal.Decimal
	GetBalanceInquiryType() inquirytype.InquiryType
	GetBillingAddress() *base.Address
	GetBills() []billing.Bill
	GetCardBrandTransactionId() string
	GetCashBackAmount() *decimal.Decimal
	GetClerkId() string
	GetClientTransactionId() string
	GetCurrency() string
	GetCustomer() *recurring.Customer
	GetCustomerId() string
	GetCustomerIpAddress() string
	GetCustomerData() *recurring.Customer
	GetCustomData() [][]string
	GetCvn() string
	GetDescription() string
	GetDecisionManager() *entities.DecisionManager
	GetDynamicDescriptor() string
	GetEcommerceInfo() *entities.EcommerceInfo
	GetFraudFilterMode() fraudfiltermode.FraudFilterMode
	GetGratuity() *decimal.Decimal
	GetHostedPaymentData() *entities.HostedPaymentData
	GetInvoiceNumber() string
	IsLevel2Request() bool
	GetOfflineAuthCode() string
	IsOneTimePayment() bool
	GetOrderId() string
	GetProductId() string
	IsRequestMultiUseToken() bool
	IsRequestUniqueToken() bool
	GetRecurringSequence() recurringsequence.RecurringSequence
	GetRecurringType() recurringtype.RecurringType
	GetScheduleId() string
	GetShippingAddress() *base.Address
	GetStoredCredential() *entities.StoredCredential
	GetSurchargeAmount() *decimal.Decimal
	GetTimestamp() string
	HasEmvFallbackData() bool
	GetConvenienceAmount() *decimal.Decimal
	GetReplacementCardValue() string
	GetReplacementCardPin() string
	GetReplacementCardType() string
	GetShippingAmount() *decimal.Decimal
	GetSupplementaryData() map[string][][]string
	GetAccountType() accounttype.AccountType
	GetEmvChipCondition() *emvchipcondition.EmvChipCondition
	GetEmvLastChipRead() *emvlastchipread.EmvLastChipRead
	GetMessageAuthenticationCode() string
	IsMultiCapture() bool
	GetMiscProductData() []base.Product
	GetTagData() string
	GetTransactionInitiator() storedcredentialinitiator.StoredCredentialInitiator
	GetFeeAmount() *decimal.Decimal
	GetFeeType() feetype.FeeType
	GetFollowOnTimestamp() string
	GetShiftNumber() string
	GetTransportData() string
	GetCardHolderLanguage() string
	GetPosSequenceNumber() string
	GetReversalReasonCode() reversalreasoncode.ReversalReasonCode
	WithReversalReasonCode(code reversalreasoncode.ReversalReasonCode)
	WithAccountType(value accounttype.AccountType)
	WithAddress(value *base.Address)
	WithAddressWithType(value *base.Address, addrType addresstype.AddressType)
	WithAlias(action aliasaction.AliasAction, value string)
	WithAllowDuplicates(value bool)
	WithGenerateReceipt(value bool)
	WithAvs(value bool)
	WithAllowPartialAuth(value bool)
	WithAmount(value *decimal.Decimal)
	WithAmountEstimated(value bool)
	WithAuthAmount(value *decimal.Decimal)
	WithBalanceInquiryType(value inquirytype.InquiryType)
	WithCardHolderLanguage(value string)
	WithBills(bills ...billing.Bill)
	WithChipCondition(value emvchipcondition.EmvChipCondition)
	WithFallbackCondition(value emvfallbackcondition.EmvFallbackCondition)
	WithLastChipRead(value emvlastchipread.EmvLastChipRead)
	WithClerkId(value string)
	WithClientTransactionId(value string)
	WithCommercialRequest(value bool)
	WithConvenienceAmt(value *decimal.Decimal)
	WithCurrency(value string)
	WithCustomer(value *recurring.Customer)
	WithCustomerId(value string)
	WithCustomerIpAddress(value string)
	WithCustomerData(value *recurring.Customer)
	WithCustomData(value ...string)
	WithCvn(value string)
	WithCardSequenceNumber(value string)
	WithDecisionManager(value *entities.DecisionManager)
	WithDynamicDescriptor(value string)
	WithEcommerceInfo(value *entities.EcommerceInfo)
	WithFraudFilter(fraudFilterMode fraudfiltermode.FraudFilterMode, fraudRules ...*entities.FraudRuleCollection)
	WithGratuity(value *decimal.Decimal)
	WithHostedPaymentData(value *entities.HostedPaymentData)
	WithIdempotencyKey(value string)
	WithInvoiceNumber(value string)
	WithMessageAuthenticationCode(value string)
	WithMultiCapture(value bool)
	WithOfflineAuthCode(value string)
	WithOneTimePayment(value bool)
	WithOrderId(value string)
	WithPosSequenceNumber(value string)
	WithMiscProductData(values []base.Product)
	WithProductId(value string)
	WithPaymentApplicationVersion(value string)
	WithPaymentMethodUsageMode(value paymentmethodusagemode.PaymentMethodUsageMode)
	WithPhoneNumber(phoneCountryCode, number string, phoneType phonenumbertype.PhoneNumberType)
	WithPaymentMethod(value pabstractions.IPaymentMethod)
	WithPriorMessageInformation(value *networkentities.PriorMessageInformation)
	WithRecurringInfo(typeValue recurringtype.RecurringType, sequence recurringsequence.RecurringSequence)
	WithRequestMultiUseToken(value bool)
	WithRequestUniqueToken(value bool)
	WithTransactionId(value string)
	WithModifier(value transactionmodifier.TransactionModifier)
	WithScheduleId(value string)
	WithShippingAmt(value *decimal.Decimal)
	WithShippingDiscount(value *decimal.Decimal)
	WithOrderDetails(value *base.OrderDetails)
	WithSimulatedHostErrors(h host.Host, errors ...hosterror.HostError)
	WithStoredCredential(value *entities.StoredCredential)
	WithSupplementaryData(t string, values ...string)
	WithSurchargeAmount(value *decimal.Decimal)
	WithTimestamp(value string)
	WithTimestampAndFollowOn(value, followOn string)
	WithTagData(value string)
	WithBatchNumber(batchNumber *int)
	WithBatchAndSequenceNumber(batchNumber, sequenceNumber *int)
	WithCardBrandStorage(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator)
	WithCardBrandStorageAndTransactionId(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator, value string)
	WithCompanyId(companyId string)
	WithFee(feeType feetype.FeeType, feeAmount *decimal.Decimal)
	WithShiftNumber(value string)
	WithSystemTraceAuditNumberAndFollowOn(original *int, followOn *int)
	WithTerminalError(value bool)
	WithTransportData(value string)
	WithTaxAmount(taxAmount *decimal.Decimal)
	WithTipAmount(tipAmount *decimal.Decimal)
	WithRemittanceReference(remittanceReferenceType remittancereferencetype.RemittanceReferenceType, remittanceReferenceValue string)
	WithBNPLShippingMethod(value bnplshippingmethod.BNPLShippingMethod) error
	WithTransactiontype(t transactiontype.TransactionType)
	WithPaymentPurposeCode(paymentPurposeCode string)
	WithEmvMaxPinEntry(emvMaxPinEntry string)
	WithGoodsSold(goodsSold string)
	WithCheckCustomerId(checkCustomerId string)
	WithRawMICRData(rawMICRData string)
	WithEWICIssuingEntity(eWICIssuingEntity string)
	WithCountry(country string)
	WithMaskedDataResponse(value *bool)
	Serialize(ctx context.Context, gateway IPaymentGateway) (string, error)
	SerializeWithConfig(ctx context.Context, configName string, gateway IPaymentGateway) (string, error)
	SetupValidations()
}
