package rebuilders

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/authorizercode"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontypeindicator"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	paymentmethods "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/shopspring/decimal"
)

type TransactionRebuilder struct {
	authCode                 string
	acquirerId               string
	messageTypeIndicator     string
	orderId                  string
	originalAmount           *decimal.Decimal
	originalApprovedAmount   *decimal.Decimal
	originalPaymentMethod    paymentmethods.IPaymentMethod
	originalProcessingCode   string
	originalTransactionTime  string
	originalTransactionDate  string
	partialApproval          bool
	paymentMethodType        paymentmethodtype.PaymentMethodType
	posDataCode              string
	systemTraceAuditNumber   string
	transactionId            string
	useAuthorizedAmount      bool
	originalTransactionType  transactiontype.TransactionType
	approvalCode             string
	authorizer               authorizercode.AuthorizerCode
	debitAuthorizer          string
	banknetRefId             string
	settlementDate           string
	visaTransactionId        string
	discoverNetworkRefId     string
	originalMessageCode      string
	transactionTypeIndicator transactiontypeindicator.TransactionTypeIndicator
	batchNumber              *int
	sequenceNumber           *int
}

func NewTransactionRebuilder(methodType paymentmethodtype.PaymentMethodType) *TransactionRebuilder {
	return &TransactionRebuilder{
		paymentMethodType: methodType,
	}
}

func (r *TransactionRebuilder) WithAuthorizationCode(value string) *TransactionRebuilder {
	r.authCode = value
	return r
}

func (r *TransactionRebuilder) WithAcquirerId(value string) *TransactionRebuilder {
	r.acquirerId = value
	return r
}

func (r *TransactionRebuilder) WithMessageTypeIndicator(value string) *TransactionRebuilder {
	r.messageTypeIndicator = value
	return r
}

func (r *TransactionRebuilder) WithOrderId(value string) *TransactionRebuilder {
	r.orderId = value
	return r
}

func (r *TransactionRebuilder) WithAmount(value *decimal.Decimal) *TransactionRebuilder {
	r.originalAmount = value
	return r
}

func (r *TransactionRebuilder) WithAuthorizedAmount(value *decimal.Decimal, useAuthorizedAmount bool) *TransactionRebuilder {
	r.originalApprovedAmount = value
	r.useAuthorizedAmount = useAuthorizedAmount
	return r
}

func (r *TransactionRebuilder) WithPartialApproval(value bool) *TransactionRebuilder {
	r.partialApproval = value
	return r
}

func (r *TransactionRebuilder) WithPaymentMethod(value paymentmethods.IPaymentMethod) *TransactionRebuilder {
	r.originalPaymentMethod = value
	return r
}

func (r *TransactionRebuilder) WithProcessingCode(value string) *TransactionRebuilder {
	r.originalProcessingCode = value
	return r
}

func (r *TransactionRebuilder) WithPosDataCode(value string) *TransactionRebuilder {
	r.posDataCode = value
	return r
}

func (r *TransactionRebuilder) WithTransactionTime(value string) *TransactionRebuilder {
	r.originalTransactionTime = value
	return r
}

func (r *TransactionRebuilder) WithPaymentMethodType(value paymentmethodtype.PaymentMethodType) *TransactionRebuilder {
	r.paymentMethodType = value
	return r
}

func (r *TransactionRebuilder) WithSystemTraceAuditNumber(value string) *TransactionRebuilder {
	r.systemTraceAuditNumber = value
	return r
}

func (r *TransactionRebuilder) WithTransactionId(value string) *TransactionRebuilder {
	r.transactionId = value
	return r
}

func (r *TransactionRebuilder) WithOriginalTransactionType(value transactiontype.TransactionType) *TransactionRebuilder {
	r.originalTransactionType = value
	return r
}

func (r *TransactionRebuilder) WithApprovalCode(value string) *TransactionRebuilder {
	r.approvalCode = value
	return r
}

func (r *TransactionRebuilder) WithAuthorizer(value authorizercode.AuthorizerCode) *TransactionRebuilder {
	r.authorizer = value
	return r
}

