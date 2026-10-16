package networkprocessingflag

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type NetworkProcessingFlag byte

const (
	NonPersistentConnection NetworkProcessingFlag = 0x00
	PersistentConnection    NetworkProcessingFlag = 0x01
)

func (npf NetworkProcessingFlag) GetByte() byte {
	return byte(npf)
}

func (npf NetworkProcessingFlag) ByteConstants() []ibyteconstant.IByteConstant {
	return []ibyteconstant.IByteConstant{
		NonPersistentConnection,
		PersistentConnection,
	}
}
