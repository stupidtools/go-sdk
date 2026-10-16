package countrycodeformat

type CountryCodeFormat string

const (
	Alpha2  CountryCodeFormat = "Alpha2"
	Alpha3  CountryCodeFormat = "Alpha3"
	Numeric CountryCodeFormat = "Numeric"
	Name    CountryCodeFormat = "Name"
	Null    CountryCodeFormat = ""
)
