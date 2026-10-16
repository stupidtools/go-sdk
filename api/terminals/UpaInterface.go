package terminals

import (
	"errors"
	"strconv"
	"unicode/utf8"

	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/currencytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/messaging"
	"github.com/globalpayments/go-sdk/api/terminals/terminalutilities"
	upabuilders "github.com/globalpayments/go-sdk/api/terminals/upa/builders"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upamessageid"
	"github.com/globalpayments/go-sdk/api/terminals/upa/responses"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaInterface struct {
	controller *UpaController
}

func NewUpaInterface(controller *UpaController) *UpaInterface {
	return &UpaInterface{
		controller: controller,
	}
}

func (ui *UpaInterface) generateLineItemMessage(leftText, rightText string) (*terminalutilities.DeviceMessage, error) {
	param := utils.NewJsonDoc()

	if leftText != "" && utf8.RuneCountInString(leftText) <= 20 {
		if utf8.RuneCountInString(leftText) > 20 {
			return nil, errors.New("Left-side text has a 20 character limit.")
		}
		param.Set("lineItemLeft", leftText, true)
	} else {
		return nil, errors.New("Left-side text is required.")
	}

	if rightText != "" && utf8.RuneCountInString(rightText) <= 10 {
		param.Set("lineItemRight", rightText, true)
	} else if utf8.RuneCountInString(rightText) > 10 {
		return nil, errors.New("Right-side text has a 10 character limit.")
	}

	body := utils.NewJsonDoc()
	body.SetJsonDoc("params", param)

	requestID := strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId())
	return terminalutilities.BuildMessage(upamessageid.LineItemDisplay, requestID, "", body)
}

func (ui *UpaInterface) AddLineItemBulk(lineItems [][2]string) (abstractions.IDeviceResponse, error) {
	var check abstractions.IDeviceResponse
	defer ui.controller.GetDeviceCommInterface().Disconnect()
	for _, li := range lineItems {
		message, err := ui.generateLineItemMessage(li[0], li[1])
		if err != nil {
			return nil, err
		}

		responseData, err := ui.controller.SendWithouDisconnect(message)
		if err != nil {
			return nil, err
		}

		responseObj, err := utils.ParseBytes(responseData)
		if err != nil {
			return nil, err
		}

		check = responses.NewUpaDeviceResponse(*responseObj, upamessageid.LineItemDisplay)

		if check.GetDeviceResponseCode() != "00" {
			return check, nil
		}
	}
	return check, nil
}

func (ui *UpaInterface) AddLineItem(leftText, rightText string) (abstractions.IDeviceResponse, error) {
	message, err := ui.generateLineItemMessage(leftText, rightText)
	if err != nil {
		return nil, err
	}

	responseData, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseData)
	if err != nil {
		return nil, err
	}

	return responses.NewUpaDeviceResponse(*responseObj, upamessageid.LineItemDisplay), nil
}

func (ui *UpaInterface) Cancel() error {
	return ui.CancelWithParam(1)
}

func (ui *UpaInterface) CancelWithParam(cancelParam int) error {
	param := utils.NewJsonDoc()
	param.SetInt("displayOption", &cancelParam)

	body := utils.NewJsonDoc()
	body.SetJsonDoc("params", param)

	message, err := terminalutilities.BuildMessage(upamessageid.CancelTransaction, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", body)

	if err != nil {
		return err
	}

	_, err2 := ui.controller.Send(message)
	return err2
}

func (ui *UpaInterface) CreditAuthWithoutAmount() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Auth, paymentmethodtype.Credit), nil
}

func (ui *UpaInterface) CreditAuth(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Auth, paymentmethodtype.Credit).WithAmount(amount), nil
}

func (ui *UpaInterface) CreditCapture(amount *decimal.Decimal) (*builders.TerminalManageBuilder, error) {
	return builders.NewTerminalManageBuilder(transactiontype.Capture, paymentmethodtype.Credit).WithAmount(amount), nil
}

func (ui *UpaInterface) CreditCaptureWithoutAmount() (*builders.TerminalManageBuilder, error) {
	return builders.NewTerminalManageBuilder(transactiontype.Capture, paymentmethodtype.Credit), nil
}

func (ui *UpaInterface) CreditRefund(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Refund, paymentmethodtype.Credit).WithAmount(amount), nil
}

func (ui *UpaInterface) CreditRefundWithoutAmount() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Refund, paymentmethodtype.Credit), nil
}

