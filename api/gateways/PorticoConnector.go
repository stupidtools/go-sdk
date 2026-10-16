package gateways

import (
	"context"
	"errors"
	"fmt"
	abstractions4 "github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders"
	abstractions5 "github.com/globalpayments/go-sdk/api/builders/abstractions"
	"github.com/globalpayments/go-sdk/api/entities"
	"github.com/globalpayments/go-sdk/api/entities/enums/aliasaction"
	"github.com/globalpayments/go-sdk/api/entities/enums/emvchipcondition"
	"github.com/globalpayments/go-sdk/api/entities/enums/mobilepaymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/paymentmethodtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/reporttype"
	"github.com/globalpayments/go-sdk/api/entities/enums/storedcredentialinitiator"
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
	"github.com/globalpayments/go-sdk/api/entities/enums/taxtype"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactionmodifier"
	"github.com/globalpayments/go-sdk/api/entities/enums/transactiontype"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/entities/transactionsummary"
	"github.com/globalpayments/go-sdk/api/gateways/helpers"
	"github.com/globalpayments/go-sdk/api/network"
	"github.com/globalpayments/go-sdk/api/paymentmethods"
	"github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	abstractions2 "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions/abstractions"
	"github.com/globalpayments/go-sdk/api/paymentmethods/references"
	"github.com/globalpayments/go-sdk/api/utils"
	"github.com/globalpayments/go-sdk/api/utils/enumutils"
	"github.com/globalpayments/go-sdk/api/utils/extrautils"
	"github.com/globalpayments/go-sdk/api/utils/packageutils"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
	"strings"
	"time"
)

type PorticoConnector struct {
	*XmlGateway
	SiteId             string
	LicenseId          string
	DeviceId           string
	Username           string
	Password           string
	DeveloperId        string
	VersionNumber      string
	SecretApiKey       string
	SdkNameVersion     string
	CardType           string
	IsSAFDataSupported bool
}

func NewPorticoConnector() *PorticoConnector {
	return &PorticoConnector{XmlGateway: NewXmlGateway()}
}

func (p *PorticoConnector) SupportsHostedPayments() bool {
	return false
}

func (p *PorticoConnector) SupportsOpenBanking() bool {
	return false
}

func (p *PorticoConnector) SerializeRequest(ctx context.Context, a abstractions4.IAuthorizationBuilder) (string, error) {
	return "", errors.New("Portico does not support hosted payments.")
}

func (p *PorticoConnector) SendKeepAlive(ctx context.Context) (*network.NetworkMessageHeader, error) {
	return nil, errors.New("Portico does not support hosted payments.")
}

