package cardissuerentrytag

type CardIssuerEntryTag string

const (
	StoredValueCards                           CardIssuerEntryTag = "1xx"
	LoyaltyCards                               CardIssuerEntryTag = "2xx"
	PrivateLabelCards                          CardIssuerEntryTag = "3xx"
	SearsProprietaryDeferDate                  CardIssuerEntryTag = "3SF"
	SearsProprietaryDelayDate                  CardIssuerEntryTag = "3SL"
	Bank_CreditCards                           CardIssuerEntryTag = "Bxx"
	Checks                                     CardIssuerEntryTag = "Cxx"
	PIN_DebitCards                             CardIssuerEntryTag = "Dxx"
	PIN_DebitAuthorizer                        CardIssuerEntryTag = "D00"
	ElectronicBenefitsTransfer                 CardIssuerEntryTag = "Exx"
	FleetCards                                 CardIssuerEntryTag = "Fxx"
	Wex_SpecVersionSupport                     CardIssuerEntryTag = "F00"
	Wex_PurchaseDeviceSequenceNumber           CardIssuerEntryTag = "F01"
	PrepaidServiceSystem                       CardIssuerEntryTag = "Gxx"
	AmountSentToIssuerOnBehalfOfPos            CardIssuerEntryTag = "IAM"
	AccountFromCardIssuer                      CardIssuerEntryTag = "IAN"
	CardIssuerAuthenticationResponseCode       CardIssuerEntryTag = "IAR"
	AccountTypeFromCardIssuer                  CardIssuerEntryTag = "IAT"
	MastercardUCAFData                         CardIssuerEntryTag = "IAU"
	Mastercard3DSCryptogram                    CardIssuerEntryTag = "IAU"
	AvsResponseCode                            CardIssuerEntryTag = "IAV"
	ChipConditionCode                          CardIssuerEntryTag = "ICC"
	CreditPlan                                 CardIssuerEntryTag = "ICP"
	CvnResponseCode                            CardIssuerEntryTag = "ICV"
	DiagnosticMessage                          CardIssuerEntryTag = "IDG"
	ExtendedExpirationDate                     CardIssuerEntryTag = "IED"
	GiftCardPurchase                           CardIssuerEntryTag = "IGS"
	UniqueDeviceId                             CardIssuerEntryTag = "IID"
	MastercardDSRPCryptogram                   CardIssuerEntryTag = "IMD"
	MastercardeCommerceIndicators              CardIssuerEntryTag = "IME"
	MastercardRemoteCommerceAcceptorIdentifier CardIssuerEntryTag = "IMU"
	MastercardWalletID                         CardIssuerEntryTag = "IMW"
	PiggyBackActionCode                        CardIssuerEntryTag = "IPA"
	ReceiptText                                CardIssuerEntryTag = "IPR"
	OriginalResponse_ActionCode                CardIssuerEntryTag = "IRA"
	CenterCallNumber                           CardIssuerEntryTag = "IRC"
	IssuerReferenceNumber                      CardIssuerEntryTag = "IRN"
	RetrievalReferenceNumber                   CardIssuerEntryTag = "IRR"
	IssuerSpecificTransactionMatchData         CardIssuerEntryTag = "ITM"
	DisplayText                                CardIssuerEntryTag = "ITX"
	Alternate_DE41                             CardIssuerEntryTag = "I41"
	Alternate_DE42                             CardIssuerEntryTag = "I42"
	DialError                                  CardIssuerEntryTag = "NDE"
	DiscoverNetworkReferenceId                 CardIssuerEntryTag = "ND2"
	NTS_MastercardBankNet_ReferenceNumber      CardIssuerEntryTag = "NM1"
	NTS_MastercardBankNet_SettlementDate       CardIssuerEntryTag = "NM2"
	NTS_POS_Capability                         CardIssuerEntryTag = "NPC"
	PetroleumSwitch                            CardIssuerEntryTag = "NPS"
	SwipeIndicator                             CardIssuerEntryTag = "NSI"
	TerminalError                              CardIssuerEntryTag = "NTE"
	NTS_System                                 CardIssuerEntryTag = "NTS"
	VisaTransactionId                          CardIssuerEntryTag = "NV1"
	HeartlandTimeRequest                       CardIssuerEntryTag = "HTR"
	MerchantActionCode                         CardIssuerEntryTag = "IAC"
	ActualStan                                 CardIssuerEntryTag = "IST"
	DeclineCategoryCode                        CardIssuerEntryTag = "ICA"
	WexAvailableProductRestrictions            CardIssuerEntryTag = "WAP"
	NIL                                        CardIssuerEntryTag = ""
)

func (e CardIssuerEntryTag) GetValue() string {
	return string(e)
}

func (e CardIssuerEntryTag) GetBytes() []byte {
	return []byte(e)
}

func (e CardIssuerEntryTag) FindPartial(value string) CardIssuerEntryTag {
	switch value[0] {
	case '1':
		return StoredValueCards
	case '2':
		return LoyaltyCards
	case '3':
		return PrivateLabelCards
	case 'B':
		return Bank_CreditCards
	case 'C':
		return Checks
	case 'D':
		return PIN_DebitCards
	case 'E':
		return ElectronicBenefitsTransfer
	case 'F':
		return FleetCards
	case 'G':
		return PrepaidServiceSystem
	case 'N':
		return NTS_System
	default:
		return NIL
	}
}
