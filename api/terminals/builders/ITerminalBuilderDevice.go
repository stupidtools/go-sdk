package builders

import (
	"context"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
)

type ITerminalBuilderDevice interface {
	ManageTransactionWithContext(ctx context.Context, tmb *TerminalManageBuilder) (terminalresponse.ITerminalResponse, error)
	ProcessTransactionWithContext(ctx context.Context, tab *TerminalAuthBuilder) (terminalresponse.ITerminalResponse, error)
}
