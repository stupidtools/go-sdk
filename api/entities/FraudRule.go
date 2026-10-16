package entities

import "github.com/globalpayments/go-sdk/api/entities/enums/fraudfiltermode"

type FraudRule struct {
	Key         string
	Mode        fraudfiltermode.FraudFilterMode
	Description string
	Result      string
}

func NewFraudRule() *FraudRule {
	return &FraudRule{}
}

func NewFraudRuleWithParams(key string, mode fraudfiltermode.FraudFilterMode) *FraudRule {
	return &FraudRule{
		Key:  key,
		Mode: mode,
	}
}

func NewFraudRuleWithAllParams(key string, mode fraudfiltermode.FraudFilterMode, description string, result string) *FraudRule {
	return &FraudRule{
		Key:         key,
		Mode:        mode,
		Description: description,
		Result:      result,
	}
}
