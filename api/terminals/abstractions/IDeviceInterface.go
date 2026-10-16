package abstractions

import (
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/messaging"
)

type IDeviceInterface interface {
	SetOnMessageSent(onMessageSent messaging.IMessageSentInterface)
	CreditRefund(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error)
	CreditSale(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error)
	CreditAuth(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error)
	CreditAuthWithoutAmount() (*builders.TerminalAuthBuilder, error)
	CreditCapture(amount *decimal.Decimal) (*builders.TerminalManageBuilder, error)
	CreditCaptureWithoutAmount() (*builders.TerminalManageBuilder, error)
	CreditRefundWithoutAmount() (*builders.TerminalAuthBuilder, error)
	CreditSaleWithoutAmount() (*builders.TerminalAuthBuilder, error)
	CreditVoid() (builders.ITerminalManageBuilder, error)
	DebitRefund(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error)
	DebitSale(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error)
	DebitRefundWithoutAmount() (*builders.TerminalAuthBuilder, error)
	DebitSaleWithoutAmount() (*builders.TerminalAuthBuilder, error)
	DebitVoid() (builders.ITerminalManageBuilder, error)
	EndOfDay() (IEODResponse, error)
	GetOpenTabDetails() (IBatchReportResponse, error)
	GiftAddValue(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error)
	Cancel() error
	Reboot() (IDeviceResponse, error)
	SafDelete(referenceNumber string, transactionNumber string) (ISAFResponse, error)
	SafSummaryReport() (ISAFResponse, error)
	SendStoreAndForward() (ISAFResponse, error)
	TipAdjust(amount *decimal.Decimal) (*builders.TerminalManageBuilder, error)
	Ping() (IDeviceResponse, error)
	PromptForSignature() (ISignatureResponse, error)
	PreAuthDelete(referenceNumber string, amount *decimal.Decimal) (IDeviceResponse, error)
	AddLineItem(leftText string, rightText string) (IDeviceResponse, error)
	AddLineItemBulk(lineItems [][2]string) (IDeviceResponse, error)
	CancelWithParam(val int) error
}
