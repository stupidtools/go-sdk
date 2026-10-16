package subgroups

import (
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/utils"
)

type RequestTransactionFields struct {
	Amount              *decimal.Decimal
	BaseAmount          *decimal.Decimal
	CommercialRequest   bool
	TaxAmount           *decimal.Decimal
	TipAmount           *decimal.Decimal
	TaxIndicator        *int
	CashBackAmount      *decimal.Decimal
	InvoiceNbr          *int
	TotalAmount         *decimal.Decimal
	TerminalRefNumber   string
	GatewayRefNumber    string
	GiftTransactionType string
}

func (r *RequestTransactionFields) SetManageBuilderParams(builder *builders.TerminalManageBuilder) {
	if builder.GetTerminalRefNumber() != "" {
		r.TerminalRefNumber = builder.GetTerminalRefNumber()
	}
	if builder.GetGratuity() != nil {
		r.TipAmount = builder.GetGratuity()
	}

	if builder.GetTransactionId() != "" {
		r.GatewayRefNumber = builder.GetTransactionId()
	}

	if builder.GetAmount() != nil {
		r.Amount = builder.GetAmount()
	}
}

func (r *RequestTransactionFields) SetAuthBuilderParams(builder *builders.TerminalAuthBuilder) {
	if builder.Amount != nil {
		if builder.GetTransactionType() == transactiontype.Refund || builder.GetTransactionType() == transactiontype.Activate {
			r.TotalAmount = builder.Amount
		} else if builder.GetTransactionType() == transactiontype.Auth {
			r.Amount = builder.Amount
		} else {
			r.BaseAmount = builder.Amount
		}
	}

	if builder.Gratuity != nil {
		r.TipAmount = builder.Gratuity
	}

	if builder.GetTransactionId() != "" {
		r.GatewayRefNumber = builder.GetTransactionId()
	}

	if builder.CashBackAmount != nil {
		r.CashBackAmount = builder.CashBackAmount
	}

	if builder.TaxAmount != nil {
		r.TaxAmount = builder.TaxAmount
	}

	if builder.CommercialRequest {
		r.CommercialRequest = builder.CommercialRequest
	}
	if builder.GiftTransactionType != transactiontype.Empty {
		r.GiftTransactionType = builder.GiftTransactionType.GetStringValue()
	}

}

func (r *RequestTransactionFields) GetElementsJson() *utils.JsonDoc {
	params := utils.NewJsonDoc()
	hasContents := false
	if r.Amount != nil {
		params.Set("amount", r.Amount.StringFixed(2), true)
		hasContents = true
	}

	if r.TerminalRefNumber != "" {
		params.Set("tranNo", r.TerminalRefNumber, true)
		hasContents = true
	}

	if r.TotalAmount != nil {
		params.Set("totalAmount", r.TotalAmount.StringFixed(2), true)
		hasContents = true
	}

	if r.BaseAmount != nil {
		params.Set("baseAmount", r.BaseAmount.StringFixed(2), true)
		hasContents = true
	}

	if r.TaxAmount != nil {
		params.Set("taxAmount", r.TaxAmount.StringFixed(2), true)
		hasContents = true
	}

	if r.TipAmount != nil {
		params.Set("tipAmount", r.TipAmount.StringFixed(2), true)
		hasContents = true
	}

	if r.TaxIndicator != nil {
		params.SetInt("taxIndicator", r.TaxIndicator)
		hasContents = true
	}

	if r.CashBackAmount != nil {
		params.Set("cashBackAmount", r.CashBackAmount.StringFixed(2), true)
		hasContents = true
	}

	if r.InvoiceNbr != nil {
		params.SetInt("invoiceNbr", r.InvoiceNbr)
		hasContents = true
	}

	if r.GatewayRefNumber != "" {
		params.Set("referenceNumber", r.GatewayRefNumber, true)
		hasContents = true
	}

	if r.CommercialRequest {
		params.Set("processCPC", "1", true)
	}

	if r.GiftTransactionType != "" {
		params.Set("transactionType", r.GiftTransactionType, true)
	}

	if hasContents {
		return params
	} else {
		return nil
	}
}

func NewRequestTransactionFields() *RequestTransactionFields {
	return &RequestTransactionFields{}
}
