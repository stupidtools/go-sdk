package entities

import "github.com/shopspring/decimal"

type DiscountDetails struct {
	DiscountName        string
	DiscountAmount      *decimal.Decimal
	DiscountPercentage  *decimal.Decimal
	DiscountType        string
	DiscountPriority    *int
	DiscountIsStackable bool
}

func NewDiscountDetails() *DiscountDetails {
	return &DiscountDetails{}
}
