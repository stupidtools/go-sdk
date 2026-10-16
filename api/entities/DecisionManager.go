package entities

import "github.com/globalpayments/go-sdk/api/entities/enums/risk"

type DecisionManager struct {
	BillToHostName                   string
	BillToHttpBrowserCookiesAccepted bool
	BillToHttpBrowserEmail           string
	BillToHttpBrowserType            string
	BillToIpNetworkAddress           string
	BusinessRulesCoreThreshold       string
	BillToPersonalId                 string
	DecisionManagerProfile           string
	InvoiceHeaderTenderType          string
	ItemHostHedge                    risk.Risk
	ItemNonsensicalHedge             risk.Risk
	ItemObscenitiesHedge             risk.Risk
	ItemPhoneHedge                   risk.Risk
	ItemTimeHedge                    risk.Risk
	ItemVelocityHedge                risk.Risk
	InvoiceHeaderIsGift              bool
	InvoiceHeaderReturnsAccepted     bool
}

func (d *DecisionManager) GetBillToHostName() string {
	return d.BillToHostName
}

func (d *DecisionManager) SetBillToHostName(value string) {
	d.BillToHostName = value
}

func (d *DecisionManager) GetBillToHttpBrowserCookiesAccepted() bool {
	return d.BillToHttpBrowserCookiesAccepted
}

func (d *DecisionManager) SetBillToHttpBrowserCookiesAccepted(value bool) {
	d.BillToHttpBrowserCookiesAccepted = value
}

func (d *DecisionManager) GetBillToHttpBrowserEmail() string {
	return d.BillToHttpBrowserEmail
}

func (d *DecisionManager) SetBillToHttpBrowserEmail(value string) {
	d.BillToHttpBrowserEmail = value
}

func (d *DecisionManager) GetBillToHttpBrowserType() string {
	return d.BillToHttpBrowserType
}

func (d *DecisionManager) SetBillToHttpBrowserType(value string) {
	d.BillToHttpBrowserType = value
}

func (d *DecisionManager) GetBillToIpNetworkAddress() string {
	return d.BillToIpNetworkAddress
}

func (d *DecisionManager) SetBillToIpNetworkAddress(value string) {
	d.BillToIpNetworkAddress = value
}

func (d *DecisionManager) GetBusinessRulesCoreThreshold() string {
	return d.BusinessRulesCoreThreshold
}

func (d *DecisionManager) SetBusinessRulesCoreThreshold(value string) {
	d.BusinessRulesCoreThreshold = value
}

func (d *DecisionManager) GetBillToPersonalId() string {
	return d.BillToPersonalId
}

func (d *DecisionManager) SetBillToPersonalId(value string) {
	d.BillToPersonalId = value
}

func (d *DecisionManager) GetDecisionManagerProfile() string {
	return d.DecisionManagerProfile
}

func (d *DecisionManager) SetDecisionManagerProfile(value string) {
	d.DecisionManagerProfile = value
}

func (d *DecisionManager) GetInvoiceHeaderTenderType() string {
	return d.InvoiceHeaderTenderType
}

func (d *DecisionManager) SetInvoiceHeaderTenderType(value string) {
	d.InvoiceHeaderTenderType = value
}

func (d *DecisionManager) GetItemHostHedge() risk.Risk {
	return d.ItemHostHedge
}

func (d *DecisionManager) SetItemHostHedge(value risk.Risk) {
	d.ItemHostHedge = value
}

func (d *DecisionManager) GetItemNonsensicalHedge() risk.Risk {
	return d.ItemNonsensicalHedge
}

func (d *DecisionManager) SetItemNonsensicalHedge(value risk.Risk) {
	d.ItemNonsensicalHedge = value
}

func (d *DecisionManager) GetItemObscenitiesHedge() risk.Risk {
	return d.ItemObscenitiesHedge
}

func (d *DecisionManager) SetItemObscenitiesHedge(value risk.Risk) {
	d.ItemObscenitiesHedge = value
}

func (d *DecisionManager) GetItemPhoneHedge() risk.Risk {
	return d.ItemPhoneHedge
}

func (d *DecisionManager) SetItemPhoneHedge(value risk.Risk) {
	d.ItemPhoneHedge = value
}

func (d *DecisionManager) GetItemTimeHedge() risk.Risk {
	return d.ItemTimeHedge
}

func (d *DecisionManager) SetItemTimeHedge(value risk.Risk) {
	d.ItemTimeHedge = value
}

func (d *DecisionManager) GetItemVelocityHedge() risk.Risk {
	return d.ItemVelocityHedge
}

func (d *DecisionManager) SetItemVelocityHedge(value risk.Risk) {
	d.ItemVelocityHedge = value
}

func (d *DecisionManager) GetInvoiceHeaderIsGift() bool {
	return d.InvoiceHeaderIsGift
}

func (d *DecisionManager) SetInvoiceHeaderIsGift(value bool) {
	d.InvoiceHeaderIsGift = value
}

func (d *DecisionManager) GetInvoiceHeaderReturnsAccepted() bool {
	return d.InvoiceHeaderReturnsAccepted
}

func (d *DecisionManager) SetInvoiceHeaderReturnsAccepted(value bool) {
	d.InvoiceHeaderReturnsAccepted = value
}
