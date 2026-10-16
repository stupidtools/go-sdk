package api

import (
	"context"
	"errors"
	"github.com/globalpayments/go-sdk/api/abstractions"
	"github.com/globalpayments/go-sdk/api/builders/rebuilders"
	"github.com/globalpayments/go-sdk/api/entities/transactions"
	"github.com/globalpayments/go-sdk/api/gateways"
	paymentmethods "github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"github.com/shopspring/decimal"
)

type AutoReversalError struct {
	msg           string
	originalError error
}

// Error implements the error interface.
func (e *AutoReversalError) Error() string {
	return e.msg
}

// Unwrap returns the original error.
func (e *AutoReversalError) Unwrap() error {
	return e.originalError
}

func NewAutoReversalError(msg string, err error) error {
	return &AutoReversalError{
		msg:           msg,
		originalError: err,
	}
}

type ReversibleGateway interface {
	abstractions.ExecutableGateway
	GetPaymentMethod() paymentmethods.IPaymentMethod
	GetAmount() *decimal.Decimal
	GetClientTransactionId() string
}

func ExecuteGateway[T abstractions.ITransaction](ctx context.Context, gw abstractions.ExecutableGateway) (*T, error) {
	return ExecuteGatewayWithName[T](ctx, "default", gw)
}

func ExecuteGatewayWithName[T abstractions.ITransaction](ctx context.Context, name string, gw abstractions.ExecutableGateway) (*T, error) {
	client, err := LoadGatewayByName(name)
	if err != nil {
		return nil, err
	}
	res, err := gw.Execute(ctx, client)
	if err != nil {
		return nil, err
	}
	trans, ok := res.(*T)
	if !ok {
		return nil, errors.New("Transaction type not correct")
	}
	return trans, nil

}

func LoadGateway() (abstractions.IPaymentGateway, error) {
	return LoadGatewayByName("default")
}

func LoadGatewayByName(name string) (abstractions.IPaymentGateway, error) {
	client, err := GetServiceContainerInstance().GetGateway(name)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ExecuteTimeoutReversibleGateway(ctx context.Context, gw ReversibleGateway) (*transactions.Transaction, error) {
	return ExecuteTimeoutReversibleGatewayWithName(ctx, "default", gw)
}

func ExecuteTimeoutReversibleGatewayWithName(ctx context.Context, name string, gw ReversibleGateway) (*transactions.Transaction, error) {
	if gw.GetClientTransactionId() == "" {
		return nil, errors.New("You must supply a client transaction id in order to make this reversable")
	}
	result, err := ExecuteGatewayWithName[transactions.Transaction](ctx, name, gw)
	if (result != nil) && (result.GetResponseCode() == "91") {
		trans := rebuilders.FromId("", gw.GetPaymentMethod().GetPaymentMethodType())
		trans.GetTransactionReference().SetClientTransactionId(gw.GetClientTransactionId())
		revTrans := trans.ReverseWithAmount(gw.GetAmount())
		_, err := ExecuteGatewayWithName[transactions.Transaction](ctx, name, revTrans)
		failed := true
		if err == nil {
			failed = false
		} else if gwErr, ok := err.(*gateways.GatewayResponseError); ok && (gwErr.GetErrorCode() == "3") {
			failed = false
		}
		if failed {
			return nil, NewAutoReversalError("The transaction failed with a gateway timeout, and was unable to be reversed.", err)
		}
		result.SetAutoReversed(true)
		return result, nil
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
