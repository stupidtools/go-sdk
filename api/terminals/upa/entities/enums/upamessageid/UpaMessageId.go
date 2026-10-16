package upamessageid

type UpaMessageId string

const (
	Sale                 UpaMessageId = "Sale"
	Void                 UpaMessageId = "Void"
	Refund               UpaMessageId = "Refund"
	BalanceInquiry       UpaMessageId = "BalanceInquiry"
	CardVerify           UpaMessageId = "CardVerify"
	TipAdjust            UpaMessageId = "TipAdjust"
	EODProcessing        UpaMessageId = "EODProcessing"
	CancelTransaction    UpaMessageId = "CancelTransaction"
	Reboot               UpaMessageId = "Reboot"
	Reversal             UpaMessageId = "Reversal"
	LineItemDisplay      UpaMessageId = "LineItemDisplay"
	SendSAF              UpaMessageId = "SendSAF"
	GetSAFReport         UpaMessageId = "GetSAFReport"
	DeleteSAF            UpaMessageId = "DeleteSAF"
	GetBatchReport       UpaMessageId = "GetBatchReport"
	GetBatchDetails      UpaMessageId = "GetBatchDetails"
	GetOpenTabDetails    UpaMessageId = "GetOpenTabDetails"
	GetSignature         UpaMessageId = "GetSignature"
	Ping                 UpaMessageId = "Ping"
	PreAuth              UpaMessageId = "PreAuth"
	DeletePreAuth        UpaMessageId = "DeletePreAuth"
	Restart              UpaMessageId = "Restart"
	AuthCompletion       UpaMessageId = "AuthCompletion"
	StartCardTransaction UpaMessageId = "StartCardTransaction"
)
