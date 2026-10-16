package fraudfiltermode

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type FraudFilterMode string

const (
	None    FraudFilterMode = "NONE"
	Off     FraudFilterMode = "OFF"
	Passive FraudFilterMode = "PASSIVE"
	Active  FraudFilterMode = "ACTIVE"
)

func (ffm FraudFilterMode) GetBytes() []byte {
	return []byte(ffm)
}

func (ffm FraudFilterMode) GetValue() string {
	return string(ffm)
}

func (ffm FraudFilterMode) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{None, Off, Passive, Active}
}

func FromString(value string) (FraudFilterMode, bool) {
	for _, ffm := range []FraudFilterMode{None, Off, Passive, Active} {
		if string(ffm) == value {
			return ffm, true
		}
	}
	return "", false // Returning an empty FraudFilterMode and false as it was not found.
}
