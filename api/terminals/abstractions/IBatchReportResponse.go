package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
)

type IBatchReportResponse interface {
	IDeviceResponse
	GetBatchSummary() entities.BatchSummary
	GetTransactionSummaries() []transactionsummary.TransactionSummary
}
