package storedcredentialinitiator

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
)

type StoredCredentialInitiator string

const (
	Nil         StoredCredentialInitiator = ""
	CardHolder  StoredCredentialInitiator = "cardholder"
	Merchant    StoredCredentialInitiator = "merchant"
	Scheduled   StoredCredentialInitiator = "scheduled"
	Installment StoredCredentialInitiator = "installment"
)

var storedCredentialInitiatorMap = map[StoredCredentialInitiator]map[target.Target]string{
	CardHolder: {
		target.Realex:  "cardholder",
		target.Portico: "C",
		target.GP_API:  "PAYER",
		target.Genius:  "UNSCHEDULEDCIT",
	},
	Merchant: {
		target.Realex:  "merchant",
		target.Portico: "M",
		target.GP_API:  "MERCHANT",
		target.Genius:  "UNSCHEDULEDMIT",
	},
	Scheduled: {
		target.Realex: "scheduled",
		target.Genius: "RECURRING",
	},
	Installment: {
		target.Genius: "INSTALLMENT",
	},
}

func (sci StoredCredentialInitiator) GetBytes(t target.Target) []byte {
	if value, ok := storedCredentialInitiatorMap[sci][t]; ok {
		return []byte(value)
	}
	return nil
}

func (sci StoredCredentialInitiator) GetValue(t target.Target) string {
	if value, ok := storedCredentialInitiatorMap[sci][t]; ok {
		return value
	}
	return ""
}
