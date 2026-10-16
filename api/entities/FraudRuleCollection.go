package entities

import "github.com/globalpayments/go-sdk/api/entities/enums/fraudfiltermode"

type FraudRuleCollection struct {
	Rules []FraudRule
}

func NewFraudRuleCollection() *FraudRuleCollection {
	return &FraudRuleCollection{
		Rules: make([]FraudRule, 0),
	}
}

func (frc *FraudRuleCollection) AddRule(key string, mode fraudfiltermode.FraudFilterMode) {
	if frc.HasRule(key) {
		return
	}
	frc.Rules = append(frc.Rules, FraudRule{Key: key, Mode: mode})
}

func (frc *FraudRuleCollection) HasRule(key string) bool {
	for _, rule := range frc.Rules {
		if rule.Key == key {
			return true
		}
	}
	return false
}
