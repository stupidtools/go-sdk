package entities

import (
	"github.com/shopspring/decimal"
)

type AdditionalDuplicateData struct {
	OriginalGatewayTxnId string
	OriginalRspDT        string
	OriginalClientTxnId  string
	OriginalAuthCode     string
	OriginalRefNbr       string
	OriginalAuthAmt      *decimal.Decimal
	OriginalCardType     string
	OriginalCardNbrLast4 string
}

func NewAdditionalDuplicateData() *AdditionalDuplicateData {
	return &AdditionalDuplicateData{}
}
