package aliasaction

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type AliasAction string

const (
	Add    AliasAction = "ADD"
	Create AliasAction = "CREATE"
	Delete AliasAction = "DELETE"
)

var allStringConstants = []istringconstant.IStringConstant{Add, Create, Delete}

func (a AliasAction) GetBytes() []byte {
	return []byte(a)
}

func (a AliasAction) GetValue() string {
	return string(a)
}

func (a AliasAction) StringConstants() []istringconstant.IStringConstant {
	return allStringConstants
}
