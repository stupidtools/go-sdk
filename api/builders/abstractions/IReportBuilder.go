package abstractions

import (
	"context"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/entities/enums/reporttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/timezoneconversion"
)

type IReportBuilder interface {
	GetReportType() reporttype.ReportType
	SetReportType(reporttype.ReportType)
	GetTimeZoneConversion() timezoneconversion.TimeZoneConversion
	SetTimeZoneConversion(timezoneconversion.TimeZoneConversion)
	Execute(ctx context.Context, gateway abstractions.IPaymentGateway) (abstractions.ITransaction, error)
}
