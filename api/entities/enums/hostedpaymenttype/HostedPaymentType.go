package hostedpaymenttype

type HostedPaymentType string

const (
	NONE                      HostedPaymentType = "NONE"
	MAKE_PAYMENT              HostedPaymentType = "MAKE_PAYMENT"
	MAKE_PAYMENT_RETURN_TOKEN HostedPaymentType = "MAKE_PAYMENT_RETURN_TOKEN"
	GET_TOKEN                 HostedPaymentType = "GET_TOKEN"
	MY_ACCOUNT                HostedPaymentType = "MY_ACCOUNT"
)
