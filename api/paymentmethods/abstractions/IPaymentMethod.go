package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
)

type IPaymentMethod interface {
	GetPaymentMethodType() paymentmethodtype.PaymentMethodType
}
