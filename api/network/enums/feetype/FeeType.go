package feetype

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type FeeType string

const (
	TransactionFee FeeType = "00"
	Surcharge      FeeType = "22"
)

func (ft FeeType) GetValue() string {
	return string(ft)
}

func (ft FeeType) GetBytes() []byte {
	return []byte(ft)
}

func (ft FeeType) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{TransactionFee, Surcharge}
}
