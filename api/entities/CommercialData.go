package entities

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/shopspring/decimal"
)

type CommercialData struct {
	AdditionalTaxDetails    *AdditionalTaxDetails
	CommercialIndicator     transactionmodifier.TransactionModifier
	CustomerVATNumber       string
	CustomerReferenceId     string
	Description             string
	DiscountAmount          *decimal.Decimal
	DutyAmount              *decimal.Decimal
	DestinationPostalCode   string
	DestinationCountryCode  string
	FreightAmount           *decimal.Decimal
	LineItems               []*CommercialLineItem
	OrderDate               string
	OriginPostalCode        string
	PONumber                string
	SupplierReferenceNumber string
	TaxAmount               *decimal.Decimal
	TaxType                 taxtype.TaxType
	SummaryCommodityCode    string
	VATInvoiceNumber        string
}

func NewCommercialData(taxType taxtype.TaxType) *CommercialData {
	return NewCommercialDataWithLevel(taxType, transactionmodifier.LevelII)
}

func NewCommercialDataWithLevel(taxType taxtype.TaxType, level transactionmodifier.TransactionModifier) *CommercialData {
	return &CommercialData{
		TaxType:             taxType,
		CommercialIndicator: level,
		LineItems:           []*CommercialLineItem{},
	}
}

func (cd *CommercialData) AddLineItems(item *CommercialLineItem) *CommercialData {
	cd.LineItems = append(cd.LineItems, item)
	return cd
}
