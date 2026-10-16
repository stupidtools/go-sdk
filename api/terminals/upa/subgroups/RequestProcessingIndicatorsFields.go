package subgroups

import (
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upaacquisitiontype"
	"github.com/globalpayments/go-sdk/api/utils"
)

type RequestProcessingIndicatorsFields struct {
	AcquisitionTypes []upaacquisitiontype.UpaAcquisitionType
	QuickChip        bool
	CheckLuhn        bool
}

func NewRequestProcessingIndicatorsFields() *RequestProcessingIndicatorsFields {
	return &RequestProcessingIndicatorsFields{}
}

func (r *RequestProcessingIndicatorsFields) GetElementsJson() *utils.JsonDoc {
	params := utils.NewJsonDoc()

	// will use the most common settings for these parameters for the time being
	// will later add builder methods to set these values
	params.Set("quickChip", "Y", true)
	params.Set("checkLuhn", "N", true)

	return params
}

func (r *RequestProcessingIndicatorsFields) GetElementString() string {
	return ""
}
