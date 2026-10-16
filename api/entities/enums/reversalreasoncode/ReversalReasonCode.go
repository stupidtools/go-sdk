package reversalreasoncode

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type ReversalReasonCode string

const (
	CustomerCancellation ReversalReasonCode = "CUSTOMERCANCELLATION"
	TerminalError        ReversalReasonCode = "TERMINALERROR"
	Timeout              ReversalReasonCode = "TIMEOUT"
	ChipCardDecline      ReversalReasonCode = "CHIPCARDDECLINE"
	MacFailure           ReversalReasonCode = "MACFAILURE"
)

func (rrc ReversalReasonCode) GetBytes() []byte {
	return []byte(rrc)
}

func (rrc ReversalReasonCode) GetValue() string {
	return string(rrc)
}

func (rrc ReversalReasonCode) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		CustomerCancellation,
		TerminalError,
		Timeout,
		ChipCardDecline,
		MacFailure,
	}
}