func (ui *UpaInterface) CreditSaleWithoutAmount() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Sale, paymentmethodtype.Credit), nil
}

func (ui *UpaInterface) CreditSale(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Sale, paymentmethodtype.Credit).WithAmount(amount), nil
}

func (ui *UpaInterface) CreditVerify() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Verify, paymentmethodtype.Credit), nil
}

func (ui *UpaInterface) CreditVoid() (builders.ITerminalManageBuilder, error) {
	return upabuilders.NewUpaTerminalManageBuilder(transactiontype.Void, paymentmethodtype.Credit), nil
}
func (ui *UpaInterface) DebitRefund(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Refund, paymentmethodtype.Debit).WithAmount(amount), nil
}

func (ui *UpaInterface) DebitRefundWithoutAmount() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Refund, paymentmethodtype.Debit), nil
}

func (ui *UpaInterface) DebitSaleWithoutAmount() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Sale, paymentmethodtype.Debit), nil
}

func (ui *UpaInterface) DebitSale(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Sale, paymentmethodtype.Debit).WithAmount(amount), nil
}

func (ui *UpaInterface) DebitVerify() (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Verify, paymentmethodtype.Debit), nil
}

func (ui *UpaInterface) DebitVoid() (builders.ITerminalManageBuilder, error) {
	return upabuilders.NewUpaTerminalManageBuilder(transactiontype.Void, paymentmethodtype.Debit), nil
}

func (ui *UpaInterface) EndOfDay() (abstractions.IEODResponse, error) {
	message, err := terminalutilities.BuildMessage(upamessageid.EODProcessing, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", nil)
	if err != nil {
		return nil, err
	}

	message.SetAwaitResponse(true)

	responseBytes, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseBytes)
	if err != nil {
		return nil, err
	}

	return responses.NewUpaEODResponse(*responseObj), nil
}

func (ui *UpaInterface) GetOpenTabDetails() (abstractions.IBatchReportResponse, error) {
	message, err := terminalutilities.BuildMessage(upamessageid.GetOpenTabDetails, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", nil)
	if err != nil {
		return nil, err
	}

	message.SetAwaitResponse(true)

	responseBytes, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseBytes)
	if err != nil {
		return nil, err
	}

	return responses.NewUpaReportResponse(responseObj), nil
}

func (ui *UpaInterface) GiftAddValue(amount *decimal.Decimal) (*builders.TerminalAuthBuilder, error) {
	return builders.NewTerminalAuthBuilder(transactiontype.Activate, paymentmethodtype.Gift).WithCurrency(currencytype.Currency).WithAmount(amount), nil
}

func (ui *UpaInterface) Ping() (abstractions.IDeviceResponse, error) {
	message, err := terminalutilities.BuildMessage(upamessageid.Ping, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseData)

	if err != nil {
		return nil, err
	}

	return responses.NewUpaDeviceResponse(*responseObj, upamessageid.Ping), nil
}

func (ui *UpaInterface) PromptForSignature() (abstractions.ISignatureResponse, error) {
	param := utils.NewJsonDoc()

	param.Set("prompt1", "Please Sign", true)

	body := utils.NewJsonDoc()
	body.SetJsonDoc("params", param)

	requestID := strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId())
	message, err := terminalutilities.BuildMessage(upamessageid.GetSignature, requestID, "", body)
	responseData, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseData)
	if err != nil {
		return nil, err
	}

	return responses.NewUpaDeviceResponse(*responseObj, upamessageid.GetSignature), nil
}

func (ui *UpaInterface) Reboot() (abstractions.IDeviceResponse, error) {
	message, err := terminalutilities.BuildMessage(upamessageid.Reboot, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", nil)
	if err != nil {
		return nil, err
	}
	responseData, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseData)

	if err != nil {
		return nil, err
	}

	return responses.NewUpaDeviceResponse(*responseObj, upamessageid.Reboot), nil
}

