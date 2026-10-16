package upaacquisitiontype

type UpaAcquisitionType string

const (
	Contact     UpaAcquisitionType = "Contact"
	Contactless UpaAcquisitionType = "Contactless"
	Swipe       UpaAcquisitionType = "Swipe"
	Manual      UpaAcquisitionType = "Manual"
	Scan        UpaAcquisitionType = "Scan"
)
