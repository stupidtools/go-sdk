package abstractions

import (
	"github.com/globalpayments/go-sdk/api/terminals/messaging"
)

type IDeviceCommInterface interface {
	Connect()
	Disconnect()
	Send(message IDeviceMessage) ([]byte, error)
	SendWithouDisconnect(message IDeviceMessage) ([]byte, error)
	SetMessageSentHandler(messageInterface messaging.IMessageSentInterface)
}
