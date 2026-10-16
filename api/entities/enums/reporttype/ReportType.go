package reporttype

type ReportType string

const (
	FindTransactions                ReportType = "FindTransactions"
	Activity                        ReportType = "Activity"
	BatchDetail                     ReportType = "BatchDetail"
	BatchHistory                    ReportType = "BatchHistory"
	BatchSummary                    ReportType = "BatchSummary"
	OpenAuths                       ReportType = "OpenAuths"
	Search                          ReportType = "Search"
	TransactionDetail               ReportType = "TransactionDetail"
	DepositDetail                   ReportType = "DepositDetail"
	DisputeDetail                   ReportType = "DisputeDetail"
	SettlementDisputeDetail         ReportType = "SettlementDisputeDetail"
	FindTransactionsPaged           ReportType = "FindTransactionsPaged"
	FindSettlementTransactionsPaged ReportType = "FindSettlementTransactionsPaged"
	FindDepositsPaged               ReportType = "FindDepositsPaged"
	FindDisputesPaged               ReportType = "FindDisputesPaged"
	FindSettlementDisputesPaged     ReportType = "FindSettlementDisputesPaged"
	StoredPaymentMethodDetail       ReportType = "StoredPaymentMethodDetail"
	FindStoredPaymentMethodsPaged   ReportType = "FindStoredPaymentMethodsPaged"
	ActionDetail                    ReportType = "ActionDetail"
	FindActionsPaged                ReportType = "FindActionsPaged"
	DocumentDisputeDetail           ReportType = "DocumentDisputeDetail"
	FindBankPayment                 ReportType = "FindBankPayment"
	PayByLinkDetail                 ReportType = "PayByLinkDetail"
	FindPayByLinkPaged              ReportType = "FindPayByLinkPaged"
	FindMerchantsPaged              ReportType = "FindMerchantsPaged"
	FindAccountsPaged               ReportType = "FindAccountsPaged"
	FindAccountDetail               ReportType = "FindAccountDetail"
)

var reportTypeValues = map[ReportType]int{
	FindTransactions:                0,
	Activity:                        1,
	BatchDetail:                     2,
	BatchHistory:                    3,
	BatchSummary:                    4,
	OpenAuths:                       5,
	Search:                          6,
	TransactionDetail:               7,
	DepositDetail:                   8,
	DisputeDetail:                   9,
	SettlementDisputeDetail:         10,
	FindTransactionsPaged:           11,
	FindSettlementTransactionsPaged: 12,
	FindDepositsPaged:               13,
	FindDisputesPaged:               14,
	FindSettlementDisputesPaged:     15,
	StoredPaymentMethodDetail:       16,
	FindStoredPaymentMethodsPaged:   17,
	ActionDetail:                    18,
	FindActionsPaged:                19,
	DocumentDisputeDetail:           20,
	FindBankPayment:                 21,
	PayByLinkDetail:                 22,
	FindPayByLinkPaged:              23,
	FindMerchantsPaged:              24,
	FindAccountsPaged:               25,
	FindAccountDetail:               26,
}

func (r ReportType) LongValue() int64 {
	return 1 << reportTypeValues[r]
}

func GetSet(value int64) map[ReportType]bool {
	flags := make(map[ReportType]bool)
	for reportType, _ := range reportTypeValues {
		flagValue := reportType.LongValue()
		if (flagValue & value) == flagValue {
			flags[reportType] = true
		}
	}
	return flags
}

func (r ReportType) Ordinal() int {
	return reportTypeValues[r]
}
