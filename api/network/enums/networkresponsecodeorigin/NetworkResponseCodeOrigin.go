package networkresponsecodeorigin

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type NetworkResponseCodeOrigin int

const (
	Default           NetworkResponseCodeOrigin = 0x00
	FrontEndProcess   NetworkResponseCodeOrigin = 0x01
	BackEndProcess    NetworkResponseCodeOrigin = 0x02
	InternalProcess   NetworkResponseCodeOrigin = 0x03
	AuthorizationHost NetworkResponseCodeOrigin = 0x04
)

var byteConstants = []ibyteconstant.IByteConstant{
	Default,
	FrontEndProcess,
	BackEndProcess,
	InternalProcess,
	AuthorizationHost,
}

func (n NetworkResponseCodeOrigin) GetByte() byte {
	return byte(n)
}

func (n NetworkResponseCodeOrigin) ByteConstants() []ibyteconstant.IByteConstant {
	return byteConstants
}
