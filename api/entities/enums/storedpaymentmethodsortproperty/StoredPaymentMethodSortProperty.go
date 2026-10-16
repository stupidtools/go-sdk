package storedpaymentmethodsortproperty

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
)

type StoredPaymentMethodSortProperty string

const (
	TimeCreated StoredPaymentMethodSortProperty = "TIME_CREATED"
)

var storedPaymentMethodSortPropertyValues = map[StoredPaymentMethodSortProperty]map[target.Target]string{
	TimeCreated: {
		target.GP_API: "TIME_CREATED",
	},
}

func (sp StoredPaymentMethodSortProperty) GetBytes() []byte {
	if val, ok := storedPaymentMethodSortPropertyValues[sp][target.GP_API]; ok {
		return []byte(val)
	}
	return nil
}

func (sp StoredPaymentMethodSortProperty) GetValue() string {
	if val, ok := storedPaymentMethodSortPropertyValues[sp][target.GP_API]; ok {
		return val
	}
	return ""
}

func (StoredPaymentMethodSortProperty) StringConstants() []istringconstant.IStringConstant {
	values := make([]istringconstant.IStringConstant, len(storedPaymentMethodSortPropertyValues))
	i := 0
	for k := range storedPaymentMethodSortPropertyValues {
		values[i] = k
		i++
	}
	return values
}
