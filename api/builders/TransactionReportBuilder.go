package builders

import (
	"context"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders/validations"
	"github.com/globalpayments/go-sdk/api/entities/enums/reporttype"
)

type TransactionReportBuilder struct {
	*ReportBuilder
	deviceId       string
	endDate        string
	startDate      string
	transactionId  string
	searchCriteria map[string]string
}

func NewTransactionReportBuilder(reportType reporttype.ReportType) *TransactionReportBuilder {
	return &TransactionReportBuilder{
		ReportBuilder: NewReportBuilder(reportType),
	}
}

func (b *TransactionReportBuilder) SetupValidations() {
	b.ReportBuilder.Validations = *validations.NewValidations()
	b.ReportBuilder.Validations.Of(reporttype.TransactionDetail.LongValue()).
		Check("transactionId").IsNotNull("Transaction ID cannot be null").
		Check("transactionId").IsNotEmpty("Transaction ID cannot be empty").
		Check("deviceId").IsNull("Device ID cannot be null").
		Check("startDate").IsNull("Start Date cann be null").
		Check("endDate").IsNull("End Date cannot be null")

	b.ReportBuilder.Validations.Of(reporttype.Activity.LongValue()).
		Check("transactionId").IsNull("Transaction ID cannot be null")
}

func (b *TransactionReportBuilder) WithDeviceId(deviceId string) *TransactionReportBuilder {
	if deviceId != "" {
		b.deviceId = deviceId
	}
	return b
}

func (b *TransactionReportBuilder) WithEndDate(endDate string) *TransactionReportBuilder {
	b.endDate = endDate
	return b
}

func (b *TransactionReportBuilder) WithStartDate(startDate string) *TransactionReportBuilder {
	b.startDate = startDate
	return b
}

func (b *TransactionReportBuilder) WithTransactionId(transactionId string) *TransactionReportBuilder {
	if transactionId != "" {
		b.transactionId = transactionId
	}
	return b
}

func (b *TransactionReportBuilder) Where(criteria string, value string) *TransactionReportBuilder {
	if criteria != "" && value != "" {
		if b.searchCriteria == nil {
			b.searchCriteria = make(map[string]string)
		}
		b.searchCriteria[criteria] = value
	}
	return b
}

func (t *TransactionReportBuilder) GetReportBuilder() *ReportBuilder {
	return t.ReportBuilder
}

func (t *TransactionReportBuilder) GetDeviceId() string {
	return t.deviceId
}

func (t *TransactionReportBuilder) GetEndDate() string {
	return t.endDate
}

func (t *TransactionReportBuilder) GetStartDate() string {
	return t.startDate
}

func (t *TransactionReportBuilder) GetTransactionId() string {
	return t.transactionId
}

func (t *TransactionReportBuilder) GetSearchCriteria() map[string]string {
	return t.searchCriteria
}

func (t *TransactionReportBuilder) Execute(ctx context.Context, gateway abstractions.IPaymentGateway) (abstractions.ITransaction, error) {
	return t.ReportBuilder.Execute(ctx, gateway, t)
}
