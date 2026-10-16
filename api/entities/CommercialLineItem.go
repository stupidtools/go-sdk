package entities

import (
	"github.com/shopspring/decimal"
)

type CommercialLineItem struct {
	AlternateTaxId  string
	CommodityCode   string
	Description     string
	ExtendedAmount  *decimal.Decimal
	Name            string
	ProductCode     string
	Quantity        *decimal.Decimal
	UnitOfMeasure   string
	UnitCost        *decimal.Decimal
	TaxAmount       *decimal.Decimal
	TaxName         *decimal.Decimal // If this is supposed to be a string, change the type to string
	Upc             string
	TaxPercentage   *decimal.Decimal
	TotalAmount     *decimal.Decimal
	DiscountDetails *DiscountDetails
}

func NewCommercialLineItem() *CommercialLineItem {
	return &CommercialLineItem{}
}
