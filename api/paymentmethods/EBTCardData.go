package paymentmethods

import (
	"fmt"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/cvnpresenceindicator"
	"github.com/globalpayments/go-sdk/api/entities/enums/ebtcardtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/manualentrymethod"
	"github.com/shopspring/decimal"
	"strconv"
)

type EBTCardData struct {
	*EBT
	ApprovalCode         string
	CardPresent          bool
	Cvn                  string
	CvnPresenceIndicator cvnpresenceindicator.CvnPresenceIndicator
	ExpMonth             *int
	ExpYear              *int
	Number               string
	ReaderPresent        bool
	SerialNumber         string
	EntryMethod          manualentrymethod.ManualEntryMethod
	TokenizationData     string
}

func NewEBTCardData() *EBTCardData {
	return &EBTCardData{
		EBT: NewEBT(),
	}
}

func NewEBTCardDataWithType(cardType ebtcardtype.EBTCardType) *EBTCardData {
	ebt := NewEBT()
	ebt.EbtCardType = cardType
	return &EBTCardData{
		EBT: ebt,
	}
}

func (e *EBTCardData) GetCardType() string {
	return string(e.EbtCardType)
}

func (e *EBTCardData) SetCardType(cardType string) {
	t, _ := ebtcardtype.ParseFromString(cardType)
	e.EbtCardType = t
}

func (e *EBTCardData) GetShortExpiry() string {
	if e.ExpYear == nil || e.ExpMonth == nil {
		return ""
	}
	expMonthStr := fmt.Sprintf("%02d", e.ExpMonth)
	expYearStr := strconv.Itoa(*e.ExpYear)[2:]
	return expMonthStr + expYearStr
}

func (e *EBTCardData) SetApprovalCode(approvalCode string) {
	e.ApprovalCode = approvalCode
}

func (e *EBTCardData) GetApprovalCode() string {
	return e.ApprovalCode
}

func (e *EBTCardData) SetCardPresent(cardPresent bool) {
	e.CardPresent = cardPresent
}

func (e *EBTCardData) IsCardPresent() bool {
	return e.CardPresent
}

func (e *EBTCardData) SetCvn(cvn string) {
	e.Cvn = cvn
}

func (e *EBTCardData) GetCvn() string {
	return e.Cvn
}

func (e *EBTCardData) SetCvnPresenceIndicator(cvnPresenceIndicator cvnpresenceindicator.CvnPresenceIndicator) {
	e.CvnPresenceIndicator = cvnPresenceIndicator
}

func (e *EBTCardData) GetCvnPresenceIndicator() cvnpresenceindicator.CvnPresenceIndicator {
	return e.CvnPresenceIndicator
}

func (e *EBTCardData) SetExpMonth(expMonth *int) {
	e.ExpMonth = expMonth
}

func (e *EBTCardData) GetExpMonth() *int {
	return e.ExpMonth
}

func (e *EBTCardData) SetExpYear(expYear *int) {
	e.ExpYear = expYear
}

func (e *EBTCardData) GetExpYear() *int {
	return e.ExpYear
}

func (e *EBTCardData) SetNumber(number string) {
	e.Number = number
}

func (e *EBTCardData) GetNumber() string {
	return e.Number
}

func (e *EBTCardData) SetReaderPresent(readerPresent bool) {
	e.ReaderPresent = readerPresent
}

func (e *EBTCardData) IsReaderPresent() bool {
	return e.ReaderPresent
}

func (e *EBTCardData) SetSerialNumber(serialNumber string) {
	e.SerialNumber = serialNumber
}

func (e *EBTCardData) GetSerialNumber() string {
	return e.SerialNumber
}

func (e *EBTCardData) SetEntryMethod(entryMethod manualentrymethod.ManualEntryMethod) {
	e.EntryMethod = entryMethod
}

func (e *EBTCardData) GetEntryMethod() manualentrymethod.ManualEntryMethod {
	return e.EntryMethod
}

func (e *EBTCardData) SetTokenizationData(tokenizationData string) {
	e.TokenizationData = tokenizationData
}

func (e *EBTCardData) GetTokenizationData() string {
	return e.TokenizationData
}

func (e *EBTCardData) ChargeWithAmount(amt *decimal.Decimal) *builders.AuthorizationBuilder {
	return e.EBT.ChargeWithAmount(amt, e)
}

func (e *EBTCardData) IBalanceInquiryWithType(inquiryType inquirytype.InquiryType) abstractions.IAuthorizationBuilder {
	return e.BalanceInquiryWithType(inquiryType)
}

func (e *EBTCardData) BalanceInquiryWithType(inquiryType inquirytype.InquiryType) *builders.AuthorizationBuilder {
	return e.EBT.BalanceInquiry(inquiryType, e)
}

func (e *EBTCardData) BalanceInquiry() *builders.AuthorizationBuilder {
	return e.EBT.BalanceInquiry(inquirytype.Foodstamp, e)
}

func (e *EBTCardData) RefundWithAmount(amt *decimal.Decimal) *builders.AuthorizationBuilder {
	return e.EBT.RefundWithAmount(amt, e)
}
