package summarytype

type SummaryType string

const (
	Approved          SummaryType = "Approved"
	PartiallyApproved SummaryType = "PartiallyApproved"
	VoidApproved      SummaryType = "VoidApproved"
	Pending           SummaryType = "Pending"
	VoidPending       SummaryType = "VoidPending"
	Declined          SummaryType = "Declined"
	VoidDeclined      SummaryType = "VoidDeclined"
	OfflineApproved   SummaryType = "OfflineApproved"
)
