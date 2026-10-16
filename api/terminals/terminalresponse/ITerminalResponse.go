package terminalresponse

import (
	"github.com/shopspring/decimal"

	"github.com/globalpayments/go-sdk/api/entities/enums/applicationcryptogramtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/cardtype"
)

type ITerminalResponse interface {
	ToString() string
	GetStatus() string
	SetStatus(status string)
	GetCommand() string
	SetCommand(command string)
	GetVersion() string
	SetVersion(version string)
	GetDeviceResponseCode() string
	SetDeviceResponseCode(deviceResponseCode string)
	GetDeviceResponseText() string
	SetDeviceResponseText(deviceResponseText string)
	GetResponseCode() string
	SetResponseCode(responseCode string)
	GetResponseText() string
	SetResponseText(responseText string)
	GetTransactionId() string
	SetTransactionId(transactionId string)
	GetTerminalRefNumber() string
	SetTerminalRefNumber(terminalRefNumber string)
	GetToken() string
	SetToken(token string)
	GetSignatureStatus() string
	SetSignatureStatus(signatureStatus string)
	GetSignatureData() []byte
	SetSignatureData(signatureData []byte)
	GetTransactionType() string
	SetTransactionType(transactionType string)
	GetMaskedCardNumber() string
	SetMaskedCardNumber(maskedCardNumber string)
	GetEntryMethod() string
	SetEntryMethod(entryMethod string)
	GetAuthorizationCode() string
	SetAuthorizationCode(authorizationCode string)
	GetApprovalCode() string
	SetApprovalCode(approvalCode string)
	GetTransactionAmount() *decimal.Decimal
	SetTransactionAmount(transactionAmount *decimal.Decimal)
	GetAmountDue() *decimal.Decimal
	SetAmountDue(amountDue *decimal.Decimal)
	GetBalanceAmount() *decimal.Decimal
	SetBalanceAmount(balanceAmount *decimal.Decimal)
	GetCardHolderName() string
	SetCardHolderName(cardHolderName string)
	GetCardBIN() string
	SetCardBIN(cardBIN string)
	IsCardPresent() bool
	SetCardPresent(cardPresent bool)
	GetCardType() cardtype.CardType
	SetCardType(cardType cardtype.CardType)
	GetExpirationDate() string
	SetExpirationDate(expirationDate string)
	GetTipAmount() *decimal.Decimal
	SetTipAmount(tipAmount *decimal.Decimal)
	GetCashBackAmount() *decimal.Decimal
	SetCashBackAmount(cashBackAmount *decimal.Decimal)
	GetAvsResponseCode() string
	SetAvsResponseCode(avsResponseCode string)
	GetAvsResponseText() string
	SetAvsResponseText(avsResponseText string)
	GetCvvResponseCode() string
	SetCvvResponseCode(cvvResponseCode string)
	GetCvvResponseText() string
	SetCvvResponseText(cvvResponseText string)
	IsTaxExempt() bool
	SetTaxExempt(taxExempt bool)
	GetTaxExemptId() string
	SetTaxExemptId(taxExemptId string)
	GetTicketNumber() string
	SetTicketNumber(ticketNumber string)
	GetPaymentType() string
	SetPaymentType(paymentType string)
	GetMerchantFee() *decimal.Decimal
	SetMerchantFee(merchantFee *decimal.Decimal)
	GetApplicationPreferredName() string
	SetApplicationPreferredName(applicationPreferredName string)
	GetApplicationLabel() string
	SetApplicationLabel(applicationLabel string)
	GetApplicationId() string
	SetApplicationId(applicationId string)
	GetApplicationCryptogramType() applicationcryptogramtype.ApplicationCryptogramType
	SetApplicationCryptogramType(applicationCryptogramType applicationcryptogramtype.ApplicationCryptogramType)
	GetApplicationCryptogram() string
	SetApplicationCryptogram(applicationCryptogram string)
	GetCustomerVerificationMethod() string
	SetCustomerVerificationMethod(customerVerificationMethod string)
	GetTerminalVerificationResults() string
	SetTerminalVerificationResults(terminalVerificationResults string)
	GetCardBrandTransactionId() string
	SetCardBrandTransactionId(cardBrandTransactionId string)
	GetUnmaskedCardNumber() string
	SetUnmaskedCardNumber(string)
	GetTransactionStatusInfo() string
	SetTransactionStatusInfo(status string)
	GetFallback() string
	SetFallback(fb string)
	GetServiceCode() string
	SetServiceCode(fb string)
}
