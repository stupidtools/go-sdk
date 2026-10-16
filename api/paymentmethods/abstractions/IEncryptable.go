package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities/base"
)

type IEncryptable interface {
	GetEncryptionData() *base.EncryptionData
	SetEncryptionData(encryptionData *base.EncryptionData)
	GetEncryptedPan() string
	SetEncryptedPan(encryptedPan string)
}
