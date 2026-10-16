package batchservice

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
)

func CloseBatch() *builders.ManagementBuilder {
	return builders.NewManagementBuilder(transactiontype.BatchClose)

}
