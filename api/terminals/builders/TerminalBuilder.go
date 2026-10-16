package builders

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
)

type TerminalBuilder struct {
	*builders.TransactionBuilder
	paymentMethodType paymentmethodtype.PaymentMethodType
	requestId         *int
	clerkId           *int
	referenceNumber   string
	EcrId             string
}

func (tb *TerminalBuilder) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	return tb.paymentMethodType
}

func (tb *TerminalBuilder) GetRequestId() *int {
	return tb.requestId
}

func (tb *TerminalBuilder) GetClerkId() *int {
	return tb.clerkId
}

func (tb *TerminalBuilder) GetReferenceNumber() string {
	return tb.referenceNumber
}

func (tb *TerminalBuilder) WithClerkId(value *int) *TerminalBuilder {
	tb.clerkId = value
	return tb
}

func (tb *TerminalBuilder) WithReferenceNumber(value string) *TerminalBuilder {
	tb.referenceNumber = value
	return tb
}

func (tb *TerminalAuthBuilder) WithEcrId(value string) *TerminalAuthBuilder {
	tb.EcrId = value
	return tb
}

func (tb *TerminalBuilder) WithRequestId(value *int) *TerminalBuilder {
	tb.requestId = value
	return tb
}

func NewTerminalBuilder(transactionType transactiontype.TransactionType, paymentMethod paymentmethodtype.PaymentMethodType) *TerminalBuilder {
	return &TerminalBuilder{
		TransactionBuilder: builders.NewTransactionBuilder(transactionType, nil),
		paymentMethodType:  paymentMethod,
	}
}
