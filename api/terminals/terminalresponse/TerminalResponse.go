package terminalresponse

import (
	"encoding/json"

	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/applicationcryptogramtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/cardtype"
)

type TerminalResponse struct {
	// Internal
	Status  string `json:"status"`
	Command string `json:"command"`
	Version string `json:"version"`

	// Functional
	EcrId              string `json:"EcrId"`
	DeviceResponseCode string `json:"deviceResponseCode"`
	DeviceResponseText string `json:"deviceResponseText"`
	ResponseCode       string `json:"responseCode"`
	ResponseText       string `json:"responseText"`
	TransactionId      string `json:"transactionId"`
	TerminalRefNumber  string `json:"terminalRefNumber"`
	Token              string `json:"token"`
	SignatureStatus    string `json:"signatureStatus"`
	SignatureData      []byte `json:"signatureData"`
	Fallback           string `json:"fallback"`
	ServiceCode        string `json:"serviceCode"`

	// Transactional
	TransactionType        string            `json:"transactionType"`
	TransactionStatusInfo  string            `json:"transactionStatusInfo"`
	MaskedCardNumber       string            `json:"maskedCardNumber"`
	EntryMethod            string            `json:"entryMethod"`
	AuthorizationCode      string            `json:"authorizationCode"`
	ApprovalCode           string            `json:"approvalCode"`
	TransactionAmount      *decimal.Decimal  `json:"transactionAmount"`
	AmountDue              *decimal.Decimal  `json:"amountDue"`
	BalanceAmount          *decimal.Decimal  `json:"balanceAmount"`
	CardHolderName         string            `json:"cardHolderName"`
	CardBIN                string            `json:"cardBIN"`
	CardPresent            bool              `json:"cardPresent"`
	CardType               cardtype.CardType `json:"cardType"`
	ExpirationDate         string            `json:"expirationDate"`
	TipAmount              *decimal.Decimal  `json:"tipAmount"`
	CashBackAmount         *decimal.Decimal  `json:"cashBackAmount"`
	AvsResponseCode        string            `json:"avsResponseCode"`
	AvsResponseText        string            `json:"avsResponseText"`
	CvvResponseCode        string            `json:"cvvResponseCode"`
	CvvResponseText        string            `json:"cvvResponseText"`
	TaxExempt              bool              `json:"taxExempt"`
	TaxExemptId            string            `json:"taxExemptId"`
	TicketNumber           string            `json:"ticketNumber"`
	PaymentType            string            `json:"paymentType"`
	MerchantFee            *decimal.Decimal  `json:"merchantFee"`
	CardBrandTransactionId string            `json:"cardBrandTransactionId"`

	// EMV
	ApplicationPreferredName    string                                              `json:"applicationPreferredName"`
	ApplicationLabel            string                                              `json:"applicationLabel"`
	ApplicationId               string                                              `json:"applicationId"`
	ApplicationCryptogramType   applicationcryptogramtype.ApplicationCryptogramType `json:"applicationCryptogramType"`
	ApplicationCryptogram       string                                              `json:"applicationCryptogram"`
	CustomerVerificationMethod  string                                              `json:"customerVerificationMethod"`
	TerminalVerificationResults string                                              `json:"terminalVerificationResults"`
	UnmaskedCardNumber          string                                              `json:"unmaskedCardNumber"`
}

