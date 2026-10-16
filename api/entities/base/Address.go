package base

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/addresstype"
	"github.com/globalpayments/go-sdk/api/entities/enums/countrycodeformat"
	"github.com/globalpayments/go-sdk/api/utils/countryutils"
)

type Address struct {
	Type        addresstype.AddressType
	StreetAddr1 string
	StreetAddr2 string
	StreetAddr3 string
	City        string
	Name        string
	Province    string
	PostalCode  string
	Country     string
	CountryCode string
}

func (a *Address) GetState() string {
	return a.Province
}

func (a *Address) SetState(province string) *Address {
	a.Province = province
	return a
}

func (a *Address) SetCountryCode(countryCode string) *Address {
	a.CountryCode = countryCode
	if a.Country == "" {
		a.Country = countryutils.GetCountryByCode(countryCode)
	}

	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.Country = country
	if a.CountryCode == "" {
		a.CountryCode = countryutils.GetCountryCodeByCountry(country, countrycodeformat.Null)
	}

	return a
}

func (address *Address) IsCountry(countryCode string) bool {
	if address.CountryCode != "" {
		return address.CountryCode == countryCode
	} else if address.Country != "" {
		code := countryutils.GetCountryCodeByCountry(address.Country, countrycodeformat.Null)
		if code != "" {
			return code == countryCode
		}
		return false
	}
	return false
}

func NewAddress(postalCode string) *Address {
	return NewAddressWithStreet(postalCode, "")
}

func NewAddressWithStreet(postalCode, streetAddress1 string) *Address {
	return &Address{
		StreetAddr1: streetAddress1,
		PostalCode:  postalCode,
	}
}
