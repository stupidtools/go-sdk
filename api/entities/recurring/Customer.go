package recurring

import (
	"github.com/globalpayments/go-sdk/api/entities/base"
)

type Customer struct {
	*RecurringEntity
	Title             string
	FirstName         string
	MiddleName        string
	LastName          string
	Company           string
	CustomerPassword  string
	DateOfBirth       string
	DomainName        string
	DeviceFingerPrint string
	Address           base.Address
	HomePhone         string
	WorkPhone         string
	Fax               string
	MobilePhone       string
	Email             string
	Note              string
	Comments          string
	Department        string
	Status            string
	Phone             base.PhoneNumber
}

func NewCustomer() *Customer {
	return &Customer{RecurringEntity: &RecurringEntity{}}
}

func NewCustomerWithID(id string) *Customer {
	c := NewCustomer()
	c.ID = id
	return c
}
