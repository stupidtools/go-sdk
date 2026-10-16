package entities

import (
	"github.com/shopspring/decimal"
)

type AdditionalTaxDetails struct {
	TaxAmount *decimal.Decimal
	TaxRate   *decimal.Decimal
	TaxType   string
}

func NewAdditionalTaxDetails() *AdditionalTaxDetails {
	return &AdditionalTaxDetails{}
}
