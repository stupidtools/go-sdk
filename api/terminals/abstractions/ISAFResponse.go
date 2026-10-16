package abstractions

import (
	"github.com/shopspring/decimal"
)

type ISAFResponse interface {
	IDeviceResponse
	GetSafTotalCount() *int
	GetSafTotalAmount() *decimal.Decimal
	GetApproved() []ISummaryResponse
	GetPending() []ISummaryResponse
	GetDeclined() []ISummaryResponse
}
