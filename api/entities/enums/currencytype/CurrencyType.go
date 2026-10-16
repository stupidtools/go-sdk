package currencytype

type CurrencyType string

const (
	Currency     CurrencyType = "USD"
	Points       CurrencyType = "POINTS"
	CashBenefits CurrencyType = "CASH_BENEFITS"
	FoodStamps   CurrencyType = "FOODSTAMPS"
	Voucher      CurrencyType = "VOUCHER"
)
