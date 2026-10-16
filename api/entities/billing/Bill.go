package billing

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/billpresentment"
	"github.com/globalpayments/go-sdk/api/entities/recurring"
	"github.com/shopspring/decimal"
	"time"
)

type Bill struct {
	billType        string
	identifier1     string
	identifier2     string
	identifier3     string
	identifier4     string
	amount          *decimal.Decimal
	customer        recurring.Customer
	billPresentment billpresentment.BillPresentment
	dueDate         time.Time
}

func (b *Bill) GetBillType() string {
	return b.billType
}

func (b *Bill) GetIdentifier1() string {
	return b.identifier1
}

func (b *Bill) GetIdentifier2() string {
	return b.identifier2
}

func (b *Bill) GetIdentifier3() string {
	return b.identifier3
}

func (b *Bill) GetIdentifier4() string {
	return b.identifier4
}

func (b *Bill) GetAmount() *decimal.Decimal {
	return b.amount
}

func (b *Bill) GetCustomer() recurring.Customer {
	return b.customer
}

func (b *Bill) GetBillPresentment() billpresentment.BillPresentment {
	return b.billPresentment
}

func (b *Bill) GetDueDate() time.Time {
	return b.dueDate
}

func (b *Bill) SetBillType(billType string) {
	b.billType = billType
}

func (b *Bill) SetIdentifier1(identifier1 string) {
	b.identifier1 = identifier1
}

func (b *Bill) SetIdentifier2(identifier2 string) {
	b.identifier2 = identifier2
}

func (b *Bill) SetIdentifier3(identifier3 string) {
	b.identifier3 = identifier3
}

func (b *Bill) SetIdentifier4(identifier4 string) {
	b.identifier4 = identifier4
}

func (b *Bill) SetAmount(amount *decimal.Decimal) {
	b.amount = amount
}

func (b *Bill) SetCustomer(customer recurring.Customer) {
	b.customer = customer
}

func (b *Bill) SetBillPresentment(billPresentment billpresentment.BillPresentment) {
	b.billPresentment = billPresentment
}

func (b *Bill) SetDueDate(dueDate time.Time) {
	b.dueDate = dueDate
}
