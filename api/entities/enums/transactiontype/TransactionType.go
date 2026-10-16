package transactiontype

type TransactionType string

const (
	Decline                TransactionType = "Decline"
	Verify                 TransactionType = "Verify"
	Capture                TransactionType = "Capture"
	Auth                   TransactionType = "Auth"
	Refund                 TransactionType = "Refund"
	Reversal               TransactionType = "Reversal"
	Sale                   TransactionType = "Sale"
	Edit                   TransactionType = "Edit"
	Void                   TransactionType = "Void"
	AddValue               TransactionType = "AddValue"
	Balance                TransactionType = "Balance"
	Activate               TransactionType = "Activate"
	Alias                  TransactionType = "Alias"
	Replace                TransactionType = "Replace"
	Reward                 TransactionType = "Reward"
	Deactivate             TransactionType = "Deactivate"
	BatchClose             TransactionType = "BatchClose"
	Create                 TransactionType = "Create"
	Delete                 TransactionType = "Delete"
	BenefitWithdrawal      TransactionType = "BenefitWithdrawal"
	Fetch                  TransactionType = "Fetch"
	Search                 TransactionType = "Search"
	Hold                   TransactionType = "Hold"
	Release                TransactionType = "Release"
	VerifyEnrolled         TransactionType = "VerifyEnrolled"
	VerifySignature        TransactionType = "VerifySignature"
	TokenUpdate            TransactionType = "TokenUpdate"
	TokenDelete            TransactionType = "TokenDelete"
	Confirm                TransactionType = "Confirm"
	InitiateAuthentication TransactionType = "InitiateAuthentication"
	DataCollect            TransactionType = "DataCollect"
	PreAuthCompletion      TransactionType = "PreAuthCompletion"
	DccRateLookup          TransactionType = "DccRateLookup"
	Increment              TransactionType = "Increment"
	Tokenize               TransactionType = "Tokenize"
	CashOut                TransactionType = "CashOut"
	SendFile               TransactionType = "SendFile"
	Payment                TransactionType = "Payment"
	CashAdvance            TransactionType = "CashAdvance"
	DisputeAcceptance      TransactionType = "DisputeAcceptance"
	DisputeChallenge       TransactionType = "DisputeChallenge"
	LoadReversal           TransactionType = "LoadReversal"
	Reauth                 TransactionType = "Reauth"
	Mail                   TransactionType = "Mail"
	PDL                    TransactionType = "PDL"
	UtilityMessage         TransactionType = "UtilityMessage"
	MagnumPDL              TransactionType = "MagnumPDL"
	EmvPdl                 TransactionType = "EmvPdl"
	PosSiteConfiguration   TransactionType = "PosSiteConfiguration"
	PayLinkUpdate          TransactionType = "PayLinkUpdate"
	RiskAssess             TransactionType = "RiskAssess"
	TimeRequest            TransactionType = "TimeRequest"
	Issue                  TransactionType = "Issue"
	RequestPendingMessages TransactionType = "RequestPendingMessages"
	Empty                  TransactionType = ""
)

func (t TransactionType) LongValue() int64 {
	return 1 << transactionTypeOrdinalMap[t]
}

func (t TransactionType) IsReversal() bool {
	return t == Reversal || t == LoadReversal
}

func (t TransactionType) GetStringValue() string {
	return string(t)
}

func GetTransactionTypeSet(value int64) []TransactionType {
	var flags []TransactionType
	for t, i := range transactionTypeOrdinalMap {
		flagValue := int64(1 << i)
		if (flagValue & value) == flagValue {
			flags = append(flags, t)
		}
	}
	return flags
}

var transactionTypeOrdinalMap = map[TransactionType]int{
	Decline:                0,
	Verify:                 1,
	Capture:                2,
	Auth:                   3,
	Refund:                 4,
	Reversal:               5,
	Sale:                   6,
	Edit:                   7,
	Void:                   8,
	AddValue:               9,
	Balance:                10,
	Activate:               11,
	Alias:                  12,
	Replace:                13,
	Reward:                 14,
	Deactivate:             15,
	BatchClose:             16,
	Create:                 17,
	Delete:                 18,
	BenefitWithdrawal:      19,
	Fetch:                  20,
	Search:                 21,
	Hold:                   22,
	Release:                23,
	VerifyEnrolled:         24,
	VerifySignature:        25,
	TokenUpdate:            26,
	TokenDelete:            27,
	Confirm:                28,
	InitiateAuthentication: 29,
	DataCollect:            30,
	PreAuthCompletion:      31,
	DccRateLookup:          32,
	Increment:              33,
	Tokenize:               34,
	CashOut:                35,
	SendFile:               36,
	Payment:                37,
	CashAdvance:            38,
	DisputeAcceptance:      39,
	DisputeChallenge:       40,
	LoadReversal:           41,
	Reauth:                 42,
	Mail:                   43,
	PDL:                    44,
	UtilityMessage:         45,
	MagnumPDL:              46,
	EmvPdl:                 47,
	PosSiteConfiguration:   48,
	PayLinkUpdate:          49,
	RiskAssess:             50,
	TimeRequest:            51,
	Issue:                  52,
	RequestPendingMessages: 53,
}
