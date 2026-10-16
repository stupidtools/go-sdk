package paymentmethods

import (
	"context"
	"errors"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/mobilepaymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodusagemode"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	abstractions2 "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/shopspring/decimal"
)

type Credit struct {
	BankName              string
	CardType              string
	EncryptionData        *base.EncryptionData
	EncryptedPan          string
	PaymentMethodType     paymentmethodtype.PaymentMethodType
	Token                 string
	MobileType            mobilepaymentmethodtype.MobilePaymentMethodType
	ThreeDSecure          *entities.ThreeDSecure
	Cryptogram            string
	FleetCard             bool
	PurchaseCard          bool
	ReadyLinkCard         bool
	PinBlock              string
	PaymentDataSourceType entities.PaymentDataSourceType
}

func NewCredit() *Credit {
	return &Credit{
		ThreeDSecure:      entities.NewThreeDSecure(),
		CardType:          "Unknown",
		PaymentMethodType: paymentmethodtype.Credit,
	}
}

func (c *Credit) IsFleet() bool {
	return c.FleetCard
}

func (c *Credit) AuthorizeWithAmount(amount *decimal.Decimal, isEstimated bool, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	if amount == nil && c.ThreeDSecure != nil {
		amount = c.ThreeDSecure.Amount
	}
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Auth, pm)
	builder.WithAmount(amount)
	builder.WithCurrency(c.ThreeDSecure.Currency)
	builder.WithOrderId(c.ThreeDSecure.OrderId)
	builder.WithAmountEstimated(isEstimated)
	return builder
}

func (c *Credit) Charge(pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	return c.ChargeWithAmount(nil, pm)
}

func (c *Credit) ChargeWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	if amount == nil && c.ThreeDSecure != nil {
		amount = c.ThreeDSecure.Amount
	}
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Sale, pm)
	builder.WithAmount(amount)
	builder.WithCurrency(c.ThreeDSecure.Currency)
	builder.WithOrderId(c.ThreeDSecure.OrderId)
	return builder
}

func (c *Credit) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	return c.PaymentMethodType
}

func (c *Credit) AddValue(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.AddValue, pm)
	builder.WithAmount(amount)
	return builder
}

func (c *Credit) BalanceInquiry(pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	return builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Balance, pm)
}

func (c *Credit) BalanceInquiryWithType(inquiry inquirytype.InquiryType, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Balance, pm)
	builder.WithBalanceInquiryType(inquiry)
	return builder
}

func (c *Credit) LoadReversal(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.LoadReversal, pm)
	builder.WithAmount(amount)
	return builder
}

func (c *Credit) Refund(pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	return c.RefundWithAmount(nil, pm)
}

func (c *Credit) RefundWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Refund, pm)
	builder.WithAmount(amount)
	return builder
}

func (c *Credit) Reverse(pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	return c.ReverseWithAmount(nil, pm)
}

func (c *Credit) ReverseWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Reversal, pm)
	builder.WithAmount(amount)
	return builder
}

func (c *Credit) Verify(pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	return builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Verify, pm)
}

func (c *Credit) Tokenize(pm abstractions2.IPaymentMethod) (*builders.AuthorizationBuilder, error) {
	return c.TokenizeWithParams(true, pm)
}

func (c *Credit) TokenizeWithParams(verifyCard bool, pm abstractions2.IPaymentMethod) (*builders.AuthorizationBuilder, error) {
	return c.tokenize(verifyCard, paymentmethodusagemode.MULTIPLE, pm)
}

func (c *Credit) tokenize(verifyCard bool, paymentMethodUsageMode paymentmethodusagemode.PaymentMethodUsageMode, pm abstractions2.IPaymentMethod) (*builders.AuthorizationBuilder, error) {

	transactionType := transactiontype.Tokenize
	if verifyCard {
		transactionType = transactiontype.Verify
	}

	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactionType, pm)
	builder.WithRequestMultiUseToken(verifyCard)
	builder.WithPaymentMethodUsageMode(paymentMethodUsageMode)
	return builder, nil
}

func (c *Credit) UpdateToken(ctx context.Context, gateway abstractions.IPaymentGateway, pm abstractions2.IPaymentMethod) (bool, error) {

	if c.Token == "" {
		return false, errors.New("Token cannot be null")
	}

	manager := builders.NewManagementBuilder(transactiontype.TokenUpdate)
	manager.WithPaymentMethod(pm)

	_, err := manager.Execute(ctx, gateway)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Credit) DeleteToken(ctx context.Context, gateway abstractions.IPaymentGateway, pm abstractions2.IPaymentMethod) (bool, error) {

	if c.Token == "" {
		return false, errors.New("Token cannot be null")
	}

	builder := builders.NewManagementBuilder(transactiontype.TokenDelete)

	builder.WithPaymentMethod(pm)
	_, err := builder.Execute(ctx, gateway)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Credit) SetEncryptionData(encryptionData *base.EncryptionData) {
	c.EncryptionData = encryptionData
}

func (c *Credit) GetEncryptionData() *base.EncryptionData {
	return c.EncryptionData
}

func (c *Credit) SetPaymentMethodType(paymentMethodType paymentmethodtype.PaymentMethodType) *Credit {
	c.PaymentMethodType = paymentMethodType
	return c
}

func (c *Credit) SetCardType(cardType string) *Credit {
	c.CardType = cardType
	return c
}

func (c *Credit) SetThreeDSecure(threeDSecure *entities.ThreeDSecure) *Credit {
	c.ThreeDSecure = threeDSecure
	return c
}

func (c *Credit) SetToken(token string) {
	c.Token = token
}

func (c *Credit) GetToken() string {
	return c.Token
}

func (c *Credit) SetMobileType(mobileType mobilepaymentmethodtype.MobilePaymentMethodType) *Credit {
	c.MobileType = mobileType
	return c
}

func (c *Credit) SetCryptogram(cryptogram string) *Credit {
	c.Cryptogram = cryptogram
	return c
}

func (c *Credit) SetFleetCard(fleetCard bool) *Credit {
	c.FleetCard = fleetCard
	return c
}

func (c *Credit) SetPurchaseCard(purchaseCard bool) *Credit {
	c.PurchaseCard = purchaseCard
	return c
}

func (c *Credit) SetReadyLinkCard(readyLinkCard bool) *Credit {
	c.ReadyLinkCard = readyLinkCard
	return c
}

func (c *Credit) SetPinBlock(pinBlock string) {
	c.PinBlock = pinBlock

}

func (c *Credit) GetPinBlock() string {
	return c.PinBlock
}

func (c *Credit) SetPaymentDataSourceType(paymentDataSourceType entities.PaymentDataSourceType) *Credit {
	c.PaymentDataSourceType = paymentDataSourceType
	return c
}

func (c *Credit) GetEncryptedPan() string {
	return c.EncryptedPan
}

func (c *Credit) SetEncryptedPan(pan string) {
	c.EncryptedPan = pan
}
