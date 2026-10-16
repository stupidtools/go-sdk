package authorizercode

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type AuthorizerCode string

const (
	InterchangeAuthorized AuthorizerCode = " "
	HostAuthorized        AuthorizerCode = "B"
	TerminalAuthorized    AuthorizerCode = "T"
	VoiceAuthorized       AuthorizerCode = "V"
	PassThrough           AuthorizerCode = "P"
	NegativeFile          AuthorizerCode = "N"
	LocalNegativeFile     AuthorizerCode = "L"
	AuthTable             AuthorizerCode = "A"
	ReservedAuthorized    AuthorizerCode = "D"
	Synchrony             AuthorizerCode = "S"
	ChaseNet              AuthorizerCode = "C"
	HeartlandGift         AuthorizerCode = "G"
)

func (ac AuthorizerCode) GetBytes() []byte {
	return []byte(ac)
}

func (ac AuthorizerCode) GetValue() string {
	return string(ac)
}

func (ac AuthorizerCode) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		InterchangeAuthorized,
		HostAuthorized,
		TerminalAuthorized,
		VoiceAuthorized,
		PassThrough,
		NegativeFile,
		LocalNegativeFile,
		AuthTable,
		ReservedAuthorized,
		Synchrony,
		ChaseNet,
		HeartlandGift,
	}
}
