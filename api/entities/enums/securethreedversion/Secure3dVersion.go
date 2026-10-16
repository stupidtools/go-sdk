package securethreedversion

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type Secure3dVersion string

const (
	NONE Secure3dVersion = "None"
	ONE  Secure3dVersion = "One"
	TWO  Secure3dVersion = "Two"
	ANY  Secure3dVersion = "Any"
)

var intMap map[Secure3dVersion]int = map[Secure3dVersion]int{
	NONE: 0,
	ONE:  1,
	TWO:  2,
	ANY:  1,
}

func (v Secure3dVersion) GetBytes() []byte {
	return []byte(v)
}

func (v Secure3dVersion) GetValue() string {
	return string(v)
}

func (v Secure3dVersion) GetInt() int {
	return intMap[v]
}

func (v Secure3dVersion) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{NONE, ONE, TWO, ANY}
}
