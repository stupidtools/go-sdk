package gateways

type GatewayResponseError struct {
	msg       string
	errorCode string
}

// Error implements the error interface.
func (e *GatewayResponseError) Error() string {
	return e.msg
}

func (e *GatewayResponseError) GetErrorCode() string {
	return e.errorCode
}

func NewGatewayResponseError(msg string, errorCode string) *GatewayResponseError {
	return &GatewayResponseError{msg: msg, errorCode: errorCode}
}

type GatewayResponse struct {
	StatusCode  int
	RawResponse string
}

func (g *GatewayResponse) GetStatusCode() int {
	return g.StatusCode
}

func (g *GatewayResponse) SetStatusCode(statusCode int) {
	g.StatusCode = statusCode
}

func (g *GatewayResponse) GetRawResponse() string {
	return g.RawResponse
}

func (g *GatewayResponse) SetRawResponse(rawResponse string) {
	g.RawResponse = rawResponse
}
