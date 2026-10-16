package actionsortproperty

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
)

type ActionSortProperty string

const (
	TimeCreated ActionSortProperty = "TIME_CREATED"
)

var actionSortPropertyValues = map[ActionSortProperty]map[target.Target]string{
	TimeCreated: {
		target.GP_API: "TIME_CREATED",
	},
}

func (ap ActionSortProperty) GetBytes() []byte {
	if val, ok := actionSortPropertyValues[ap][target.GP_API]; ok {
		return []byte(val)
	}
	return nil
}

func (ap ActionSortProperty) GetValue() string {
	if val, ok := actionSortPropertyValues[ap][target.GP_API]; ok {
		return val
	}
	return ""
}

func (ActionSortProperty) StringConstants() []istringconstant.IStringConstant {
	values := make([]istringconstant.IStringConstant, 0, len(actionSortPropertyValues))
	for k := range actionSortPropertyValues {
		values = append(values, k)
	}
	return values
}
