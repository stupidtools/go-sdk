package entities

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/ecommercechannel"
	"time"
)

type EcommerceInfo struct {
	channel   ecommercechannel.EcommerceChannel
	shipDay   *int
	shipMonth *int
}

func NewEcommerceInfo(c ecommercechannel.EcommerceChannel) *EcommerceInfo {
	tomorrow := time.Now().AddDate(0, 0, 1)
	day := tomorrow.Day()
	month := int(tomorrow.Month())
	return &EcommerceInfo{
		channel:   c,
		shipDay:   &day,
		shipMonth: &month,
	}
}

func (e *EcommerceInfo) GetChannel() ecommercechannel.EcommerceChannel {
	return e.channel
}

func (e *EcommerceInfo) SetChannel(channel ecommercechannel.EcommerceChannel) {
	e.channel = channel
}

func (e *EcommerceInfo) GetShipDay() *int {
	return e.shipDay
}

func (e *EcommerceInfo) SetShipDay(shipDay int) {
	e.shipDay = &shipDay
}

func (e *EcommerceInfo) GetShipMonth() *int {
	return e.shipMonth
}

func (e *EcommerceInfo) SetShipMonth(shipMonth int) {
	e.shipMonth = &shipMonth
}
