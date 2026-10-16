package paymentmethods

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	abstractions2 "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/shopspring/decimal"
)

type Debit struct {
	EncryptionData    *base.EncryptionData
	PaymentMethodType paymentmethodtype.PaymentMethodType
	PinBlock          string
	CardType          string
	TokenizedData     string
}

func NewDebit() *Debit {
	return &Debit{
		CardType:          "Unknown",
		PaymentMethodType: paymentmethodtype.Debit,
	}
}

func (d *Debit) SetPaymentMethodType(pa paymentmethodtype.PaymentMethodType) {
	d.PaymentMethodType = pa
}

func (d *Debit) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	return d.PaymentMethodType
}

func (d *Debit) SetCardType(cardType string) {
	d.CardType = cardType
}

func (d *Debit) GetCardType() string {
	return d.CardType
}

func (d *Debit) SetEncryptionData(encryptionData *base.EncryptionData) {
	d.EncryptionData = encryptionData
}

func (d *Debit) GetEncryptionData() *base.EncryptionData {
	return d.EncryptionData
}

func (d *Debit) SetPinBlock(pinBlock string) {
	d.PinBlock = pinBlock
}

func (d *Debit) GetPinBlock() string {
	return d.PinBlock
}

func (d *Debit) SetTokenizedData(tokenizedData string) {
	d.TokenizedData = tokenizedData
}

func (d *Debit) GetTokenizedData() string {
	return d.TokenizedData
}

//
//func (d *Debit) AddValue(amount *decimal.Decimal) *builders.AuthorizationBuilder {
//	builder := builders.NewAuthorizationBuilder(transactiontype.AddValue)
//	builder.WithAmount(amount)
//	return builder
//}
//

func (d *Debit) AuthorizeWithAmount(amount *decimal.Decimal, isEstimated bool, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Auth, pm)
	builder.WithAmount(amount)
	builder.WithAmountEstimated(isEstimated)
	return builder
}

func (d *Debit) ChargeWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Sale, pm)
	builder.WithAmount(amount)
	return builder
}

func (d *Debit) RefundWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Refund, pm)
	builder.WithAmount(amount)
	return builder
}

func (d *Debit) ReverseWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Reversal, pm)
	builder.WithAmount(amount)
	return builder
}
