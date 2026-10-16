package timezoneconversion

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type TimeZoneConversion string

const (
	UTC        TimeZoneConversion = "UTC"
	Merchant   TimeZoneConversion = "Merchant"
	Datacenter TimeZoneConversion = "Datacenter"
)

func (t TimeZoneConversion) GetBytes() []byte {
	return []byte(t)
}

func (t TimeZoneConversion) GetValue() string {
	return string(t)
}

func (t TimeZoneConversion) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{UTC, Merchant, Datacenter}
}
