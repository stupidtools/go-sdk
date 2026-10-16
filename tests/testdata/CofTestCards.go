package testdata

import "github.com/globalpayments/go-sdk/api/paymentmethods"

// AmexManual creates a test American Express card
func CofA1Card(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("374101000000608")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("1234")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

func CofM1Card(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("5114610000004778")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

func CofM2Card(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("5425230000004415")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

func CofV1Card(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("4263970000005262")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

func CofV2Card(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("4000120000001154")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}

func CofD1Card(cardPresent bool, readerPresent bool) *paymentmethods.CreditCardData {
	card := paymentmethods.NewCreditCardData()
	card.SetNumber("6011000990191250")
	exMonth := 12
	exYear := 2025
	card.SetExpMonth(&exMonth)
	card.SetExpYear(&exYear)
	card.SetCvn("123")
	card.SetCardPresent(cardPresent)
	card.SetReaderPresent(readerPresent)
	return card
}