func (p *PorticoConnector) ProcessAuthorization(ctx context.Context, builder abstractions4.IAuthorizationBuilder) (abstractions4.ITransaction, error) {
	et := utils.NewElementTree(nil)
	transactionType := builder.GetTransactionType()
	paymentType := builder.GetPaymentMethod().GetPaymentMethodType()
	modifier := builder.GetTransactionModifier()

	// build request
	transactionRaw, err := mapTransactionType(builder)
	if err != nil {
		return nil, err
	}
	transaction, err := et.Element(transactionRaw)
	if err != nil {
		return nil, err
	}

	block1, err := et.SubElement(transaction, "Block1")
	if err != nil {
		return nil, err
	}

	if transactionType == transactiontype.Sale || transactionType == transactiontype.Auth {
		if paymentType != paymentmethodtype.Gift && paymentType != paymentmethodtype.ACH {
			et.SubElementWithCdataValue(block1, "AllowDup", stringutils.BoolToString(builder.IsAllowDuplicates()))
			if modifier == transactionmodifier.None && paymentType != paymentmethodtype.EBT && paymentType != paymentmethodtype.Recurring {
				et.SubElementWithCdataValue(block1, "AllowPartialAuth", stringutils.BoolToString(builder.IsAllowPartialAuth()))
			}
		}
	}

	et.SubElementWithCdataValue(block1, "Amt", stringutils.ToCurrencyString(builder.GetAmount()))
	et.SubElementWithCdataValue(block1, "GratuityAmtInfo", stringutils.ToCurrencyString(builder.GetGratuity()))
	et.SubElementWithCdataValue(block1, "ConvenienceAmtInfo", stringutils.ToCurrencyString(builder.GetConvenienceAmount()))
	et.SubElementWithCdataValue(block1, "ShippingAmtInfo", stringutils.ToCurrencyString(builder.GetShippingAmount()))

	// surcharge
	if builder.GetSurchargeAmount() != nil {
		et.SubElementWithCdataValue(block1, "SurchargeAmtInfo", stringutils.ToCurrencyString(builder.GetSurchargeAmount()))
	}

	// because plano...
	et.SubElementWithCdataValue(block1, extrautils.IfThenElse(paymentType == paymentmethodtype.Debit, "CashbackAmtInfo", "CashBackAmount"), stringutils.ToCurrencyString(builder.GetCashBackAmount()))

	// offline auth code
	et.SubElementWithCdataValue(block1, "OfflineAuthCode", builder.GetOfflineAuthCode())

	// alias action
	if transactionType == transactiontype.Alias {
		et.SubElementWithCdataValue(block1, "Action", builder.GetAliasAction().GetValue())
		et.SubElementWithCdataValue(block1, "Alias", builder.GetAlias())
	}

	isCheck := paymentType == paymentmethodtype.ACH
	if isCheck || builder.GetBillingAddress() != nil || !stringutils.IsNullOrEmpty(builder.GetCardHolderLanguage()) {
		holder, err := et.SubElement(block1, extrautils.IfThenElse(isCheck, "ConsumerInfo", "CardHolderData"))
		if err != nil {
			return nil, err
		}

		address := builder.GetBillingAddress()
		if address != nil {
			et.SubElementWithCdataValue(holder, extrautils.IfThenElse(isCheck, "Address1", "CardHolderAddr"), address.StreetAddr1)
			et.SubElementWithCdataValue(holder, extrautils.IfThenElse(isCheck, "City", "CardHolderCity"), address.City)
			et.SubElementWithCdataValue(holder, extrautils.IfThenElse(isCheck, "State", "CardHolderState"), address.Province)
			et.SubElementWithCdataValue(holder, extrautils.IfThenElse(isCheck, "Zip", "CardHolderZip"), address.PostalCode)
		}

		//if isCheck {
		//	check, ok := builder.GetPaymentMethod().(eCheck)
		//	if ok && !stringutils.IsNullOrEmpty(check.GetCheckHolderName()) {
		//		names := strings.Split(check.GetCheckHolderName(), " ", 2)
		//		et.SubElementWithCdataValue(holder, "FirstName", names[0])
		//		et.SubElementWithCdataValue(holder, "LastName", names[1])
		//	}
		//
		//	et.SubElementWithCdataValue(holder, "CheckName", check.GetCheckName())
		//	et.SubElementWithCdataValue(holder, "PhoneNumber", check.GetPhoneNumber())
		//	et.SubElementWithCdataValue(holder, "DLNumber", check.GetDriversLicenseNumber())
		//	et.SubElementWithCdataValue(holder, "DLState", check.GetDriversLicenseState())
		//
		//	if !stringutils.IsNullOrEmpty(check.GetSsnLast4()) || check.GetBirthYear() != 0 {
		//		identity, err := et.SubElementWithCdataValue(holder, "IdentityInfo")
		//		if err != nil {
		//			return Transaction{}, err
		//		}
		//		et.SubElementWithCdataValue(identity, "SSNL4", check.GetSsnLast4())
		//		et.SubElementWithCdataValue(identity, "DOBYear", check.GetBirthYear())
		//	}
		//} else if card, ok := builder.GetPaymentMethod().(CreditCardData); ok {
		if card, ok := builder.GetPaymentMethod().(paymentmethods.CreditCardData); ok {
			if !stringutils.IsNullOrEmpty(card.GetCardHolderName()) {
				names := strings.Split(card.GetCardHolderName(), " ")
				et.SubElementWithCdataValue(holder, "CardHolderFirstName", names[0])
				if len(names) > 1 {
					et.SubElementWithCdataValue(holder, "CardHolderLastName", names[1])
				}
			}
		}

		// card holder language
		if !isCheck && !stringutils.IsNullOrEmpty(builder.GetCardHolderLanguage()) {
			et.SubElementWithCdataValue(holder, "CardHolderLanguage", builder.GetCardHolderLanguage())
		}
	}

	// card data
	tokenValue := GetToken(builder.GetPaymentMethod())
	hasToken := !stringutils.IsNullOrEmpty(tokenValue)

	// because debit is weird (Ach too)
	cardData := block1
	if paymentType != paymentmethodtype.Debit && paymentType != paymentmethodtype.ACH {
		cardData, err = et.Element("CardData")
		if err != nil {
			return nil, err
		}
	}

	if card, ok := builder.GetPaymentMethod().(abstractions.ICardData); ok {
		// card on File
		if builder.GetTransactionInitiator() != storedcredentialinitiator.Nil || !stringutils.IsNullOrEmpty(builder.GetCardBrandTransactionId()) {
			cardOnFileData, err := et.SubElement(block1, "CardOnFileData")
			if err != nil {
				return nil, err
			}

			if builder.GetTransactionInitiator() == storedcredentialinitiator.CardHolder {
				et.SubElementWithCdataValue(cardOnFileData, "CardOnFile", storedcredentialinitiator.CardHolder.GetValue(target.Portico))
			} else {
				et.SubElementWithCdataValue(cardOnFileData, "CardOnFile", storedcredentialinitiator.Merchant.GetValue(target.Portico))
			}
			et.SubElementWithCdataValue(cardOnFileData, "CardBrandTxnId", builder.GetCardBrandTransactionId())
		}

		manualEntry, err := et.SubElement(cardData, extrautils.IfThenElse(hasToken, "TokenData", "ManualEntry"))
		if err != nil {
			return nil, err
		}

		et.SubElementWithCdataValue(manualEntry, extrautils.IfThenElse(hasToken, "TokenValue", "CardNbr"), extrautils.IfThenElse(tokenValue != "", tokenValue, card.GetNumber()))
		var expMonth, expYear string
		if month := card.GetExpMonth(); month != nil {
			expMonth = stringutils.IntToString(*month)
		}
		if year := card.GetExpYear(); year != nil {
			expYear = stringutils.IntToString(*year)
		}

		et.SubElementWithCdataValue(manualEntry, "ExpMonth", expMonth)
		et.SubElementWithCdataValue(manualEntry, "ExpYear", expYear)
		et.SubElementWithCdataValue(manualEntry, "CVV2", card.GetCvn())
		et.SubElementWithCdataValue(manualEntry, "ReaderPresent", stringutils.BoolToString(card.IsReaderPresent()))
		et.SubElementWithCdataValue(manualEntry, "CardPresent", stringutils.BoolToString(card.IsCardPresent()))
		block1.Append(cardData)

		// secure 3d
		if creditCardData, ok := card.(*paymentmethods.CreditCardData); ok {
			secureEcom := creditCardData.ThreeDSecure

			if secureEcom != nil {
				// 3d Secure Element
				if !stringutils.IsNullOrEmpty(secureEcom.Eci) && !IsAppleOrGooglePay(secureEcom.PaymentDataSource) {
					secure3D, err := et.SubElement(block1, "Secure3D")
					if err != nil {
						return nil, err
					}
					et.SubElementWithInt(secure3D, "Version", secureEcom.Version.GetInt())
					et.SubElementWithCdataValue(secure3D, "AuthenticationValue", secureEcom.Cavv)
					et.SubElementWithCdataValue(secure3D, "ECI", secureEcom.Eci)
					et.SubElementWithCdataValue(secure3D, "DirectoryServerTxnId", secureEcom.Xid)
				}

				// WalletData Element
				if IsAppleOrGooglePay(secureEcom.PaymentDataSource) {
					walletData, err := et.SubElement(block1, "WalletData")
					if err != nil {
						return nil, err
					}
					et.SubElementWithCdataValue(walletData, "PaymentSource", secureEcom.PaymentDataSource.GetValue())
					et.SubElementWithCdataValue(walletData, "Cryptogram", secureEcom.Cavv)
					et.SubElementWithCdataValue(walletData, "ECI", secureEcom.Eci)
				}
			}

			//WalletData Element
			if (creditCardData.MobileType == mobilepaymentmethodtype.APPLEPAY || creditCardData.MobileType == mobilepaymentmethodtype.GOOGLEPAY) && IsAppleOrGooglePay(creditCardData.PaymentDataSourceType) {
				walletData, err := et.SubElement(block1, "WalletData")
				if err != nil {
					return nil, err
				}
				et.SubElementWithCdataValue(walletData, "PaymentSource", creditCardData.PaymentDataSourceType.GetValue())
				et.SubElementWithCdataValue(walletData, "Cryptogram", extrautils.IfThenElse(secureEcom != nil, secureEcom.Cavv, ""))
				et.SubElementWithCdataValue(walletData, "ECI", creditCardData.Eci)
				if creditCardData.MobileType != mobilepaymentmethodtype.NIL {
					et.SubElementWithCdataValue(walletData, "DigitalPaymentToken", creditCardData.GetToken())
					if block1.Has("CardData") {
						block1.Remove("CardData")
					}
					if block1.Has("CardHolderData") {
						block1.Remove("CardHolderData")
					}
				}
			}

		}

		// recurring data
		if builder.GetTransactionModifier() == transactionmodifier.Recurring {
			recurring, err := et.SubElement(block1, "RecurringData")
			if err != nil {
				return nil, err
			}
			et.SubElementWithCdataValue(recurring, "ScheduleID", builder.GetScheduleId())
			et.SubElementWithCdataValue(recurring, "OneTime", stringutils.BoolToString(builder.IsOneTimePayment()))
		}
	} else if track, ok := builder.GetPaymentMethod().(abstractions.ITrackData); ok {
		trackData, err := et.SubElement(cardData, extrautils.IfThenElse(hasToken, "TokenData", "TrackData"))
		if err != nil {
			return nil, err
		}

		if !hasToken {
			trackData.Text(track.GetValue())
			trackData.Set("method", track.GetEntryMethod().GetValue())
			if paymentType == paymentmethodtype.Credit {
				// tag data
				if !stringutils.IsNullOrEmpty(builder.GetTagData()) {
					tagData, err := et.SubElement(block1, "TagData")
					if err != nil {
						return nil, err
					}
					tagValues, err := et.SubElementWithCdataValue(tagData, "TagValues", builder.GetTagData())
					if err != nil {
						return nil, err
					}
					tagValues.Set("source", "chip")
				}
				emvData, err := et.SubElement(block1, "EMVData")
				if err != nil {
					return nil, err
				}
				var chipCondition string
				if builder.GetEmvChipCondition() != nil {
					chipCondition = extrautils.IfThenElse(*builder.GetEmvChipCondition() == emvchipcondition.ChipFailPreviousSuccess, "CHIP_FAILED_PREV_SUCCESS", "CHIP_FAILED_PREV_FAILED")
					et.SubElementWithCdataValue(emvData, "EMVChipCondition", chipCondition)
				}
				if pinProtected, ok := builder.GetPaymentMethod().(abstractions.IPinProtected); ok && pinProtected.GetPinBlock() != "" {
					et.SubElementWithCdataValue(emvData, "PINBlock", pinProtected.GetPinBlock())
				}
				if !block1.Has("EMVChipCondition") && !block1.Has("PINBlock") {
					block1.Remove("EMVData")
				}
			}
			if paymentType == paymentmethodtype.Debit {
				var chipCondition string
				if builder.GetEmvChipCondition() != nil {
					chipCondition = extrautils.IfThenElse(*builder.GetEmvChipCondition() == emvchipcondition.ChipFailPreviousSuccess, "CHIP_FAILED_PREV_SUCCESS", "CHIP_FAILED_PREV_FAILED")
				}
				et.SubElementWithCdataValue(block1, "AccountType", enumutils.GetMapping(target.Portico, builder.GetAccountType()))
				et.SubElementWithCdataValue(block1, "EMVChipCondition", chipCondition)
				et.SubElementWithCdataValue(block1, "MessageAuthenticationCode", builder.GetMessageAuthenticationCode())
				et.SubElementWithCdataValue(block1, "PosSequenceNbr", builder.GetPosSequenceNumber())
				et.SubElementWithCdataValue(block1, "ReversalReasonCode", builder.GetReversalReasonCode().GetValue())
				if !stringutils.IsNullOrEmpty(builder.GetTagData()) {
					tagData, err := et.SubElement(block1, "TagData")
					if err != nil {
						return nil, err
					}
					el, err := et.SubElementWithCdataValue(tagData, "TagValues", builder.GetTagData())
					if err != nil {
						return nil, err
					}
					el.Set("source", "chip")
				}
			} else {
				block1.Append(cardData)
			}
		} else {
			et.SubElementWithCdataValue(trackData, "TokenValue", tokenValue)
		}
	} else if card, ok := builder.GetPaymentMethod().(*paymentmethods.GiftCard); ok {
		// currency
		et.SubElementWithCdataValue(block1, "Currency", builder.GetCurrency())

		// if it's replace put the new card and change the card data name to be old card data
		if builder.GetTransactionType() == transactiontype.Replace {
			newCardData, err := et.SubElement(block1, "NewCardData")
			if err != nil {
				return nil, err
			}
			et.SubElementWithCdataValue(newCardData, builder.GetReplacementCardType(), builder.GetReplacementCardValue())
			et.SubElementWithCdataValue(newCardData, "PIN", builder.GetReplacementCardPin())

			cardData, err = et.Element("OldCardData")
			if err != nil {
				return nil, err
			}
		}
		et.SubElementWithCdataValue(cardData, card.GetValueType(), card.GetValue())
		et.SubElementWithCdataValue(cardData, "PIN", card.GetPIN())

		if builder.GetAliasAction() != aliasaction.Create {
			block1.Append(cardData)
		}
	} else if reference, ok := builder.GetPaymentMethod().(*references.TransactionReference); ok {
		et.SubElementWithCdataValue(block1, "GatewayTxnId", reference.TransactionId)
		et.SubElementWithCdataValue(block1, "ClientTxnId", reference.ClientTransactionId)
	}

	// pin block
	if pinProtected, ok := builder.GetPaymentMethod().(abstractions.IPinProtected); ok {
		if transactionType != transactiontype.Reversal {
			et.SubElementWithCdataValue(block1, "PinBlock", pinProtected.GetPinBlock())
		}
	}

	// encryption
	if encryptable, ok := builder.GetPaymentMethod().(abstractions.IEncryptable); ok {
		encryptionData := encryptable.GetEncryptionData()

		if encryptionData != nil {
			enc, err := et.SubElement(cardData, "EncryptionData")
			if err != nil {
				return nil, err
			}
			et.SubElementWithCdataValue(enc, "Version", encryptionData.GetVersion())
			et.SubElementWithCdataValue(enc, "EncryptedTrackNumber", encryptionData.GetTrackNumber())
			et.SubElementWithCdataValue(enc, "KTB", encryptionData.GetKtb())
			et.SubElementWithCdataValue(enc, "KSN", encryptionData.GetKsn())
		}
	}

	// set token flag
	if _, ok := builder.GetPaymentMethod().(abstractions2.ITokenizable); ok && builder.GetPaymentMethod().GetPaymentMethodType() != paymentmethodtype.ACH {
		et.SubElementWithCdataValue(cardData, "TokenRequest", stringutils.BoolToString(builder.IsRequestMultiUseToken()))

	}

	if builder.IsRequestUniqueToken() {
		uniqueTokenData, err := et.SubElement(cardData, "TokenParameters")
		if err != nil {
			return nil, err
		}
		et.SubElementWithCdataValue(uniqueTokenData, "Mapping", "UNIQUE")
	}

	// balance inquiry type
	if _, ok := builder.GetPaymentMethod().(abstractions2.IBalanceable); ok {
		et.SubElementWithCdataValue(block1, "BalanceInquiryType", builder.GetBalanceInquiryType().GetValue())
	}

	// cpc request
	if builder.IsLevel2Request() {
		et.SubElementWithCdataValue(block1, "CPCReq", "Y")
	}

	// details
	if !stringutils.IsNullOrEmpty(builder.GetCustomerId()) || !stringutils.IsNullOrEmpty(builder.GetDescription()) || !stringutils.IsNullOrEmpty(builder.GetInvoiceNumber()) {
		addons, err := et.SubElement(block1, "AdditionalTxnFields")
		if err != nil {
			return nil, err
		}
		et.SubElementWithCdataValue(addons, "CustomerID", builder.GetCustomerId())
		et.SubElementWithCdataValue(addons, "Description", builder.GetDescription())
		et.SubElementWithCdataValue(addons, "InvoiceNbr", builder.GetInvoiceNumber())
	}

	// ecommerce info
	if builder.GetEcommerceInfo() != nil {
		ecom := builder.GetEcommerceInfo()
		et.SubElementWithCdataValue(block1, "Ecommerce", ecom.GetChannel().GetValue())
		if !stringutils.IsNullOrEmpty(builder.GetInvoiceNumber()) || ecom.GetShipMonth() != nil {
			direct, err := et.SubElement(block1, "DirectMktData")
			if err != nil {
				return nil, err
			}
			et.SubElementWithCdataValue(direct, "DirectMktInvoiceNbr", builder.GetInvoiceNumber())
			et.SubElementWithCdataValue(direct, "DirectMktShipDay", stringutils.IntToString(*ecom.GetShipDay()))
			et.SubElementWithCdataValue(direct, "DirectMktShipMonth", stringutils.IntToString(*ecom.GetShipMonth()))
		}
	}

	// dynamic descriptor
	et.SubElementWithCdataValue(block1, "TxnDescriptor", builder.GetDynamicDescriptor())

	request, err := p.BuildEnvelope(et, transaction, builder.GetClientTransactionId())
	if err != nil {
		return nil, err
	}
	response, err := p.DoTransaction(ctx, request)
	if err != nil {
		return nil, err
	}
	return MapResponse(response, builder.GetPaymentMethod())
}

