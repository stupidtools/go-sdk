package countryutils

import (
	"math"
	"strings"

	"github.com/globalpayments/go-sdk/api/entities/countrydata"
	"github.com/globalpayments/go-sdk/api/entities/enums/countrycodeformat"
)

const (
	significantCountryMatch = 6
	significantCodeMatch    = 3
)

func GetCountryByCode(countryCode string) string {
	if countryCode == "" {
		return ""
	}
	output := ""

	if isAlpha2(countryCode) {
		output = convertFromAlpha2(countryCode, countrycodeformat.Name)
	} else if isAlpha3(countryCode) {
		output = convertFromAlpha3(countryCode, countrycodeformat.Name)
	} else if isNumeric(countryCode) {
		output = convertFromNumeric(countryCode, countrycodeformat.Name)
	}
	if output != "" {
		return output
	} else {
		if len(countryCode) > 3 {
			return ""
		}
		return fuzzyMatch(countrydata.GetCountryByAlpha2Code(), countryCode, significantCodeMatch)
	}
}

func GetCountryCodeByCountry(country string, format countrycodeformat.CountryCodeFormat) string {
	if format == "" {
		format = countrycodeformat.Alpha2
	}

	output := ""

	if country == "" {
		return ""
	}

	if isCountryName(country) {
		output = convertFromName(country, format)
	} else if isAlpha2(country) {
		output = convertFromAlpha2(country, format)
	} else if isAlpha3(country) {
		output = convertFromAlpha3(country, format)
	} else if isNumeric(country) {
		output = convertFromNumeric(country, format)
	}

	if output != "" {
		return output
	} else {
		return fuzzyByFormat(format, country)
	}
}

func fuzzyByFormat(format countrycodeformat.CountryCodeFormat, country string) string {
	fuzzyCountryMatch := ""
	output := ""

	if format == countrycodeformat.Alpha2 {
		fuzzyCountryMatch = fuzzyMatch(countrydata.GetAlpha2CodeByCountry(), country, significantCountryMatch)
	} else if format == countrycodeformat.Alpha3 {
		fuzzyCountryMatch = fuzzyMatch(countrydata.GetAlpha3CodeByCountry(), country, significantCountryMatch)
	} else if format == countrycodeformat.Numeric {
		fuzzyCountryMatch = fuzzyMatch(countrydata.GetNumericCodeByCountry(), country, significantCountryMatch)
	}

	if fuzzyCountryMatch != "" {
		return fuzzyCountryMatch
	}

	if len(country) > 3 {
		return ""
	}

	var fuzzyCodeMatch string
	if format == countrycodeformat.Alpha2 {
		fuzzyCodeMatch = fuzzyMatch(countrydata.GetCountryByAlpha2Code(), country, significantCodeMatch)
		if fuzzyCodeMatch != "" {
			output = countrydata.GetAlpha2CodeByCountry()[fuzzyCodeMatch]
		}
	} else if format == countrycodeformat.Alpha3 {
		fuzzyCodeMatch = fuzzyMatch(countrydata.GetCountryByAlpha3Code(), country, significantCodeMatch)
		if fuzzyCodeMatch != "" {
			output = countrydata.GetAlpha3CodeByCountry()[fuzzyCodeMatch]
		}
	} else if format == countrycodeformat.Numeric {
		fuzzyCodeMatch = fuzzyMatch(countrydata.GetCountryByNumericCode(), country, significantCodeMatch)
		if fuzzyCodeMatch != "" {
			output = countrydata.GetNumericCodeByCountry()[fuzzyCodeMatch]
		}
	}

	return output
}

func fuzzyMatch(dict map[string]string, query string, significantMatch int) string {
	var rvalue string
	matches := make(map[string]string)

	highScore := -1
	for key, value := range dict {
		score := fuzzyScore(key, query)
		if score > significantMatch && score > highScore {
			matches = make(map[string]string)

			highScore = score
			rvalue = value
			matches[key] = rvalue
		} else if score == highScore {
			matches[key] = value
		}
	}

	if len(matches) > 1 {
		return ""
	}
	return rvalue
}

func fuzzyScore(term, query string) int {
	if term == "" || query == "" {
		panic("Strings must not be null")
	}

	termLowerCase := strings.ToLower(term)
	queryLowerCase := strings.ToLower(query)

	score := 0
	termIndex := 0
	previousMatchingCharacterIndex := math.MinInt32

	for queryIndex := 0; queryIndex < len(queryLowerCase); queryIndex++ {
		queryChar := queryLowerCase[queryIndex]

		termCharacterMatchFound := false
		for termIndex < len(termLowerCase) && !termCharacterMatchFound {
			termChar := termLowerCase[termIndex]

			if queryChar == termChar {
				score++

				if previousMatchingCharacterIndex+1 == termIndex {
					score += 2
				}

				previousMatchingCharacterIndex = termIndex
				termCharacterMatchFound = true
			}
			termIndex++
		}
	}
	return score
}

// GetNumericCodeByCountry returns the numeric code for the given country
func GetNumericCodeByCountry(country string) string {
	numericCodeByCountry := countrydata.GetNumericCodeByCountry()
	numericByAlpha2CountryCode := countrydata.GetNumericByAlpha2CountryCode()
	numericByAlpha3CountryCode := countrydata.GetNumericByAlpha3CountryCode()

	if isCountryName(country) {
		if code, exists := numericCodeByCountry[country]; exists {
			return code
		}
	} else if isAlpha2(country) {
		if code, exists := numericByAlpha2CountryCode[country]; exists {
			return code
		}
	} else if isAlpha3(country) {
		if code, exists := numericByAlpha3CountryCode[country]; exists {
			return code
		}
	} else if isNumeric(country) {
		return country
	}

	return ""
}

