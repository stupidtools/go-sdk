package cvnpresenceindicator

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type CvnPresenceIndicator string

const (
	Present      CvnPresenceIndicator = "1"
	Illegible    CvnPresenceIndicator = "2"
	NotOnCard    CvnPresenceIndicator = "3"
	NotRequested CvnPresenceIndicator = "4"
)

// GetBytes returns the byte representation of the CVN presence indicator.
func (c CvnPresenceIndicator) GetBytes() []byte {
	return []byte(c)
}

// GetValue returns the string value of the CVN presence indicator.
func (c CvnPresenceIndicator) GetValue() string {
	return string(c)
}

// StringConstants returns a slice of all CVN presence indicators.
func (c CvnPresenceIndicator) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{Present, Illegible, NotOnCard, NotRequested}
}
