package responses

import (
	"encoding/json"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upasaftype"
	"github.com/globalpayments/go-sdk/api/utils"
	"github.com/shopspring/decimal"
)

type UpaSafResponse struct {
	Approved           []abstractions.ISummaryResponse
	Command            string
	Declined           []abstractions.ISummaryResponse
	DeviceResponseCode string
	DeviceResponseText string
	Pending            []abstractions.ISummaryResponse
	Status             string
	SafTotalAmount     *decimal.Decimal
	SafTotalCount      *int
	Version            string
	TransactionType    string
	MultipleMessage    *int
	SafType            string
	Transactions       []transactionsummary.TransactionSummary
	TransactionTime    string
	HostTimeout        *int
}

func NewUpaSafResponse(responseObj *utils.JsonDoc) *UpaSafResponse {
	res := &UpaSafResponse{}
	responseData := responseObj.Get("data")
	res.SafTotalAmount = &decimal.Zero
	sc := 0
	res.SafTotalCount = &sc
	res.Approved = make([]abstractions.ISummaryResponse, 0)
	res.Declined = make([]abstractions.ISummaryResponse, 0)
	res.Pending = make([]abstractions.ISummaryResponse, 0)
	res.Transactions = make([]transactionsummary.TransactionSummary, 0)
	if responseData != nil {
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

		innerData := responseData.Get("data")
		if innerData != nil {
			if res.TransactionType == "DeleteSAF" {
				tr := parseDelSaf(innerData)
				res.Transactions = append(res.Transactions, *tr)
			}
			res.MultipleMessage = innerData.GetInt("multipleMessage")
			safDetails := innerData.GetArray("SafDetails")
			if safDetails != nil {
				for _, safDetail := range safDetails {

					safType := mapSafType(safDetail.GetString("SafType"))
					stc := *res.SafTotalCount
					stc += *safDetail.GetInt("SafCount")
					res.SafTotalCount = &stc
					res.SafTotalAmount.Add(*safDetail.GetDecimal("SafTotal"))
					summaryResponse := UpaSummaryResponse{}
					summaryResponse.Amount = safDetail.GetDecimal("SafTotal")
					summaryResponse.Count = safDetail.GetInt("SafCount")
					summaryResponse.Transactions = make([]transactionsummary.TransactionSummary, 0)
					if safDetail.GetArray("SafRecords") != nil {
						for _, safRecord := range safDetail.GetArray("SafRecords") {
							transactionSummary := transactionsummary.TransactionSummary{}
							transactionSummary.SafTotal = safRecord.GetDecimal("totalAmount")
							transactionSummary.AuthorizedAmount = safRecord.GetDecimal("authorizedAmount")
							transactionSummary.TransactionId = safRecord.GetString("tranNo")

							transactionSummary.TransactionDate = safRecord.GetString("transactionTime")

							transactionSummary.TransactionType = safRecord.GetString("transactionType")
							transactionSummary.MaskedCardNumber = safRecord.GetString("maskedPan")
							transactionSummary.CardType = safRecord.GetString("cardType")
							transactionSummary.EntryMode = safRecord.GetString("cardAcquisition")
							transactionSummary.GatewayResponseCode = safRecord.GetString("responseCode")
							transactionSummary.GatewayResponseMessage = safRecord.GetString("responseText")
							transactionSummary.ReferenceNumber = safRecord.GetString("referenceNumber")
							transactionSummary.SafReferenceNumber = safRecord.GetString("safReferenceNumber")
							res.HostTimeout = safRecord.GetInt("hostTimeout")
							if safRecord.GetDecimal("baseAmount") != nil {
								transactionSummary.BaseAmount = safRecord.GetDecimal("baseAmount")
							}
							if safRecord.GetDecimal("taxAmount") != nil {
								transactionSummary.TaxAmount = safRecord.GetDecimal("taxAmount")
							}
							if safRecord.GetDecimal("tipAmount") != nil {
								transactionSummary.GratuityAmount = safRecord.GetDecimal("tipAmount")
							}
							if safRecord.GetDecimal("requestAmount") != nil {
								transactionSummary.Amount = safRecord.GetDecimal("requestAmount")
							}
							transactionSummary.InvoiceNumber = safRecord.GetString("invoiceNbr")
							transactionSummary.ClerkId = safRecord.GetString("clerkId")

							res.Transactions = append(res.Transactions, transactionSummary)
							summaryResponse.Transactions = append(summaryResponse.Transactions, transactionSummary)

						}
					}
					if safType == upasaftype.Approved.GetValue() {
						res.Approved = append(res.Approved, summaryResponse)
					} else if safType == upasaftype.Pending.GetValue() {
						res.Pending = append(res.Pending, summaryResponse)
					} else if safType == upasaftype.Failed.GetValue() {
						res.Declined = append(res.Declined, summaryResponse)
					}
				}
			}
		}
	}

	return res
}

// getters
func (u *UpaSafResponse) GetApproved() []abstractions.ISummaryResponse {
	return u.Approved
}

func (u *UpaSafResponse) GetCommand() string {
	return u.Command
}

func (u *UpaSafResponse) GetDeclined() []abstractions.ISummaryResponse {
	return u.Declined
}

func (u *UpaSafResponse) GetDeviceResponseCode() string {
	return u.DeviceResponseCode
}

