package serviceendpoints

type ServiceEndpoints string

const (
	GLOBAL_ECOM_PRODUCTION     ServiceEndpoints = "https://api.realexpayments.com/epage-remote.cgi"
	GLOBAL_ECOM_TEST           ServiceEndpoints = "https://api.sandbox.realexpayments.com/epage-remote.cgi"
	PORTICO_PRODUCTION         ServiceEndpoints = "https://api2.heartlandportico.com"
	PORTICO_TEST               ServiceEndpoints = "https://cert.api2.heartlandportico.com"
	THREE_DS_AUTH_PRODUCTION   ServiceEndpoints = "https://api.globalpay-ecommerce.com/3ds2/"
	THREE_DS_AUTH_TEST         ServiceEndpoints = "https://api.sandbox.globalpay-ecommerce.com/3ds2/"
	PAYROLL_PRODUCTION         ServiceEndpoints = "https://taapi.heartlandpayrollonlinetest.com/PosWebUI"
	PAYROLL_TEST               ServiceEndpoints = "https://taapi.heartlandpayrollonlinetest.com/PosWebUI/Test/Test"
	TABLE_SERVICE_PRODUCTION   ServiceEndpoints = "https://www.freshtxt.com/api31/"
	TABLE_SERVICE_TEST         ServiceEndpoints = "https://www.freshtxt.com/api31/"
	GP_API_PRODUCTION          ServiceEndpoints = "https://apis.globalpay.com/ucp"
	GP_API_TEST                ServiceEndpoints = "https://apis.sandbox.globalpay.com/ucp"
	BILLPAY_TEST               ServiceEndpoints = "https://testing.heartlandpaymentservices.net"
	BILLPAY_CERTIFICATION      ServiceEndpoints = "https://staging.heartlandpaymentservices.net"
	BILLPAY_PRODUCTION         ServiceEndpoints = "https://heartlandpaymentservices.net"
	TRANSACTION_API_PRODUCTION ServiceEndpoints = "https://api.paygateway.com/transactions"
	TRANSACTION_API_TEST       ServiceEndpoints = "https://api.pit.paygateway.com/transactions"
	OPEN_BANKING_TEST          ServiceEndpoints = "https://api.sandbox.globalpay-ecommerce.com/openbanking"
	OPEN_BANKING_PRODUCTION    ServiceEndpoints = "https://api.globalpay-ecommerce.com/openbanking"
	GENIUS_API_PRODUCTION      ServiceEndpoints = "https://ps1.merchantware.net/Merchantware/ws/RetailTransaction/v46/Credit.asmx"
	GENIUS_API_TEST            ServiceEndpoints = "https://ps1.merchantware.net/Merchantware/ws/RetailTransaction/v46/Credit.asmx"
	GENIUS_TERMINAL_PRODUCTION ServiceEndpoints = "https://transport.merchantware.net/v4/transportService.asmx"
	GENIUS_TERMINAL_TEST       ServiceEndpoints = "https://transport.merchantware.net/v4/transportService.asmx"
	PROPAY_TEST                ServiceEndpoints = "https://xmltest.propay.com/API/PropayAPI.aspx"
	PROPAY_TEST_CANADIAN       ServiceEndpoints = "https://xmltestcanada.propay.com/API/PropayAPI.aspx"
	PROPAY_PRODUCTION          ServiceEndpoints = "https://epay.propay.com/API/PropayAPI.aspx"
	PROPAY_PRODUCTION_CANADIAN ServiceEndpoints = "https://www.propaycanada.ca/API/PropayAPI.aspx"
)

func (se ServiceEndpoints) GetValue() string {
	return string(se)
}

func (se ServiceEndpoints) GetBytes() []byte {
	return []byte(se)
}