func (p *PorticoConnector) ManageTransaction(ctx context.Context, builder abstractions4.IManagementBuilder) (abstractions4.ITransaction, error) {
	et := utils.NewElementTree(nil)
	transactionType := builder.GetTransactionType()
	modifier := builder.GetTransactionModifier()
	paymentMethod := builder.GetPaymentMethod()
	paymentMethodType := paymentmethodtype.Nil
	if paymentMethod != nil {
		paymentMethodType = paymentMethod.GetPaymentMethodType()
		if reference, ok := paymentMethod.(*references.TransactionReference); ok {
			paymentMethod = reference.OriginalPaymentMethod
		}

	}

	transactionRaw, err := mapTransactionType(builder)
	if err != nil {
		return nil, err
	}
	transaction, err := et.Element(transactionRaw)
	if err != nil {
		return nil, err
	}

	if transactionType != transactiontype.BatchClose {

		var root *utils.Element
		if transactionType == transactiontype.Reversal ||
			transactionType == transactiontype.Refund ||
			paymentMethodType == paymentmethodtype.Gift ||
			paymentMethodType == paymentmethodtype.ACH ||
			transactionType == transactiontype.Increment {
			root, err = et.SubElement(transaction, "Block1")
			if err != nil {
				return nil, err
			}
		} else {
			root = transaction
		}

		if builder.GetAmount() != nil {
			et.SubElementWithCdataValue(root, "Amt", stringutils.ToCurrencyString(builder.GetAmount()))
		}
		if builder.GetTransactionInitiator() != storedcredentialinitiator.Nil || !stringutils.IsNullOrEmpty(builder.GetCardBrandTransactionId()) {
			cardOnFileData, err := et.SubElement(root, "CardOnFileData")
			if err != nil {
				return nil, err
			}

			if builder.GetTransactionInitiator() == storedcredentialinitiator.CardHolder {
				et.SubElementWithCdataValue(cardOnFileData, "CardOnFile", storedcredentialinitiator.CardHolder.GetValue(target.Portico))
			} else {
				et.SubElementWithCdataValue(cardOnFileData, "CardOnFile", storedcredentialinitiator.Merchant.GetValue(target.Portico))
			}
			et.SubElementWithCdataValue(cardOnFileData, "CardBrandTxnId", builder.GetCardBrandTransactionId())
		}

		if builder.GetAuthAmount() != nil {
			et.SubElementWithCdataValue(root, "AuthAmt", stringutils.ToCurrencyString(builder.GetAuthAmount()))
		}

		if builder.GetGratuity() != nil {
			et.SubElementWithCdataValue(root, "GratuityAmtInfo", stringutils.ToCurrencyString(builder.GetGratuity()))
		}

		if builder.GetSurchargeAmount() != nil {
			et.SubElementWithCdataValue(root, "SurchargeAmtInfo", stringutils.ToCurrencyString(builder.GetSurchargeAmount()))
		}

		et.SubElementWithCdataValue(root, "GatewayTxnId", builder.GetTransactionId())

		if transactionType == transactiontype.Reversal || (paymentMethodType != paymentmethodtype.Nil && paymentMethodType == paymentmethodtype.ACH) {
			// client transaction id
			if transRef, ok := builder.GetPaymentMethod().(*references.TransactionReference); ok {
				_, err := et.SubElementWithCdataValue(root, "ClientTxnId", transRef.GetClientTransactionId())
				if err != nil {
					return nil, err
				}
			}

			// reversal reason code & PosSequenceNumber
			if paymentMethodType == paymentmethodtype.Debit {
				if builder.GetEmvChipCondition() != nil {
					chipCondition := "CHIP_FAILED_PREV_FAILED"
					if *builder.GetEmvChipCondition() == emvchipcondition.ChipFailPreviousSuccess {
						chipCondition = "CHIP_FAILED_PREV_SUCCESS"
					}
					_, err := et.SubElementWithCdataValue(root, "EMVChipCondition", chipCondition)
					if err != nil {
						return nil, err
					}
				}
				_, err := et.SubElementWithCdataValue(root, "ReversalReasonCode", builder.GetReversalReasonCode().GetValue())
				if err != nil {
					return nil, err
				}
				_, err = et.SubElementWithCdataValue(root, "PosSequenceNbr", builder.GetPosSequenceNumber())
				if err != nil {
					return nil, err
				}
				accountType := enumutils.GetMapping(target.Portico, builder.GetAccountType())
				if err != nil {
					return nil, err
				}
				_, err = et.SubElementWithCdataValue(root, "AccountType", accountType)
				if err != nil {
					return nil, err
				}

				// track data
				if paymentMethod != nil {
					track, ok := paymentMethod.(*paymentmethods.DebitTrackData)
					if ok {
						if transactionType == transactiontype.Reversal {
							_, err := et.SubElementWithCdataValue(root, "TrackData", track.GetValue())
							if err != nil {
								return nil, err
							}
						}
						_, err := et.SubElementWithCdataValue(root, "PinBlock", track.GetPinBlock())
						if err != nil {
							return nil, err
						}

						encryptionData := track.GetEncryptionData()
						if encryptionData != nil {
							enc, err := et.SubElement(root, "EncryptionData")
							if err != nil {
								return nil, err
							}
							_, err = et.SubElementWithCdataValue(enc, "Version", encryptionData.GetVersion())
							if err != nil {
								return nil, err
							}
							_, err = et.SubElementWithCdataValue(enc, "EncryptedTrackNumber", encryptionData.GetTrackNumber())
							if err != nil {
								return nil, err
							}
							_, err = et.SubElementWithCdataValue(enc, "KTB", encryptionData.GetKtb())
							if err != nil {
								return nil, err
							}
							_, err = et.SubElementWithCdataValue(enc, "KSN", encryptionData.GetKsn())
							if err != nil {
								return nil, err
							}
						}
					}
				}
			}

			// tag data
			if !stringutils.IsNullOrEmpty(builder.GetTagData()) {
				tagData, err := et.SubElement(root, "TagData")
				if err != nil {
					return nil, err
				}
				el, err := et.SubElementWithCdataValue(tagData, "TagValues", builder.GetTagData())
				if err != nil {
					return nil, err
				}
				el.Set("source", "chip")
			}
		}

		if builder.GetCommercialData() != nil {
			cd := builder.GetCommercialData()
			if modifier == transactionmodifier.LevelII || modifier == transactionmodifier.LevelIII {
				cpc, err := et.SubElement(root, "CPCData")
				if err != nil {
					return nil, err
				}
				et.SubElementWithCdataValue(cpc, "CardHolderPONbr", cd.PONumber)
				et.SubElementWithCdataValue(cpc, "TaxType", hydrateTaxType(cd.TaxType))
				et.SubElementWithCdataValue(cpc, "TaxAmt", stringutils.ToCurrencyString(cd.TaxAmount))
			}

			if modifier == transactionmodifier.LevelIII && paymentMethodType == paymentmethodtype.Credit {
				cdc, err := et.SubElement(root, "CorporateData")
				if err != nil {
					return nil, err
				}
				isVisa := builder.GetCardType() == "Visa"
				dataElementName := "Visa"
				if !isVisa {
					dataElementName = "MC"
				}
				data, err := et.SubElement(cdc, dataElementName)
				if err != nil {
					return nil, err
				}
				if cd.LineItems != nil && len(cd.LineItems) > 0 {
					lineItems, _ := et.SubElement(data, "LineItems")
					for _, lineItem := range cd.LineItems {
						lineItemElement, err := et.SubElement(lineItems, "LineItemDetail")
						if err != nil {
							return nil, err
						}

						et.SubElementWithCdataValue(lineItemElement, "ItemDescription", lineItem.Description)
						et.SubElementWithCdataValue(lineItemElement, "ProductCode", lineItem.ProductCode)
						et.SubElementWithCdataValue(lineItemElement, "Quantity", lineItem.Quantity.StringFixed(0))
						et.SubElementWithCdataValue(lineItemElement, "ItemTotalAmt", stringutils.ToCurrencyString(lineItem.TotalAmount))
						et.SubElementWithCdataValue(lineItemElement, "UnitOfMeasure", lineItem.UnitOfMeasure)

						if !isVisa {
							continue
						}

						et.SubElementWithCdataValue(lineItemElement, "ItemCommodityCode", lineItem.CommodityCode)
						et.SubElementWithCdataValue(lineItemElement, "UnitCost", stringutils.ToCurrencyString(lineItem.UnitCost))
						et.SubElementWithCdataValue(lineItemElement, "VATTaxAmt", stringutils.ToCurrencyString(lineItem.TaxAmount))
						et.SubElementWithCdataValue(lineItemElement, "DiscountAmt", stringutils.ToCurrencyString(lineItem.DiscountDetails.DiscountAmount))
					}
				}

				if isVisa {
					et.SubElementWithCdataValue(data, "SummaryCommodityCode", cd.SummaryCommodityCode)
					et.SubElementWithCdataValue(data, "DiscountAmt", stringutils.ToCurrencyString(cd.DiscountAmount))
					et.SubElementWithCdataValue(data, "FreightAmt", stringutils.ToCurrencyString(cd.FreightAmount))
					et.SubElementWithCdataValue(data, "DutyAmt", stringutils.ToCurrencyString(cd.DutyAmount))
					et.SubElementWithCdataValue(data, "DestinationPostalZipCode", cd.DestinationPostalCode)
					et.SubElementWithCdataValue(data, "ShipFromPostalZipCode", cd.OriginPostalCode)
					et.SubElementWithCdataValue(data, "DestinationCountryCode", cd.DestinationCountryCode)
					et.SubElementWithCdataValue(data, "InvoiceRefNbr", cd.CustomerReferenceId)

					taxAmount := cd.TaxAmount
					if cd.AdditionalTaxDetails != nil && cd.AdditionalTaxDetails.TaxAmount != nil {
						taxAmount = cd.AdditionalTaxDetails.TaxAmount
					}
					et.SubElementWithCdataValue(data, "VATTaxAmtFreight", stringutils.ToCurrencyString(taxAmount))

					if cd.OrderDate != "" {
						et.SubElementWithCdataValue(data, "OrderDate", stringutils.DateStringFormatted(cd.OrderDate, "2006-01-02T15:04:05"))
					}

					if cd.AdditionalTaxDetails != nil {
						et.SubElementWithCdataValue(data, "VATTaxRateFreight", cd.AdditionalTaxDetails.TaxRate.StringFixed(4))
					}
				}
			}
		} else if builder.GetTransactionType() == transactiontype.Edit &&
			builder.GetTransactionModifier() == transactionmodifier.LevelII {
			cpc, err := et.SubElement(root, "CPCData")
			if err != nil {
				return nil, err
			}
			et.SubElementWithCdataValue(cpc, "CardHolderPONbr", builder.GetPoNumber())
			et.SubElementWithCdataValue(cpc, "TaxType", hydrateTaxType(builder.GetTaxType()))
			et.SubElementWithCdataValue(cpc, "TaxAmt", stringutils.ToCurrencyString(builder.GetTaxAmount()))
		}

		if transactionType == transactiontype.TokenUpdate ||
			transactionType == transactiontype.TokenDelete ||
			transactionType == transactiontype.Capture {

			if token, ok := builder.GetPaymentMethod().(abstractions2.ITokenizable); ok {

				et.SubElementWithCdataValue(root, "TokenValue", token.GetToken())

				tokenActions, err := et.SubElement(root, "TokenActions")
				if err != nil {
					return nil, err
				}
				if transactionType == transactiontype.TokenUpdate {
					if card, ok := builder.GetPaymentMethod().(*paymentmethods.CreditCardData); ok {
						setElement, err := et.SubElement(tokenActions, "Set")
						if err != nil {
							return nil, err
						}
						expMonthElement, err := et.SubElement(setElement, "Attribute")
						if err != nil {
							return nil, err
						}
						et.SubElementWithCdataValue(expMonthElement, "Name", "ExpMonth")
						et.SubElementWithCdataValue(expMonthElement, "Value", stringutils.IntToString(*card.GetExpMonth()))

						expYearElement, err := et.SubElement(setElement, "Attribute")
						if err != nil {
							return nil, err
						}
						et.SubElementWithCdataValue(expYearElement, "Name", "ExpYear")
						et.SubElementWithCdataValue(expYearElement, "Value", stringutils.IntToString(*card.GetExpYear()))
					}
				} else {
					et.SubElement(tokenActions, "Delete")
				}
			}
		}

		if !stringutils.IsNullOrEmpty(builder.GetCustomerId()) || !stringutils.IsNullOrEmpty(builder.GetDescription()) || !stringutils.IsNullOrEmpty(builder.GetInvoiceNumber()) {
			addons, err := et.SubElement(root, "AdditionalTxnFields")
			if err != nil {
				return nil, err
			}
			et.SubElementWithCdataValue(addons, "CustomerID", builder.GetCustomerId())
			et.SubElementWithCdataValue(addons, "Description", builder.GetDescription())
			et.SubElementWithCdataValue(addons, "InvoiceNbr", builder.GetInvoiceNumber())
		}
	}

	request, err := p.BuildEnvelope(et, transaction, builder.GetClientTransactionId())
	if err != nil {
		return nil, err
	}
	response, err := p.DoTransaction(ctx, request)
	if err != nil {
		return nil, err
	}
	return MapResponse(response, builder.GetPaymentMethod())
}

