package myJson

import (
	"fmt"
	"reflect"
)

type JsonData struct {
	data interface{}
}

func NewJsonData(data interface{}) *JsonData {
	return &JsonData{data: data}
}

func (j *JsonData) EncodeValue() string {
	return simpleEncode(reflect.ValueOf(j.data))
}

func simpleEncode(value reflect.Value) string {
	switch value.Kind() {
	case reflect.String:
		return fmt.Sprintf(`"%s"`, value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", value.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", value.Bool())
	case reflect.Slice, reflect.Array:
		str := "["
		for index := 0; index < value.Len(); index++ {
			if index > 0 {
				str += ","
			}
			str += simpleEncode(value.Index(index))
		}
		str += "]"
		return str
	case reflect.Map:
		str := "{"
		for index, key := range value.MapKeys() {
			if index > 0 {
				str += ","
			}

			str += fmt.Sprintf(`"%s":%s`, key.String(), simpleEncode(value.MapIndex(key)))
		}
		str += "}"
		return str
	case reflect.Struct:
		str := "{"
		for index := 0; index < value.NumField(); index++ {
			if index > 0 {
				str += ","
			}
			str += fmt.Sprintf(`"%s":%s`, value.Type().Field(index).Name, simpleEncode(value.Field(index)))
		}
		str += "}"
		return str
	case reflect.Ptr:
		return simpleEncode(value.Elem())
	default:
		return ""
	}
}
