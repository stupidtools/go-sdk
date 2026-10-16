package networkresponsecode

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type NetworkResponseCode byte

const (
	Success                     NetworkResponseCode = 0x00
	FailedConnection            NetworkResponseCode = 0x01
	Timeout                     NetworkResponseCode = 0x02
	FormatError_Originator      NetworkResponseCode = 0x03
	StoreAndForward             NetworkResponseCode = 0x04
	UnsupportedTransaction      NetworkResponseCode = 0x05
	UnsupportedServiceProvider  NetworkResponseCode = 0x06
	FormatError_ServiceProvider NetworkResponseCode = 0x07
)

func (n NetworkResponseCode) GetByte() byte {
	return byte(n)
}

func (n NetworkResponseCode) ByteConstants() []ibyteconstant.IByteConstant {
	return []ibyteconstant.IByteConstant{
		Success,
		FailedConnection,
		Timeout,
		FormatError_Originator,
		StoreAndForward,
		UnsupportedTransaction,
		UnsupportedServiceProvider,
		FormatError_ServiceProvider,
	}
}