func mapTransactionType(builder abstractions4.ITransactionBuilder) (string, error) {
	modifier := builder.GetTransactionModifier()
	var paymentMethodType *paymentmethodtype.PaymentMethodType
	if builder.GetPaymentMethod() != nil {
		pmt := builder.GetPaymentMethod().GetPaymentMethodType()
		paymentMethodType = &pmt
	}

	switch builder.GetTransactionType() {
	case transactiontype.BatchClose:
		return "BatchClose", nil
	case transactiontype.Decline:
		if modifier == transactionmodifier.ChipDecline {
			return "ChipCardDecline", nil
		} else if modifier == transactionmodifier.FraudDecline {
			return "OverrideFraudDecline", nil
		}
		return "", errors.New("unsupported transaction exception")
	case transactiontype.Verify:
		if modifier == transactionmodifier.EncryptedMobile {
			return "", errors.New("transaction not supported for this payment method")
		}
		return "CreditAccountVerify", nil
	case transactiontype.Capture:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				return "CreditAddToBatch", nil
			case paymentmethodtype.Debit:
				return "DebitAddToBatch", nil
			}
			return "", errors.New("transaction not supported for this payment method")
		}
	case transactiontype.Auth:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				switch modifier {
				case transactionmodifier.Additional:
					return "CreditAdditionalAuth", nil
				case transactionmodifier.Incremental:
					return "CreditIncrementalAuth", nil
				case transactionmodifier.Offline:
					return "CreditOfflineAuth", nil
				case transactionmodifier.Recurring:
					return "RecurringBillingAuth", nil
				}
				return "CreditAuth", nil
			case paymentmethodtype.Recurring:
				return "RecurringBillingAuth", nil
			case paymentmethodtype.Debit:
				return "DebitAuth", nil
			}
			return "", errors.New("transaction not supported for this payment method")
		}
	case transactiontype.Sale:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				switch modifier {
				case transactionmodifier.Offline:
					return "CreditOfflineSale", nil
				case transactionmodifier.Recurring:
					return "RecurringBilling", nil
				}
				return "CreditSale", nil
			case paymentmethodtype.Recurring:
				// Cast to RecurringPaymentMethod must be done in the actual implementation
				// if ((RecurringPaymentMethod)builder.GetPaymentMethod()).GetPaymentType() == "ACH" {
				// 	return "CheckSale", nil
				// }
				return "RecurringBilling", nil
			case paymentmethodtype.Debit:
				return "DebitSale", nil
			case paymentmethodtype.Cash:
				return "CashSale", nil
			case paymentmethodtype.ACH:
				return "CheckSale", nil
			case paymentmethodtype.EBT:
				switch modifier {
				case transactionmodifier.CashBack:
					return "EBTCashBackPurchase", nil
				case transactionmodifier.Voucher:
					return "EBTVoucherPurchase", nil
				}
				return "EBTFSPurchase", nil
			case paymentmethodtype.Gift:
				return "GiftCardSale", nil
			}
			return "", errors.New("unsupported transaction exception")
		}
	case transactiontype.Refund:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				return "CreditReturn", nil
			case paymentmethodtype.Debit:
				// Cast to TransactionReference must be done in the actual implementation
				// if builder.GetPaymentMethod() instanceof TransactionReference {
				// 	return "", errors.New("unsupported transaction exception")
				// }
				return "DebitReturn", nil
			case paymentmethodtype.Cash:
				return "CashReturn", nil
			case paymentmethodtype.EBT:
				// Cast to TransactionReference must be done in the actual implementation
				// if builder.GetPaymentMethod() instanceof TransactionReference {
				// 	return "", errors.New("unsupported transaction exception")
				// }
				return "EBTFSReturn", nil
			}
			return "", errors.New("unsupported transaction exception")
		}
	case transactiontype.Reversal:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				return "CreditReversal", nil
			case paymentmethodtype.Debit:
				return "DebitReversal", nil
			case paymentmethodtype.Gift:
				return "GiftCardReversal", nil
			case paymentmethodtype.EBT:
				return "EBTFSReversal", nil
			}
			return "", errors.New("unsupported transaction exception")
		}
	case transactiontype.Edit:
		if modifier == transactionmodifier.LevelII || modifier == transactionmodifier.LevelIII {
			return "CreditCPCEdit", nil
		}
		return "CreditTxnEdit", nil
	case transactiontype.Void:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				return "CreditVoid", nil
			case paymentmethodtype.Debit:
				return "DebitVoid", nil
			case paymentmethodtype.Gift:
				return "GiftCardVoid", nil
			}
			return "", errors.New("unsupported transaction exception")
		}
	case transactiontype.AddValue:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.Credit:
				return "PrePaidAddValue", nil
			case paymentmethodtype.Debit:
				return "DebitAddValue", nil
			case paymentmethodtype.Gift:
				return "GiftCardAddValue", nil
			}
			return "", errors.New("unsupported transaction exception")
		}

	case transactiontype.Balance:
		if paymentMethodType != nil {
			switch *paymentMethodType {
			case paymentmethodtype.EBT:
				return "EBTBalanceInquiry", nil
			case paymentmethodtype.Gift:
				return "GiftCardBalance", nil
			case paymentmethodtype.Credit:
				return "PrePaidBalanceInquiry", nil
			}
			return "", errors.New("unsupported transaction exception")
		}

	case transactiontype.BenefitWithdrawal:
		return "EBTCashBenefitWithdrawal", nil

	case transactiontype.Activate:
		return "GiftCardActivate", nil
	case transactiontype.Alias:
		return "GiftCardAlias", nil
	case transactiontype.Deactivate:
		return "GiftCardDeactivate", nil
	case transactiontype.Replace:
		return "GiftCardReplace", nil
	case transactiontype.Reward:
		return "GiftCardReward", nil
	case transactiontype.Increment:
		return "CreditIncrementalAuth", nil
	case transactiontype.Tokenize:
		return "Tokenize", nil
	case transactiontype.TokenUpdate, transactiontype.TokenDelete:
		return "ManageTokens", nil
	default:
		return "", errors.New("unsupported transaction exception")
	}
	return "", errors.New("unsupported transaction exception")
}

