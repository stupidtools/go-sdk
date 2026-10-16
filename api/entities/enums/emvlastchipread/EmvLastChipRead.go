package emvlastchipread

type EmvLastChipRead string

const (
	SUCCESSFUL             EmvLastChipRead = "SUCCESSFUL"
	FAILED                 EmvLastChipRead = "FAILED"
	NOT_A_CHIP_TRANSACTION EmvLastChipRead = "NOT_A_CHIP_TRANSACTION"
	UNKNOWN                EmvLastChipRead = "UNKNOWN"
)
