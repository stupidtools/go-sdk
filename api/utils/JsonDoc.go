package utils

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"math"
	"strconv"

	"github.com/globalpayments/go-sdk/api/entities/enums/devicetype"
)

type IRequestEncoder interface {
	Encode(interface{}) interface{}
	Decode(interface{}) interface{}
}

type JsonDoc struct {
	dict    map[string]interface{}
	encoder IRequestEncoder
}

func NewJsonDoc() *JsonDoc {
	return &JsonDoc{
		dict:    make(map[string]interface{}),
		encoder: nil,
	}
}

func NewJsonDocWithEncoder(encoder IRequestEncoder) *JsonDoc {
	return &JsonDoc{
		dict:    make(map[string]interface{}),
		encoder: encoder,
	}
}

func NewJsonDocWithValues(values map[string]interface{}) *JsonDoc {
	return &JsonDoc{
		dict:    values,
		encoder: nil,
	}
}

func NewJsonDocWithValuesAndEncoder(values map[string]interface{}, encoder IRequestEncoder) *JsonDoc {
	return &JsonDoc{
		dict:    values,
		encoder: encoder,
	}
}

func (jd *JsonDoc) Remove(key string) *JsonDoc {
	delete(jd.dict, key)
	return jd
}

func (jd *JsonDoc) Set(key string, value string, force bool) *JsonDoc {
	if value != "" || force {
		if jd.encoder != nil {
			jd.dict[key] = jd.encoder.Encode(value)
		} else {
			jd.dict[key] = value
		}
	}
	return jd
}

func (jd *JsonDoc) SetEnum(key string, value devicetype.DeviceType, force bool) *JsonDoc {
	if value != "" || force {
		if jd.encoder != nil {
			jd.dict[key] = jd.encoder.Encode(value)
		} else {
			jd.dict[key] = value
		}
	}
	return jd
}

func (jd *JsonDoc) SetInt(key string, value *int) *JsonDoc {
	if jd.encoder != nil {
		jd.dict[key] = jd.encoder.Encode(&value)
	} else {
		jd.dict[key] = &value
	}
	return jd
}

func (jd *JsonDoc) SetStringArray(key string, values []string) *JsonDoc {
	if values != nil {
		jd.dict[key] = values
	}
	return jd
}

func (jd *JsonDoc) SetIntArray(key string, values []int) *JsonDoc {
	if values != nil {
		jd.dict[key] = values
	}
	return jd
}

func (jd *JsonDoc) SetBool(key string, value bool) *JsonDoc {
	if jd.encoder != nil {
		jd.dict[key] = jd.encoder.Encode(value)
	} else {
		jd.dict[key] = value
	}
	return jd
}

func (jd *JsonDoc) SetJsonDoc(key string, value *JsonDoc) *JsonDoc {
	if value != nil {
		jd.dict[key] = value.dict
	}
	return jd
}

func (jd *JsonDoc) ToString() (string, error) {
	jsonBytes, err := json.Marshal(jd.dict)
	return string(jsonBytes), err
}

func (jd *JsonDoc) Get(name string) *JsonDoc {
	if value, ok := jd.dict[name]; ok {
		if data, ok := value.(map[string]interface{}); ok {
			return NewJsonDocWithValues(data)
		}
	}
	return nil
}

func (jd *JsonDoc) GetArray(name string) []*JsonDoc {
	output := make([]*JsonDoc, 0)
	if value, ok := jd.dict[name]; ok {
		if data, ok := value.([]interface{}); ok {
			for _, d := range data {
				if inner, ok := d.(map[string]interface{}); ok {
					output = append(output, NewJsonDocWithValues(inner))
				}
			}
		}
	}
	return output
}

func (jd *JsonDoc) GetString(name string) string {
	if value, ok := jd.dict[name]; ok {
		if jd.encoder != nil {
			return jd.encoder.Decode(value).(string)
		} else {
			return value.(string)
		}
	}
	return ""
}

func (jd *JsonDoc) GetDecimal(name string) *decimal.Decimal {
	value := jd.GetString(name)
	if value != "" {
		decimalValue, err := decimal.NewFromString(value)
		if err != nil {
			return nil
		}
		return &decimalValue
	}
	return nil
}

func (jd *JsonDoc) GetBool(name string) bool {
	if value, ok := jd.dict[name]; ok {
		return value.(bool)
	}
	return false
}

func (jd *JsonDoc) GetInt(name string) *int {
	if value, ok := jd.dict[name]; ok {
		if i, ok := value.(float64); ok {
			intValue := int(math.Round(i))
			return &intValue
		} else if strval, ok := value.(string); ok {
			if intValue, err := strconv.Atoi(strval); err == nil {
				return &intValue
			}
		}
	}
	return nil
}

func (jd *JsonDoc) GetIntOrStringAsString(name string) string {
	if value, ok := jd.dict[name]; ok {
		if s, ok := value.(string); ok {
			return s
		} else {
			i := jd.GetInt(name)
			if i != nil {
				return strconv.Itoa(*i)
			}
		}
	}
	return ""
}

func (jd *JsonDoc) GetStringArray(name string) []string {
	if value, ok := jd.dict[name]; ok {
		return value.([]string)
	}
	return nil
}

func (jd *JsonDoc) Has(name string) bool {
	_, ok := jd.dict[name]
	return ok
}

func Parse(jsonStr string) (*JsonDoc, error) {
	values := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &values)
	if err != nil {
		return nil, err
	}
	return NewJsonDocWithValues(values), nil
}

func ParseBytes(jsonBytes []byte) (*JsonDoc, error) {
	values := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &values)
	if err != nil {
		return nil, err
	}
	return NewJsonDocWithValues(values), nil
}

func ParseWithEncoder(jsonStr string, encoder IRequestEncoder) (*JsonDoc, error) {
	values := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &values)
	if err != nil {
		return nil, err
	}
	return NewJsonDocWithValuesAndEncoder(values, encoder), nil
}

func ParseSingleValue(jsonStr string, name string) (string, error) {
	doc, err := Parse(jsonStr)
	if err != nil {
		return "", err
	}
	return doc.GetString(name), nil
}

func ParseSingleValueWithEncoder(jsonStr string, name string, encoder IRequestEncoder) (string, error) {
	doc, err := ParseWithEncoder(jsonStr, encoder)
	if err != nil {
		return "", err
	}
	return doc.GetString(name), nil
}
