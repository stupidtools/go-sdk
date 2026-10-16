package stringutils

import (
	"bytes"
	"github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
	"reflect"
	"strconv"
)

type StringParser struct {
	position int
	buffer   string
}

func NewStringParserFromBytes(buffer []byte) *StringParser {
	return &StringParser{
		buffer: string(buffer),
	}
}

func NewStringParser(str string) *StringParser {
	return &StringParser{
		buffer: str,
	}
}

func (sp *StringParser) ReadBoolean() *bool {
	return sp.ReadBooleanWithIndicator("1")
}

func (sp *StringParser) ReadBooleanWithIndicator(indicator string) *bool {
	value := sp.ReadString(1)
	if value == "" {
		return nil
	}
	b := value == indicator
	return &b
}

func (sp *StringParser) ReadByteConstant(clazz ibyteconstant.IByteConstant) ibyteconstant.IByteConstant {
	value := sp.ReadString(1)
	if value != "" {
		return ReverseByteEnumMapParse(value[0], clazz)
	}
	return nil
}

func (sp *StringParser) ReadSingleByte() *byte {
	bytes := sp.ReadBytes(1)
	if len(bytes) > 0 {
		return &bytes[0]
	}
	return nil
}

func (sp *StringParser) ReadBytes(length int) []byte {
	value := sp.ReadString(length)
	return []byte(value)
}

func (sp *StringParser) ReadInt(length int) *int {
	value := sp.ReadString(length)
	if value != "" {
		i, err := strconv.Atoi(value)
		if err == nil {
			return &i
		}
	}
	return nil
}

func (sp *StringParser) ReadLVAR() string {
	return sp.readVar(1)
}

func (sp *StringParser) ReadLLVAR() string {
	return sp.readVar(2)
}

func (sp *StringParser) ReadLLLVAR() string {
	return sp.readVar(3)
}

func (sp *StringParser) ReadRemaining() string {
	if sp.position < len(sp.buffer) {
		value := sp.buffer[sp.position:]
		sp.position = len(sp.buffer)
		return value
	}
	return ""
}

func (sp *StringParser) ReadRemainingBytes() []byte {
	return []byte(sp.ReadRemaining())
}

func (sp *StringParser) ReadString(length int) string {
	index := sp.position + length
	if index > len(sp.buffer) {
		return ""
	}
	value := sp.buffer[sp.position:index]
	sp.position += length
	return value
}

func (sp *StringParser) ReadStringConstant(length int, clazz istringconstant.IStringConstant) istringconstant.IStringConstant {
	value := sp.ReadString(length)
	return ReverseStringEnumMapParse(value, clazz)
}

func (sp *StringParser) ReadToChar(c rune) string {
	return sp.ReadToCharWithRemove(c, true)
}

func (sp *StringParser) ReadToCharWithRemove(c rune, remove bool) string {
	index := bytes.IndexRune([]byte(sp.buffer[sp.position:]), c)
	if index < 0 {
		return sp.ReadRemaining()
	}
	value := sp.buffer[sp.position : sp.position+index]
	sp.position += index
	if remove {
		sp.position++
	}
	if value == "" {
		return ""
	}
	return value
}

func (sp *StringParser) readVar(length int) string {
	actual := sp.ReadInt(length)
	if actual != nil {
		return sp.ReadString(*actual)
	}
	return ""
}

func (sp *StringParser) Buffer() string {
	return sp.buffer
}

func (sp *StringParser) RemainingLength() int {
	return len(sp.buffer) - sp.position
}

type ReverseByteEnumMap struct {
	mapData map[byte]ibyteconstant.IByteConstant
}

func NewReverseByteEnumMap(valueType reflect.Type) *ReverseByteEnumMap {
	mapper := &ReverseByteEnumMap{mapData: make(map[byte]ibyteconstant.IByteConstant)}

	for i := 0; i < valueType.NumMethod(); i++ {
		method := valueType.Method(i)
		if method.Name == "GetByte" {
			instance := reflect.New(valueType).Interface().(ibyteconstant.IByteConstant)
			mapper.mapData[instance.GetByte()] = instance
		}
	}

	return mapper
}

func (rbem *ReverseByteEnumMap) Get(value byte) ibyteconstant.IByteConstant {
	return rbem.mapData[value]
}

func ReverseByteEnumMapParse(value byte, enum ibyteconstant.IByteConstant) ibyteconstant.IByteConstant {
	for _, option := range enum.ByteConstants() {
		if value == option.GetByte() {
			return option
		}
	}
	return nil
}

func ReverseStringEnumMapParse(s string, enum istringconstant.IStringConstant) istringconstant.IStringConstant {
	for _, option := range enum.StringConstants() {
		if s == option.GetValue() {
			return option
		}
	}
	return nil
}
