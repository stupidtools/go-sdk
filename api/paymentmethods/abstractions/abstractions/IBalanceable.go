package abstractions

import (
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
)

type IBalanceable interface {
	IBalanceInquiryWithType(inquiry inquirytype.InquiryType) abstractions.IAuthorizationBuilder
}
