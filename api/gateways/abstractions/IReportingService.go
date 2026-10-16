package abstractions

import (
	"context"
	"github.com/globalpayments/go-sdk/api/abstractions"
	abstractions2 "github.com/globalpayments/go-sdk/api/builders/abstractions"
)

type IReportingService interface {
	ProcessReport(ctx context.Context, builder abstractions2.IReportBuilder) (abstractions.ITransaction, error)
}
