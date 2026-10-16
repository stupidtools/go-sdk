package abstractions

type IPinProtected interface {
	GetPinBlock() string
	SetPinBlock(pinBlock string)
}
