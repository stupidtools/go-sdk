package validations

import (
	"fmt"
	"reflect"
)

type ValidationClause struct {
	parent       *Validations
	target       *ValidationTarget
	callback     func(interface{}) bool
	message      string
	precondition bool
}

func NewValidationClause(parent *Validations, target *ValidationTarget, precondition bool) *ValidationClause {
	return &ValidationClause{
		parent:       parent,
		target:       target,
		precondition: precondition,
	}
}

func (vc *ValidationClause) IsNotNull(message string) *ValidationTarget {
	vc.callback = func(builder interface{}) bool {
		value := reflect.ValueOf(builder).Elem().FieldByName(vc.target.property)
		return value.IsValid() && !value.IsNil()
	}
	if message != "" {
		vc.message = message
	} else {
		vc.message = vc.target.property + " cannot be null for this transaction type."
	}

	if vc.precondition {
		return vc.target
	}

	return vc.parent.OfGeneral(vc.target.enumName, vc.target.typ).With(vc.target.constraint)
}

func (vc *ValidationClause) IsNull(message string) *ValidationTarget {
	vc.callback = func(builder interface{}) bool {
		value := reflect.ValueOf(builder).Elem().FieldByName(vc.target.property)
		return !value.IsValid() || value.IsNil()
	}
	if message != "" {
		vc.message = message
	} else {
		vc.message = vc.target.property + " cannot be set for this transaction type."
	}

	if vc.precondition {
		return vc.target
	}

	return vc.parent.OfGeneral(vc.target.enumName, vc.target.typ).With(vc.target.constraint)
}

func (vc *ValidationClause) IsNotEmpty(message string) *ValidationTarget {
	vc.callback = func(builder interface{}) bool {
		value := reflect.ValueOf(builder).Elem().FieldByName(vc.target.property)
		return !value.IsZero()
	}
	if message != "" {
		vc.message = message
	} else {
		vc.message = vc.target.property + " cannot be empty for this transaction type."
	}

	if vc.precondition {
		return vc.target
	}

	return vc.parent.OfGeneral(vc.target.enumName, vc.target.typ).With(vc.target.constraint)
}

func (vc *ValidationClause) IsEqual(expected interface{}, message string) *ValidationTarget {
	vc.callback = func(builder interface{}) bool {
		value := reflect.ValueOf(builder).Elem().FieldByName(vc.target.property)
		return value.Interface() == expected
	}
	if message != "" {
		vc.message = message
	} else {
		vc.message = fmt.Sprintf("%s was not the expected value %v", vc.target.property, expected)
	}

	if vc.precondition {
		return vc.target
	}

	return vc.parent.OfGeneral(vc.target.enumName, vc.target.typ).With(vc.target.constraint)
}

func (vc *ValidationClause) IsNotEqual(expected interface{}, message string) *ValidationTarget {
	vc.callback = func(builder interface{}) bool {
		value := reflect.ValueOf(builder).Elem().FieldByName(vc.target.property)
		return value.Interface() != expected
	}
	if message != "" {
		vc.message = message
	} else {
		vc.message = fmt.Sprintf("%s cannot be the value %v.", vc.target.property, expected)
	}

	if vc.precondition {
		return vc.target
	}

	return vc.parent.OfGeneral(vc.target.enumName, vc.target.typ).With(vc.target.constraint)
}
