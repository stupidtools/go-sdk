package mobilepaymentmethodtype

type MobilePaymentMethodType string

const (
	NIL          MobilePaymentMethodType = ""
	APPLEPAY     MobilePaymentMethodType = "apple-pay"
	GOOGLEPAY    MobilePaymentMethodType = "pay-with-google"
	CLICK_TO_PAY MobilePaymentMethodType = "click-to-pay"
)

func (m MobilePaymentMethodType) String() string {
	return string(m)
}
