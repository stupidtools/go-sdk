package validations

import "github.com/globalpayments/go-sdk/api/entities/enums/iflag"

type ValidationTarget struct {
	parent       *Validations
	typ          int64
	property     string
	clause       *ValidationClause
	constraint   iflag.IFlag
	enumName     string
	precondition *ValidationClause
}

func NewValidationTarget(parent *Validations, enumName string, typ int64) *ValidationTarget {
	return &ValidationTarget{
		parent:   parent,
		typ:      typ,
		enumName: enumName,
	}
}

func (v *ValidationTarget) With(property iflag.IFlag) *ValidationTarget {
	v.constraint = property
	return v
}

func (v *ValidationTarget) Check(targetProperty string) *ValidationClause {
	v.property = targetProperty
	v.clause = NewValidationClause(v.parent, v, false)
	return v.clause
}

func (v *ValidationTarget) When(targetProperty string) *ValidationClause {
	v.property = targetProperty
	v.precondition = NewValidationClause(v.parent, v, true)
	return v.precondition
}
