package messagetype

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type MessageType byte

const (
	NoMessage                           MessageType = 0x00
	Heartland_Z01                       MessageType = 0x01
	Auth_Quick_Credit                   MessageType = 0x02
	Remote_Entry                        MessageType = 0x03
	FDMS_PassThrough                    MessageType = 0x04
	TCOP_AuthType_PassThrough           MessageType = 0x05
	TCOP_ApplicationType_PassThrough    MessageType = 0x06
	TCOP_Legacy_Auth_PassThrough        MessageType = 0x07
	TCOP_Legacy_QuickCredit_PassThrough MessageType = 0x08
	Discover_ISO_8583_PassThrough       MessageType = 0x09
	Combo_KSM_NPC                       MessageType = 0x0A
	Geobridge_3DES_NWS                  MessageType = 0x0B
	HSM_Fixed                           MessageType = 0x10
	KSM_Delimited                       MessageType = 0x11
	WEX_Processing                      MessageType = 0x12
	Site_Configured_Layout              MessageType = 0x1F
	DEPS_8583                           MessageType = 0x21
	Heartland_8583                      MessageType = 0x22
	Heartland_POS_8583                  MessageType = 0x23
	BIC_ISO                             MessageType = 0x24
	Heartland_NTS                       MessageType = 0x25
	ATT_Format_III                      MessageType = 0x26
	ATT_Format_IV                       MessageType = 0x27
	CPNI                                MessageType = 0x28
	Heartland_Prepaid_XML               MessageType = 0x29
	Heartland_Prepaid_Host_to_Host      MessageType = 0x2A
	NPC_Processing                      MessageType = 0x30
	IFCS_Processing                     MessageType = 0x31
	FedChex_Processing                  MessageType = 0x32
	JCP_Private_Label                   MessageType = 0x33
	SOPUS_8583                          MessageType = 0x34
	JSON                                MessageType = 0x35
)

func (m MessageType) GetByte() byte {
	return byte(m)
}

func (m MessageType) ByteConstants() []ibyteconstant.IByteConstant {
	return []ibyteconstant.IByteConstant{
		NoMessage,
		Heartland_Z01,
		Auth_Quick_Credit,
		Remote_Entry,
		FDMS_PassThrough,
		TCOP_AuthType_PassThrough,
		TCOP_ApplicationType_PassThrough,
		TCOP_Legacy_Auth_PassThrough,
		TCOP_Legacy_QuickCredit_PassThrough,
		Discover_ISO_8583_PassThrough,
		Combo_KSM_NPC,
		Geobridge_3DES_NWS,
		HSM_Fixed,
		KSM_Delimited,
		WEX_Processing,
		Site_Configured_Layout,
		DEPS_8583,
		Heartland_8583,
		Heartland_POS_8583,
		BIC_ISO,
		Heartland_NTS,
		ATT_Format_III,
		ATT_Format_IV,
		CPNI,
		Heartland_Prepaid_XML,
		Heartland_Prepaid_Host_to_Host,
		NPC_Processing,
		IFCS_Processing,
		FedChex_Processing,
		JCP_Private_Label,
		SOPUS_8583,
		JSON,
	}
}
