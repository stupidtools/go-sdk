package terminalutilities

import (
	"strings"

	"github.com/globalpayments/go-sdk/api/entities/enums/controlcodes"
	"github.com/globalpayments/go-sdk/api/utils/enumutils"
)

type DeviceMessage struct {
	buffer        []byte
	keepAlive     bool
	awaitResponse bool
}

func NewDeviceMessage(buffer []byte) *DeviceMessage {
	return &DeviceMessage{
		buffer: buffer,
	}
}

func (dm *DeviceMessage) GetSendBuffer() []byte {
	return dm.buffer
}

func (dm *DeviceMessage) IsKeepAlive() bool {
	return dm.keepAlive
}

func (dm *DeviceMessage) SetKeepAlive(keepAlive bool) {
	dm.keepAlive = keepAlive
}

func (dm *DeviceMessage) IsAwaitResponse() bool {
	return dm.awaitResponse
}

func (dm *DeviceMessage) SetAwaitResponse(awaitResponse bool) {
	dm.awaitResponse = awaitResponse
}

func (dm *DeviceMessage) ToString() string {
	var sb strings.Builder
	for _, b := range dm.buffer {
		if enumutils.IsDefined(controlcodes.ControlCodes(b), b) {
			code := controlcodes.ControlCodes(b).String()
			sb.WriteString(code)
		} else {
			sb.WriteByte(b)
		}
	}

	return sb.String()
}
