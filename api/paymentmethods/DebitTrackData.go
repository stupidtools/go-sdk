package paymentmethods

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/entrymethod"
	"github.com/globalpayments/go-sdk/api/entities/enums/tracknumber"
	"github.com/globalpayments/go-sdk/api/utils/cardutils"
	"github.com/shopspring/decimal"
	"strings"
)

type DebitTrackData struct {
	*Debit
	DiscretionaryData string
	EncryptedPan      string
	EntryMethod       entrymethod.EntryMethod
	Expiry            string
	Pan               string
	TrackNumber       tracknumber.TrackNumber
	TrackData         string
	Value             string
}

func NewDebitTrackData() *DebitTrackData {
	return &DebitTrackData{
		Debit:       NewDebit(),
		EntryMethod: entrymethod.Swipe,
	}
}

func (d *DebitTrackData) GetDiscretionaryData() string {
	return d.DiscretionaryData
}

func (d *DebitTrackData) SetDiscretionaryData(discretionaryData string) {
	d.DiscretionaryData = discretionaryData
}

func (d *DebitTrackData) GetEncryptedPan() string {
	return d.EncryptedPan
}

func (d *DebitTrackData) SetEncryptedPan(encryptedPan string) {
	d.EncryptedPan = encryptedPan
}

func (d *DebitTrackData) GetEntryMethod() entrymethod.EntryMethod {
	return d.EntryMethod
}

func (d *DebitTrackData) SetEntryMethod(entryMethod entrymethod.EntryMethod) {
	d.EntryMethod = entryMethod
}

func (d *DebitTrackData) GetExpiry() string {
	return d.Expiry
}

func (d *DebitTrackData) SetExpiry(expiry string) {
	d.Expiry = expiry
}

func (d *DebitTrackData) GetPan() string {
	return d.Pan
}

func (d *DebitTrackData) SetPan(pan string) {
	if pan != d.GetPan() {
		d.Pan = pan
		cardutils.ParseTrackData(d)
		d.CardType = cardutils.MapCardType(d.Pan)
	}
}

func (d *DebitTrackData) GetTrackNumber() tracknumber.TrackNumber {
	return d.TrackNumber
}

func (d *DebitTrackData) SetTrackNumber(trackNumber tracknumber.TrackNumber) {
	d.TrackNumber = trackNumber
}

func (d *DebitTrackData) GetTrackData() string {
	return d.TrackData
}

func (d *DebitTrackData) SetTrackData(trackData string) {
	if d.Value == "" {
		d.SetValue(trackData)
	} else {
		d.TrackData = trackData
	}
}

func (d *DebitTrackData) GetTruncatedTrackData() string {
	if d.DiscretionaryData != "" {
		return strings.Replace(d.TrackData, d.DiscretionaryData, "", -1)
	}
	return d.TrackData
}

func (d *DebitTrackData) GetValue() string {
	return d.Value
}

func (d *DebitTrackData) SetValue(value string) {
	d.Value = value
	cardutils.ParseTrackData(d)
	d.CardType = cardutils.MapCardType(d.Pan)
}

func (d *DebitTrackData) ChargeWithAmount(amt *decimal.Decimal) *builders.AuthorizationBuilder {
	return d.Debit.ChargeWithAmount(amt, d)
}

func (d *DebitTrackData) AuthorizeWithAmount(amt *decimal.Decimal, isEstimate bool) *builders.AuthorizationBuilder {
	return d.Debit.AuthorizeWithAmount(amt, isEstimate, d)
}

func (d *DebitTrackData) RefundWithAmount(amt *decimal.Decimal) *builders.AuthorizationBuilder {
	return d.Debit.RefundWithAmount(amt, d)
}

func (d *DebitTrackData) ReverseWithAmount(amt *decimal.Decimal) *builders.AuthorizationBuilder {
	return d.Debit.ReverseWithAmount(amt, d)
}
