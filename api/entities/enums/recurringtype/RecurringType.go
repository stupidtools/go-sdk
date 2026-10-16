package recurringtype

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type RecurringType string

const (
	Fixed    RecurringType = "Fixed"
	Variable RecurringType = "Variable"
)

func (r RecurringType) GetBytes() []byte {
	return []byte(r)
}

func (r RecurringType) GetValue() string {
	return string(r)
}

func (r RecurringType) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{Fixed, Variable}
}
