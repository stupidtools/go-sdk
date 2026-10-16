package responses

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/applicationcryptogramtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/cardtype"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaTransactionResponse struct {
	terminalresponse.TerminalResponse
}

func NewUpaTransactionResponse(responseData *utils.JsonDoc) *UpaTransactionResponse {
	res := &UpaTransactionResponse{}
	res.TerminalResponse = terminalresponse.TerminalResponse{}
	res.DeviceResponseCode = "99" // Default error code
	res.DeviceResponseText = "Unknown error"
	cmdResult := responseData.Get("cmdResult")
	if cmdResult != nil {
		res.Status = cmdResult.GetString("result")
		if res.Status == "Success" {
			res.DeviceResponseCode = "00"
		} else {
			res.DeviceResponseCode = cmdResult.GetString("errorCode")
		}
		res.DeviceResponseText = cmdResult.GetString("errorMessage")
	}
	res.TransactionType = responseData.GetString("response")
	res.EcrId = responseData.GetString("EcrId")
	data := responseData.Get("data")
	if data != nil {
		host := data.Get("host")
		if host != nil {
			res.AmountDue = host.GetDecimal("balanceDue")
			res.ApprovalCode = host.GetString("approvalCode")
			res.AvsResponseCode = host.GetString("AvsResultCode")
			res.AvsResponseText = host.GetString("AvsResultText")
			res.BalanceAmount = host.GetDecimal("availableBalance")
			res.CardBrandTransactionId = host.GetString("cardBrandTransId")
			res.ResponseCode = host.GetString("responseCode")
			res.ResponseText = host.GetString("responseText")
			res.MerchantFee = host.GetDecimal("surcharge")
			res.TerminalRefNumber = host.GetString("tranNo")
			res.Token = host.GetString("tokenValue")
			res.TransactionId = host.GetIntOrStringAsString("referenceNumber")
			res.TransactionAmount = host.GetDecimal("totalAmount")
		}

		payment := data.Get("payment")
		if payment != nil {
			res.CardHolderName = payment.GetString("cardHolderName")
			cardTypeStr := payment.GetString("cardType")
			if cardTypeStr != "" {
				switch cardTypeStr {
				case "VISA":
					res.CardType = cardtype.VISA
				case "MASTERCARD":
					res.CardType = cardtype.MC
				case "DISCOVER":
					res.CardType = cardtype.DISC
				case "AMERICAN EXPRESS":
					res.CardType = cardtype.AMEX
				}
				res.EntryMethod = payment.GetString("cardAcquisition")
				res.MaskedCardNumber = payment.GetString("maskedPan")
				res.PaymentType = payment.GetString("cardGroup")
			}
		}

		transaction := data.Get("transaction")
		if transaction != nil {
			if transaction.GetDecimal("totalAmount") != nil {
				res.TransactionAmount = transaction.GetDecimal("totalAmount")
			}

			if transaction.GetDecimal("tipAmount") != nil {
				res.TipAmount = transaction.GetDecimal("tipAmount")
			}
		}

		emv := data.Get("emv")
		if emv != nil {
			res.ApplicationCryptogram = emv.GetString("9F26")
			cryptogramTypeStr := emv.GetString("9F27")
			if cryptogramTypeStr != "" {
				switch cryptogramTypeStr {
				case "0":
					res.ApplicationCryptogramType = applicationcryptogramtype.AAC
				case "40":
					res.ApplicationCryptogramType = applicationcryptogramtype.TC
				case "80":
					res.ApplicationCryptogramType = applicationcryptogramtype.ARQC
				}
			}
			res.TransactionStatusInfo = emv.GetString("9B")
			res.ApplicationId = emv.GetString("9F06")
			res.ApplicationLabel = emv.GetString("50")
			res.ApplicationPreferredName = emv.GetString("9F12")
		}

		pan := data.Get("PAN")
		if pan != nil {
			res.UnmaskedCardNumber = pan.GetString("clearPAN")
		}
		fallback := data.GetString("fallback")
		if fallback != "" {
			res.Fallback = fallback
		}
		serviceCode := data.GetString("serviceCode")
		if serviceCode != "" {
			res.ServiceCode = serviceCode
		}
		expirationDate := data.GetString("expiryDate")
		if expirationDate != "" {
			res.ExpirationDate = expirationDate
		}

	}

	return res
}
