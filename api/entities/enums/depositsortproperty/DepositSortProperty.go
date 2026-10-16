package depositsortproperty

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type DepositSortProperty string

const (
	TimeCreated DepositSortProperty = "TIME_CREATED"
	Status      DepositSortProperty = "STATUS"
	Type        DepositSortProperty = "TYPE"
	DepositId   DepositSortProperty = "DEPOSIT_ID"
)

// GetValue returns the string value of the DepositSortProperty
func (dsp DepositSortProperty) GetValue() string {
	return string(dsp)
}

// GetBytes returns the byte slice representation of the DepositSortProperty
func (dsp DepositSortProperty) GetBytes() []byte {
	return []byte(dsp)
}

// StringConstants returns a slice of all DepositSortProperty values
func (DepositSortProperty) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		TimeCreated,
		Status,
		Type,
		DepositId,
	}
}
