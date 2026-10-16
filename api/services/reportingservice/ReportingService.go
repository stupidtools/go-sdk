package reportingservice

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/reporttype"
)

func FindTransactions() *builders.TransactionReportBuilder {
	return builders.NewTransactionReportBuilder(reporttype.FindTransactions)
}

func FindTransactionsWithID(transactionId string) *builders.TransactionReportBuilder {
	return builders.NewTransactionReportBuilder(reporttype.FindTransactions).
		WithTransactionId(transactionId)
}

func Activity() *builders.TransactionReportBuilder {
	return builders.NewTransactionReportBuilder(reporttype.Activity)
}
