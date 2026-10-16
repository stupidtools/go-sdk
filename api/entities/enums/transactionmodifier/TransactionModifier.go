package transactionmodifier

type TransactionModifier string

const (
	None                     TransactionModifier = "None"
	Incremental              TransactionModifier = "Incremental"
	Additional               TransactionModifier = "Additional"
	Offline                  TransactionModifier = "Offline"
	LevelII                  TransactionModifier = "LevelII"
	FraudDecline             TransactionModifier = "FraudDecline"
	ChipDecline              TransactionModifier = "ChipDecline"
	CashBack                 TransactionModifier = "CashBack"
	Voucher                  TransactionModifier = "Voucher"
	Secure3D                 TransactionModifier = "Secure3D"
	HostedRequest            TransactionModifier = "HostedRequest"
	Recurring                TransactionModifier = "Recurring"
	EncryptedMobile          TransactionModifier = "EncryptedMobile"
	Fallback                 TransactionModifier = "Fallback"
	LevelIII                 TransactionModifier = "LevelIII"
	DecryptedMobile          TransactionModifier = "DecryptedMobile"
	AlternativePaymentMethod TransactionModifier = "AlternativePaymentMethod"
	OfflineDecline           TransactionModifier = "OfflineDecline"
	DeletePreAuth            TransactionModifier = "DeletePreAuth"
	BankPayment              TransactionModifier = "BankPayment"
	BuyNowPayLater           TransactionModifier = "BuyNowPayLater"
	Merchant                 TransactionModifier = "Merchant"
)

var transactionModifierValues = map[TransactionModifier]int{
	None:                     0,
	Incremental:              1,
	Additional:               2,
	Offline:                  3,
	LevelII:                  4,
	FraudDecline:             5,
	ChipDecline:              6,
	CashBack:                 7,
	Voucher:                  8,
	Secure3D:                 9,
	HostedRequest:            10,
	Recurring:                11,
	EncryptedMobile:          12,
	Fallback:                 13,
	LevelIII:                 14,
	DecryptedMobile:          15,
	AlternativePaymentMethod: 16,
	OfflineDecline:           17,
	DeletePreAuth:            18,
	BankPayment:              19,
	BuyNowPayLater:           20,
	Merchant:                 21,
}

func (t TransactionModifier) LongValue() int64 {
	return 1 << transactionModifierValues[t]
}

func (t TransactionModifier) GetSet(value int64) map[TransactionModifier]bool {
	flags := make(map[TransactionModifier]bool)
	for _, flag := range []TransactionModifier{
		None, Incremental, Additional, Offline, LevelII, FraudDecline, ChipDecline, CashBack, Voucher,
		Secure3D, HostedRequest, Recurring, EncryptedMobile, Fallback, LevelIII, DecryptedMobile,
		AlternativePaymentMethod, OfflineDecline, DeletePreAuth, BankPayment, BuyNowPayLater, Merchant,
	} {
		flagValue := flag.LongValue()
		if (flagValue & value) == flagValue {
			flags[flag] = true
		}
	}
	return flags
}

func (t TransactionModifier) Ordinal() int {
	return transactionModifierValues[t]
}
