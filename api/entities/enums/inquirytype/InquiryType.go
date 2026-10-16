package inquirytype

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type InquiryType string

const (
	Foodstamp InquiryType = "FOODSTAMP"
	Cash      InquiryType = "CASH"
)

func (i InquiryType) GetBytes() []byte {
	return []byte(i)
}

func (i InquiryType) GetValue() string {
	return string(i)
}

func (i InquiryType) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{Foodstamp, Cash}
}
