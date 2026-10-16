package transactionsortproperty

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

// TransactionSortProperty represents the sorting properties for transactions.
type TransactionSortProperty string

const (
	Id          TransactionSortProperty = "ID"
	TimeCreated TransactionSortProperty = "TIME_CREATED"
	Status      TransactionSortProperty = "STATUS"
	Type        TransactionSortProperty = "TYPE"
	DepositId   TransactionSortProperty = "DEPOSIT_ID"
)

// GetValue returns the string value of the TransactionSortProperty.
func (t TransactionSortProperty) GetValue() string {
	return string(t)
}

// GetBytes returns the byte representation of the TransactionSortProperty value.
func (t TransactionSortProperty) GetBytes() []byte {
	return []byte(t)
}

// StringConstants returns a slice of all TransactionSortProperty values.
func (t TransactionSortProperty) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{Id, TimeCreated, Status, Type, DepositId}
}
