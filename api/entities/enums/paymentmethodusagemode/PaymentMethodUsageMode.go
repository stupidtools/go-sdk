package paymentmethodusagemode

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type PaymentMethodUsageMode string

const (
	SINGLE   PaymentMethodUsageMode = "SINGLE"
	MULTIPLE PaymentMethodUsageMode = "MULTIPLE"
)

func (pmum PaymentMethodUsageMode) GetBytes() []byte {
	return []byte(pmum)
}

func (pmum PaymentMethodUsageMode) GetValue() string {
	return string(pmum)
}

func (pmum PaymentMethodUsageMode) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		SINGLE,
		MULTIPLE,
	}
}
