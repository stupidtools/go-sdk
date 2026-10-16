package abstractions

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/entrymethod"
	"github.com/globalpayments/go-sdk/api/entities/enums/tracknumber"
)

type ITrackData interface {
	GetExpiry() string
	SetExpiry(string)

	GetPan() string
	SetPan(string)

	GetTrackNumber() tracknumber.TrackNumber
	SetTrackNumber(number tracknumber.TrackNumber)

	GetTrackData() string
	SetTrackData(string)

	GetDiscretionaryData() string
	SetDiscretionaryData(string)

	GetValue() string
	SetValue(string)

	GetEntryMethod() entrymethod.EntryMethod
	SetEntryMethod(entrymethod.EntryMethod)

	GetTruncatedTrackData() string
}
