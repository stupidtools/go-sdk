package abstractions

type IRecurringEntity interface {
	GetID() string
	SetID(value string)
	GetKey() string
	SetKey(value string)

	Create(configName string) (IRecurringEntity, error)
	Delete() error
	ForceDelete(force bool) error
	SaveChanges() error
}
