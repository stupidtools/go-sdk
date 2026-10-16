package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/summarytype"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/shopspring/decimal"
)

type ISummaryResponse interface {
	GetAmount() *decimal.Decimal
	GetAmountDue() *decimal.Decimal
	GetAuthorizedAmount() *decimal.Decimal
	GetCount() *int
	GetSummaryType() summarytype.SummaryType
	GetTotalAmount() *decimal.Decimal
	GetTransactions() []transactionsummary.TransactionSummary
}
