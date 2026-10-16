package paymentmethods

import (
	"github.com/globalpayments/go-sdk/api/builders"
	"github.com/globalpayments/go-sdk/api/entities/enums/ebtcardtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/inquirytype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	abstractions2 "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/shopspring/decimal"
)

type EBT struct {
	EbtCardType    ebtcardtype.EBTCardType
	PinBlock       string
	CardHolderName string
}

func NewEBT() *EBT {
	return &EBT{}
}

func (e *EBT) SetEbtCardType(ebtCardType ebtcardtype.EBTCardType) {
	e.EbtCardType = ebtCardType
}

func (e *EBT) GetEbtCardType() ebtcardtype.EBTCardType {
	return e.EbtCardType
}

func (e *EBT) GetPaymentMethodType() paymentmethodtype.PaymentMethodType {
	return paymentmethodtype.EBT
}

func (e *EBT) SetPinBlock(pinBlock string) {
	e.PinBlock = pinBlock
}

func (e *EBT) GetPinBlock() string {
	return e.PinBlock
}

func (e *EBT) SetCardHolderName(cardHolderName string) {
	e.CardHolderName = cardHolderName
}

func (e *EBT) GetCardHolderName() string {
	return e.CardHolderName
}

func (e *EBT) AuthorizeWithAmount(amount *decimal.Decimal, isEstimated bool, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Auth, pm)
	builder.WithAmount(amount)
	builder.WithAmountEstimated(isEstimated)
	return builder
}

func (e *EBT) ChargeWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Sale, pm)
	builder.WithAmount(amount)
	return builder
}

func (e *EBT) BalanceInquiry(inquiryType inquirytype.InquiryType, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Balance, pm)
	builder.WithBalanceInquiryType(inquiryType)
	val := decimal.NewFromInt(0)
	builder.WithAmount(&val)
	return builder
}

func (e *EBT) BenefitWithdrawalWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.BenefitWithdrawal, pm)
	builder.WithAmount(amount)
	val := decimal.NewFromInt(0)
	builder.WithCashBackAmount(&val)
	return builder
}

func (e *EBT) RefundWithAmount(amount *decimal.Decimal, pm abstractions2.IPaymentMethod) *builders.AuthorizationBuilder {
	builder := builders.NewAuthorizationBuilderWithPaymentMethod(transactiontype.Refund, pm)
	builder.WithAmount(amount)
	return builder
}
