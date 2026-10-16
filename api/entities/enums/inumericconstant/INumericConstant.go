package inumericconstant

type INumericConstant interface {
	GetValue() int
	NumericConstants() []INumericConstant
}
