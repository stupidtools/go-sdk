package entities

type DebitMac struct {
	TransactionCode           string
	TransmissionNumber        string
	BankResponseCode          string
	MacKey                    string
	PinKey                    string
	FieldKey                  string
	TraceNumber               string
	MessageAuthenticationCode string
}

func NewDebitMac() *DebitMac {
	return &DebitMac{}
}

func (d *DebitMac) GetTransactionCode() string {
	return d.TransactionCode
}

func (d *DebitMac) SetTransactionCode(transactionCode string) {
	d.TransactionCode = transactionCode
}

func (d *DebitMac) GetTransmissionNumber() string {
	return d.TransmissionNumber
}

func (d *DebitMac) SetTransmissionNumber(transmissionNumber string) {
	d.TransmissionNumber = transmissionNumber
}

func (d *DebitMac) GetBankResponseCode() string {
	return d.BankResponseCode
}

func (d *DebitMac) SetBankResponseCode(bankResponseCode string) {
	d.BankResponseCode = bankResponseCode
}

func (d *DebitMac) GetMacKey() string {
	return d.MacKey
}

func (d *DebitMac) SetMacKey(macKey string) {
	d.MacKey = macKey
}

func (d *DebitMac) GetPinKey() string {
	return d.PinKey
}

func (d *DebitMac) SetPinKey(pinKey string) {
	d.PinKey = pinKey
}

func (d *DebitMac) GetFieldKey() string {
	return d.FieldKey
}

func (d *DebitMac) SetFieldKey(fieldKey string) {
	d.FieldKey = fieldKey
}

func (d *DebitMac) GetTraceNumber() string {
	return d.TraceNumber
}

func (d *DebitMac) SetTraceNumber(traceNumber string) {
	d.TraceNumber = traceNumber
}

func (d *DebitMac) GetMessageAuthenticationCode() string {
	return d.MessageAuthenticationCode
}

func (d *DebitMac) SetMessageAuthenticationCode(messageAuthenticationCode string) {
	d.MessageAuthenticationCode = messageAuthenticationCode
}
