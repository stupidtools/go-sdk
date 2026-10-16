package stringutils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func IsNullOrEmpty(value string) bool {
	return value == "" || strings.TrimSpace(value) == ""
}

func PadLeft(input interface{}, totalLength int, paddingCharacter rune) string {
	var str string
	if input == nil {
		str = ""
	} else {
		str = input.(string)
	}

	for len(str) < totalLength {
		str = string(paddingCharacter) + str
	}
	return str
}

func PadRight(input string, totalLength int, paddingCharacter rune) string {
	for len(input) < totalLength {
		input = input + string(paddingCharacter)
	}
	return input
}

func ToDecimalAmount(str string) (*decimal.Decimal, error) {
	if IsNullOrEmpty(str) {
		return nil, errors.New("invalid amount")
	}

	amount, err := decimal.NewFromString(str)
	if err != nil {
		return nil, errors.New("invalid amount")
	}
	return &amount, nil
}

func ToAmount(str string) (*decimal.Decimal, error) {
	if IsNullOrEmpty(str) {
		return nil, errors.New("invalid amount")
	}

	amount, err := decimal.NewFromString(str)
	if err != nil {
		return nil, errors.New("invalid amount")
	}
	res := amount.Div(decimal.NewFromInt(100))
	return &res, nil
}

func ToFractionalAmount(str string) *decimal.Decimal {
	if IsNullOrEmpty(str) {
		return nil
	}

	numDecimals, _ := strconv.Atoi(string(str[0]))
	shiftValue, _ := strconv.Atoi(PadRight("1", numDecimals+1, '0'))

	qty, _ := decimal.NewFromString(str[1:])
	res := qty.Div(decimal.NewFromInt(int64(shiftValue)))
	return &res
}

func ToString(value string, decimalPlace int) *decimal.Decimal {
	value = strings.TrimSpace(value)
	value = value[:len(value)-decimalPlace] + "." + value[len(value)-decimalPlace:]
	floatValue, _ := decimal.NewFromString(value)
	return &floatValue
}

func ToNumeric(str string) string {
	return ExtractDigits(str)
}

func ToNumericFromAmount(amount *decimal.Decimal) string {
	if amount == nil {
		return ""
	} else if amount.Cmp(decimal.NewFromInt(0)) == 0 {
		return "000"
	}

	currency := amount.Mul(decimal.NewFromInt(100)).String()
	return TrimStart(ExtractDigits(currency), "0")
}

func ToNumericWithLength(amount *decimal.Decimal, length int) string {
	rvalue := ToNumericFromAmount(amount)
	return PadLeft(rvalue, length, '0')
}

func ToDecimal(amount *decimal.Decimal, length int) string {
	if amount == nil {
		return ""
	}

	formattedAmount := amount.StringFixed(3)
	formattedAmount = strings.ReplaceAll(formattedAmount, ".", "")
	formattedAmount = TrimStart(formattedAmount, "0")
	return PadLeft(formattedAmount, length, '0')
}

func TrimStart(str, trimString string) string {
	rvalue := str
	for strings.HasPrefix(rvalue, trimString) {
		rvalue = rvalue[len(trimString):]
	}
	return rvalue
}

func ExtractDigits(str string) string {
	regex := regexp.MustCompile("[^0-9]")
	return regex.ReplaceAllString(str, "")
}

func ToCurrencyString(amount *decimal.Decimal) string {
	if amount == nil {
		return ""
	}
	if amount.Cmp(decimal.NewFromInt(0)) == 0 {
		return "0"
	}

	formatted := amount.StringFixed(2)
	parts := strings.Split(formatted, ".")
	dollars := parts[0]
	cents := parts[1]

	var buf bytes.Buffer
	if amount.Sign() < 0 {
		buf.WriteRune('-')
		dollars = dollars[1:]
	}
	for i := len(dollars) - 1; i >= 0; i-- {
		//if (len(dollars)-i)%3 == 0 && i != len(dollars)-1 {
		//	buf.WriteRune(',')
		//}
		buf.WriteRune(rune(dollars[i]))
	}
	dollarsStr := reverseString(buf.String())
	return fmt.Sprintf("%s.%s", dollarsStr, cents)
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func BoolToString(b bool) string {
	if b {
		return "Y"
	} else {
		return "N"
	}
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToIntPointer(s string) *int {
	res, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &res
}

func DateStringFormatted(d string, format string) string {
	t, err := dateparse.ParseAny(d)
	if err != nil {
		return ""
	}
	return t.Format(format)
}

func ToStandardDateString(d time.Time) string {
	return d.Format(time.RFC3339)
}
