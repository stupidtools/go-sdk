package responses

import (
	"encoding/json"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaReportResponse struct {
	BatchSummary       entities.BatchSummary
	DeviceResponseCode string
	DeviceResponseText string
	Status             string
	Transactions       []transactionsummary.TransactionSummary
	TransactionType    string
}

func NewUpaReportResponse(responseObj *utils.JsonDoc) *UpaReportResponse {
	res := &UpaReportResponse{}
	responseData := responseObj.Get("data")

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
			batchRecord := innerData.Get("batchRecord")

			if batchRecord != nil {
				batchTime := batchRecord.GetString("openUtcDateTime")
				res.BatchSummary = entities.BatchSummary{
					BatchId:           batchRecord.GetInt("batchId"),
					SequenceNumber:    batchRecord.GetString("batchSeqNbr"),
					Status:            batchRecord.GetString("batchStatus"),
					OpenTime:          batchTime,
					OpenTransactionId: batchRecord.GetString("openTxnId"),
					TotalAmount:       batchRecord.GetDecimal("totalAmount"),
					TransactionCount:  batchRecord.GetInt("totalCnt"),
				}

				batchDetailRecords := batchRecord.GetArray("batchDetailRecords")
				if batchDetailRecords != nil {
					for _, n := range batchDetailRecords {
						trans := transactionsummary.TransactionSummary{
							AmountDue:         n.GetDecimal("balanceDue"),
							AuthCode:          n.GetString("approvalCode"),
							AuthorizedAmount:  n.GetDecimal("authorizedAmount"),
							BaseAmount:        n.GetDecimal("baseAmount"),
							CardSwiped:        n.GetString("cardSwiped"),
							CardType:          n.GetString("cardType"),
							CashBackAmount:    n.GetDecimal("cashbackAmount"),
							ClerkId:           n.GetString("clerkId"),
							InvoiceNumber:     n.GetString("invoiceNbr"),
							MaskedCardNumber:  n.GetString("maskedCardNumber"),
							SettlementAmount:  n.GetDecimal("settleAmount"),
							TaxAmount:         n.GetDecimal("taxAmount"),
							GratuityAmount:    n.GetDecimal("tipAmount"),
							Amount:            n.GetDecimal("totalAmount"),
							TransactionId:     n.GetString("gatewayTxnId"),
							TransactionStatus: n.GetString("transactionStatus"),
							TransactionType:   n.GetString("transactionType"),
						}
						res.Transactions = append(res.Transactions, trans)
					}
				}
			}

			openTabDetails := innerData.GetArray("OpenTabDetails")
			if openTabDetails != nil {
				for _, n := range openTabDetails {
					trans := transactionsummary.TransactionSummary{
						AuthorizedAmount: n.GetDecimal("authorizedAmount"),
						CardType:         n.GetString("cardType"),
						ClerkId:          n.GetString("clerkId"),
						MaskedCardNumber: n.GetString("maskedPan"),
						TransactionId:    n.GetString("referenceNumber"),
					}
					res.Transactions = append(res.Transactions, trans)
				}
			}
		}
	}

	return res
}

func (res *UpaReportResponse) GetTransactionType() string {
	return res.TransactionType
}

func (res *UpaReportResponse) GetStatus() string {
	return res.Status
}

func (res *UpaReportResponse) GetDeviceResponseCode() string {
	return res.DeviceResponseCode
}

func (res *UpaReportResponse) GetDeviceResponseText() string {
	return res.DeviceResponseText
}

func (res *UpaReportResponse) GetTransactionSummaries() []transactionsummary.TransactionSummary {
	return res.Transactions
}

func (res *UpaReportResponse) GetBatchSummary() entities.BatchSummary {
	return res.BatchSummary
}

func (res *UpaReportResponse) SetDeviceResponseCode(deviceResponseCode string) {
	res.DeviceResponseCode = deviceResponseCode
}

func (res *UpaReportResponse) SetDeviceResponseText(deviceResponseMessage string) {
	res.DeviceResponseText = deviceResponseMessage
}

func (res *UpaReportResponse) GetVersion() string {
	return ""
}

func (res *UpaReportResponse) SetVersion(version string) {

}

func (res *UpaReportResponse) SetStatus(status string) {
	res.Status = status
}

func (res *UpaReportResponse) GetCommand() string {
	return ""
}

func (res *UpaReportResponse) SetCommand(command string) {

}

func (res *UpaReportResponse) ToString() string {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
