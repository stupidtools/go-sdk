package billpresentment

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type BillPresentment string

const (
	FULL BillPresentment = "FULL"
)

func (bp BillPresentment) GetBytes() []byte {
	return []byte(bp)
}

func (bp BillPresentment) GetValue() string {
	return string(bp)
}

// This method is added based on the provided IStringConstant interface in Go
func (bp BillPresentment) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{FULL}
}
