package enumutils

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/inumericconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
)

type ReverseByteEnumMap struct {
	Map map[byte]ibyteconstant.IByteConstant
}

func NewReverseByteEnumMap(valueType ibyteconstant.IByteConstant) *ReverseByteEnumMap {
	mapper := &ReverseByteEnumMap{
		Map: make(map[byte]ibyteconstant.IByteConstant),
	}
	for _, v := range valueType.ByteConstants() {
		mapper.Map[v.GetByte()] = v
	}
	return mapper
}

type ReverseIntEnumMap struct {
	Map map[int]inumericconstant.INumericConstant
}

func NewReverseIntEnumMap(valueType inumericconstant.INumericConstant) *ReverseIntEnumMap {
	mapper := &ReverseIntEnumMap{
		Map: make(map[int]inumericconstant.INumericConstant),
	}
	for _, v := range valueType.NumericConstants() {
		mapper.Map[v.GetValue()] = v
	}
	return mapper
}

func (mapper *ReverseIntEnumMap) Get(value int) inumericconstant.INumericConstant {
	return mapper.Map[value]
}

type ReverseStringEnumMap struct {
	Map map[string]istringconstant.IStringConstant
}

func NewReverseStringEnumMap(valueType istringconstant.IStringConstant) *ReverseStringEnumMap {
	mapper := &ReverseStringEnumMap{
		Map: make(map[string]istringconstant.IStringConstant),
	}
	for _, v := range valueType.StringConstants() {
		mapper.Map[v.GetValue()] = v
	}
	return mapper
}

func (mapper *ReverseStringEnumMap) Get(value string) istringconstant.IStringConstant {
	return mapper.Map[value]
}
