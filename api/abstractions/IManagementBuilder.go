package abstractions

import (
	"context"
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
	"github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/shopspring/decimal"
)

type IManagementBuilder interface {
	ITransactionBuilder
	Execute(ctx context.Context, gateway IPaymentGateway) (ITransaction, error)
	WithMiscProductData(values []base.Product)
	GetAlternativePaymentType() alternativepaymenttype.AlternativePaymentType
	GetAccountType() accounttype.AccountType
	GetAmount() *decimal.Decimal
	GetAuthAmount() *decimal.Decimal
	GetBills() []billing.Bill
	GetCashBackAmount() *decimal.Decimal
	GetClerkId() string
	GetClientTransactionId() string
	GetTransactionId() string
	GetCommercialData() *entities.CommercialData
	GetConvenienceAmount() *decimal.Decimal
	GetCurrency() string
	GetCustomerId() string
	GetCustomerIpAddress() string
	IsCustomerInitiated() bool
	GetDescription() string
	IsForceToHost() bool
	GetEmvChipCondition() *emvchipcondition.EmvChipCondition
	GetGratuity() *decimal.Decimal
	GetInvoiceNumber() string
	GetOrderId() string
	GetPayerAuthenticationResponse() string
	GetPoNumber() string
	GetPosSequenceNumber() string
	GetProductId() string
	GetReferenceNumber() string
	GetReversalReasonCode() reversalreasoncode.ReversalReasonCode
	GetShiftNumber() string
	GetSupplementaryData() map[string][][2]string
	GetSurchargeAmount() *decimal.Decimal
	GetTagData() string
	GetTaxAmount() *decimal.Decimal
	GetTaxType() taxtype.TaxType
	GetTimestamp() string
	GetTransportData() string
	GetTransactionCount() *int
	GetTotalCredits() *decimal.Decimal
	GetTotalDebits() *decimal.Decimal
	GetTransactionInitiator() storedcredentialinitiator.StoredCredentialInitiator
	GetCardBrandTransactionId() string
	GetCardType() string
	WithCommercialData(cd *entities.CommercialData)
	WithCardType(s string)
	WithSettlementAmount(value *decimal.Decimal)
	WithPaymentPurposeCode(paymentPurposeCode string)
	WithAccountType(value accounttype.AccountType)
	WithApprovalCode(value string)
	WithDataCollectResponseCode(value int)
	WithAlternativePaymentType(value alternativepaymenttype.AlternativePaymentType)
	WithAmount(value *decimal.Decimal)
	WithChipCondition(value emvchipcondition.EmvChipCondition)
	WithAuthAmount(value *decimal.Decimal)
	WithBatchTotals(transactionCount int, totalDebits, totalCredits *decimal.Decimal)
	WithBatchTotalsAmount(totalAmount, totalDebits, totalCredits *decimal.Decimal)
	WithBatchTotalTransaction(transactionCount int, totalSales, totalReturns *decimal.Decimal)
	WithBatchReference(value string)
	WithBills(bills ...billing.Bill)
	WithMultiCapture(sequence *int)
	WithMultiCaptureAndPaymentCount(sequence, paymentCount *int)
	WithCardBrandStorage(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator)
	WithCardBrandStorageAndTransactionId(transactionInitiator storedcredentialinitiator.StoredCredentialInitiator, value string)
	WithCashBackAmount(value *decimal.Decimal)
	WithClerkId(value string)
	WithClientTransactionId(value string)
	WithConvenienceAmt(value *decimal.Decimal)
	WithCurrency(value string)
	WithCustomerId(value string)
	WithCustomerIpAddress(value string)
	WithCustomerInitiated(value bool)
	WithDisputeId(value string)
	WithDynamicDescriptor(value string)
	WithForceToHost(value bool)
	WithGratuity(value *decimal.Decimal)
	WithInvoiceNumber(value string)
	WithIdempotencyKey(value string)
	WithPaymentMethod(value abstractions.IPaymentMethod)
	WithPayerAuthenticationResponse(value string)
	WithPoNumber(value string)
	WithPosSequenceNumber(value string)
	WithProductId(value string)
	WithReferenceNumber(value string)
	WithShiftNumber(value string)
	WithEcommerceInfo(value entities.EcommerceInfo)
	WithStoredCredential(value entities.StoredCredential)
	WithSupplementaryData(typeKey string, values ...string)
	WithSurchargeAmount(value *decimal.Decimal)
	WithTagData(value string)
	WithPaymentMethodUsageMode(value paymentmethodusagemode.PaymentMethodUsageMode)
	WithTaxAmount(value *decimal.Decimal)
	WithTaxType(value taxtype.TaxType)
	WithTimestamp(value string)
	WithTransportData(value string)
	WithReference(value string)
	WithReversalReasonCode(code reversalreasoncode.ReversalReasonCode)
}
