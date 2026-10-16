package sortdirection

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type SortDirection string

const (
	Ascending  SortDirection = "ASC"
	Descending SortDirection = "DESC"
)

// GetBytes returns the byte representation of the SortDirection value
func (sd SortDirection) GetBytes() []byte {
	return []byte(sd)
}

// GetValue returns the string value of the SortDirection
func (sd SortDirection) GetValue() string {
	return string(sd)
}

// StringConstants returns a slice of all SortDirection constants
func (SortDirection) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{Ascending, Descending}
}
