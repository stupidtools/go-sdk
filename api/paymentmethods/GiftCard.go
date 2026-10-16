package paymentmethods

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/aliasaction"
	"github.com/globalpayments/go-sdk/api/entities/enums/entrymethod"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/tracknumber"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/shopspring/decimal"
)

type GiftCard struct {
	Alias       string
	Expiry      string
	Number      string
	PAN         string
	PIN         string
	Token       string
	TrackData   string
	TrackNumber tracknumber.TrackNumber
	Value       string
	ValueType   string
	EntryMethod entrymethod.EntryMethod
}

func NewGiftCard() *GiftCard {
	return &GiftCard{}
}

func (g *GiftCard) SetAlias(alias string) {
	g.Alias = alias
	g.Value = alias
	if g.Value == "" {
		g.SetValue(alias)
		g.ValueType = "Alias"
	}

}

func (g *GiftCard) SetExpiry(expiry string) {
	g.Expiry = expiry
}

func (g *GiftCard) SetNumber(number string) {
	if g.Value == "" {
		g.SetValue(number)
		g.ValueType = "CardNbr"
	}
	g.Number = number
}

func (g *GiftCard) SetPAN(pan string) {
	g.PAN = pan
}

func (g *GiftCard) SetPIN(pin string) {
	g.PIN = pin
}

func (g *GiftCard) SetToken(token string) {
	g.Token = token
	g.Value = token
	g.ValueType = "TokenValue"
}

func (g *GiftCard) SetTrackData(trackData string) {
	g.TrackData = trackData
	if g.Value == "" {
		g.SetValue(trackData)
		g.ValueType = "TrackData"
	}
}

func (g *GiftCard) SetTrackNumber(trackNumber tracknumber.TrackNumber) {
	g.TrackNumber = trackNumber
}

func (g *GiftCard) SetValue(value string) {
	g.Value = value

}

func (g *GiftCard) SetEntryMethod(method entrymethod.EntryMethod) {
	g.EntryMethod = method
}

func (g *GiftCard) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	return paymentmethodtype.Gift
}

func (g *GiftCard) AuthorizeWithAmount(amount *decimal.Decimal, isEstimated bool) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Auth, g)
	builder.WithAmount(amount)
	builder.WithAmountEstimated(isEstimated)
	return builder
}

func (g *GiftCard) AddValueWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.AddValue, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) BalanceInquiry(inquiryType inquirytype.InquiryType) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Balance, g)
	builder.WithBalanceInquiryType(inquiryType)
	return builder
}

func (g *GiftCard) ChargeWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Sale, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) Deactivate() *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Deactivate, g)
	return builder
}

func (g *GiftCard) RefundWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Refund, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) ReplaceWith(newCard *GiftCard) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Replace, g)
	builder.WithReplacementCard(newCard.Value, newCard.PIN, newCard.ValueType)
	return builder
}

func (g *GiftCard) ReverseWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Reversal, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) RewardsWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Reward, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) IssueWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Issue, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) ActivateWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Activate, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) CashOut() *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.CashOut, g)
	return builder
}

func (g *GiftCard) CaptureWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Capture, g)
	builder.WithAmount(amount)
	return builder
}

func (g *GiftCard) AddAlias(phoneNumber string) *builders.AuthorizationBuilder {
	g.SetAlias(phoneNumber)
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Alias, g)
	builder.WithAlias(aliasaction.Add, phoneNumber)
	return builder
}

func (g *GiftCard) RemoveAlias(phoneNumber string) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Alias, g)
	builder.WithAlias(aliasaction.Delete, phoneNumber)
	return builder
}

func CreateGiftCard(phoneNumber string) *builders.AuthorizationBuilder {
	card := NewGiftCard()
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Alias, card)
	builder.WithAlias(aliasaction.Create, phoneNumber)
	return builder
}

func (g *GiftCard) GetAlias() string {
	return g.Alias
}

func (g *GiftCard) GetExpiry() string {
	return g.Expiry
}

func (g *GiftCard) GetNumber() string {
	return g.Number
}

func (g *GiftCard) GetPAN() string {
	return g.PAN
}

func (g *GiftCard) GetPIN() string {
	return g.PIN
}

func (g *GiftCard) GetToken() string {
	return g.Token
}

func (g *GiftCard) GetTrackData() string {
	return g.TrackData
}

func (g *GiftCard) GetTrackNumber() tracknumber.TrackNumber {
	return g.TrackNumber
}

func (g *GiftCard) GetValue() string {
	return g.Value
}

func (g *GiftCard) GetValueType() string {
	return g.ValueType
}

func (g *GiftCard) GetEntryMethod() entrymethod.EntryMethod {
	return g.EntryMethod
}
