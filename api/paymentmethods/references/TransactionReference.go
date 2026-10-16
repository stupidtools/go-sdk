package references

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/authorizercode"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontypeindicator"
	"github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/shopspring/decimal"
	"reflect"
)

type TransactionReference struct {
	AlternativePaymentType           string
	AcquiringInstitutionId           string
	authorizer                       authorizercode.AuthorizerCode
	AuthCode                         string
	BatchNumber                      *int
	ClientTransactionId              string
	MessageTypeIndicator             string
	OrderId                          string
	OriginalAmount                   *decimal.Decimal
	OriginalApprovedAmount           *decimal.Decimal
	OriginalPaymentMethod            abstractions.IPaymentMethod
	OriginalProcessingCode           string
	OriginalTransactionTime          string
	PartialApproval                  bool
	PaymentMethodType                paymentmethodtype.PaymentMethodType
	PosDataCode                      string
	SequenceNumber                   *int
	SystemTraceAuditNumber           string
	TransactionId                    string
	OriginalTransactionDate          string
	OriginalTransactionTypeIndicator transactiontypeindicator.TransactionTypeIndicator
	ResponseCode                     string
	UseAuthorizedAmount              bool
	TransactionIdentifier            string
	OriginalInvoiceNumber            string
	OriginalTransactionInfo          string
	OriginalPOSEntryMode             string
	OriginalTransactionType          transactiontype.TransactionType
	ApprovalCode                     string
	OriginalMessageCode              string
	MastercardBanknetRefNo           string
	MastercardBanknetSettlementDate  string
	DebitAuthorizer                  string
	VisaTransactionId                string
	DiscoverNetworkRefId             string
}

func (t *TransactionReference) GetOriginalTransactionTypeIndicator() transactiontypeindicator.TransactionTypeIndicator {
	return t.OriginalTransactionTypeIndicator
}

func (t *TransactionReference) SetOriginalTransactionTypeIndicator(indicator transactiontypeindicator.TransactionTypeIndicator) *TransactionReference {
	t.OriginalTransactionTypeIndicator = indicator
	return t
}

func (t *TransactionReference) GetAuthorizer() authorizercode.AuthorizerCode {
	return t.authorizer
}

func (t *TransactionReference) SetAuthorizer(authorizer authorizercode.AuthorizerCode) *TransactionReference {
	t.authorizer = authorizer
	return t
}

func (t *TransactionReference) GetOriginalApprovedAmount() *decimal.Decimal {
	if t.OriginalApprovedAmount != nil {
		return t.OriginalApprovedAmount
	}
	return t.OriginalAmount
}

func (t *TransactionReference) SetOriginalApprovedAmount(amount *decimal.Decimal) {
	t.OriginalApprovedAmount = amount
}

func (t *TransactionReference) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	if t.OriginalPaymentMethod != nil {
		return t.OriginalPaymentMethod.GetPaymentMethodType()
	}
	return t.PaymentMethodType
}

func (t *TransactionReference) GetToken() string {
	return ""
}

func (t *TransactionReference) SetPaymentMethodType(p paymentmethodtype.PaymentMethodType) *TransactionReference {
	t.PaymentMethodType = p
	return t
}

func NewTransactionReference() *TransactionReference {
	return &TransactionReference{}
}

func IsItATransactionReference(i interface{}) bool {
	t := reflect.TypeOf(i)
	typeString := t.String()
	return typeString == "paymentmethods.TransactionReference"
}
func (t *TransactionReference) GetOriginalAmount() *decimal.Decimal {
	return t.OriginalAmount
}

func (t *TransactionReference) GetOriginalPaymentMethod() abstractions.IPaymentMethod {
	return t.OriginalPaymentMethod
}

func (t *TransactionReference) GetOriginalProcessingCode() string {
	return t.OriginalProcessingCode
}

func (t *TransactionReference) GetOriginalTransactionTime() string {
	return t.OriginalTransactionTime
}

func (t *TransactionReference) IsPartialApproval() bool {
	return t.PartialApproval
}

func (t *TransactionReference) GetPosDataCode() string {
	return t.PosDataCode
}

func (t *TransactionReference) GetSequenceNumber() *int {
	return t.SequenceNumber
}

func (t *TransactionReference) GetSystemTraceAuditNumber() string {
	return t.SystemTraceAuditNumber
}

func (t *TransactionReference) GetTransactionId() string {
	return t.TransactionId
}

func (t *TransactionReference) GetOriginalTransactionDate() string {
	return t.OriginalTransactionDate
}

func (t *TransactionReference) GetResponseCode() string {
	return t.ResponseCode
}

func (t *TransactionReference) IsUseAuthorizedAmount() bool {
	return t.UseAuthorizedAmount
}

func (t *TransactionReference) GetTransactionIdentifier() string {
	return t.TransactionIdentifier
}

func (t *TransactionReference) GetOriginalInvoiceNumber() string {
	return t.OriginalInvoiceNumber
}

func (t *TransactionReference) GetOriginalTransactionInfo() string {
	return t.OriginalTransactionInfo
}

func (t *TransactionReference) GetOriginalPOSEntryMode() string {
	return t.OriginalPOSEntryMode
}

func (t *TransactionReference) GetOriginalTransactionType() transactiontype.TransactionType {
	return t.OriginalTransactionType
}

func (t *TransactionReference) GetApprovalCode() string {
	return t.ApprovalCode
}

func (t *TransactionReference) GetOriginalMessageCode() string {
	return t.OriginalMessageCode
}

func (t *TransactionReference) GetMastercardBanknetRefNo() string {
	return t.MastercardBanknetRefNo
}

