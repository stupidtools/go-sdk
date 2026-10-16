package abstractions

import (
	"context"
	"github.com/globalpayments/go-sdk/api/network"
)

type IPaymentGateway interface {
	ProcessAuthorization(ctx context.Context, builder IAuthorizationBuilder) (ITransaction, error)
	ManageTransaction(ctx context.Context, builder IManagementBuilder) (ITransaction, error)
	SerializeRequest(ctx context.Context, builder IAuthorizationBuilder) (string, error)
	SendKeepAlive(ctx context.Context) (*network.NetworkMessageHeader, error)
	SupportsHostedPayments() bool
	SupportsOpenBanking() bool
}

type ExecutableGateway interface {
	Execute(ctx context.Context, gateway IPaymentGateway) (ITransaction, error)
}
