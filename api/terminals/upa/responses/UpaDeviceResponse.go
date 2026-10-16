package responses

import (
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upamessageid"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaDeviceResponse struct {
	terminalresponse.TerminalResponse
	MessageId upamessageid.UpaMessageId
}

func NewUpaDeviceResponse(responseObj utils.JsonDoc, messageId upamessageid.UpaMessageId) *UpaDeviceResponse {
	deviceResponse := &UpaDeviceResponse{
		MessageId:        messageId,
		TerminalResponse: terminalresponse.TerminalResponse{},
	}
	data := responseObj.Get("data")
	if data != nil {
		deviceResponse.SignatureData = []byte(data.GetString("signatureData"))
		cmdResult := data.Get("cmdResult")
		if cmdResult != nil {
			deviceResponse.Status = cmdResult.GetString("result")
			if deviceResponse.Status == "Success" {
				deviceResponse.DeviceResponseCode = "00"
			} else {
				deviceResponse.DeviceResponseCode = cmdResult.GetString("errorCode")
			}
			deviceResponse.DeviceResponseText = cmdResult.GetString("errorMessage")
		}
	}
	return deviceResponse
}
