package validations

import (
	"errors"
	"reflect"
)

type IRuleSet map[string]map[int64][]*ValidationTarget

type Validations struct {
	rules IRuleSet
}

func NewValidations() *Validations {
	return &Validations{
		rules: make(IRuleSet),
	}
}

func (v *Validations) OfGeneral(enumProperty string, typeVal int64) *ValidationTarget {
	if _, ok := v.rules[enumProperty]; !ok {
		v.rules[enumProperty] = make(map[int64][]*ValidationTarget)
	}

	if _, ok := v.rules[enumProperty][typeVal]; !ok {
		v.rules[enumProperty][typeVal] = []*ValidationTarget{}
	}

	target := NewValidationTarget(v, enumProperty, typeVal)
	v.rules[enumProperty][typeVal] = append(v.rules[enumProperty][typeVal], target)
	return target
}

func (v *Validations) Of(typeVal int64) *ValidationTarget {
	return v.OfGeneral("TransactionType", typeVal)
}

func (v *Validations) Validate(builder interface{}) error {
	for enumName, rules := range v.rules {
		value, ok := reflect.ValueOf(builder).Elem().FieldByName(enumName).Interface().(int64)
		if !ok && reflect.ValueOf(builder).Elem().FieldByName("paymentMethod").IsValid() {
			value, _ = reflect.ValueOf(builder).Elem().FieldByName("paymentMethod").Interface().(map[string]int64)[enumName]
		}

		for iKey, validationRules := range rules {
			if (iKey & value) != value {
				continue
			}

			for _, validation := range validationRules {
				if validation.clause == nil {
					continue
				}

				if validation.constraint.LongValue() != 0 && validation.constraint != reflect.ValueOf(builder).Elem().FieldByName(validation.constraint.GetStringValue()).Interface() {
					continue
				}

				if !validation.clause.callback(builder) {
					return errors.New(validation.clause.message)
				}
			}
		}
	}
	return nil
}
