package builders

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
)

type UpaTerminalManageBuilder struct {
	*builders.TerminalManageBuilder
}

func NewUpaTerminalManageBuilder(t transactiontype.TransactionType, p paymentmethodtype.PaymentMethodType) *UpaTerminalManageBuilder {
	return &UpaTerminalManageBuilder{
		TerminalManageBuilder: builders.NewTerminalManageBuilder(t, p),
	}
}

func (b *UpaTerminalManageBuilder) SetupValidations() {
	b.Validations.Of(transactiontype.Capture.LongValue() | transactiontype.Void.LongValue()).Check("terminalRefNumber").IsNotNull("Terminal Reference cannot be empty")
	b.Validations.Of(paymentmethodtype.Gift.LongValue()).Check("currency").IsNotNull("Currency cannot be empty")
}
