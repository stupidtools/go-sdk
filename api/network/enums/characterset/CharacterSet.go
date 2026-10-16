package characterset

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type CharacterSet byte

const (
	ASCII  CharacterSet = 0x01
	EBCDIC CharacterSet = 0x02
)

func (cs CharacterSet) GetByte() byte {
	return byte(cs)
}

func (cs CharacterSet) ByteConstants() []ibyteconstant.IByteConstant {
	return []ibyteconstant.IByteConstant{ASCII, EBCDIC}
}
