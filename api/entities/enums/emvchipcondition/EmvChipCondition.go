package emvchipcondition

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type EmvChipCondition string

const (
	ChipFailPreviousSuccess EmvChipCondition = "CHIP_FAILED_PREV_SUCCESS"
	ChipFailPreviousFail    EmvChipCondition = "CHIP_FAILED_PREV_FAILED"
)

func (e EmvChipCondition) GetBytes() []byte {
	return []byte(e)
}

func (e EmvChipCondition) GetValue() string {
	return string(e)
}

func (e EmvChipCondition) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		ChipFailPreviousSuccess,
		ChipFailPreviousFail,
	}
}
