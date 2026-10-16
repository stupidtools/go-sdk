package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/globalpayments/go-sdk/api/entities/enums/controlcodes"
	"github.com/globalpayments/go-sdk/api/entities/exceptions"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/messaging"
	"github.com/globalpayments/go-sdk/api/terminals/terminalutilities"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/constants"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaTcpInterface struct {
	client                *net.TCPConn
	settings              IConnectionConfig
	onMessageSent         messaging.IMessageSentInterface
	data                  *utils.MessageWriter
	responseMessageString string
	readyReceived         bool
}

func NewUpaTcpInterface(settings IConnectionConfig) *UpaTcpInterface {
	return &UpaTcpInterface{
		settings: settings,
	}
}

func (upa *UpaTcpInterface) SetMessageSentHandler(onMessageSent messaging.IMessageSentInterface) {
	upa.onMessageSent = onMessageSent
}

func (upa *UpaTcpInterface) Connect() {
	if upa.client == nil {
		address := fmt.Sprintf("%s:%d", upa.settings.GetIpAddress(), upa.settings.GetPort())
		tcpAddress, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			return
		}
		client, err := net.DialTCP("tcp", nil, tcpAddress)
		if err == nil {
			upa.client = client
			upa.client.SetKeepAlive(true)
			upa.client.SetKeepAlivePeriod(time.Duration(upa.settings.GetTimeout()) * time.Millisecond)
			upa.client.SetNoDelay(true)
			upa.client.SetReadBuffer(1024)
			upa.client.SetWriteBuffer(1024)
		}
	}
}

func (upa *UpaTcpInterface) Disconnect() {
	if upa.client != nil {
		_ = upa.client.Close()
		//necessary to make sure connection closes properly
		time.Sleep(250 * time.Millisecond)
		upa.client = nil
	}
}

func (upa *UpaTcpInterface) awaitResponse(ch chan error) {
	timeOfSend := time.Now()
	var errOuter error
	errOuter = nil
	for !upa.readyReceived {
		err := upa.getTerminalResponse()
		if err != nil {
			errOuter = err
			break
		}

		if time.Since(timeOfSend) > time.Duration(upa.settings.GetTimeout())*time.Millisecond {
			errOuter = errors.New("Terminal did not respond in the given timeout.")
			break
		}

		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	ch <- errOuter
}

func (upa *UpaTcpInterface) SendWithouDisconnect(message abstractions.IDeviceMessage) ([]byte, error) {
	upa.Connect()
	if upa.client == nil {
		return nil, exceptions.NewApiException("Api operations failed", exceptions.NewMessageException("Unable to connect with device."))
	}

	upa.readyReceived = false
	sendBuffer := message.GetSendBuffer()

	if upa.onMessageSent != nil {
		currentTime := time.Now()
		upa.onMessageSent.MessageSent(fmt.Sprintf("%s:\n%s", currentTime, string(sendBuffer)))
	}

	if upa.settings.GetRequestLogger() != nil {
		formMsg := string(sendBuffer)
		upa.settings.GetRequestLogger().RequestSent(formMsg)
	}
	resChannel := make(chan error)

	go upa.awaitResponse(resChannel)

	_, err := upa.client.Write(sendBuffer)
	if err != nil {
		return nil, exceptions.NewApiException("Api send operations failed", err)
	}

	err2 := <-resChannel

	if err2 != nil {
		return nil, exceptions.NewApiException("Api read operations failed", err2)
	}

	return []byte(upa.responseMessageString), nil
}

func (upa *UpaTcpInterface) Send(message abstractions.IDeviceMessage) ([]byte, error) {
	defer upa.Disconnect()
	return upa.SendWithouDisconnect(message)
}

func (upa *UpaTcpInterface) getTerminalResponse() error {
	err := upa.validateResponsePacket()
	if err != nil {
		return err
	}

	buffer := upa.data.ToArray()

	if len(buffer) > 0 {
		responseObj := make(map[string]interface{})
		err := json.Unmarshal(buffer, &responseObj)
		if err != nil {
			return err
		}
		message, ok := responseObj["message"].(string)
		if !ok {
			return errors.New("Message field not found in API Response.")
		}

		if upa.settings.GetRequestLogger() != nil {
			formMsg := string(buffer)
			upa.settings.GetRequestLogger().ResponseReceived(formMsg)
		}

		switch message {
		case constants.ACK_MESSAGE, constants.NAK_MESSAGE, constants.TIMEOUT_MESSAGE:
			break
		case constants.BUSY_MESSAGE:
			return errors.New("Device is busy")
		case constants.DATA_MESSAGE:
			upa.responseMessageString = string(buffer)
			eval, ok := responseObj["data"].(map[string]interface{})["response"].(string)
			if ok && eval == "Reboot" {
				upa.readyReceived = true
			}
			return upa.sendAckMessageToDevice()
		case constants.READY_MESSAGE:
			upa.readyReceived = true
			break
		default:
			return errors.New("Message field value is unknown in API Response.")
		}
	}
	return nil
}

func (upa *UpaTcpInterface) validateResponsePacket() error {
	stx := controlcodes.STX.GetByte()
	etx := controlcodes.ETX.GetByte()
	lf := controlcodes.LF.GetByte()

	upa.data = utils.NewMessageWriter()

	buffer := make([]byte, 0, 1024)

	for {
		buf := make([]byte, 1)
		n, err := upa.client.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if n == 0 {
			break
		}
		buffer = append(buffer, buf...)
		buffLen := len(buffer)
		if buffLen > 3 && buffer[buffLen-2] == etx && (buffer[buffLen-3]&buffer[buffLen-1] == lf) {
			break
		}
	}

	for i := 0; i < len(buffer); i++ {
		if i < 2 {
			if buffer[i] != stx && buffer[i+1] != lf {
				return errors.New("The bytes of the start response packet are not the expected bytes.")
			}
			i++
			continue
		}

		if buffer[i] == etx {
			if buffer[i-1]&buffer[i+1] != lf {
				return errors.New("The bytes of the end response packet are not the expected bytes.")
			}
			break
		} else if buffer[i] != lf {
			upa.data.Add(buffer[i])
		}
	}
	return nil
}

func (upa *UpaTcpInterface) sendAckMessageToDevice() error {
	jsonObj := make(map[string]interface{})
	jsonObj["data"] = ""
	jsonObj["message"] = "ACK"
	body, err := json.Marshal(jsonObj)
	if err != nil {
		return err
	}
	message := terminalutilities.CompileMessage(string(body))
	sendBuffer := message.GetSendBuffer()

	if upa.settings.GetRequestLogger() != nil {
		formMsg := string(sendBuffer)
		upa.settings.GetRequestLogger().RequestSent(formMsg)
	}

	if upa.onMessageSent != nil {
		currentTime := time.Now()
		upa.onMessageSent.MessageSent(fmt.Sprintf("%s:\n%s", currentTime, string(sendBuffer)))
	}

	_, err = upa.client.Write(sendBuffer)
	if err != nil {
		return err
	}

	return nil
}
