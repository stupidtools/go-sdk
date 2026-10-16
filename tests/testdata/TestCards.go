package testdata

import "github.com/globalpayments/go-sdk/api/paymentmethods"

func VisaManual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("4012002000060016")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// MasterCard5Manual creates a test Mastercard (series 5) card
func MasterCard5Manual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("5473500000000014")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// MasterCard5Manual2 creates an alternate test Mastercard (series 5) card
func MasterCard5Manual2(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("5569992222222222225")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// MasterCard2Manual creates a test Mastercard (series 2) card
func MasterCard2Manual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("2223000010005780")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("900")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// MasterCard2Manual2 creates an alternate test Mastercard (series 2) card
func MasterCard2Manual2(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("2223000010005798")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("988")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// AmexManual creates a test American Express card
func AmexManual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("372700699251018")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("1234")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// DiscoverManual creates a test Discover card
func DiscoverManual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("6011000990156527")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// JcbManual creates a test JCB card
func JcbManual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("3566007770007321")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

// DiscoverPaypalManual creates a test Discover PayPal card
func DiscoverPaypalManual(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("6506001000010029")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}
