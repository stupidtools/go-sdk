package abstractions

import "github.com/globalpayments/go-sdk/api/builders/validations"

type IBaseBuilder interface {
	GetValidations() validations.Validations
	SetValidations(validations validations.Validations)
}
