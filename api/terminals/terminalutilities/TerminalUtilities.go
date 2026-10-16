package terminalutilities

import (
	"bytes"

	"github.com/globalpayments/go-sdk/api/entities/enums/controlcodes"
	"github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/paxmsgid"
	"github.com/globalpayments/go-sdk/api/terminals/abstractions"
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upamessageid"
	"github.com/globalpayments/go-sdk/api/utils"
)

const version string = "1.35"

func GetElementString(elements ...interface{}) string {
	var sb bytes.Buffer
	for _, element := range elements {
		switch element.(type) {
		case controlcodes.ControlCodes:
			sb.WriteByte(byte(element.(controlcodes.ControlCodes)))
		case abstractions.IRequestSubGroup:
			sb.WriteString(element.(abstractions.IRequestSubGroup).GetElementString())
		case []string:
			for _, sub_element := range element.([]string) {
				sb.WriteByte(byte(controlcodes.FS))
				sb.WriteString(sub_element)
			}
		case istringconstant.IStringConstant:
			sb.WriteString(element.(istringconstant.IStringConstant).GetValue())
		case ibyteconstant.IByteConstant:
			sb.WriteByte(element.(ibyteconstant.IByteConstant).GetByte())
		default:
			sb.WriteString(element.(string))
		}
	}

	return sb.String()
}

func buildMessage(messageId paxmsgid.PaxMsgId, message string) DeviceMessage {
	buffer := utils.NewMessageWriter()

	// Begin Message
	buffer.Add(byte(controlcodes.STX))

	// Add Message Id
	buffer.AddRange(messageId.GetBytes())
	buffer.Add(byte(controlcodes.FS))

	// Add Version
	buffer.AddRange([]byte(version))
	buffer.Add(byte(controlcodes.FS))

	// Add Message
	buffer.AddRange([]byte(message))

	// End the message
	buffer.Add(byte(controlcodes.ETX))

	lrc := CalculateLRC(buffer.ToArray())
	buffer.Add(lrc)

	return *NewDeviceMessage(buffer.ToArray())
}

func BuildMessage(messageType upamessageid.UpaMessageId, requestId string, ecrId string, body *utils.JsonDoc) (*DeviceMessage, error) {
	data := utils.NewJsonDoc()
	json := utils.NewJsonDoc()
	if ecrId != "" {
		data.Set("EcrId", ecrId, true)
	} else {
		data.Set("EcrId", "13", false)
	}
	data.Set("requestId", requestId, true)
	data.Set("command", string(messageType), true)

	if body != nil {
		data.SetJsonDoc("data", body)
	}

	json.SetJsonDoc("data", data)
	json.Set("message", "MSG", false)
	msg, err := json.ToString()
	if err != nil {
		return nil, err
	}
	output := CompileMessage(msg)
	return &output, nil
}

func BuildRequestByBytes(message []byte) DeviceMessage {
	buffer := utils.NewMessageWriter()

	// beginning sentinel
	buffer.Add(byte(controlcodes.STX))

	// put message
	buffer.AddRange(message)

	// ending sentinel
	buffer.Add(byte(controlcodes.ETX))

	lrc := CalculateLRC(buffer.ToArray())
	buffer.Add(lrc)

	return *NewDeviceMessage(buffer.ToArray())
}

func CalculateLRC(buffer []byte) byte {
	length := len(buffer)
	if buffer[length-1] != byte(controlcodes.ETX) {
		length--
	}

	lrc := byte(0x00)
	for i := 1; i < length; i++ {
		lrc ^= buffer[i]
	}
	return lrc
}

func CheckLRC(message string) bool {
	messageBuffer := []byte(message)

	expected := messageBuffer[len(messageBuffer)-1]
	actual := CalculateLRC([]byte(message[:len(message)-1]))

	return expected == actual
}

func CompileMessage(body string) DeviceMessage {
	buffer := utils.NewMessageWriter()

	buffer.Add(byte(controlcodes.STX))
	buffer.Add(byte(controlcodes.LF))
	buffer.AddRange([]byte(body))
	buffer.Add(byte(controlcodes.LF))
	buffer.Add(byte(controlcodes.ETX))
	buffer.Add(byte(controlcodes.LF))

	return *NewDeviceMessage(buffer.ToArray())
}
