package alternativepaymenttype

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type AlternativePaymentType string

const (
	ASTROPAY_DIRECT                           AlternativePaymentType = "astropaydirect"
	AURA                                      AlternativePaymentType = "aura"
	BALOTO_CASH                               AlternativePaymentType = "baloto"
	BANAMEX                                   AlternativePaymentType = "banamex"
	BANCA_AV_VILLAS                           AlternativePaymentType = "bancaavvillas"
	BANCA_CAJA_SOCIAL                         AlternativePaymentType = "bancacajasocial"
	BANCO_GNB_SUDAMERIS                       AlternativePaymentType = "bancagnbsudameris"
	BANCO_CONSORCIO                           AlternativePaymentType = "bancoconsorcio"
	BANCO_COOPERATIVO_COOPCENTRAL             AlternativePaymentType = "bancocooperativocoopcentral"
	BANCO_CORPBANCA                           AlternativePaymentType = "bancocorpbanca"
	BANCO_DE_BOGOTA                           AlternativePaymentType = "bancodebogota"
	BANCO_DE_CHILE_EDWARDS_CITI               AlternativePaymentType = "bancodechile"
	BANCO_DE_CHILE_CASH                       AlternativePaymentType = "bancodechilecash"
	BANCO_DE_OCCIDENTE                        AlternativePaymentType = "bancodeoccidente"
	BANCO_DE_OCCIDENTE_CASH                   AlternativePaymentType = "bancodeoccidentecash"
	BANCO_DO_BRASIL                           AlternativePaymentType = "bancodobrasil"
	BANCO_FALABELLA_Chile                     AlternativePaymentType = "bancofalabellachile"
	BANCO_FALABELLA_Columbia                  AlternativePaymentType = "bancofalabellacolumbia"
	BANCO_INTERNATIONAL                       AlternativePaymentType = "bancointernational"
	BANCO_PICHINCHA                           AlternativePaymentType = "bancopichincha"
	BANCO_POPULAR                             AlternativePaymentType = "bancopopular"
	BANCO_PROCREDIT                           AlternativePaymentType = "bancoprocredit"
	BANCO_RIPLEY                              AlternativePaymentType = "bancoripley"
	BANCO_SANTANDER                           AlternativePaymentType = "bancosantander"
	BANCO_SANTANDER_BANEFE                    AlternativePaymentType = "bancosantanderbanefe"
	BANCO_SECURITY                            AlternativePaymentType = "bancosecurity"
	BANCOBICE                                 AlternativePaymentType = "bancobice"
	BANCOESTADO                               AlternativePaymentType = "bancoestado"
	BANCOLOMBIA                               AlternativePaymentType = "bancolombia"
	BANCOMER                                  AlternativePaymentType = "bancomer"
	BANCONTACT_MR_CASH                        AlternativePaymentType = "bancontact"
	BANCOOMEVA                                AlternativePaymentType = "bancoomeva"
	BANK_ISLAM                                AlternativePaymentType = "bankislam"
	BANK_TRANSFER                             AlternativePaymentType = "banktransfer"
	BBVA_Chile                                AlternativePaymentType = "bbvachile"
	BBVA_Columbia                             AlternativePaymentType = "bbvacolumbia"
	BCI_TBANC                                 AlternativePaymentType = "bcitbanc"
	BITPAY                                    AlternativePaymentType = "bitpay"
	BOLETO_BANCARIO                           AlternativePaymentType = "boletobancario_"
	BRADESCO                                  AlternativePaymentType = "bradesco"
	CABAL                                     AlternativePaymentType = "cabal_"
	CARTAO_MERCADOLIVRE                       AlternativePaymentType = "cartaomercadolivre"
	CARULLA                                   AlternativePaymentType = "carulla"
	CENCOSUD                                  AlternativePaymentType = "cencosud"
	CHINA_UNION_PAY                           AlternativePaymentType = "unionpay"
	CIMB_CLICKS                               AlternativePaymentType = "cimbclicks"
	CITIBANK                                  AlternativePaymentType = "citibank"
	CMR                                       AlternativePaymentType = "cmr"
	COLPATRIA                                 AlternativePaymentType = "colpatria"
	COOPEUCH                                  AlternativePaymentType = "coopeuch"
	CORPBANCA                                 AlternativePaymentType = "corpbanca"
	DANSKE_BANK                               AlternativePaymentType = "danskebank"
	DAVIVIENDA                                AlternativePaymentType = "davivienda"
	DRAGONPAY                                 AlternativePaymentType = "dragonpay"
	EASYPAY                                   AlternativePaymentType = "easypay"
	EFECTY                                    AlternativePaymentType = "efecty"
	ELO                                       AlternativePaymentType = "elo"
	EMPRESA_DE_ENERGIA_DEL_QUINDIO            AlternativePaymentType = "empresadeenergia"
	ENETS                                     AlternativePaymentType = "enets"
	ENTERCASH                                 AlternativePaymentType = "entercash"
	E_PAY_PETRONAS                            AlternativePaymentType = "epaypetronas"
	EPS                                       AlternativePaymentType = "EPS"
	ESTONIAN_ONLINE_BANK_TRANSFER             AlternativePaymentType = "estonianbanks"
	FINLAND_ONLINE_BANK_TRANSFER              AlternativePaymentType = "finlandonlinebt"
	GIROPAY                                   AlternativePaymentType = "giropay"
	HANDELSBANKEN                             AlternativePaymentType = "handelsbanken"
	HELM_BANK                                 AlternativePaymentType = "helm"
	HIPERCARD                                 AlternativePaymentType = "hipercard"
	HONG_LEONG_BANK                           AlternativePaymentType = "hongleongbank"
	IDEAL                                     AlternativePaymentType = "ideal"
	INDONESIA_ATM                             AlternativePaymentType = "indonesiaatm"
	INSTANT_TRANSFER                          AlternativePaymentType = "instanttransfer"
	INTERNATIONAL_PAY_OUT                     AlternativePaymentType = "intpayout"
	ITAU_BRAZIL                               AlternativePaymentType = "itaubrazil"
	ITAU_CHILE                                AlternativePaymentType = "itauchile"
	LATVIANBT                                 AlternativePaymentType = "latvianbt"
	LINK                                      AlternativePaymentType = "link"
	LITHUANIAN_ONLINE_BANK_TRANSFER           AlternativePaymentType = "lituanianbt"
	MAGNA                                     AlternativePaymentType = "magna"
	MAXIMA                                    AlternativePaymentType = "maxima"
	MAYBANK2U                                 AlternativePaymentType = "maybank2u"
	MULTIBANCO                                AlternativePaymentType = "multibanco"
	MYBANK                                    AlternativePaymentType = "mybank"
	MYCLEAR_FPX                               AlternativePaymentType = "myclearfpx"
	NARANJA                                   AlternativePaymentType = "naranja"
	NARVESEN_LIETUVOS_SPAUDA                  AlternativePaymentType = "narvesen"
	NATIVA                                    AlternativePaymentType = "nativa"
	NORDEA                                    AlternativePaymentType = "nordea"
	OSUUSPANKKI                               AlternativePaymentType = "osuuspankki"
	OXXO                                      AlternativePaymentType = "oxxo"
	PAGO_FACIL                                AlternativePaymentType = "pagofacil"
	PAYPAL                                    AlternativePaymentType = "paypal"
	PAYPOST_LIETUVOS_PASTAS                   AlternativePaymentType = "paypost"
	PAYSAFECARD                               AlternativePaymentType = "paysafecard"
	PAYSBUY_CASH                              AlternativePaymentType = "paysbuy"
	PAYSERA                                   AlternativePaymentType = "paysera"
	PAYU                                      AlternativePaymentType = "payu"
	PERLAS                                    AlternativePaymentType = "perlas"
	POLI                                      AlternativePaymentType = "poli"
	POLISH_PAYOUT                             AlternativePaymentType = "polishpayout"
	POP_PANKKI                                AlternativePaymentType = "poppankki"
	POSTFINANCE                               AlternativePaymentType = "postfinance"
	PRESTO                                    AlternativePaymentType = "presto"
	PROVINCIA_NET                             AlternativePaymentType = "provincianet"
	PRZELEWY24                                AlternativePaymentType = "p24"
	PSE                                       AlternativePaymentType = "pse"
	QIWI                                      AlternativePaymentType = "qiwi"
	QIWI_PAYOUT                               AlternativePaymentType = "qiwipayout"
	RAPI_PAGO                                 AlternativePaymentType = "rapipago"
	REDPAGOS                                  AlternativePaymentType = "redpagos"
	RHB_BANK                                  AlternativePaymentType = "rhbbank"
	SAASTOPANKKI                              AlternativePaymentType = "sasstopankki"
	SAFETYPAY                                 AlternativePaymentType = "safetypay"
	SANTANDER_BRAZIL                          AlternativePaymentType = "santanderbr"
	SANTANDER_MEXICO                          AlternativePaymentType = "santandermx"
	SANTANDER_RIO                             AlternativePaymentType = "santanderrio"
	SCOTIABANK                                AlternativePaymentType = "scotiabank"
	SEPA_DIRECTDEBIT_MERCHANT_MANDATE_MODEL_C AlternativePaymentType = "sepamm"
	SEPA_DIRECTDEBIT_PPPRO_MANDATE_MODEL_A    AlternativePaymentType = "sepapm"
	SEPA_PAYOUT                               AlternativePaymentType = "sepapayout"
	SERVIPAG                                  AlternativePaymentType = "servipag"
	SINGPOST                                  AlternativePaymentType = "singpost"
	SKRILL                                    AlternativePaymentType = "skrill"
	SOFORTUBERWEISUNG                         AlternativePaymentType = "sofort"
	S_PANKKI                                  AlternativePaymentType = "spankki"
	SURTIMAX                                  AlternativePaymentType = "surtimax"
	TARJETA_SHOPPING                          AlternativePaymentType = "tarjeta"
	TELEINGRESO                               AlternativePaymentType = "teleingreso"
	TESTPAY                                   AlternativePaymentType = "testpay"
	TRUSTLY                                   AlternativePaymentType = "trustly"
	TRUSTPAY                                  AlternativePaymentType = "trustpay"
	WEBMONEY                                  AlternativePaymentType = "webmoney"
	WEBPAY                                    AlternativePaymentType = "webpay"
	WECHAT_PAY                                AlternativePaymentType = "wechatpay"
	ZIMPLER                                   AlternativePaymentType = "zimpler"
	UK_DIRECT_DEBIT                           AlternativePaymentType = "ukdirectdebit"
	PAYBYBANKAPP                              AlternativePaymentType = "paybybankapp"
)