// GetPhoneCodesByCountry returns the phone code for the given country
func GetPhoneCodesByCountry(country string) string {
	phoneCodeByCountry := countrydata.GetPhoneCodeByCountry()
	countryByNumericCode := countrydata.GetCountryByNumericCode()
	numericByAlpha2CountryCode := countrydata.GetNumericByAlpha2CountryCode()
	numericByAlpha3CountryCode := countrydata.GetNumericByAlpha3CountryCode()

	if isCountryName(country) {
		if code, exists := phoneCodeByCountry[country]; exists {
			return code
		}
	} else if isNumeric(country) {
		if countryCode, exists := countryByNumericCode[country]; exists {
			if code, exists := phoneCodeByCountry[countryCode]; exists {
				return code
			}
		}
	} else if isAlpha2(country) {
		if numericCode, exists := numericByAlpha2CountryCode[country]; exists {
			if countryCode, exists := countryByNumericCode[numericCode]; exists {
				if code, exists := phoneCodeByCountry[countryCode]; exists {
					return code
				}
			}
		}
	} else if isAlpha3(country) {
		if numericCode, exists := numericByAlpha3CountryCode[country]; exists {
			if countryCode, exists := countryByNumericCode[numericCode]; exists {
				if code, exists := phoneCodeByCountry[countryCode]; exists {
					return code
				}
			}
		}
	}

	return ""
}

func convertFromName(input string, countryCodeFormat countrycodeformat.CountryCodeFormat) string {
	if countryCodeFormat == countrycodeformat.Alpha2 && countrydata.GetAlpha2CodeByCountry()[input] != "" {
		return countrydata.GetAlpha2CodeByCountry()[input]
	} else if countryCodeFormat == countrycodeformat.Alpha3 && countrydata.GetAlpha3CodeByCountry()[input] != "" {
		return countrydata.GetAlpha3CodeByCountry()[input]
	} else if countryCodeFormat == countrycodeformat.Numeric && countrydata.GetNumericCodeByCountry()[input] != "" {
		return countrydata.GetNumericCodeByCountry()[input]
	} else if countryCodeFormat == countrycodeformat.Name {
		return input
	}
	return ""
}

func convertFromAlpha2(input string, countryCodeFormat countrycodeformat.CountryCodeFormat) string {
	if countryCodeFormat == countrycodeformat.Numeric && countrydata.GetNumericByAlpha2CountryCode()[input] != "" {
		return countrydata.GetNumericByAlpha2CountryCode()[input]
	} else if countryCodeFormat == countrycodeformat.Alpha3 && countrydata.GetAlpha3CodeByAlpha2Code()[input] != "" {
		return countrydata.GetAlpha3CodeByAlpha2Code()[input]
	} else if countryCodeFormat == countrycodeformat.Alpha2 {
		return input
	} else if countryCodeFormat == countrycodeformat.Name && countrydata.GetCountryByAlpha2Code()[input] != "" {
		return countrydata.GetCountryByAlpha2Code()[input]
	}
	return ""
}

func convertFromAlpha3(input string, countryCodeFormat countrycodeformat.CountryCodeFormat) string {
	if countryCodeFormat == countrycodeformat.Alpha2 && countrydata.GetAlpha2CodeByAlpha3Code()[input] != "" {
		return countrydata.GetAlpha2CodeByAlpha3Code()[input]
	} else if countryCodeFormat == countrycodeformat.Numeric && countrydata.GetNumericByAlpha3CountryCode()[input] != "" {
		return countrydata.GetNumericByAlpha3CountryCode()[input]
	} else if countryCodeFormat == countrycodeformat.Alpha3 {
		return input
	} else if countryCodeFormat == countrycodeformat.Name && countrydata.GetCountryByAlpha3Code()[input] != "" {
		return countrydata.GetCountryByAlpha3Code()[input]
	}
	return ""
}

func convertFromNumeric(input string, countryCodeFormat countrycodeformat.CountryCodeFormat) string {
	if countryCodeFormat == countrycodeformat.Alpha2 && countrydata.GetAlpha2CountryCodeByNumeric()[input] != "" {
		return countrydata.GetAlpha2CountryCodeByNumeric()[input]
	} else if countryCodeFormat == countrycodeformat.Alpha3 && countrydata.GetAlpha3CountryCodeByNumeric()[input] != "" {
		return countrydata.GetAlpha3CountryCodeByNumeric()[input]
	} else if countryCodeFormat == countrycodeformat.Numeric {
		return input
	} else if countryCodeFormat == countrycodeformat.Name && countrydata.GetCountryByNumericCode()[input] != "" {
		return countrydata.GetCountryByNumericCode()[input]
	}
	return ""
}

func isCountryName(input string) bool {
	return countrydata.GetAlpha2CodeByCountry()[input] != ""
}

func isAlpha2(input string) bool {
	return countrydata.GetCountryByAlpha2Code()[input] != ""
}

func isAlpha3(input string) bool {
	return countrydata.GetAlpha2CodeByAlpha3Code()[input] != ""
}

func isNumeric(input string) bool {
	return countrydata.GetAlpha2CountryCodeByNumeric()[input] != ""
}
