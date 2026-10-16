package entities

import (
	"github.com/shopspring/decimal"
)

// BatchSummary details a closed batch.
type BatchSummary struct {
	BatchId           *int
	TransactionCount  *int
	TotalAmount       *decimal.Decimal
	SequenceNumber    string
	Status            string
	BatchReference    string
	HostBatchNbr      *int
	HostTotalCnt      *int
	HostTotalAmt      *decimal.Decimal
	ProcessedDeviceId *int
	OpenTime          string
	OpenTransactionId string
}

func NewBatchSummary() *BatchSummary {
	return &BatchSummary{}
}

// SetId sets the batch's Id.
func (b *BatchSummary) SetId(id *int) {
	b.BatchId = id
}

// GetId returns the batch's Id.
func (b *BatchSummary) GetId() *int {
	return b.BatchId
}

// SetTransactionCount sets the batch's transaction count.
func (b *BatchSummary) SetTransactionCount(count *int) {
	b.TransactionCount = count
}

// GetTransactionCount returns the batch's transaction count.
func (b *BatchSummary) GetTransactionCount() *int {
	return b.TransactionCount
}

// SetTotalAmount sets the batch's total amount to be settled.
func (b *BatchSummary) SetTotalAmount(amount *decimal.Decimal) {
	b.TotalAmount = amount
}

// GetTotalAmount returns the batch's total amount to be settled.
func (b *BatchSummary) GetTotalAmount() *decimal.Decimal {
	return b.TotalAmount
}

// SetSequenceNumber sets the batch's sequence number.
func (b *BatchSummary) SetSequenceNumber(sequenceNumber string) {
	b.SequenceNumber = sequenceNumber
}

// GetSequenceNumber returns the batch's sequence number.
func (b *BatchSummary) GetSequenceNumber() string {
	return b.SequenceNumber
}

// SetStatus sets the batch's status.
func (b *BatchSummary) SetStatus(status string) {
	b.Status = status
}

// GetStatus returns the batch's status.
func (b *BatchSummary) GetStatus() string {
	return b.Status
}

// SetBatchReference sets the batch's reference.
func (b *BatchSummary) SetBatchReference(reference string) {
	b.BatchReference = reference
}

// GetBatchReference returns the batch's reference.
func (b *BatchSummary) GetBatchReference() string {
	return b.BatchReference
}

// SetHostBatchNbr sets the host batch number.
func (b *BatchSummary) SetHostBatchNbr(batchNbr *int) {
	b.HostBatchNbr = batchNbr
}

// GetHostBatchNbr returns the host batch number.
func (b *BatchSummary) GetHostBatchNbr() *int {
	return b.HostBatchNbr
}

// SetHostTotalCnt sets the host total count.
func (b *BatchSummary) SetHostTotalCnt(totalCnt *int) {
	b.HostTotalCnt = totalCnt
}

// GetHostTotalCnt returns the host total count.
func (b *BatchSummary) GetHostTotalCnt() *int {
	return b.HostTotalCnt
}

// SetHostTotalAmt sets the host total amount.
func (b *BatchSummary) SetHostTotalAmt(totalAmt *decimal.Decimal) {
	b.HostTotalAmt = totalAmt
}

// GetHostTotalAmt returns the host total amount.
func (b *BatchSummary) GetHostTotalAmt() *decimal.Decimal {
	return b.HostTotalAmt
}

// SetProcessedDeviceId sets the processed device Id.
func (b *BatchSummary) SetProcessedDeviceId(deviceId *int) {
	b.ProcessedDeviceId = deviceId
}

// GetProcessedDeviceId returns the processed device Id.
func (b *BatchSummary) GetProcessedDeviceId() *int {
	return b.ProcessedDeviceId
}
