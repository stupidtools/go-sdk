package base

import (
	"github.com/shopspring/decimal"
)

type OrderDetails struct {
	InsuranceAmount *decimal.Decimal
	HasInsurance    bool
	HandlingAmount  *decimal.Decimal
	Description     string
}
