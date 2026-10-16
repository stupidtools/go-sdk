package subgroups

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upaacquisitiontype"
	"github.com/globalpayments/go-sdk/api/utils"
	"strconv"
)

type RequestParamFields struct {
	acquisitionTypes       []upaacquisitiontype.UpaAcquisitionType
	acquisitionTypesString string
	cardBrandStorage       storedcredentialinitiator.StoredCredentialInitiator
	cardBrandTransactionId string
	clerkId                string
	tokenRequest           *int
	tokenValue             string
	showIfEmpty            bool
}

func NewRequestParamFields() *RequestParamFields {
	return &RequestParamFields{}
}

func (rpf *RequestParamFields) SetManageBuilderParams(builder *builders.TerminalManageBuilder) {
	if (builder.GetTransactionType() == transactiontype.Edit) && (builder.GetGratuity() != nil) {
		rpf.showIfEmpty = true
	}
}

func (rpf *RequestParamFields) SetAuthBuilderParams(builder *builders.TerminalAuthBuilder) {

	if builder.RequestMultiUseToken {
		request := 1
		rpf.tokenRequest = &request
	}

	if builder.TokenValue != "" {
		rpf.tokenValue = builder.TokenValue
	}

	if builder.CardBrandStorage != "" {
		rpf.cardBrandStorage = builder.CardBrandStorage
	}

	if builder.CardBrandTransactionId != "" {
		rpf.cardBrandTransactionId = builder.CardBrandTransactionId
	}

	if builder.GetClerkId() != nil {
		cid := *builder.GetClerkId()
		rpf.clerkId = strconv.Itoa(cid)
	}

	if builder.GetTransactionType() == transactiontype.Activate {
		if len(rpf.acquisitionTypes) != 0 {
			// handle integration-supplied list
		} else {
			rpf.acquisitionTypes = []upaacquisitiontype.UpaAcquisitionType{
				upaacquisitiontype.Contact,
				upaacquisitiontype.Contactless,
				upaacquisitiontype.Manual,
				upaacquisitiontype.Scan,
				upaacquisitiontype.Swipe,
			}
		}
	}
}

func (rpf *RequestParamFields) GetElementsJson() *utils.JsonDoc {
	params := utils.NewJsonDoc()
	hasContents := false
	if rpf.showIfEmpty {
		hasContents = true
	}
	if rpf.cardBrandStorage != "" {
		// only two values supported per v1.30 integrator's guide
		if rpf.cardBrandStorage == storedcredentialinitiator.Merchant {
			params.Set("cardOnFileIndicator", "M", true)
		}
		if rpf.cardBrandStorage == storedcredentialinitiator.CardHolder {
			params.Set("cardOnFileIndicator", "C", true)
		}
		hasContents = true
	}

	if rpf.cardBrandTransactionId != "" {
		params.Set("cardBrandTransId", rpf.cardBrandTransactionId, true)
		hasContents = true
	}

	if rpf.clerkId != "" {
		params.Set("clerkId", rpf.clerkId, true)
		hasContents = true
	}

	if rpf.tokenRequest != nil {
		params.SetInt("tokenRequest", rpf.tokenRequest)
		hasContents = true
	}

	if rpf.tokenValue != "" {
		params.Set("tokenValue", rpf.tokenValue, true)
		hasContents = true
	}

	if rpf.acquisitionTypes != nil && len(rpf.acquisitionTypes) > 0 {
		rpf.acquisitionTypesString = ""

		for _, at := range rpf.acquisitionTypes {
			rpf.acquisitionTypesString += string(at)
			rpf.acquisitionTypesString += "|"
		}

		rpf.acquisitionTypesString = rpf.acquisitionTypesString[:len(rpf.acquisitionTypesString)-1]
		params.Set("acquisitionTypes", rpf.acquisitionTypesString, true)
		hasContents = true
	}

	if !hasContents {
		return nil
	}

	return params
}

func (rpf *RequestParamFields) GetElementString() string {
	return ""
}
