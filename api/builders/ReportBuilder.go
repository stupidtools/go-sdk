package builders

import (
	"context"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/entities/enums/reporttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/timezoneconversion"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	abstractions2 "github.com/globalpayments/go-sdk/api/gateways/abstractions"
)

type ReportBuilder struct {
	*BaseBuilder
	reportType         reporttype.ReportType
	timeZoneConversion timezoneconversion.TimeZoneConversion
}

func NewReportBuilder(reportType reporttype.ReportType) *ReportBuilder {
	return &ReportBuilder{
		reportType: reportType,
	}
}

func (b *ReportBuilder) GetReportType() reporttype.ReportType {
	return b.reportType
}

func (b *ReportBuilder) SetReportType(reportType reporttype.ReportType) {
	b.reportType = reportType
}

func (b *ReportBuilder) GetTimeZoneConversion() timezoneconversion.TimeZoneConversion {
	return b.timeZoneConversion
}

func (b *ReportBuilder) SetTimeZoneConversion(timeZoneConversion timezoneconversion.TimeZoneConversion) {
	b.timeZoneConversion = timeZoneConversion
}

func (b *ReportBuilder) Execute(ctx context.Context, gateway abstractions.IPaymentGateway, builder IReportBuilder) (abstractions.ITransaction, error) {

	if reportingService, ok := gateway.(abstractions2.IReportingService); ok {
		res, err := reportingService.ProcessReport(ctx, builder)
		return res, err
	}

	return nil, exceptions.NewApiException("Invalid reporting service client", nil)
}
