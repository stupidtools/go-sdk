package gateways

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/globalpayments/go-sdk/api/entities/abstractions"
	"github.com/globalpayments/go-sdk/api/logging"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Gateway struct {
	ContentType            string
	EnableLogging          bool
	RequestLogger          logging.IRequestLogger
	LogEntry               strings.Builder
	LSChar                 string
	Headers                map[string]string
	DynamicHeaders         map[string]string
	Timeout                int
	ServiceUrl             string
	AuthorizationHeaderKey string
	WebProxy               abstractions.IWebProxy
}

func NewGateway(contentType string) *Gateway {
	return &Gateway{
		ContentType:            contentType,
		Headers:                make(map[string]string),
		DynamicHeaders:         make(map[string]string),
		LSChar:                 "\n",
		AuthorizationHeaderKey: "Authorization",
	}
}

func (g *Gateway) SendRequest(ctx context.Context, verb string, endpoint string, data string, queryStringParams map[string]string) (*GatewayResponse, error) {
	queryString := g.BuildQueryString(queryStringParams)
	client := &http.Client{
		Timeout: time.Duration(g.Timeout) * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, verb, g.ServiceUrl+endpoint+queryString, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", fmt.Sprintf("%s; charset=UTF-8", g.ContentType))
	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}
	for key, value := range g.DynamicHeaders {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawResponse, err := g.GetRawResponse(resp)

	if err != nil {
		return nil, err
	}
	return &GatewayResponse{
		StatusCode:  resp.StatusCode,
		RawResponse: rawResponse,
	}, nil
}

func (g *Gateway) GetRawResponse(responseStream *http.Response) (string, error) {
	var reader *gzip.Reader
	var err error

	if g.Headers["Accept-Encoding"] == "gzip" {
		reader, err = gzip.NewReader(responseStream.Body)
		if err != nil {
			return "", err
		}
		defer reader.Close()
		rawResponse, err := ioutil.ReadAll(reader)
		if err != nil {
			return "", err
		}
		return string(rawResponse), nil

	} else {
		rawResponse, err := ioutil.ReadAll(responseStream.Body)
		if err != nil {
			return "", err
		}
		return string(rawResponse), nil
	}

}

func (g *Gateway) BuildQueryString(queryStringParams map[string]string) string {
	if queryStringParams == nil {
		return ""
	}

	params := url.Values{}
	for key, value := range queryStringParams {
		params.Add(key, value)
	}

	return "?" + params.Encode()
}

func (g *Gateway) SendRequestWithMultipart(endpoint string, content *multipart.Writer, body bytes.Buffer) (*GatewayResponse, error) {
	client := &http.Client{
		Timeout: time.Duration(g.Timeout) * time.Second,
	}

	req, err := http.NewRequest("POST", g.ServiceUrl+endpoint, &body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", content.FormDataContentType())
	req.Header.Set("Content-Length", strconv.Itoa(body.Len()))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawResponse, err := g.GetRawResponse(resp)
	if err != nil {
		return nil, err
	}

	return &GatewayResponse{
		StatusCode:  resp.StatusCode,
		RawResponse: rawResponse,
	}, nil
}

func (g *Gateway) GenerateRequestLog() {
	if g.EnableLogging {
		if g.RequestLogger == nil {
			fmt.Println(g.LogEntry.String())
		} else {
			err := g.RequestLogger.RequestSent(g.LogEntry.String())
			if err != nil {
				panic(err)
			}
		}
	} else {
		if g.RequestLogger != nil {
			err := g.RequestLogger.RequestSent(g.LogEntry.String())
			if err != nil {
				panic(err)
			}
		}
	}
	g.LogEntry.Reset()
}

func (g *Gateway) GenerateResponseLog() {
	if g.EnableLogging {
		if g.RequestLogger == nil {
			fmt.Println(g.LogEntry.String())
		} else {
			err := g.RequestLogger.ResponseReceived(g.LogEntry.String())
			if err != nil {
				panic(err)
			}
		}
	} else {
		if g.RequestLogger != nil {
			err := g.RequestLogger.ResponseReceived(g.LogEntry.String())
			if err != nil {
				panic(err)
			}
		}
	}
	g.LogEntry.Reset()
}