func (ui *UpaInterface) Reset() (abstractions.IDeviceResponse, error) {
	message, err := terminalutilities.BuildMessage(upamessageid.Reboot, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", nil)
	if err != nil {
		return nil, err
	}
	responseData, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(responseData)

	if err != nil {
		return nil, err
	}

	return responses.NewUpaDeviceResponse(*responseObj, upamessageid.Restart), nil
}

func (ui *UpaInterface) Reverse() (*upabuilders.UpaTerminalManageBuilder, error) {
	return upabuilders.NewUpaTerminalManageBuilder(transactiontype.Reversal, paymentmethodtype.Credit), nil
}

func (ui *UpaInterface) SendStoreAndForward() (abstractions.ISAFResponse, error) {
	message, err := terminalutilities.BuildMessage(upamessageid.SendSAF, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", nil)
	if err != nil {
		return nil, err
	}

	resp, err := ui.controller.Send(message)
	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(resp)
	if err != nil {
		return nil, err
	}
	return responses.NewUpaSafResponse(responseObj), nil
}

func (ui *UpaInterface) SafDelete(referenceNumber string, transactionNumber string) (abstractions.ISAFResponse, error) {
	body := utils.NewJsonDoc()
	transaction := utils.NewJsonDoc()

	if transactionNumber != "" {
		transaction.Set("tranNo", transactionNumber, true)
	}
	if referenceNumber != "" {
		transaction.Set("referenceNumber", referenceNumber, true)
	}
	body.SetJsonDoc("transaction", transaction)

	message, err := terminalutilities.BuildMessage(upamessageid.DeleteSAF, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", body)
	if err != nil {
		return nil, err
	}

	message.SetAwaitResponse(true)
	resp, err := ui.controller.Send(message)

	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(resp)
	if err != nil {
		return nil, err
	}
	return responses.NewUpaSafResponse(responseObj), nil
}

func (ui *UpaInterface) PreAuthDelete(referenceNumber string, amount *decimal.Decimal) (abstractions.IDeviceResponse, error) {
	body := utils.NewJsonDoc()
	transaction := utils.NewJsonDoc()

	if referenceNumber != "" {
		transaction.Set("referenceNumber", referenceNumber, true)
	}
	if amount != nil {
		transaction.Set("preAuthAmount", amount.StringFixed(2), true)
	}
	body.SetJsonDoc("transaction", transaction)

	message, err := terminalutilities.BuildMessage(upamessageid.DeletePreAuth, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", body)
	if err != nil {
		return nil, err
	}

	message.SetAwaitResponse(true)
	resp, err := ui.controller.Send(message)

	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(resp)
	if err != nil {
		return nil, err
	}
	return responses.NewUpaDeviceResponse(*responseObj, upamessageid.DeletePreAuth), nil
}

func (ui *UpaInterface) SafSummaryReport() (abstractions.ISAFResponse, error) {
	body := utils.NewJsonDoc()
	param := utils.NewJsonDoc()

	param.Set("reportOutput", "ReturnData", true)
	body.SetJsonDoc("params", param)

	message, err := terminalutilities.BuildMessage(upamessageid.GetSAFReport, strconv.Itoa(ui.controller.GetRequestIdProvider().GetRequestId()), "", body)

	if err != nil {
		return nil, err
	}

	message.SetAwaitResponse(true)
	resp, err := ui.controller.Send(message)

	if err != nil {
		return nil, err
	}

	responseObj, err := utils.ParseBytes(resp)

	if err != nil {
		return nil, err
	}
	return responses.NewUpaSafResponse(responseObj), nil
}

func (ui *UpaInterface) TipAdjust(amount *decimal.Decimal) (*builders.TerminalManageBuilder, error) {
	return upabuilders.NewUpaTerminalManageBuilder(transactiontype.Edit, paymentmethodtype.Credit).WithGratuity(amount), nil

}

func (ui *UpaInterface) CloseLane() (abstractions.IDeviceResponse, error) {
	return nil, exceptions.NewUnsupportedTransactionException("This transaction is unsupported on this device")
}

func (ui *UpaInterface) Initialize() (abstractions.IInitializeResponse, error) {
	return nil, exceptions.NewUnsupportedTransactionException("This transaction is unsupported on this device")
}

func (ui *UpaInterface) Dispose() {
	// unused in UPA
}

func (ui *UpaInterface) SetOnMessageSent(onMessageSent messaging.IMessageSentInterface) {
	ui.controller.SetOnMessageSent(onMessageSent)
}

func (ui *UpaInterface) DisableHostResponseBeep() (abstractions.IDeviceResponse, error) {
	return nil, exceptions.NewUnsupportedTransactionException("This transaction is unsupported on this device")
}

func (ui *UpaInterface) OpenLane() (abstractions.IDeviceResponse, error) {
	return nil, exceptions.NewUnsupportedTransactionException("This transaction is unsupported on this device")
}

func (ui *UpaInterface) StartCard(paymentMethodType paymentmethodtype.PaymentMethodType) (abstractions.IDeviceResponse, error) {
	return nil, exceptions.NewUnsupportedTransactionException("This transaction is unsupported on this device")
}
