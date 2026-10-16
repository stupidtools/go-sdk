package paymentmethodtype

import (
	"math"
)

type PaymentMethodType string

const (
	Nil         PaymentMethodType = ""
	Reference   PaymentMethodType = "Reference"
	Credit      PaymentMethodType = "Credit"
	Debit       PaymentMethodType = "Debit"
	EBT         PaymentMethodType = "EBT"
	Cash        PaymentMethodType = "Cash"
	ACH         PaymentMethodType = "ACH"
	Gift        PaymentMethodType = "Gift"
	Recurring   PaymentMethodType = "Recurring"
	Other       PaymentMethodType = "Other"
	APM         PaymentMethodType = "APM"
	Ewic        PaymentMethodType = "Ewic"
	BankPayment PaymentMethodType = "BankPayment"
	BNPL        PaymentMethodType = "BNPL"
)

var paymentMethodTypeOrdinal = map[PaymentMethodType]int{
	Reference:   0,
	Credit:      1,
	Debit:       2,
	EBT:         3,
	Cash:        4,
	ACH:         5,
	Gift:        6,
	Recurring:   7,
	Other:       8,
	APM:         9,
	Ewic:        10,
	BankPayment: 11,
	BNPL:        12,
}

func (t PaymentMethodType) LongValue() int64 {
	return int64(math.Pow(2, float64(paymentMethodTypeOrdinal[t])))
}

func (t PaymentMethodType) GetStringValue() string {
	return string(t)
}

func (t PaymentMethodType) GetPaymentMethodType() PaymentMethodType {
	return t
}

func GetSet(value int64) map[PaymentMethodType]bool {
	flags := make(map[PaymentMethodType]bool)
	for _, flag := range []PaymentMethodType{
		Reference,
		Credit,
		Debit,
		EBT,
		Cash,
		ACH,
		Gift,
		Recurring,
		Other,
		APM,
		Ewic,
		BankPayment,
		BNPL,
	} {
		flagValue := flag.LongValue()
		if (flagValue & value) == flagValue {
			flags[flag] = true
		}
	}
	return flags
}
