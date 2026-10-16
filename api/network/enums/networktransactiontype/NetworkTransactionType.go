package networktransactiontype

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type NetworkTransactionType string

const (
	KeepAlive   NetworkTransactionType = "NT"
	Transaction NetworkTransactionType = "EH"
)

func (n NetworkTransactionType) GetBytes() []byte {
	return []byte(n)
}

func (n NetworkTransactionType) GetValue() string {
	return string(n)
}

func (n NetworkTransactionType) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{KeepAlive, Transaction}
}
