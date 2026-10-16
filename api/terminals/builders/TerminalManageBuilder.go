package builders

import (
	"context"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/currencytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
)

type TerminalManageBuilder struct {
	*TerminalBuilder
	amount            *decimal.Decimal
	currency          currencytype.CurrencyType
	gratuity          *decimal.Decimal
	transactionId     string
	terminalRefNumber string
}

func (builder *TerminalManageBuilder) GetAmount() *decimal.Decimal {
	return builder.amount
}

func (builder *TerminalManageBuilder) GetTransactionId() string {
	return builder.transactionId
}

func (builder *TerminalManageBuilder) GetCurrency() currencytype.CurrencyType {
	return builder.currency
}

func (builder *TerminalManageBuilder) GetGratuity() *decimal.Decimal {
	return builder.gratuity
}

func (builder *TerminalManageBuilder) GetTerminalRefNumber() string {
	return builder.terminalRefNumber
}

func (builder *TerminalManageBuilder) WithAmount(value *decimal.Decimal) *TerminalManageBuilder {
	builder.amount = value
	return builder
}

func (builder *TerminalManageBuilder) WithCurrency(value currencytype.CurrencyType) *TerminalManageBuilder {
	builder.currency = value
	return builder
}

func (builder *TerminalManageBuilder) WithGratuity(value *decimal.Decimal) *TerminalManageBuilder {
	builder.gratuity = value
	return builder
}

func (builder *TerminalManageBuilder) WithTerminalRefNumber(value string) *TerminalManageBuilder {
	builder.terminalRefNumber = value
	return builder
}

func (builder *TerminalManageBuilder) WithTransactionId(value string) *TerminalManageBuilder {
	paymentMethod := builder.GetPaymentMethod()
	if !references.IsItATransactionReference(paymentMethod) {
		paymentMethod = references.NewTransactionReference()
	}
	tr := paymentMethod.(*references.TransactionReference)
	tr.SetTransactionId(value)
	builder.SetPaymentMethod(tr)
	builder.transactionId = value
	return builder
}

func NewTerminalManageBuilder(transactionType transactiontype.TransactionType, paymentType paymentmethodtype.PaymentMethodType) *TerminalManageBuilder {
	return &TerminalManageBuilder{
		TerminalBuilder: NewTerminalBuilder(transactionType, paymentType),
	}
}

func (builder *TerminalManageBuilder) ExecuteWithName(ctx context.Context, configName string, device ITerminalBuilderDevice) (terminalresponse.ITerminalResponse, error) {
	if err := builder.TerminalBuilder.Execute(configName); err != nil {
		return nil, err
	}

	return device.ManageTransactionWithContext(ctx, builder)

}

func (builder *TerminalManageBuilder) SetupValidations() {

}
