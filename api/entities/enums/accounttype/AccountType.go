package accounttype

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
)

type AccountType struct {
	value map[target.Target]string
}

// The defined values
var (
	Checking = AccountType{
		value: map[target.Target]string{
			target.DEFAULT: "CHECKING",
			target.GP_API:  "CHECKING",
		},
	}

	Savings = AccountType{
		value: map[target.Target]string{
			target.DEFAULT: "SAVINGS",
			target.GP_API:  "SAVING",
		},
	}

	Credit = AccountType{
		value: map[target.Target]string{
			target.DEFAULT: "CREDIT",
			target.GP_API:  "CREDIT",
		},
	}
)

// GetValue provides the value based on the provided target.
func (at AccountType) GetValue(t target.Target) string {
	if val, ok := at.value[t]; ok {
		return val
	}
	if val, ok := at.value[target.DEFAULT]; ok {
		return val
	}
	return ""
}
