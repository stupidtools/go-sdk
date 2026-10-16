package connectiontype

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type ConnectionType byte

const (
	NotSpecified          ConnectionType = 0x00
	Service800            ConnectionType = 0x01
	LeasedLine            ConnectionType = 0x02
	Connect950            ConnectionType = 0x03
	DirectDial            ConnectionType = 0x04
	VSAT                  ConnectionType = 0x05
	ISDN                  ConnectionType = 0x06
	eCommerce             ConnectionType = 0x07
	FrameRelay            ConnectionType = 0x08
	FixedWireless         ConnectionType = 0x09
	MobileWireless        ConnectionType = 0x0A
	BlackBox_Presidia     ConnectionType = 0x19
	TNS_Internet          ConnectionType = 0x30
	Datawire_Internet     ConnectionType = 0x31
	Echosat               ConnectionType = 0x32
	Accel                 ConnectionType = 0x33
	MobileWireless_2      ConnectionType = 0x34
	Tech_Pilot            ConnectionType = 0x35
	Hughes_VSAT_Broadband ConnectionType = 0x37
	Hughes_DSL_Broadband  ConnectionType = 0x37
	EchoSat_Smartlink     ConnectionType = 0x41
	MPLS                  ConnectionType = 0x42
	SSL_Gateway           ConnectionType = 0x43
	Native_SSL            ConnectionType = 0x44
	Sagenet               ConnectionType = 0x45
	Cybera                ConnectionType = 0x46
)

func (c ConnectionType) GetByte() byte {
	return byte(c)
}

func (c ConnectionType) ByteConstants() []ibyteconstant.IByteConstant {
	return []ibyteconstant.IByteConstant{
		NotSpecified,
		Service800,
		LeasedLine,
		Connect950,
		DirectDial,
		VSAT,
		ISDN,
		eCommerce,
		FrameRelay,
		FixedWireless,
		MobileWireless,
		BlackBox_Presidia,
		TNS_Internet,
		Datawire_Internet,
		Echosat,
		Accel,
		MobileWireless_2,
		Tech_Pilot,
		Hughes_VSAT_Broadband,
		Hughes_DSL_Broadband,
		EchoSat_Smartlink,
		MPLS,
		SSL_Gateway,
		Native_SSL,
		Sagenet,
		Cybera,
	}
}
