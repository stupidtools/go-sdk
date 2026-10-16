package portico

import (
	"context"
	"github.com/globalpayments/go-sdk/api"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/globalpayments/go-sdk/api/serviceconfigs"
	"github.com/globalpayments/go-sdk/api/services/reportingservice"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"testing"
	"time"
)

func TestPorticoReportingTests(t *testing.T) {
	config := serviceconfigs.NewPorticoConfig()
	config.SecretApiKey = "skapi_cert_MTeSAQAfG1UA9qQDrzl-kz4toXvARyieptFwSKP24w"
	config.ServiceUrl = "https://cert.api2.heartlandportico.com"
	config.EnableLogging = true

	err := api.ConfigureService(config, "default")
	if err != nil {
		panic("Failed to configure service: " + err.Error())
	}
	reportActivity(t)
	reportTransactionDetail(t)
}

func reportActivity(t *testing.T) {
	ctx := context.Background()
	startDate := time.Now().AddDate(0, 0, -7)
	endDate := time.Now().AddDate(0, 0, -1)

	activity := reportingservice.Activity().
		WithStartDate(stringutils.ToStandardDateString(startDate)).
		WithEndDate(stringutils.ToStandardDateString(endDate))

	summary, err := api.ExecuteGateway[transactionsummary.ActivityReport](ctx, activity)

	if err != nil {
		t.Errorf("ReportActivity failed with error: %s", err.Error())
		return
	}

	if summary == nil || len(summary.TransactionSummaries) == 0 {
		t.Errorf("No transactions found")
	}
}

func reportTransactionDetail(t *testing.T) {
	ctx := context.Background()
	startDate := time.Now().AddDate(0, 0, -7)
	endDate := time.Now()

	activity := reportingservice.Activity().
		WithStartDate(stringutils.ToStandardDateString(startDate)).
		WithEndDate(stringutils.ToStandardDateString(endDate))

	activityReport, err := api.ExecuteGateway[transactionsummary.ActivityReport](ctx, activity)
	if err != nil {
		t.Errorf("ReportActivity failed with error: %s", err.Error())
		return
	}

	if activityReport == nil || len(activityReport.TransactionSummaries) == 0 {
		t.Errorf("No activities found")
		return
	}

	transactionDetail := reportingservice.FindTransactionsWithID(activityReport.TransactionSummaries[0].TransactionId)
	report, err := api.ExecuteGateway[transactionsummary.ActivityReport](ctx, transactionDetail)
	if err != nil {
		t.Errorf("ReportTransactionDetail failed with error: %s", err.Error())
		return
	}

	if report == nil || len(report.TransactionSummaries) == 0 {
		t.Errorf("No transaction detail found")
		return
	}
	detail := report.TransactionSummaries[0]
	if detail.TransactionId != activityReport.TransactionSummaries[0].TransactionId {
		t.Errorf("Transactions do not match")
		return
	}

	if detail.GatewayResponseCode != "00" {
		t.Errorf("Gateway response code is not '00', got %s", detail.GatewayResponseCode)
	}
}
