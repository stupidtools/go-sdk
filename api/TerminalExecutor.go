package api

import (
	"context"
	"github.com/globalpayments/go-sdk/api/terminals/builders"
	"github.com/globalpayments/go-sdk/api/terminals/terminalresponse"
)

type ExecutableTerminal interface {
	ExecuteWithName(ctx context.Context, name string, device builders.ITerminalBuilderDevice) (terminalresponse.ITerminalResponse, error)
}

func ExecuteTerminal(terminal ExecutableTerminal) (terminalresponse.ITerminalResponse, error) {
	return ExecuteTerminalWithName("default", terminal)
}

func ExecuteTerminalWithName(name string, terminal ExecutableTerminal) (terminalresponse.ITerminalResponse, error) {
	ctx := context.Background()
	device, err := GetServiceContainerInstance().GetDeviceController(name)
	if err != nil {
		return nil, err
	}
	return terminal.ExecuteWithName(ctx, name, device)
}

// Context Execute will cancel terminal operations if any cancellation signal is detected from the context

func ExecuteTerminalWithNameAndContext(ctx context.Context, name string, terminal ExecutableTerminal) (terminalresponse.ITerminalResponse, error) {
	device, err := GetServiceContainerInstance().GetDeviceController(name)
	if err != nil {
		return nil, err
	}
	return terminal.ExecuteWithName(ctx, name, device)
}

func ExecuteTerminalWithContext(ctx context.Context, terminal ExecutableTerminal) (terminalresponse.ITerminalResponse, error) {
	return ExecuteTerminalWithNameAndContext(ctx, "default", terminal)
}