func (u *UpaSafResponse) GetDeviceResponseText() string {
	return u.DeviceResponseText
}

func (u *UpaSafResponse) GetPending() []abstractions.ISummaryResponse {
	return u.Pending
}

func (u *UpaSafResponse) GetStatus() string {
	return u.Status
}

func (u *UpaSafResponse) GetSafTotalAmount() *decimal.Decimal {
	return u.SafTotalAmount
}

func (u *UpaSafResponse) GetSafTotalCount() *int {
	return u.SafTotalCount
}

func (u *UpaSafResponse) GetVersion() string {
	return u.Version
}

func (u *UpaSafResponse) GetTransactionType() string {
	return u.TransactionType
}

func (u *UpaSafResponse) GetMultipleMessage() *int {
	return u.MultipleMessage
}

func (u *UpaSafResponse) GetSafType() string {
	return u.SafType
}

func (u *UpaSafResponse) GetTransactions() []transactionsummary.TransactionSummary {
	return u.Transactions
}

func (u *UpaSafResponse) GetTransactionTime() string {
	return u.TransactionTime
}

func (u *UpaSafResponse) GetHostTimeout() *int {
	return u.HostTimeout
}

// setters
func (u *UpaSafResponse) SetApproved(approved []abstractions.ISummaryResponse) {
	u.Approved = approved
}

func (u *UpaSafResponse) SetCommand(command string) {
	u.Command = command
}

func (u *UpaSafResponse) SetDeclined(declined []abstractions.ISummaryResponse) {
	u.Declined = declined
}

func (u *UpaSafResponse) SetDeviceResponseCode(deviceResponseCode string) {
	u.DeviceResponseCode = deviceResponseCode
}

func (u *UpaSafResponse) SetDeviceResponseText(deviceResponseText string) {
	u.DeviceResponseText = deviceResponseText
}

func (u *UpaSafResponse) SetPending(pending []abstractions.ISummaryResponse) {
	u.Pending = pending
}

func (u *UpaSafResponse) SetStatus(status string) {
	u.Status = status
}

func (u *UpaSafResponse) SetSafTotalAmount(totalAmount *decimal.Decimal) {
	u.SafTotalAmount = totalAmount
}

func (u *UpaSafResponse) SetSafTotalCount(totalCount *int) {
	u.SafTotalCount = totalCount
}

func (u *UpaSafResponse) SetVersion(version string) {
	u.Version = version
}

func (u *UpaSafResponse) SetTransactionType(transactionType string) {
	u.TransactionType = transactionType
}

func (u *UpaSafResponse) SetMultipleMessage(multipleMessage *int) {
	u.MultipleMessage = multipleMessage
}

func (u *UpaSafResponse) SetSafType(safType string) {
	u.SafType = safType
}

func (u *UpaSafResponse) SetTransactions(transactions []transactionsummary.TransactionSummary) {
	u.Transactions = transactions
}

func (u *UpaSafResponse) SetTransactionTime(transactionTime string) {
	u.TransactionTime = transactionTime
}

func (u *UpaSafResponse) SetHostTimeout(hostTimeout *int) {
	u.HostTimeout = hostTimeout
}

func (res *UpaSafResponse) ToString() string {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

func mapSafType(safDetails string) string {
	switch safDetails {
	case upasaftype.Approved.GetValue():
		return upasaftype.Approved.GetValue()
	case upasaftype.Pending.GetValue():
		return upasaftype.Pending.GetValue()
	case upasaftype.Failed.GetValue():
		return upasaftype.Failed.GetValue()
	default:
		return upasaftype.Failed.GetValue()
	}
}

func parseDelSaf(innerData *utils.JsonDoc) *transactionsummary.TransactionSummary {
	transactionSummary := &transactionsummary.TransactionSummary{}

	if innerData.GetDecimal("tipAmount") != nil {
		transactionSummary.SetGratuityAmount(innerData.GetDecimal("tipAmount"))
	}
	if innerData.GetDecimal("taxAmount") != nil {
		transactionSummary.SetTaxAmount(innerData.GetDecimal("taxAmount"))
	}
	if innerData.GetDecimal("surcharge") != nil {
		transactionSummary.SetSurchargeAmount(innerData.GetDecimal("surcharge"))
	}

	transactionSummary.SetInvoiceNumber(innerData.GetString("invoiceNbr"))
	transactionSummary.SetClerkId(innerData.GetString("clerkId"))
	transactionSummary.SetTransactionType(innerData.GetString("transactionType"))
	transactionSummary.SetAmount(innerData.GetDecimal("totalAmount"))
	transactionSummary.SetMaskedCardNumber(innerData.GetString("maskedPan"))
	transactionSummary.SetReferenceNumber(innerData.GetString("referenceNumber"))
	transactionSummary.SetSafReferenceNumber(innerData.GetString("safReferenceNumber"))
	transactionSummary.SetGatewayResponseMessage(innerData.GetString("responseText"))
	transactionSummary.SetCardType(innerData.GetString("cardType"))
	transactionSummary.SetTransactionId(innerData.GetString("tranNo"))
	transactionSummary.SetTransactionDate(innerData.GetString("transactionTime"))

	if innerData.GetDecimal("baseAmount") != nil {
		transactionSummary.SetBaseAmount(innerData.GetDecimal("baseAmount"))
	}
	transactionSummary.SetGatewayResponseCode("responseCode")

	return transactionSummary
}
