package base

import (
	"fmt"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
)

type PhoneNumber struct {
	CountryCode string
	AreaCode    string
	Number      string
	Extension   string
}

func (p *PhoneNumber) ToString() string {
	if stringutils.IsNullOrEmpty(p.CountryCode) {
		p.CountryCode = "1"
	}

	result := "+" + p.CountryCode

	if !stringutils.IsNullOrEmpty(p.AreaCode) {
		result += fmt.Sprintf("(%s)", p.AreaCode)
	}

	result += p.Number

	if !stringutils.IsNullOrEmpty(p.Extension) {
		result += fmt.Sprintf("EXT: %s", p.Extension)
	}

	return result
}
