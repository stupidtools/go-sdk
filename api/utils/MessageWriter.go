package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
)

type MessageWriter struct {
	buffer         bytes.Buffer
	messageRequest strings.Builder
}

func NewMessageWriter() *MessageWriter {
	return &MessageWriter{}
}

func NewMessageWriterWithBytes(bytes []byte) *MessageWriter {
	mw := &MessageWriter{}
	mw.buffer.Write(bytes)
	return mw
}

func (mw *MessageWriter) GetMessageRequest() *strings.Builder {
	return &mw.messageRequest
}

func (mw *MessageWriter) SetMessageRequest(messageRequest strings.Builder) {
	mw.messageRequest = messageRequest
}

func (mw *MessageWriter) Add(b byte) {
	mw.buffer.WriteByte(b)
}

func (mw *MessageWriter) AddByteConstant(constant ibyteconstant.IByteConstant) {
	mw.buffer.WriteByte(constant.GetByte())
}

func (mw *MessageWriter) AddStringConstant(constant istringconstant.IStringConstant) {
	mw.buffer.Write(constant.GetBytes())
}

func (mw *MessageWriter) AddRange(bytes []byte) {
	mw.buffer.Write(bytes)
}

func (mw *MessageWriter) Pop() {
	buf := mw.buffer.Bytes()
	mw.buffer.Reset()
	mw.buffer.Write(buf[:len(buf)-1])
}

func (mw *MessageWriter) ToArray() []byte {
	return mw.buffer.Bytes()
}

func (mw *MessageWriter) AddInt(value int) {
	mw.AddRange(intToBytes(value))
}

func (mw *MessageWriter) AddIntWithLength(value int, length int) {
	mw.AddRange(formatInteger(int64(value), length))
}

func (mw *MessageWriter) AddString(value string) {
	mw.AddRange([]byte(value))
}

func intToBytes(data int) []byte {
	return []byte{
		byte((data >> 24) & 0xff),
		byte((data >> 16) & 0xff),
		byte((data >> 8) & 0xff),
		byte((data >> 0) & 0xff),
	}
}

func formatInteger(value int64, length int) []byte {
	offsets := []int{0, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	if length == 1 {
		return []byte{byte(value & 0xFF)}
	}

	byteCount := abs(int(value)) / 8
	baseLength := byteCount * 2
	if baseLength > length {
		baseLength = length
	}

	inputBuffer := NewMessageWriter()
	for i := 0; i < baseLength; i++ {
		offset := offsets[baseLength-1-i]
		inputBuffer.Add(byte(value >> offset))
	}
	input := inputBuffer.ToArray()

	output := make([]byte, length)
	copy(output[length-baseLength:], input)

	return output
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (mw *MessageWriter) AddRangePadded(fieldValue string, digitCount int) {
	if fieldValue != "" && digitCount > 0 {
		if len(fieldValue) == digitCount {
			mw.messageRequest.WriteString(fieldValue)
		} else if len(fieldValue) > digitCount {
			mw.messageRequest.WriteString(fieldValue[:digitCount])
		} else {
			paddedValue := fmt.Sprintf("%0*s", digitCount, fieldValue)
			mw.messageRequest.WriteString(paddedValue)
		}
	}
}

func (mw *MessageWriter) AddRangePaddedInt(fieldValue int, digitCount int) {
	mw.AddRangePadded(fmt.Sprint(fieldValue), digitCount)
}

func (mw *MessageWriter) String() string {
	buf := mw.buffer.Bytes()
	chars := make([]byte, 2*len(buf))
	hexChars := "0123456789abcdef"

	for i := 0; i < len(buf); i++ {
		chars[2*i] = hexChars[(buf[i]&0xF0)>>4]
		chars[2*i+1] = hexChars[buf[i]&0x0F]
	}

	return string(chars)
}

func (mw *MessageWriter) Length() int {
	return mw.buffer.Len()
}
