package controlcodes

import "github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"

type ControlCodes byte

const (
	STX   ControlCodes = 0x02 // Denotes the beginning of a message frame
	ETX   ControlCodes = 0x03 // Denotes the ending of a message frame
	EOT   ControlCodes = 0x04 // Indicates communication session terminated
	ENQ   ControlCodes = 0x05 // Begin Session sent from the host to the POS
	ACK   ControlCodes = 0x06 // Acknowledge of message received
	NAK   ControlCodes = 0x15 // Indicates invalid message received
	FS    ControlCodes = 0x1C // Field separator
	GS    ControlCodes = 0x1D // Message ID follows (for non-PIN entry prompts)
	RS    ControlCodes = 0x1E // Message ID follows (for PIN entry prompts)
	US    ControlCodes = 0x1F
	COMMA ControlCodes = 0x2C
	COLON ControlCodes = 0x3A
	PTGS  ControlCodes = 0x7C
	LF    ControlCodes = 0x0A
)

var controlCodesToString = map[ControlCodes]string{
	STX:   "STX",
	ETX:   "ETX",
	EOT:   "EOT",
	ENQ:   "ENQ",
	ACK:   "ACK",
	NAK:   "NAK",
	FS:    "FS",
	GS:    "GS",
	RS:    "RS",
	US:    "US",
	COMMA: "COMMA",
	COLON: "COLON",
	PTGS:  "PTGS",
	LF:    "LF",
}

func (c ControlCodes) ByteConstants() []ibyteconstant.IByteConstant {
	var output []ibyteconstant.IByteConstant
	output = make([]ibyteconstant.IByteConstant, 0)
	for k := range controlCodesToString {
		output = append(output, k)
	}
	return output
}

func (c ControlCodes) String() string {
	if str, ok := controlCodesToString[c]; ok {
		return str
	}
	return ""
}

func (c ControlCodes) GetByte() byte {
	return byte(c)
}
