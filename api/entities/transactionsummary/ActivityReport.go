package transactionsummary

type ActivityReport struct {
	TransactionSummaries []TransactionSummary
}

func NewActivityReport() *ActivityReport {
	summaries := make([]TransactionSummary, 0)
	return &ActivityReport{TransactionSummaries: summaries}
}

func (ap *ActivityReport) AddSummary(t TransactionSummary) {
	ap.TransactionSummaries = append(ap.TransactionSummaries, t)
}
