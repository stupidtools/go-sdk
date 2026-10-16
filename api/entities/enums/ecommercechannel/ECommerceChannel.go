package ecommercechannel

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type EcommerceChannel string

const (
	Ecom EcommerceChannel = "ECOM"
	Moto EcommerceChannel = "MOTO"
)

func (e EcommerceChannel) GetBytes() []byte {
	return []byte(e)
}

func (e EcommerceChannel) GetValue() string {
	return string(e)
}

func (e EcommerceChannel) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{Ecom, Moto}
}
