package iflag

type IFlag interface {
	LongValue() int64
	GetStringValue() string
}
