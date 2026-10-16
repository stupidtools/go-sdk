package responses

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/summarytype"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/shopspring/decimal"
)

type UpaSummaryResponse struct {
	Amount           *decimal.Decimal
	AmountDue        *decimal.Decimal
	AuthorizedAmount *decimal.Decimal
	Count            *int
	Type             summarytype.SummaryType
	TotalAmount      *decimal.Decimal
	Transactions     []transactionsummary.TransactionSummary
}

func (s UpaSummaryResponse) GetAmount() *decimal.Decimal {
	return s.Amount
}

func (s UpaSummaryResponse) GetAmountDue() *decimal.Decimal {
	return s.AmountDue
}

func (s UpaSummaryResponse) GetAuthorizedAmount() *decimal.Decimal {
	return s.AuthorizedAmount
}

func (s UpaSummaryResponse) GetCount() *int {
	return s.Count
}

func (s UpaSummaryResponse) GetSummaryType() summarytype.SummaryType {
	return s.Type
}

func (s UpaSummaryResponse) GetTotalAmount() *decimal.Decimal {
	return s.TotalAmount
}

func (s UpaSummaryResponse) GetTransactions() []transactionsummary.TransactionSummary {
	return s.Transactions
}
