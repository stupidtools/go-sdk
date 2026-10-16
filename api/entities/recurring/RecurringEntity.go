package recurring

import (
	"github.com/globalpayments/go-sdk/api/entities/abstractions"
)

type RecurringEntity struct {
	ID              string
	Key             string
	ResponseCode    string
	ResponseMessage string
}

func (r *RecurringEntity) GetID() string {
	return r.ID
}

func (r *RecurringEntity) SetID(value string) {
	r.ID = value
}

func (r *RecurringEntity) GetKey() string {
	if r.Key != "" {
		return r.Key
	}
	return r.ID
}

func (r *RecurringEntity) SetKey(value string) {
	r.Key = value
}

func (r *RecurringEntity) Create() (abstractions.IRecurringEntity, error) {
	return r.CreateWithConfig("default")
}

func (r *RecurringEntity) CreateWithConfig(configName string) (abstractions.IRecurringEntity, error) {
	return r.CreateWithConfig(configName)
}
