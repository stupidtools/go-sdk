package entities

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type PaymentDataSourceType string

const (
	APPLEPAY         PaymentDataSourceType = "ApplePay"
	APPLEPAYAPP      PaymentDataSourceType = "ApplePayApp"
	APPLEPAYWEB      PaymentDataSourceType = "ApplePayWeb"
	GOOGLEPAYAPP     PaymentDataSourceType = "GooglePayApp"
	GOOGLEPAYWEB     PaymentDataSourceType = "GooglePayWeb"
	DISCOVER3DSECURE PaymentDataSourceType = "Discover 3DSecure"
)

func (p PaymentDataSourceType) GetBytes() []byte {
	return []byte(p)
}

func (p PaymentDataSourceType) GetValue() string {
	return string(p)
}

func (PaymentDataSourceType) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{APPLEPAY, APPLEPAYAPP, APPLEPAYWEB, GOOGLEPAYAPP, GOOGLEPAYWEB, DISCOVER3DSECURE}
}
