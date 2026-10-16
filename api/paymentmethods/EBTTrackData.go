package paymentmethods

import (
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/entrymethod"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/tracknumber"
	"github.com/globalpayments/go-sdk/api/utils/cardutils"
	"strings"
)

type EBTTrackData struct {
	*EBT
	DiscretionaryData string
	EncryptionData    *base.EncryptionData
	EncryptedPan      string
	EntryMethod       entrymethod.EntryMethod
	Expiry            string
	Pan               string
	TrackNumber       tracknumber.TrackNumber
	TrackData         string
	Value             string
}

func NewEBTTrackData() *EBTTrackData {
	return &EBTTrackData{
		EBT: NewEBT(),
	}
}

func (e *EBTTrackData) GetDiscretionaryData() string {
	return e.DiscretionaryData
}

func (e *EBTTrackData) SetDiscretionaryData(discretionaryData string) {
	e.DiscretionaryData = discretionaryData
}

func (e *EBTTrackData) GetEncryptionData() *base.EncryptionData {
	return e.EncryptionData
}

func (e *EBTTrackData) SetEncryptionData(encryptionData *base.EncryptionData) {
	e.EncryptionData = encryptionData
}

func (e *EBTTrackData) GetEncryptedPan() string {
	return e.EncryptedPan
}

func (e *EBTTrackData) SetEncryptedPan(encryptedPan string) {
	e.EncryptedPan = encryptedPan
}

func (e *EBTTrackData) GetEntryMethod() entrymethod.EntryMethod {
	return e.EntryMethod
}

func (e *EBTTrackData) SetEntryMethod(entryMethod entrymethod.EntryMethod) {
	e.EntryMethod = entryMethod
}

func (e *EBTTrackData) GetExpiry() string {
	return e.Expiry
}

func (e *EBTTrackData) SetExpiry(expiry string) {
	e.Expiry = expiry
}

func (e *EBTTrackData) GetPan() string {
	return e.Pan
}

func (e *EBTTrackData) SetPan(pan string) {
	if pan != e.GetPan() {
		e.Pan = pan
		cardutils.ParseTrackData(e) // Assuming cardutils.ParseTrackData exists in Go
	}
}

func (e *EBTTrackData) GetTrackNumber() tracknumber.TrackNumber {
	return e.TrackNumber
}

func (e *EBTTrackData) SetTrackNumber(trackNumber tracknumber.TrackNumber) {
	e.TrackNumber = trackNumber
}

func (e *EBTTrackData) GetTrackData() string {
	return e.TrackData
}

func (e *EBTTrackData) SetTrackData(trackData string) {
	if e.Value == "" {
		e.SetValue(trackData)
	} else {
		e.TrackData = trackData
	}
}

func (e *EBTTrackData) GetTruncatedTrackData() string {
	if e.DiscretionaryData != "" {
		return strings.Replace(e.TrackData, e.DiscretionaryData, "", -1)
	}
	return e.TrackData
}

func (e *EBTTrackData) GetValue() string {
	return e.Value
}

func (e *EBTTrackData) SetValue(value string) {
	e.Value = value
	cardutils.ParseTrackData(e)
}

func (e *EBTTrackData) IBalanceInquiryWithType(inquiryType inquirytype.InquiryType) abstractions.IAuthorizationBuilder {
	return e.BalanceInquiryWithType(inquiryType)
}

func (e *EBTTrackData) BalanceInquiryWithType(inquiryType inquirytype.InquiryType) *builders.AuthorizationBuilder {
	return e.EBT.BalanceInquiry(inquiryType, e)
}

func (e *EBTTrackData) BalanceInquiry() *builders.AuthorizationBuilder {
	return e.EBT.BalanceInquiry(inquirytype.Foodstamp, e)
}
