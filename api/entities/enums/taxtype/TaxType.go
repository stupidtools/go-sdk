package taxtype

type TaxType string

const (
	NotUsed   TaxType = "NOTUSED"
	SalesTax  TaxType = "SALESTAX"
	TaxExempt TaxType = "TAXEXEMPT"
)

func (tt TaxType) GetValue() string {
	return string(tt)
}

func (tt TaxType) GetBytes() []byte {
	return []byte(tt)
}
