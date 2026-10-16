package entrymethod

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type EntryMethod string

const (
	Manual                                              EntryMethod = "manual"
	Swipe                                               EntryMethod = "swipe"
	Proximity                                           EntryMethod = "proximity"
	Unspecified                                         EntryMethod = "00"
	MagneticStripeAndMSRFallback                        EntryMethod = "02"
	EMVIntegratedChipCard                               EntryMethod = "05"
	EmvContactlessCard                                  EntryMethod = "07"
	CredentialOnFile                                    EntryMethod = "10"
	TechnicalFallback                                   EntryMethod = "80"
	ProximityVisaPayWaveMsdORPayPassMagORAmexExpressPay EntryMethod = "91"
	ContactlessEMV                                      EntryMethod = "contactlessEMV"
	ContactEMV                                          EntryMethod = "contactEMV"
	ContactlessRFID                                     EntryMethod = "contactlessRFID"
	QrCode                                              EntryMethod = "QrCode"
	ContactlessRfidRingTechnology                       EntryMethod = "contactlessRfidRingTechnology"
	BarCode                                             EntryMethod = "3"
	ManualDriverLicense                                 EntryMethod = "4"
	NoTrackData                                         EntryMethod = "G"
	ECommerce                                           EntryMethod = "ecommerce"
	SecureEcommerce                                     EntryMethod = "secureEcommerce"
	CardOnFileEcommerce                                 EntryMethod = "cardOnFileEcommerce"
)

func (em EntryMethod) GetBytes() []byte {
	return []byte(em)
}

func (em EntryMethod) GetValue() string {
	return string(em)
}

func (em EntryMethod) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		Manual,
		Swipe,
		Proximity,
		Unspecified,
		MagneticStripeAndMSRFallback,
		EMVIntegratedChipCard,
		EmvContactlessCard,
		CredentialOnFile,
		TechnicalFallback,
		ProximityVisaPayWaveMsdORPayPassMagORAmexExpressPay,
		ContactlessEMV,
		ContactEMV,
		ContactlessRFID,
		QrCode,
		ContactlessRfidRingTechnology,
		BarCode,
		ManualDriverLicense,
		NoTrackData,
		ECommerce,
		SecureEcommerce,
		CardOnFileEcommerce,
	}
}
