package storedcredentialreason

import (
	"errors"
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
)

type StoredCredentialReason struct {
	value map[target.Target]string
}

var (
	Incremental     = StoredCredentialReason{value: map[target.Target]string{target.GP_API: "INCREMENTAL"}}
	Resubmission    = StoredCredentialReason{value: map[target.Target]string{target.GP_API: "RESUBMISSION"}}
	Reauthorization = StoredCredentialReason{value: map[target.Target]string{target.GP_API: "REAUTHORIZATION"}}
	Delayed         = StoredCredentialReason{value: map[target.Target]string{target.GP_API: "DELAYED"}}
	NoShow          = StoredCredentialReason{value: map[target.Target]string{target.GP_API: "NO_SHOW"}}
)

func (scr StoredCredentialReason) GetValue(target target.Target) string {
	if val, ok := scr.value[target]; ok {
		return val
	}
	return ""
}

func (scr StoredCredentialReason) GetBytes(target target.Target) ([]byte, error) {
	if val, ok := scr.value[target]; ok {
		return []byte(val), nil
	}
	return nil, errors.New("target not found")
}
