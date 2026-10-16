package imappedconstant

import "github.com/globalpayments/go-sdk/api/entities/enums/target"

type IMappedConstant interface {
	GetValue(Target target.Target) string
}
