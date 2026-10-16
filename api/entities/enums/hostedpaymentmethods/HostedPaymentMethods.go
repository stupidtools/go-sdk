package hostedpaymentmethods

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type HostedPaymentMethods string

const (
	OB    HostedPaymentMethods = "ob"
	CARDS HostedPaymentMethods = "cards"
)

func (h HostedPaymentMethods) GetBytes() []byte {
	return []byte(h)
}

func (h HostedPaymentMethods) GetValue() string {
	return string(h)
}

func (h HostedPaymentMethods) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{OB, CARDS}
}
