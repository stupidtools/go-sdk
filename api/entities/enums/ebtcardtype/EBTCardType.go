package ebtcardtype

import (
	"errors"
)

type EBTCardType string

const (
	CashBenefit EBTCardType = "CashBenefit"
	FoodStamp   EBTCardType = "FoodStamp"
)

func StringConstants() []EBTCardType {
	return []EBTCardType{CashBenefit, FoodStamp}
}

func ParseFromString(s string) (EBTCardType, error) {
	for _, c := range StringConstants() {
		if string(c) == s {
			return c, nil
		}
	}
	return CashBenefit, errors.New("unsupported type")
}
