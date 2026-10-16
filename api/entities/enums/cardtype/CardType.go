package cardtype

type CardType string

const (
	VISA            CardType = "VISA"
	MC              CardType = "MC"
	DISC            CardType = "DISC"
	AMEX            CardType = "AMEX"
	GIFTCARD        CardType = "GIFTCARD"
	PAYPALECOMMERCE CardType = "PAYPALECOMMERCE"
)
