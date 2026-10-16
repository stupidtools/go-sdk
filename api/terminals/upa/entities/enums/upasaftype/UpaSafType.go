package upasaftype

type UpaSafType string

const (
	Approved UpaSafType = "AUTHORIZED TRANSACTIONS"
	Pending  UpaSafType = "PENDING TRANSACTIONS"
	Failed   UpaSafType = "FAILED TRANSACTIONS"
)

// Define a method for UpaSafType to implement the IStringConstant interface
func (safType UpaSafType) GetBytes() []byte {
	return nil
}

// Define a method for UpaSafType to implement the IStringConstant interface
func (safType UpaSafType) GetValue() string {
	return string(safType)
}
