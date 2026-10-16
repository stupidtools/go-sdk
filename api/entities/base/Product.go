package base

import "github.com/shopspring/decimal"

type Product struct {
	ProductId        string
	ProductName      string
	Description      string
	Quantity         int
	UnitPrice        *decimal.Decimal
	NetUnitPrice     *decimal.Decimal
	Gift             bool
	UnitCurrency     string
	Type             string
	Risk             string
	TaxAmount        *decimal.Decimal
	TaxPercentage    *decimal.Decimal
	NetUnitAmount    *decimal.Decimal
	DiscountAmount   *decimal.Decimal
	GiftCardCurrency string
	Url              string
	ImageUrl         string
}

// ProductId
func (p *Product) GetProductId() string {
	return p.ProductId
}

func (p *Product) SetProductId(productId string) {
	p.ProductId = productId
}

// ProductName
func (p *Product) GetProductName() string {
	return p.ProductName
}

func (p *Product) SetProductName(productName string) {
	p.ProductName = productName
}

// Description
func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) SetDescription(description string) {
	p.Description = description
}

// Quantity
func (p *Product) GetQuantity() int {
	return p.Quantity
}

func (p *Product) SetQuantity(quantity int) {
	p.Quantity = quantity
}

// UnitPrice
func (p *Product) GetUnitPrice() *decimal.Decimal {
	return p.UnitPrice
}

func (p *Product) SetUnitPrice(unitPrice *decimal.Decimal) {
	p.UnitPrice = unitPrice
}

// NetUnitPrice
func (p *Product) GetNetUnitPrice() *decimal.Decimal {
	return p.NetUnitPrice
}

func (p *Product) SetNetUnitPrice(netUnitPrice *decimal.Decimal) {
	p.NetUnitPrice = netUnitPrice
}

// Gift
func (p *Product) GetGift() bool {
	return p.Gift
}

func (p *Product) SetGift(gift bool) {
	p.Gift = gift
}

// UnitCurrency
func (p *Product) GetUnitCurrency() string {
	return p.UnitCurrency
}

func (p *Product) SetUnitCurrency(unitCurrency string) {
	p.UnitCurrency = unitCurrency
}

// Type
func (p *Product) GetType() string {
	return p.Type
}

func (p *Product) SetType(typeValue string) {
	p.Type = typeValue
}

// Risk
func (p *Product) GetRisk() string {
	return p.Risk
}

func (p *Product) SetRisk(risk string) {
	p.Risk = risk
}

// TaxAmount
func (p *Product) GetTaxAmount() *decimal.Decimal {
	return p.TaxAmount
}

func (p *Product) SetTaxAmount(taxAmount *decimal.Decimal) {
	p.TaxAmount = taxAmount
}

// TaxPercentage
func (p *Product) GetTaxPercentage() *decimal.Decimal {
	return p.TaxPercentage
}

func (p *Product) SetTaxPercentage(taxPercentage *decimal.Decimal) {
	p.TaxPercentage = taxPercentage
}

// NetUnitAmount
func (p *Product) GetNetUnitAmount() *decimal.Decimal {
	return p.NetUnitAmount
}

func (p *Product) SetNetUnitAmount(netUnitAmount *decimal.Decimal) {
	p.NetUnitAmount = netUnitAmount
}

// DiscountAmount
func (p *Product) GetDiscountAmount() *decimal.Decimal {
	return p.DiscountAmount
}

func (p *Product) SetDiscountAmount(discountAmount *decimal.Decimal) {
	p.DiscountAmount = discountAmount
}

// GiftCardCurrency
func (p *Product) GetGiftCardCurrency() string {
	return p.GiftCardCurrency
}

func (p *Product) SetGiftCardCurrency(giftCardCurrency string) {
	p.GiftCardCurrency = giftCardCurrency
}

// Url
func (p *Product) GetUrl() string {
	return p.Url
}

func (p *Product) SetUrl(url string) {
	p.Url = url
}

// ImageUrl
func (p *Product) GetImageUrl() string {
	return p.ImageUrl
}

func (p *Product) SetImageUrl(imageUrl string) {
	p.ImageUrl = imageUrl
}
