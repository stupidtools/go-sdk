package storedcredentialtype

import (
	"errors"
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
)

type StoredCredentialType struct {
	value map[target.Target]string
}

func NewStoredCredentialType(values map[target.Target]string) *StoredCredentialType {
	return &StoredCredentialType{
		value: values,
	}
}

var (
	OneOff                      = NewStoredCredentialType(map[target.Target]string{target.Realex: "oneoff"})
	Installment                 = NewStoredCredentialType(map[target.Target]string{target.Realex: "installment", target.GP_API: "INSTALLMENT"})
	Recurring                   = NewStoredCredentialType(map[target.Target]string{target.Realex: "recurring", target.GP_API: "RECURRING"})
	Unscheduled                 = NewStoredCredentialType(map[target.Target]string{target.GP_API: "UNSCHEDULED"})
	Subscription                = NewStoredCredentialType(map[target.Target]string{target.GP_API: "SUBSCRIPTION"})
	MaintainPaymentMethod       = NewStoredCredentialType(map[target.Target]string{target.GP_API: "MAINTAIN_PAYMENT_METHOD"})
	MaintainPaymentVerification = NewStoredCredentialType(map[target.Target]string{target.GP_API: "MAINTAIN_PAYMENT_VERIFICATION"})
)

func (sct *StoredCredentialType) GetValue(target target.Target) string {
	if val, ok := sct.value[target]; ok {
		return val
	}
	return ""
}

func (sct *StoredCredentialType) GetBytes(target target.Target) ([]byte, error) {
	if val, ok := sct.value[target]; ok {
		return []byte(val), nil
	}
	return nil, errors.New("target not found")
}