func (apt AlternativePaymentType) GetBytes() []byte {
	return []byte(apt)
}

func (apt AlternativePaymentType) GetValue() string {
	return string(apt)
}

func (apt AlternativePaymentType) StringConstants() []istringconstant.IStringConstant {
	output := make([]istringconstant.IStringConstant, 0)
	for _, val := range allAlternativePaymentTypes() {
		output = append(output, val)
	}
	return output
}

func FromValue(value string) *AlternativePaymentType {
	for _, apt := range allAlternativePaymentTypes() {
		if apt.GetValue() == value {
			return &apt
		}
	}
	return nil
}

func allAlternativePaymentTypes() []AlternativePaymentType {
	return []AlternativePaymentType{
		ASTROPAY_DIRECT,
		AURA,
		BALOTO_CASH,
		BANAMEX,
		BANCA_AV_VILLAS,
		BANCA_CAJA_SOCIAL,
		BANCO_GNB_SUDAMERIS,
		BANCO_CONSORCIO,
		BANCO_COOPERATIVO_COOPCENTRAL,
		BANCO_CORPBANCA,
		BANCO_DE_BOGOTA,
		BANCO_DE_CHILE_EDWARDS_CITI,
		BANCO_DE_CHILE_CASH,
		BANCO_DE_OCCIDENTE,
		BANCO_DE_OCCIDENTE_CASH,
		BANCO_DO_BRASIL,
		BANCO_FALABELLA_Chile,
		BANCO_FALABELLA_Columbia,
		BANCO_INTERNATIONAL,
		BANCO_PICHINCHA,
		BANCO_POPULAR,
		BANCO_PROCREDIT,
		BANCO_RIPLEY,
		BANCO_SANTANDER,
		BANCO_SANTANDER_BANEFE,
		BANCO_SECURITY,
		BANCOBICE,
		BANCOESTADO,
		BANCOLOMBIA,
		BANCOMER,
		BANCONTACT_MR_CASH,
		BANCOOMEVA,
		BANK_ISLAM,
		BANK_TRANSFER,
		BBVA_Chile,
		BBVA_Columbia,
		BCI_TBANC,
		BITPAY,
		BOLETO_BANCARIO,
		BRADESCO,
		CABAL,
		CARTAO_MERCADOLIVRE,
		CARULLA,
		CENCOSUD,
		CHINA_UNION_PAY,
		CIMB_CLICKS,
		CITIBANK,
		CMR,
		COLPATRIA,
		COOPEUCH,
		CORPBANCA,
		DANSKE_BANK,
		DAVIVIENDA,
		DRAGONPAY,
		EASYPAY,
		EFECTY,
		ELO,
		EMPRESA_DE_ENERGIA_DEL_QUINDIO,
		ENETS,
		ENTERCASH,
		E_PAY_PETRONAS,
		EPS,
		ESTONIAN_ONLINE_BANK_TRANSFER,
		FINLAND_ONLINE_BANK_TRANSFER,
		GIROPAY,
		HANDELSBANKEN,
		HELM_BANK,
		HIPERCARD,
		HONG_LEONG_BANK,
		IDEAL,
		INDONESIA_ATM,
		INSTANT_TRANSFER,
		INTERNATIONAL_PAY_OUT,
		ITAU_BRAZIL,
		ITAU_CHILE,
		LATVIANBT,
		LINK,
		LITHUANIAN_ONLINE_BANK_TRANSFER,
		MAGNA,
		MAXIMA,
		MAYBANK2U,
		MULTIBANCO,
		MYBANK,
		MYCLEAR_FPX,
		NARANJA,
		NARVESEN_LIETUVOS_SPAUDA,
		NATIVA,
		NORDEA,
		OSUUSPANKKI,
		OXXO,
		PAGO_FACIL,
		PAYPAL,
		PAYPOST_LIETUVOS_PASTAS,
		PAYSAFECARD,
		PAYSBUY_CASH,
		PAYSERA,
		PAYU,
		PERLAS,
		POLI,
		POLISH_PAYOUT,
		POP_PANKKI,
		POSTFINANCE,
		PRESTO,
		PROVINCIA_NET,
		PRZELEWY24,
		PSE,
		QIWI,
		QIWI_PAYOUT,
		RAPI_PAGO,
		REDPAGOS,
		RHB_BANK,
		SAASTOPANKKI,
		SAFETYPAY,
		SANTANDER_BRAZIL,
		SANTANDER_MEXICO,
		SANTANDER_RIO,
		SCOTIABANK,
		SEPA_DIRECTDEBIT_MERCHANT_MANDATE_MODEL_C,
		SEPA_DIRECTDEBIT_PPPRO_MANDATE_MODEL_A,
		SEPA_PAYOUT,
		SERVIPAG,
		SINGPOST,
		SKRILL,
		SOFORTUBERWEISUNG,
		S_PANKKI,
		SURTIMAX,
		TARJETA_SHOPPING,
		TELEINGRESO,
		TESTPAY,
		TRUSTLY,
		TRUSTPAY,
		WEBMONEY,
		WEBPAY,
		WECHAT_PAY,
		ZIMPLER,
		UK_DIRECT_DEBIT,
		PAYBYBANKAPP,
	}
}
