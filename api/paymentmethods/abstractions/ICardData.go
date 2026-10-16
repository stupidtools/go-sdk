package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/cvnpresenceindicator"
	"github.com/globalpayments/go-sdk/api/entities/enums/manualentrymethod"
)

type ICardData interface {
	IsCardPresent() bool
	SetCardPresent(cardPresent bool)

	GetCardType() string
	SetCardType(cardType string)

	GetCardHolderName() string
	SetCardHolderName(cardHolderName string)

	GetCvn() string
	SetCvn(cvn string)

	GetCvnPresenceIndicator() cvnpresenceindicator.CvnPresenceIndicator
	SetCvnPresenceIndicator(cvnPresenceIndicator cvnpresenceindicator.CvnPresenceIndicator)

	GetNumber() string
	SetNumber(number string)

	GetExpMonth() *int
	SetExpMonth(expMonth *int)

	GetExpYear() *int
	SetExpYear(expYear *int)

	IsReaderPresent() bool
	SetReaderPresent(readerPresent bool)

	GetShortExpiry() string

	GetEntryMethod() manualentrymethod.ManualEntryMethod
	SetEntryMethod(manualEntryMethod manualentrymethod.ManualEntryMethod)

	GetTokenizationData() string
	SetTokenizationData(s string)
}
