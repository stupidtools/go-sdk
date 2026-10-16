package entities

type TransactionMatchingData struct {
	OriginalBatchNumber string
	OriginalDate        string
}

func (t *TransactionMatchingData) GetElementData() string {
	if t.OriginalBatchNumber != "" && t.OriginalDate != "" {
		return t.OriginalBatchNumber + t.OriginalDate
	}
	return ""
}

func NewTransactionMatchingData(batchNumber string, date string) *TransactionMatchingData {
	return &TransactionMatchingData{
		OriginalBatchNumber: batchNumber,
		OriginalDate:        date,
	}
}