func (p *PorticoConnector) ProcessReport(ctx context.Context, builder abstractions5.IReportBuilder) (abstractions4.ITransaction, error) {
	et := utils.NewElementTree(nil)
	reportRequestType, err := mapReportType(builder.GetReportType())
	if err != nil {
		return nil, err
	}
	transaction, err := et.Element(reportRequestType)
	if err != nil {
		return nil, err
	}

	if builder.GetTimeZoneConversion() != "" {
		_, err := et.SubElementWithCdataValue(transaction, "TzConversion", builder.GetTimeZoneConversion().GetValue())
		if err != nil {
			return nil, err
		}
	}

	if trb, ok := builder.(*builders.TransactionReportBuilder); ok {

		if trb.GetDeviceId() != "" {
			_, err := et.SubElementWithCdataValue(transaction, "DeviceId", trb.GetDeviceId())
			if err != nil {
				return nil, err
			}
		}

		if trb.GetStartDate() != "" {
			_, err := et.SubElementWithCdataValue(transaction, "RptStartUtcDT", stringutils.DateStringFormatted(trb.GetStartDate(), time.RFC3339))
			if err != nil {
				return nil, err
			}
		}

		if trb.GetEndDate() != "" {
			_, err := et.SubElementWithCdataValue(transaction, "RptEndUtcDT", stringutils.DateStringFormatted(trb.GetEndDate(), time.RFC3339))
			if err != nil {
				return nil, err
			}
		}

		if trb.GetTransactionId() != "" {
			_, err := et.SubElementWithCdataValue(transaction, "TxnId", trb.GetTransactionId())
			if err != nil {
				return nil, err
			}
		}

		if trb.GetSearchCriteria() != nil {
			criteriaNode, err := et.SubElement(transaction, "Criteria")
			if err != nil {
				return nil, err
			}

			for property, value := range trb.GetSearchCriteria() {
				_, err := et.SubElementWithCdataValue(criteriaNode, property, value)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	request, err := p.BuildEnvelope(et, transaction, "")
	if err != nil {
		return nil, err
	}
	response, err := p.DoTransaction(ctx, request)
	if err != nil {
		return nil, err
	}

	return MapReportResponse(response, builder.GetReportType())
}

func MapReportResponse(rawResponse string, reportType reporttype.ReportType) (abstractions4.ITransaction, error) {
	root, err := utils.ParseXml([]byte(rawResponse))
	if err != nil {
		return nil, err
	}

	posResponse := root.Get("PosResponse")
	if posResponse == nil {
		return nil, errors.New("Unable to parse response xml")
	}
	acceptedCodes := []string{"00", "0"}

	gatewayRspCode := helpers.NormalizeResponse(posResponse.GetString("GatewayRspCode"))
	gatewayRspText := posResponse.GetString("GatewayRspMsg")

	found := false
	for _, code := range acceptedCodes {
		if code == gatewayRspCode {
			found = true
			break
		}
	}
	if !found {
		return nil, NewGatewayResponseError(
			fmt.Sprintf("Unexpected Gateway Response: %s - %s", gatewayRspCode, gatewayRspText), gatewayRspCode,
		)
	}

	rString, err := mapReportType(reportType)

	if err != nil {
		return nil, err
	}

	doc := root.Get(rString)

	switch reportType {
	case reporttype.Activity:
		report := transactionsummary.NewActivityReport() // Suppose this creates an instance of ActivityReport.
		for _, detail := range doc.GetAllByTag("Details") {
			report.AddSummary(*transactionsummary.HydrateTransactionSummary(detail)) // Implement HydrateTransactionSummary accordingly.
		}
		return report, nil

	case reporttype.FindTransactions:
		report := transactionsummary.NewActivityReport() // Suppose this creates an instance of ActivityReport.
		for _, detail := range doc.GetAllByTag("Transactions") {
			report.AddSummary(*transactionsummary.HydrateTransactionSummary(detail)) // Implement HydrateTransactionSummary accordingly.
		}
		return report, nil
	case reporttype.TransactionDetail:
		return transactionsummary.HydrateTransactionSummary(doc), nil
	}

	// Fallback for unhandled report types.
	return nil, errors.New("Unsupported Report Type")
}

func mapReportType(t reporttype.ReportType) (string, error) {
	switch t {
	case reporttype.Activity:
		return "ReportActivity", nil
	case reporttype.TransactionDetail:
		return "ReportTxnDetail", nil
	case reporttype.FindTransactions:
		return "FindTransactions", nil
	default:
		return "", errors.New("unsupported transaction type")
	}
}

func IsAppleOrGooglePay(paymentDataSource entities.PaymentDataSourceType) bool {
	return paymentDataSource == entities.APPLEPAY ||
		paymentDataSource == entities.APPLEPAYAPP ||
		paymentDataSource == entities.APPLEPAYWEB ||
		paymentDataSource == entities.GOOGLEPAYAPP ||
		paymentDataSource == entities.GOOGLEPAYWEB
}

func (p *PorticoConnector) BuildEnvelope(et *utils.ElementTree, transaction *utils.Element, clientTransactionId string) (string, error) {
	et.AddNamespace("soap", "http://schemas.xmlsoap.org/soap/envelope/")
	et.AddNamespace("xsi", "http://www.w3.org/2001/XMLSchema-instance")
	et.AddNamespace("xsd", "http://www.w3.org/2001/XMLSchema")

	envelope, err := et.Element("soap:Envelope")
	if err != nil {
		return "", err
	}
	envelope.Set("xmlns:soap", "http://schemas.xmlsoap.org/soap/envelope/")

	body, err := et.SubElement(envelope, "soap:Body")
	if err != nil {
		return "", err
	}

	request, err := et.SubElement(body, "PosRequest")
	if err != nil {
		return "", err
	}
	request.Set("xmlns", "http://Hps.Exchange.PosGateway")

	version1, err := et.SubElement(request, "Ver1.0")
	if err != nil {
		return "", err
	}

	// header
	header, err := et.SubElement(version1, "Header")
	if err != nil {
		return "", err
	}
	_, _ = et.SubElementWithCdataValue(header, "SecretAPIKey", p.SecretApiKey)
	if p.SiteId != "" {
		_, _ = et.SubElementWithCdataValue(header, "SiteId", p.SiteId)
	}
	if p.LicenseId != "" {
		_, _ = et.SubElementWithCdataValue(header, "LicenseId", p.LicenseId)
	}
	if p.DeviceId != "" {
		_, _ = et.SubElementWithCdataValue(header, "DeviceId", p.DeviceId)
	}
	_, _ = et.SubElementWithCdataValue(header, "UserName", p.Username)
	_, _ = et.SubElementWithCdataValue(header, "Password", p.Password)
	_, _ = et.SubElementWithCdataValue(header, "DeveloperID", p.DeveloperId)
	_, _ = et.SubElementWithCdataValue(header, "VersionNbr", p.VersionNumber)
	_, _ = et.SubElementWithCdataValue(header, "ClientTxnId", clientTransactionId)
	_, _ = et.SubElementWithCdataValue(header, "PosReqDT", GetPosReqDT()) // Assuming getPosReqDT is a function that returns a date string
	sdkNameVersionValue := p.SdkNameVersion
	if p.SdkNameVersion == "" {
		sdkNameVersionValue = "go;version=" + packageutils.GetPackageVersion()
	}
	_, _ = et.SubElementWithCdataValue(header, "SDKNameVersion", sdkNameVersionValue)

	safData, err := et.SubElement(header, "SAFData")
	if err != nil {
		return "", err
	}
	_, _ = et.SubElementWithCdataValue(safData, "SAFIndicator", extrautils.IfThenElse(p.IsSAFDataSupported, "Y", "N"))
	_, _ = et.SubElementWithCdataValue(safData, "SAFOrigDT", time.Now().Format(time.RFC3339))

	// Transaction
	trans, err := et.SubElement(version1, "Transaction")
	if err != nil {
		return "", err
	}
	trans.Append(transaction)
	res, err := et.ToString(envelope)
	if err != nil {
		return "", err
	}
	return "<?xml version=\"1.0\" encoding=\"utf-8\"?>" + res, nil
}

func GetPosReqDT() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}

func MapResponse(rawResponse string, paymentMethod abstractions.IPaymentMethod) (*transactions.Transaction, error) {
	result := transactions.NewTransaction()

	root, err := utils.ParseXml([]byte(rawResponse))
	if err != nil {
		return nil, err
	}

	posResponse := root.Get("PosResponse")
	if posResponse == nil {
		return nil, errors.New("Unable to parse response xml")
	}
	acceptedCodes := []string{"00", "0", "85", "10", "02", "2"}

	gatewayRspCode := helpers.NormalizeResponse(posResponse.GetString("GatewayRspCode"))
	gatewayRspText := posResponse.GetString("GatewayRspMsg")
	cardType := posResponse.GetString("CardType")

	found := false
	for _, code := range acceptedCodes {
		if code == gatewayRspCode {
			found = true
			break
		}
	}
	if !found {
		return nil, NewGatewayResponseError(
			fmt.Sprintf("Unexpected Gateway Response: %s - %s", gatewayRspCode, gatewayRspText), gatewayRspCode,
		)
	}

	result.AuthorizedAmount = posResponse.GetDecimal("AuthAmt")
	result.AvailableBalance = posResponse.GetDecimal("AvailableBalance")
	result.AvsResponseCode = posResponse.GetString("AVSRsltCode")
	result.AvsResponseMessage = posResponse.GetString("AVSRsltText")
	result.BalanceAmount = posResponse.GetDecimal("BalanceAmt")
	result.CardType = cardType
	result.CardLast4 = posResponse.GetString("TokenPANLast4")
	result.CavvResponseCode = posResponse.GetString("CAVVResultCode")
	result.CommercialIndicator = posResponse.GetString("CPCInd")
	result.CvnResponseCode = posResponse.GetString("CVVRsltCode")
	result.CvnResponseMessage = posResponse.GetString("CVVRsltText")
	result.EmvIssuerResponse = posResponse.GetString("EMVIssuerResp")
	result.PointsBalanceAmount = posResponse.GetDecimal("PointsBalanceAmt")
	result.RecurringDataCode = posResponse.GetString("RecurringDataCode")
	result.ReferenceNumber = posResponse.GetString("RefNbr")
	result.CardBrandTransactionId = posResponse.GetString("CardBrandTxnId")

	responseCode := helpers.NormalizeResponse(posResponse.GetString("RspCode"))
	responseText := posResponse.GetString("RspText", "RspMessage")
	if responseCode != "" {
		result.ResponseCode = responseCode
	} else {
		result.ResponseCode = gatewayRspCode
	}
	if responseText != "" {
		result.ResponseMessage = responseText
	} else {
		result.ResponseMessage = gatewayRspText
	}
	result.TransactionDescriptor = posResponse.GetString("TxnDescriptor")
	result.ResponseDate = posResponse.GetDateWithFormat("yyyy-MM-dd'T'HH:mm:ss", "RspDT")
	result.HostResponseDate = posResponse.GetDateWithFormat("yyyy-MM-dd'T'HH:mm:ss", "HostRspDT")

	if paymentMethod != nil {
		reference := references.NewTransactionReference()
		reference.PaymentMethodType = paymentMethod.GetPaymentMethodType()
		reference.TransactionId = posResponse.GetString("GatewayTxnId")
		reference.ClientTransactionId = posResponse.GetString("ClientTxnId")
		reference.AuthCode = posResponse.GetString("AuthCode")

		result.TransactionReference = reference
	}

	if posResponse.Has("CardData") {
		card := paymentmethods.NewGiftCard()
		card.SetAlias(posResponse.GetString("Alias"))
		card.SetNumber(posResponse.GetString("CardNbr"))
		card.SetPIN(posResponse.GetString("PIN"))
		result.SetGiftCard(card)

	}

	if posResponse.Has("TokenData") {
		result.Token = posResponse.GetString("TokenValue")
	}

	if posResponse.Has("BatchId") {
		summary := entities.NewBatchSummary()

		summary.BatchId = posResponse.GetInt("BatchId")
		summary.TransactionCount = posResponse.GetInt("TxnCnt")
		summary.TotalAmount = posResponse.GetDecimal("TotalAmt")
		summary.SequenceNumber = posResponse.GetString("BatchSeqNbr")
		result.BatchSummary = summary
	}

	if posResponse.Has("DebitMac") {
		debitMac := entities.NewDebitMac()

		debitMac.TransactionCode = posResponse.GetString("TransactionCode")
		debitMac.TransmissionNumber = posResponse.GetString("TransmissionNumber")
		debitMac.BankResponseCode = posResponse.GetString("BankResponseCode")
		debitMac.MacKey = posResponse.GetString("MacKey")
		debitMac.PinKey = posResponse.GetString("PinKey")
		debitMac.FieldKey = posResponse.GetString("FieldKey")
		debitMac.TraceNumber = posResponse.GetString("TraceNumber")
		debitMac.MessageAuthenticationCode = posResponse.GetString("MessageAuthenticationCode")
		result.DebitMac = debitMac

		// Add the track data for debit interact
		result.TransactionReference.OriginalPaymentMethod = paymentMethod
	}

	if posResponse.Has("AdditionalDuplicateData") {
		additionalDuplicateData := entities.NewAdditionalDuplicateData()

		additionalDuplicateData.OriginalGatewayTxnId = posResponse.GetString("OriginalGatewayTxnId")
		additionalDuplicateData.OriginalRspDT = posResponse.GetDate("yyyy-MM-dd'T'HH:mm:ss", "OriginalRspDT")
		additionalDuplicateData.OriginalClientTxnId = posResponse.GetString("OriginalClientTxnId")
		additionalDuplicateData.OriginalAuthCode = posResponse.GetString("OriginalAuthCode")
		additionalDuplicateData.OriginalRefNbr = posResponse.GetString("OriginalRefNbr")
		additionalDuplicateData.OriginalAuthAmt = posResponse.GetDecimal("OriginalAuthAmt")
		additionalDuplicateData.OriginalCardType = posResponse.GetString("OriginalCardType")
		additionalDuplicateData.OriginalCardNbrLast4 = posResponse.GetString("OriginalCardNbrLast4")

		result.AdditionalDuplicateData = additionalDuplicateData
	}

	return result, nil
}

func GetToken(paymentMethod abstractions.IPaymentMethod) string {
	if tokenizable, ok := paymentMethod.(abstractions2.ITokenizable); ok {
		tokenValue := tokenizable.GetToken()
		if tokenValue != "" {
			return tokenValue
		}
		return ""
	}
	return ""
}

func hydrateTaxType(taxType taxtype.TaxType) string {
	switch taxType {
	case taxtype.NotUsed:
		return "NOTUSED"
	case taxtype.SalesTax:
		return "SALESTAX"
	case taxtype.TaxExempt:
		return "TAXEXEMPT"
	default:
		return ""
	}
}
