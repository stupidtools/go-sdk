package entities

import (
	"github.com/globalpayments/go-sdk/api/entities/base"
	"github.com/globalpayments/go-sdk/api/entities/billing"
	"github.com/globalpayments/go-sdk/api/entities/enums/alternativepaymenttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/challengerequest"
	"github.com/globalpayments/go-sdk/api/entities/enums/hostedpaymentmethods"
	"github.com/globalpayments/go-sdk/api/entities/enums/hostedpaymenttype"
)

type HostedPaymentData struct {
	AddressesMatch              bool
	Bills                       []billing.Bill
	ChallengeRequestIndicator   challengerequest.ChallengeRequest
	CustomerExists              bool
	CustomerIsEditable          bool
	CustomerAddress             base.Address
	CustomerEmail               string
	CustomerKey                 string
	CustomerNumber              string
	CustomerCountry             string
	CustomerFirstName           string
	CustomerLastName            string
	CustomerPhoneMobile         string
	HostedPaymentType           hostedpaymenttype.HostedPaymentType
	OfferToSaveCard             bool
	PaymentKey                  string
	ProductID                   string
	CaptureAddress              bool
	ReturnAddress               bool
	EnableExemptionOptimization bool
	PresetPaymentMethods        []alternativepaymenttype.AlternativePaymentType
	SupplementaryData           map[string]string
	TransactionStatusURL        string
	CancelURL                   string
	MerchantResponseURL         string
	HostedPaymentMethods        []hostedpaymentmethods.HostedPaymentMethods
}

func (h *HostedPaymentData) GetAddressesMatch() bool {
	return h.AddressesMatch
}

func (h *HostedPaymentData) SetAddressesMatch(addressesMatch bool) {
	h.AddressesMatch = addressesMatch
}

func (h *HostedPaymentData) GetBills() []billing.Bill {
	return h.Bills
}

func (h *HostedPaymentData) SetBills(bills []billing.Bill) {
	h.Bills = bills
}

func (h *HostedPaymentData) GetChallengeRequestIndicator() challengerequest.ChallengeRequest {
	return h.ChallengeRequestIndicator
}

func (h *HostedPaymentData) SetChallengeRequestIndicator(challengeRequestIndicator challengerequest.ChallengeRequest) {
	h.ChallengeRequestIndicator = challengeRequestIndicator
}

func (h *HostedPaymentData) GetCustomerEmail() string {
	return h.CustomerEmail
}

func (h *HostedPaymentData) SetCustomerEmail(customerEmail string) {
	h.CustomerEmail = customerEmail
}

func (h *HostedPaymentData) IsCustomerExists() bool {
	return h.CustomerExists
}

func (h *HostedPaymentData) SetCustomerExists(customerExists bool) {
	h.CustomerExists = customerExists
}

func (h *HostedPaymentData) IsCustomerEditable() bool {
	return h.CustomerIsEditable
}

func (h *HostedPaymentData) SetCustomerIsEditable(customerIsEditable bool) {
	h.CustomerIsEditable = customerIsEditable
}

func (h *HostedPaymentData) GetCustomerAddress() base.Address {
	return h.CustomerAddress
}

func (h *HostedPaymentData) SetCustomerAddress(address base.Address) {
	h.CustomerAddress = address
}

func (h *HostedPaymentData) GetCustomerKey() string {
	return h.CustomerKey
}

func (h *HostedPaymentData) SetCustomerKey(customerKey string) {
	h.CustomerKey = customerKey
}

func (h *HostedPaymentData) GetCustomerNumber() string {
	return h.CustomerNumber
}

func (h *HostedPaymentData) SetCustomerNumber(customerNumber string) {
	h.CustomerNumber = customerNumber
}

func (h *HostedPaymentData) GetCustomerCountry() string {
	return h.CustomerCountry
}

func (h *HostedPaymentData) SetCustomerCountry(customerCountry string) {
	h.CustomerCountry = customerCountry
}

func (h *HostedPaymentData) GetCustomerFirstName() string {
	return h.CustomerFirstName
}

func (h *HostedPaymentData) SetCustomerFirstName(customerFirstName string) {
	h.CustomerFirstName = customerFirstName
}

func (h *HostedPaymentData) GetCustomerLastName() string {
	return h.CustomerLastName
}

func (h *HostedPaymentData) SetCustomerLastName(customerLastName string) {
	h.CustomerLastName = customerLastName
}

func (h *HostedPaymentData) GetCustomerPhoneMobile() string {
	return h.CustomerPhoneMobile
}

func (h *HostedPaymentData) SetCustomerPhoneMobile(customerPhoneMobile string) {
	h.CustomerPhoneMobile = customerPhoneMobile
}

func (h *HostedPaymentData) GetMerchantResponseURL() string {
	return h.MerchantResponseURL
}

func (h *HostedPaymentData) SetMerchantResponseURL(merchantResponseURL string) {
	h.MerchantResponseURL = merchantResponseURL
}

func (h *HostedPaymentData) GetHostedPaymentType() hostedpaymenttype.HostedPaymentType {
	return h.HostedPaymentType
}

func (h *HostedPaymentData) SetHostedPaymentType(hostedPaymentType hostedpaymenttype.HostedPaymentType) {
	h.HostedPaymentType = hostedPaymentType
}

func (h *HostedPaymentData) IsOfferToSaveCard() bool {
	return h.OfferToSaveCard
}

func (h *HostedPaymentData) SetOfferToSaveCard(offerToSaveCard bool) {
	h.OfferToSaveCard = offerToSaveCard
}

func (h *HostedPaymentData) GetPaymentKey() string {
	return h.PaymentKey
}

func (h *HostedPaymentData) SetPaymentKey(paymentKey string) {
	h.PaymentKey = paymentKey
}

func (h *HostedPaymentData) GetProductID() string {
	return h.ProductID
}

func (h *HostedPaymentData) SetProductID(productID string) {
	h.ProductID = productID
}

func (h *HostedPaymentData) GetPresetPaymentMethods() []alternativepaymenttype.AlternativePaymentType {
	return h.PresetPaymentMethods
}

func (h *HostedPaymentData) SetPresetPaymentMethods(paymentTypes ...alternativepaymenttype.AlternativePaymentType) {
	h.PresetPaymentMethods = paymentTypes
}

func (h *HostedPaymentData) GetSupplementaryData() map[string]string {
	return h.SupplementaryData
}

func (h *HostedPaymentData) SetSupplementaryData(supplementaryData map[string]string) {
	h.SupplementaryData = supplementaryData
}

func (h *HostedPaymentData) GetTransactionStatusURL() string {
	return h.TransactionStatusURL
}

func (h *HostedPaymentData) SetTransactionStatusURL(transactionStatusURL string) {
	h.TransactionStatusURL = transactionStatusURL
}

func NewHostedPaymentData() *HostedPaymentData {
	return &HostedPaymentData{
		SupplementaryData: make(map[string]string),
	}
}