func (r *TransactionRebuilder) WithDebitAuthorizer(value string) *TransactionRebuilder {
	r.debitAuthorizer = value
	return r
}

func (r *TransactionRebuilder) WithSettlementDate(value string) *TransactionRebuilder {
	r.settlementDate = value
	return r
}

func (r *TransactionRebuilder) WithBanknetRefId(value string) *TransactionRebuilder {
	r.banknetRefId = value
	return r
}

func (r *TransactionRebuilder) WithVisaTransactionId(value string) *TransactionRebuilder {
	r.visaTransactionId = value
	return r
}

func (r *TransactionRebuilder) WithDiscoverNetworkRefId(value string) *TransactionRebuilder {
	r.discoverNetworkRefId = value
	return r
}

func (r *TransactionRebuilder) WithOriginalTransactionDate(value string) *TransactionRebuilder {
	r.originalTransactionDate = value
	return r
}

func (r *TransactionRebuilder) WithOriginalMessageCode(value string) *TransactionRebuilder {
	r.originalMessageCode = value
	return r
}

func (r *TransactionRebuilder) WithTransactionTypeIndicator(value transactiontypeindicator.TransactionTypeIndicator) *TransactionRebuilder {
	r.transactionTypeIndicator = value
	return r
}

func (r *TransactionRebuilder) WithBatchNumber(batchNumber *int) *TransactionRebuilder {
	r.batchNumber = batchNumber
	return r
}

func (r *TransactionRebuilder) WithSequenceNumber(sequenceNumber *int) *TransactionRebuilder {
	r.sequenceNumber = sequenceNumber
	return r
}

func (r *TransactionRebuilder) Build() *transactions.Transaction {
	reference := references.NewTransactionReference()
	reference.SetAcquiringInstitutionId(r.acquirerId)
	reference.SetAuthCode(r.authCode)
	reference.SetMessageTypeIndicator(r.messageTypeIndicator)
	reference.SetOriginalAmount(r.originalAmount)
	reference.SetOrderId(r.orderId)
	reference.SetOriginalApprovedAmount(r.originalApprovedAmount)
	reference.SetOriginalPaymentMethod(r.originalPaymentMethod)
	reference.SetOriginalProcessingCode(r.originalProcessingCode)
	reference.SetOriginalTransactionTime(r.originalTransactionTime)
	reference.SetPartialApproval(r.partialApproval)
	reference.SetPaymentMethodType(r.paymentMethodType)
	reference.SetPosDataCode(r.posDataCode)
	reference.SetSystemTraceAuditNumber(r.systemTraceAuditNumber)
	reference.SetTransactionId(r.transactionId)
	reference.SetUseAuthorizedAmount(r.useAuthorizedAmount)
	reference.SetOriginalTransactionType(r.originalTransactionType)
	reference.SetAuthorizer(r.authorizer)
	reference.SetDebitAuthorizer(r.debitAuthorizer)
	reference.SetOriginalTransactionDate(r.originalTransactionDate)
	reference.SetApprovalCode(r.approvalCode)
	reference.SetMastercardBanknetRefNo(r.banknetRefId)
	reference.SetMastercardBanknetSettlementDate(r.settlementDate)
	reference.SetVisaTransactionId(r.visaTransactionId)
	reference.SetDiscoverNetworkRefId(r.discoverNetworkRefId)
	reference.SetOriginalMessageCode(r.originalMessageCode)
	reference.SetOriginalTransactionTypeIndicator(r.transactionTypeIndicator)
	if r.batchNumber != nil {
		reference.SetBatchNumber(r.batchNumber)
	}
	if r.sequenceNumber != nil {
		reference.SetSequenceNumber(r.sequenceNumber)
	}

	trans := transactions.NewTransaction()
	trans.TransactionReference = reference

	return trans
}

func FromId(transactionId string, paymentMethodType paymentmethodtype.PaymentMethodType) *transactions.Transaction {
	rebuilder := NewTransactionRebuilder(paymentMethodType)
	rebuilder.WithTransactionId(transactionId)
	rebuilder.WithPaymentMethodType(paymentMethodType)

	return rebuilder.Build()
}
