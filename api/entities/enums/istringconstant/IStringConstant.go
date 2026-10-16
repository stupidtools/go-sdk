package istringconstant

type IStringConstant interface {
	GetBytes() []byte
	GetValue() string
	StringConstants() []IStringConstant
}
