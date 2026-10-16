package storedcredentialsequence

import "github.com/globalpayments/go-sdk/api/entities/enums/target"

type StoredCredentialSequence struct {
	value map[target.Target]string
}

var (
	First = StoredCredentialSequence{
		value: map[target.Target]string{
			target.Realex: "first",
			target.GP_API: "FIRST",
		},
	}
	Subsequent = StoredCredentialSequence{
		value: map[target.Target]string{
			target.Realex: "subsequent",
			target.GP_API: "SUBSEQUENT",
		},
	}
	Last = StoredCredentialSequence{
		value: map[target.Target]string{
			target.GP_API: "LAST",
		},
	}
)

func (sc StoredCredentialSequence) GetValue(target target.Target) string {
	if val, ok := sc.value[target]; ok {
		return val
	}
	return ""
}
