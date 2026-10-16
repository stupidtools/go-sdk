package builders

import (
	"context"
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/currencytype"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
)

type ITerminalManageBuilder interface {
	GetAmount() *decimal.Decimal
	GetCurrency() currencytype.CurrencyType
	GetGratuity() *decimal.Decimal
	GetTransactionId() string
	GetTerminalRefNumber() string
	WithAmount(value *decimal.Decimal) *TerminalManageBuilder
	WithCurrency(value currencytype.CurrencyType) *TerminalManageBuilder
	WithGratuity(value *decimal.Decimal) *TerminalManageBuilder
	WithTerminalRefNumber(value string) *TerminalManageBuilder
	ExecuteWithName(ctx context.Context, configName string, device ITerminalBuilderDevice) (terminalresponse.ITerminalResponse, error)
	SetupValidations()
}
