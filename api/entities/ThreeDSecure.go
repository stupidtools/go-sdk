package entities

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/securethreedversion"
	"github.com/shopspring/decimal"
)

type ThreeDSecure struct {
	AcsTransactionId             string
	AcsEndVersion                string
	AcsStartVersion              string
	AcsInfoIndicator             []string
	AcsInterface                 string
	AcsUiTemplate                string
	AcsReferenceNumber           string
	Algorithm                    int
	Amount                       *decimal.Decimal
	AuthenticationSource         string
	AuthenticationType           string
	AuthenticationValue          string
	CardHolderResponseInfo       string
	Cavv                         string
	ChallengeMandated            bool
	Currency                     string
	DecoupledResponseIndicator   string
	DirectoryServerTransactionId string
	DirectoryServerEndVersion    string
	DirectoryServerStartVersion  string
	Eci                          string
	Enrolled                     bool
	EnrolledStatus               string
	ExemptReason                 string
	IssuerAcsUrl                 string
	ChallengeReturnUrl           string
	SessionDataFieldName         string
	MessageType                  string
	MessageCategory              string
	MessageVersion               string
	OrderId                      string
	PayerAuthenticationRequest   string
	PaymentDataSource            PaymentDataSourceType
	PaymentDataType              string
	SdkInterface                 string
	SdkUiType                    string
	ServerTransactionId          string
	Status                       string
	StatusReason                 string
	Version                      securethreedversion.Secure3dVersion
	WhitelistStatus              string
	Xid                          string
	LiabilityShift               string
	ProviderServerTransRef       string
}

func NewThreeDSecure() *ThreeDSecure {
	return &ThreeDSecure{}
}
