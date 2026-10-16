package abstractions

import "github.com/shopspring/decimal"

type ISAFSummaryReport interface {
	IDeviceResponse
	GetSafTotalAmount() *decimal.Decimal
	SetSafTotalAmount(*decimal.Decimal)
	GetSafTotalCount() *int
	SetSafTotalCount(int)
}
