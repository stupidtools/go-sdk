package risk

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type Risk string

const (
	High Risk = "High"
	Low  Risk = "Low"
)

func (r Risk) GetBytes() []byte {
	return []byte(r)
}

func (r Risk) GetValue() string {
	return string(r)
}

func (r Risk) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{High, Low}
}
