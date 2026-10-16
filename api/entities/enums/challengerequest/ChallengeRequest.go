package challengerequest

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type ChallengeRequest string

const (
	NoPreference         ChallengeRequest = "NO_PREFERENCE"
	NoChallengeRequested ChallengeRequest = "NO_CHALLENGE_REQUESTED"
	ChallengePreferred   ChallengeRequest = "CHALLENGE_PREFERRED"
	ChallengeMandated    ChallengeRequest = "CHALLENGE_MANDATED"
)

func (c ChallengeRequest) GetBytes() []byte {
	return []byte(c)
}

func (c ChallengeRequest) GetValue() string {
	return string(c)
}

func (c ChallengeRequest) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{NoPreference, NoChallengeRequested, ChallengePreferred, ChallengeMandated}
}
