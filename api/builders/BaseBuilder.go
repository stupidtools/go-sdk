package builders

import "github.com/globalpayments/go-sdk/api/builders/validations"

type BaseBuilder struct {
	Validations validations.Validations
}

func NewBaseBuilder() *BaseBuilder {
	return &BaseBuilder{
		Validations: *validations.NewValidations(),
	}
}

func (b *BaseBuilder) GetValidations() validations.Validations {
	return b.Validations
}

func (b *BaseBuilder) SetValidations(validations validations.Validations) {
	b.Validations = validations
}

func (b *BaseBuilder) Execute(configName string) error {
	err := b.Validations.Validate(*b)
	if err != nil {
		return err
	}
	return nil
}
