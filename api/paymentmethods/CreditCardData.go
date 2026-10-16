package paymentmethods

import (
	"context"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/cvnpresenceindicator"
	"github.com/globalpayments/go-sdk/api/entities/enums/manualentrymethod"
	"github.com/globalpayments/go-sdk/api/utils/cardutils"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"github.com/shopspring/decimal"
	"strconv"
)

type CreditCardData struct {
	*Credit
	CardHolderName       string
	CardPresent          bool
	cvn                  string
	CvnPresenceIndicator cvnpresenceindicator.CvnPresenceIndicator
	Eci                  string
	EntryMethod          manualentrymethod.ManualEntryMethod
	ExpMonth             *int
	expYear              *int
	number               string
	readerPresent        bool
	TokenizationData     string
}

func NewCreditCardData() *CreditCardData {
	return &CreditCardData{
		Credit:               NewCredit(),
		CvnPresenceIndicator: cvnpresenceindicator.NotRequested,
	}
}

func NewCreditCardDataWithToken(token string) *CreditCardData {
	data := NewCreditCardData()
	data.Token = token
	return data
}

func (c *CreditCardData) GetCvn() string {
	return c.cvn
}

func (c *CreditCardData) SetCvn(cvn string) {
	if cvn != "" {
		c.cvn = cvn
		c.CvnPresenceIndicator = cvnpresenceindicator.Present
	}
}

func (c *CreditCardData) GetExpYear() *int {
	return c.expYear
}

func (c *CreditCardData) GetCardHolderName() string {
	return c.CardHolderName
}

func (c *CreditCardData) SetExpYear(expYear *int) {
	c.expYear = expYear
}

func (c *CreditCardData) GetNumber() string {
	return c.number
}

func (c *CreditCardData) SetNumber(number string) {
	c.number = number
	c.CardType = cardutils.MapCardType(number)
	c.FleetCard = cardutils.IsFleet(c.CardType, number)
}

func (c *CreditCardData) GetShortExpiry() string {
	if c.ExpMonth != nil && c.expYear != nil {
		month := strconv.Itoa(*c.ExpMonth)
		year := strconv.Itoa(*c.expYear)[2:4]
		shortExpiry := stringutils.PadLeft(month, 2, '0') + year
		return shortExpiry
	}
	return ""
}

func (c *CreditCardData) IsReaderPresent() bool {
	return c.readerPresent
}

func (c *CreditCardData) SetReaderPresent(readerPresent bool) {
	c.readerPresent = readerPresent
}

func (c *CreditCardData) IsCardPresent() bool {
	return c.CardPresent
}

func (c *CreditCardData) SetCardPresent(cardPresent bool) {
	c.CardPresent = cardPresent
}

func (c *CreditCardData) GetCardType() string {
	return c.CardType
}

func (c *CreditCardData) SetCardType(cardType string) {
	c.CardType = cardType
}

func (c *CreditCardData) SetCardHolderName(cardHolderName string) {
	c.CardHolderName = cardHolderName
}

func (c *CreditCardData) GetCvnPresenceIndicator() cvnpresenceindicator.CvnPresenceIndicator {
	return c.CvnPresenceIndicator
}

func (c *CreditCardData) SetCvnPresenceIndicator(cvnPresenceIndicator cvnpresenceindicator.CvnPresenceIndicator) {
	c.CvnPresenceIndicator = cvnPresenceIndicator
}

func (c *CreditCardData) GetExpMonth() *int {
	return c.ExpMonth
}

func (c *CreditCardData) SetExpMonth(expMonth *int) {
	c.ExpMonth = expMonth
}

func (c *CreditCardData) GetEntryMethod() manualentrymethod.ManualEntryMethod {
	return c.EntryMethod
}

func (c *CreditCardData) SetEntryMethod(entryMethod manualentrymethod.ManualEntryMethod) {
	c.EntryMethod = entryMethod
}

func (c *CreditCardData) GetTokenizationData() string {
	return c.TokenizationData
}

func (c *CreditCardData) SetTokenizationData(tokenizationData string) {
	c.TokenizationData = tokenizationData
}

func (c *CreditCardData) Charge() *builders.AuthorizationBuilder {
	return c.ChargeWithAmount(nil)
}

func (c *CreditCardData) ChargeWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	return c.Credit.ChargeWithAmount(amount, c)
}

func (c *CreditCardData) Verify() *builders.AuthorizationBuilder {
	return c.Credit.Verify(c)
}

func (c *CreditCardData) AuthorizeWithAmount(amount *decimal.Decimal, isEstimated bool) *builders.AuthorizationBuilder {
	return c.Credit.AuthorizeWithAmount(amount, isEstimated, c)
}

func (c *CreditCardData) ReverseWithAmount(amount *decimal.Decimal) *builders.AuthorizationBuilder {
	return c.Credit.ReverseWithAmount(amount, c)
}

func (c *CreditCardData) Tokenize() (abstractions.ExecutableGateway, error) {
	return c.Credit.TokenizeWithParams(true, c)
}

func (c *CreditCardData) TokenizeWithParams(verifyCard bool) (abstractions.ExecutableGateway, error) {
	return c.Credit.TokenizeWithParams(verifyCard, c)
}

func (c *CreditCardData) UpdateToken(ctx context.Context, gateway abstractions.IPaymentGateway) (bool, error) {
	return c.Credit.UpdateToken(ctx, gateway, c)
}

func (c *CreditCardData) DeleteToken(ctx context.Context, gateway abstractions.IPaymentGateway) (bool, error) {
	return c.Credit.DeleteToken(ctx, gateway, c)
}
