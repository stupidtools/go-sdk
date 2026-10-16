package recurringsequence

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type RecurringSequence string

const (
	First      RecurringSequence = "First"
	Subsequent RecurringSequence = "Subsequent"
	Last       RecurringSequence = "Last"
)

func (r RecurringSequence) GetBytes() []byte {
	return []byte(r)
}

func (r RecurringSequence) GetValue() string {
	return string(r)
}

func (r RecurringSequence) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{First, Subsequent, Last}
}