func (t *TransactionReference) GetMastercardBanknetSettlementDate() string {
	return t.MastercardBanknetSettlementDate
}

func (t *TransactionReference) GetDebitAuthorizer() string {
	return t.DebitAuthorizer
}

func (t *TransactionReference) GetVisaTransactionId() string {
	return t.VisaTransactionId
}

func (t *TransactionReference) GetDiscoverNetworkRefId() string {
	return t.DiscoverNetworkRefId
}

func (t *TransactionReference) SetOriginalAmount(amount *decimal.Decimal) *TransactionReference {
	t.OriginalAmount = amount
	return t
}

func (t *TransactionReference) SetOriginalPaymentMethod(pm abstractions.IPaymentMethod) *TransactionReference {
	t.OriginalPaymentMethod = pm
	return t
}

func (t *TransactionReference) SetOriginalProcessingCode(code string) *TransactionReference {
	t.OriginalProcessingCode = code
	return t
}

func (t *TransactionReference) SetOriginalTransactionTime(time string) *TransactionReference {
	t.OriginalTransactionTime = time
	return t
}

func (t *TransactionReference) SetPartialApproval(pa bool) *TransactionReference {
	t.PartialApproval = pa
	return t
}

func (t *TransactionReference) SetPosDataCode(code string) *TransactionReference {
	t.PosDataCode = code
	return t
}

func (t *TransactionReference) SetSequenceNumber(number *int) *TransactionReference {
	t.SequenceNumber = number
	return t
}

func (t *TransactionReference) SetSystemTraceAuditNumber(number string) *TransactionReference {
	t.SystemTraceAuditNumber = number
	return t
}

func (t *TransactionReference) SetTransactionId(id string) *TransactionReference {
	t.TransactionId = id
	return t
}

func (t *TransactionReference) SetOriginalTransactionDate(date string) *TransactionReference {
	t.OriginalTransactionDate = date
	return t
}

func (t *TransactionReference) SetResponseCode(code string) *TransactionReference {
	t.ResponseCode = code
	return t
}

func (t *TransactionReference) SetUseAuthorizedAmount(use bool) *TransactionReference {
	t.UseAuthorizedAmount = use
	return t
}

func (t *TransactionReference) SetTransactionIdentifier(identifier string) *TransactionReference {
	t.TransactionIdentifier = identifier
	return t
}

func (t *TransactionReference) SetOriginalInvoiceNumber(invoiceNumber string) *TransactionReference {
	t.OriginalInvoiceNumber = invoiceNumber
	return t
}

func (t *TransactionReference) SetOriginalTransactionInfo(info string) *TransactionReference {
	t.OriginalTransactionInfo = info
	return t
}

func (t *TransactionReference) SetOriginalPOSEntryMode(entryMode string) *TransactionReference {
	t.OriginalPOSEntryMode = entryMode
	return t
}

func (t *TransactionReference) SetOriginalTransactionType(tType transactiontype.TransactionType) *TransactionReference {
	t.OriginalTransactionType = tType
	return t
}

func (t *TransactionReference) SetApprovalCode(code string) *TransactionReference {
	t.ApprovalCode = code
	return t
}

func (t *TransactionReference) SetAuthCode(code string) *TransactionReference {
	t.AuthCode = code
	return t
}

func (t *TransactionReference) SetOriginalMessageCode(code string) *TransactionReference {
	t.OriginalMessageCode = code
	return t
}

func (t *TransactionReference) SetMastercardBanknetRefNo(refNo string) *TransactionReference {
	t.MastercardBanknetRefNo = refNo
	return t
}

func (t *TransactionReference) SetMastercardBanknetSettlementDate(date string) *TransactionReference {
	t.MastercardBanknetSettlementDate = date
	return t
}

func (t *TransactionReference) SetDebitAuthorizer(authorizer string) *TransactionReference {
	t.DebitAuthorizer = authorizer
	return t
}

func (t *TransactionReference) SetVisaTransactionId(transactionId string) *TransactionReference {
	t.VisaTransactionId = transactionId
	return t
}

func (t *TransactionReference) SetDiscoverNetworkRefId(refId string) *TransactionReference {
	t.DiscoverNetworkRefId = refId
	return t
}

func (t *TransactionReference) GetAlternativePaymentType() string {
	return t.AlternativePaymentType
}

func (t *TransactionReference) SetAlternativePaymentType(paymentType string) *TransactionReference {
	t.AlternativePaymentType = paymentType
	return t
}

func (t *TransactionReference) GetAcquiringInstitutionId() string {
	return t.AcquiringInstitutionId
}

func (t *TransactionReference) SetAcquiringInstitutionId(id string) *TransactionReference {
	t.AcquiringInstitutionId = id
	return t
}

func (t *TransactionReference) GetAuthCode() string {
	return t.AuthCode
}

func (t *TransactionReference) GetBatchNumber() *int {
	return t.BatchNumber
}

func (t *TransactionReference) SetBatchNumber(number *int) *TransactionReference {
	t.BatchNumber = number
	return t
}

func (t *TransactionReference) GetClientTransactionId() string {
	return t.ClientTransactionId
}

func (t *TransactionReference) SetClientTransactionId(id string) *TransactionReference {
	t.ClientTransactionId = id
	return t
}

func (t *TransactionReference) GetMessageTypeIndicator() string {
	return t.MessageTypeIndicator
}

func (t *TransactionReference) SetMessageTypeIndicator(indicator string) *TransactionReference {
	t.MessageTypeIndicator = indicator
	return t
}

func (t *TransactionReference) GetOrderId() string {
	return t.OrderId
}

func (t *TransactionReference) SetOrderId(id string) *TransactionReference {
	t.OrderId = id
	return t
}
