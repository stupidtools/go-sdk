package ibyteconstant

type IByteConstant interface {
	GetByte() byte
	ByteConstants() []IByteConstant
}
