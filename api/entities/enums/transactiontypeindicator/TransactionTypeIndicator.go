package transactiontypeindicator

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type TransactionTypeIndicator string

const (
	ActivateCancellation       TransactionTypeIndicator = "CAN-ACT"
	ActivateReversal           TransactionTypeIndicator = "REV-ACT"
	BalanceInquiry             TransactionTypeIndicator = "BALINQRY"
	CardActivation             TransactionTypeIndicator = "ACTIVATE"
	CardIssue                  TransactionTypeIndicator = "ISSUE"
	IssueCancellation          TransactionTypeIndicator = "CAN-ISS"
	IssueReversal              TransactionTypeIndicator = "REV-ISS"
	MerchandiseReturn          TransactionTypeIndicator = "RETURN"
	MerchandiseReturnReversal  TransactionTypeIndicator = "REV-RTRN"
	PreAuthorization           TransactionTypeIndicator = "PRE-AUTH"
	PreAuthorizationCompletion TransactionTypeIndicator = "PRE-COMP"
	PreAuthorizationReversal   TransactionTypeIndicator = "REV-PRE"
	Purchase                   TransactionTypeIndicator = "PURCHASE"
	PurchaseCancellation       TransactionTypeIndicator = "CAN-PRCH"
	PurchaseReversal           TransactionTypeIndicator = "REV-PRCH"
	RechargeCardBalance        TransactionTypeIndicator = "RECHARGE"
	RechargeReversal           TransactionTypeIndicator = "REV-RCHG"
)

var allTransactionTypeIndicators = []TransactionTypeIndicator{
	ActivateCancellation,
	ActivateReversal,
	BalanceInquiry,
	CardActivation,
	CardIssue,
	IssueCancellation,
	IssueReversal,
	MerchandiseReturn,
	MerchandiseReturnReversal,
	PreAuthorization,
	PreAuthorizationCompletion,
	PreAuthorizationReversal,
	Purchase,
	PurchaseCancellation,
	PurchaseReversal,
	RechargeCardBalance,
	RechargeReversal,
}

func (tti TransactionTypeIndicator) GetBytes() []byte {
	return []byte(tti)
}

func (tti TransactionTypeIndicator) GetValue() string {
	return string(tti)
}

func (tti TransactionTypeIndicator) StringConstants() []istringconstant.IStringConstant {
	stringConstants := make([]istringconstant.IStringConstant, len(allTransactionTypeIndicators))
	for i, v := range allTransactionTypeIndicators {
		stringConstants[i] = v
	}
	return stringConstants
}
