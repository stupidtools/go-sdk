package gateways

import (
	"context"
	"fmt"
)

type XmlGateway struct {
	*Gateway
}

func NewXmlGateway() *XmlGateway {
	return &XmlGateway{
		Gateway: NewGateway("text/xml"),
	}
}

func (x *XmlGateway) DoTransaction(ctx context.Context, request string) (string, error) {
	response, err := x.SendRequest(ctx, "POST", "", request, nil)
	if err != nil {
		return "", err
	}
	if response.GetStatusCode() != 200 {
		return "", fmt.Errorf("Unexpected http status code [%d]:%s", response.GetStatusCode(), response.GetRawResponse())
	}
	return response.GetRawResponse(), nil
}
