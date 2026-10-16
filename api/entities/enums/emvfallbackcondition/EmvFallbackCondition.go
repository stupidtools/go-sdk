package emvfallbackcondition

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type EmvFallbackCondition string

// Defined EmvFallbackCondition values
const (
	ChipReadFailure EmvFallbackCondition = "ICC_TERMINAL_ERROR"
	NoCandidateList EmvFallbackCondition = "NO_CANDIDATE_LIST"
)

// GetBytes returns the byte slice of the EmvFallbackCondition string
func (e EmvFallbackCondition) GetBytes() []byte {
	return []byte(e)
}

// GetValue returns the string value of the EmvFallbackCondition
func (e EmvFallbackCondition) GetValue() string {
	return string(e)
}

// StringConstants returns a slice of IStringConstant that includes
// all the defined values of EmvFallbackCondition
func (e EmvFallbackCondition) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{ChipReadFailure, NoCandidateList}
}