func (tr TerminalResponse) ToString() string {
	jsonBytes, err := json.Marshal(tr)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func (tr *TerminalResponse) GetStatus() string {
	return tr.Status
}

func (tr *TerminalResponse) SetStatus(status string) {
	tr.Status = status
}

func (tr *TerminalResponse) GetTransactionStatusInfo() string {
	return tr.TransactionStatusInfo
}

func (tr *TerminalResponse) SetTransactionStatusInfo(status string) {
	tr.TransactionStatusInfo = status
}

func (tr *TerminalResponse) GetCommand() string {
	return tr.Command
}

func (tr *TerminalResponse) SetCommand(command string) {
	tr.Command = command
}

func (tr *TerminalResponse) GetVersion() string {
	return tr.Version
}

func (tr *TerminalResponse) SetVersion(version string) {
	tr.Version = version
}

func (tr *TerminalResponse) GetEcrId() string {
	return tr.EcrId
}

func (tr *TerminalResponse) SetEcrId(EcrId string) {
	tr.EcrId = EcrId
}

func (tr *TerminalResponse) GetDeviceResponseCode() string {
	return tr.DeviceResponseCode
}

func (tr *TerminalResponse) SetDeviceResponseCode(deviceResponseCode string) {
	tr.DeviceResponseCode = deviceResponseCode
}

func (tr *TerminalResponse) GetDeviceResponseText() string {
	return tr.DeviceResponseText
}

func (tr *TerminalResponse) SetDeviceResponseText(deviceResponseText string) {
	tr.DeviceResponseText = deviceResponseText
}

func (tr *TerminalResponse) GetResponseCode() string {
	return tr.ResponseCode
}

func (tr *TerminalResponse) SetResponseCode(responseCode string) {
	tr.ResponseCode = responseCode
}

func (tr *TerminalResponse) GetResponseText() string {
	return tr.ResponseText
}

func (tr *TerminalResponse) SetResponseText(responseText string) {
	tr.ResponseText = responseText
}

func (tr *TerminalResponse) GetTransactionId() string {
	return tr.TransactionId
}

func (tr *TerminalResponse) SetTransactionId(transactionId string) {
	tr.TransactionId = transactionId
}

func (tr *TerminalResponse) GetTerminalRefNumber() string {
	return tr.TerminalRefNumber
}

func (tr *TerminalResponse) SetTerminalRefNumber(terminalRefNumber string) {
	tr.TerminalRefNumber = terminalRefNumber
}

func (tr *TerminalResponse) GetToken() string {
	return tr.Token
}

func (tr *TerminalResponse) SetToken(token string) {
	tr.Token = token
}

func (tr *TerminalResponse) GetSignatureStatus() string {
	return tr.SignatureStatus
}

func (tr *TerminalResponse) SetSignatureStatus(signatureStatus string) {
	tr.SignatureStatus = signatureStatus
}

func (tr *TerminalResponse) GetSignatureData() []byte {
	return tr.SignatureData
}

func (tr *TerminalResponse) SetSignatureData(signatureData []byte) {
	tr.SignatureData = signatureData
}

func (tr *TerminalResponse) GetTransactionType() string {
	return tr.TransactionType
}

func (tr *TerminalResponse) SetTransactionType(transactionType string) {
	tr.TransactionType = transactionType
}

func (tr *TerminalResponse) GetMaskedCardNumber() string {
	return tr.MaskedCardNumber
}

func (tr *TerminalResponse) SetMaskedCardNumber(maskedCardNumber string) {
	tr.MaskedCardNumber = maskedCardNumber
}

func (tr *TerminalResponse) GetEntryMethod() string {
	return tr.EntryMethod
}

func (tr *TerminalResponse) SetEntryMethod(entryMethod string) {
	tr.EntryMethod = entryMethod
}

func (tr *TerminalResponse) GetApprovalCode() string {
	return tr.ApprovalCode
}

func (tr *TerminalResponse) SetApprovalCode(approvalCode string) {
	tr.ApprovalCode = approvalCode
}

func (tr *TerminalResponse) GetTransactionAmount() *decimal.Decimal {
	return tr.TransactionAmount
}

func (tr *TerminalResponse) SetTransactionAmount(transactionAmount *decimal.Decimal) {
	tr.TransactionAmount = transactionAmount
}

func (tr *TerminalResponse) GetAmountDue() *decimal.Decimal {
	return tr.AmountDue
}

func (tr *TerminalResponse) SetAmountDue(amountDue *decimal.Decimal) {
	tr.AmountDue = amountDue
}

func (tr *TerminalResponse) GetCardHolderName() string {
	return tr.CardHolderName
}

func (tr *TerminalResponse) SetCardHolderName(cardHolderName string) {
	tr.CardHolderName = cardHolderName
}

func (tr *TerminalResponse) GetCardBIN() string {
	return tr.CardBIN
}

func (tr *TerminalResponse) SetCardBIN(cardBIN string) {
	tr.CardBIN = cardBIN
}

func (tr *TerminalResponse) IsCardPresent() bool {
	return tr.CardPresent
}

func (tr *TerminalResponse) SetCardPresent(cardPresent bool) {
	tr.CardPresent = cardPresent
}

func (tr *TerminalResponse) GetExpirationDate() string {
	return tr.ExpirationDate
}

func (tr *TerminalResponse) SetExpirationDate(expirationDate string) {
	tr.ExpirationDate = expirationDate
}

func (tr *TerminalResponse) GetTipAmount() *decimal.Decimal {
	return tr.TipAmount
}

func (tr *TerminalResponse) SetTipAmount(tipAmount *decimal.Decimal) {
	tr.TipAmount = tipAmount
}

func (tr *TerminalResponse) GetCashBackAmount() *decimal.Decimal {
	return tr.CashBackAmount
}

func (tr *TerminalResponse) SetCashBackAmount(cashBackAmount *decimal.Decimal) {
	tr.CashBackAmount = cashBackAmount
}

func (tr *TerminalResponse) GetAvsResponseCode() string {
	return tr.AvsResponseCode
}

func (tr *TerminalResponse) SetAvsResponseCode(avsResponseCode string) {
	tr.AvsResponseCode = avsResponseCode
}

func (tr *TerminalResponse) GetAvsResponseText() string {
	return tr.AvsResponseText
}

func (tr *TerminalResponse) SetAvsResponseText(avsResponseText string) {
	tr.AvsResponseText = avsResponseText
}

func (tr *TerminalResponse) GetCvvResponseCode() string {
	return tr.CvvResponseCode
}

func (tr *TerminalResponse) SetCvvResponseCode(cvvResponseCode string) {
	tr.CvvResponseCode = cvvResponseCode
}

func (tr *TerminalResponse) GetCvvResponseText() string {
	return tr.CvvResponseText
}

func (tr *TerminalResponse) SetCvvResponseText(cvvResponseText string) {
	tr.CvvResponseText = cvvResponseText
}

func (tr *TerminalResponse) IsTaxExempt() bool {
	return tr.TaxExempt
}

func (tr *TerminalResponse) SetTaxExempt(taxExempt bool) {
	tr.TaxExempt = taxExempt
}

func (tr *TerminalResponse) GetTaxExemptId() string {
	return tr.TaxExemptId
}

func (tr *TerminalResponse) SetTaxExemptId(taxExemptId string) {
	tr.TaxExemptId = taxExemptId
}

func (tr *TerminalResponse) GetTicketNumber() string {
	return tr.TicketNumber
}

func (tr *TerminalResponse) SetTicketNumber(ticketNumber string) {
	tr.TicketNumber = ticketNumber
}

func (tr *TerminalResponse) GetPaymentType() string {
	return tr.PaymentType
}

func (tr *TerminalResponse) SetPaymentType(paymentType string) {
	tr.PaymentType = paymentType
}

func (tr *TerminalResponse) GetMerchantFee() *decimal.Decimal {
	return tr.MerchantFee
}

func (tr *TerminalResponse) SetMerchantFee(merchantFee *decimal.Decimal) {
	tr.MerchantFee = merchantFee
}

func (tr *TerminalResponse) GetApplicationPreferredName() string {
	return tr.ApplicationPreferredName
}

func (tr *TerminalResponse) SetApplicationPreferredName(applicationPreferredName string) {
	tr.ApplicationPreferredName = applicationPreferredName
}

func (tr *TerminalResponse) GetApplicationLabel() string {
	return tr.ApplicationLabel
}

func (tr *TerminalResponse) SetApplicationLabel(applicationLabel string) {
	tr.ApplicationLabel = applicationLabel
}

func (tr *TerminalResponse) GetApplicationId() string {
	return tr.ApplicationId
}

func (tr *TerminalResponse) SetApplicationId(applicationId string) {
	tr.ApplicationId = applicationId
}

func (tr *TerminalResponse) GetApplicationCryptogramType() applicationcryptogramtype.ApplicationCryptogramType {
	return tr.ApplicationCryptogramType
}

func (tr *TerminalResponse) SetApplicationCryptogramType(applicationCryptogramType applicationcryptogramtype.ApplicationCryptogramType) {
	tr.ApplicationCryptogramType = applicationCryptogramType
}

func (tr *TerminalResponse) GetApplicationCryptogram() string {
	return tr.ApplicationCryptogram
}

func (tr *TerminalResponse) SetApplicationCryptogram(applicationCryptogram string) {
	tr.ApplicationCryptogram = applicationCryptogram
}

func (tr *TerminalResponse) GetCustomerVerificationMethod() string {
	return tr.CustomerVerificationMethod
}

func (tr *TerminalResponse) SetCustomerVerificationMethod(customerVerificationMethod string) {
	tr.CustomerVerificationMethod = customerVerificationMethod
}

func (tr *TerminalResponse) GetTerminalVerificationResults() string {
	return tr.TerminalVerificationResults
}

func (tr *TerminalResponse) SetTerminalVerificationResults(terminalVerificationResults string) {
	tr.TerminalVerificationResults = terminalVerificationResults
}

func (tr *TerminalResponse) GetAuthorizationCode() string {
	return tr.AuthorizationCode
}

func (tr *TerminalResponse) SetAuthorizationCode(authorizationCode string) {
	tr.AuthorizationCode = authorizationCode
}

func (tr *TerminalResponse) GetBalanceAmount() *decimal.Decimal {
	return tr.BalanceAmount
}

func (tr *TerminalResponse) SetBalanceAmount(balanceAmount *decimal.Decimal) {
	tr.BalanceAmount = balanceAmount
}

func (tr *TerminalResponse) GetCardType() cardtype.CardType {
	return tr.CardType
}

func (tr *TerminalResponse) SetCardType(cardType cardtype.CardType) {
	tr.CardType = cardType
}

func (tr *TerminalResponse) GetFallback() string {
	return tr.Fallback
}

func (tr *TerminalResponse) SetFallback(fb string) {
	tr.Fallback = fb
}

func (tr *TerminalResponse) GetServiceCode() string {
	return tr.ServiceCode
}

func (tr *TerminalResponse) SetServiceCode(fb string) {
	tr.ServiceCode = fb
}

func (tr *TerminalResponse) GetCardBrandTransactionId() string {
	return tr.CardBrandTransactionId
}

func (tr *TerminalResponse) SetCardBrandTransactionId(cardBrandTransactionId string) {
	tr.CardBrandTransactionId = cardBrandTransactionId
}

func (tr *TerminalResponse) GetUnmaskedCardNumber() string {
	return tr.UnmaskedCardNumber
}

func (tr *TerminalResponse) SetUnmaskedCardNumber(unmaskedCardNumber string) {
	tr.UnmaskedCardNumber = unmaskedCardNumber
}
