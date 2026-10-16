package protocoltype

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type ProtocolType byte

const (
	NotSpecified ProtocolType = 0x00
	TCP_IP       ProtocolType = 0x01
	UDP_IP       ProtocolType = 0x02
	X25          ProtocolType = 0x03
	SNA          ProtocolType = 0x04
	Async        ProtocolType = 0x05
	// For asynchronous protocol two options exist, without link level and with link level support
	// (protocol types 5 and 7 respectively). Use protocol type value of 7 when type message value is
	// 35 or 37. Use protocol type value of 5 when type message value is 01. For other type message
	// values consult with your Heartland representative.
	Bisync_3270 ProtocolType = 0x06
)

func (p ProtocolType) GetByte() byte {
	return byte(p)
}

func (p ProtocolType) ByteConstants() []ibyteconstant.IByteConstant {
	return []ibyteconstant.IByteConstant{
		NotSpecified,
		TCP_IP,
		UDP_IP,
		X25,
		SNA,
		Async,
		Bisync_3270,
	}
}
