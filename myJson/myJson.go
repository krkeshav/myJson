package myJson

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/krkeshav/myJson/encrypt"
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
		strValue := value.String()
		strValue = strings.ReplaceAll(strValue, "\"", "\\\"") // This is also hackish and not recommended
		sb := strings.Builder{}
		sb.WriteString(`"`)
		sb.WriteString(strValue)
		sb.WriteString(`"`)
		// The below commented ones are probably not required since ideally
		// the default json library preserves everything and not cleans it
		// strValue = strings.ReplaceAll(strValue, "\n", "")
		// strValue = strings.ReplaceAll(strValue, "\t", "")
		return sb.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal := value.Int()
		strIntVal := strconv.FormatInt(intVal, 10)
		return strIntVal
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal := value.Uint()
		strUintVal := strconv.FormatUint(uintVal, 10)
		return strUintVal
	case reflect.Float32, reflect.Float64:
		floatVal := value.Float()
		strFloatVal := strconv.FormatFloat(floatVal, 'f', -1, 64)
		return strFloatVal
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.Slice, reflect.Array:
		sb := strings.Builder{}
		sb.WriteString("[")
		for index := 0; index < value.Len(); index++ {
			if index > 0 {
				sb.WriteString(",")
			}
			strVal := simpleEncode(value.Index(index))
			sb.WriteString(strVal)
		}
		sb.WriteString("]")
		return sb.String()
	case reflect.Map:
		sb := strings.Builder{}
		sb.WriteString("{")
		for index, key := range value.MapKeys() {
			if index > 0 {
				sb.WriteString(",")
			}
			keyValue := getMapKey(key)
			sb.WriteString(`"`)
			sb.WriteString(keyValue)
			sb.WriteString(`"`)
			sb.WriteString(":")
			mpValue := simpleEncode(value.MapIndex(key))
			sb.WriteString(mpValue)
		}
		sb.WriteString("}")
		return sb.String()
	case reflect.Struct:
		sb := strings.Builder{}
		sb.WriteString("{")
		for index := 0; index < value.NumField(); index++ {
			valueType := value.Type().Field(index)
			jsonTagName := valueType.Tag.Get("json")
			jsonTagValues := strings.Split(jsonTagName, ",")
			if len(jsonTagValues) > 0 {
				jsonTagName = jsonTagValues[0]
				if jsonTagName == "-" {
					continue
				}
			}
			isOmitEmpty := false
			if len(jsonTagValues) > 1 {
				isOmitEmpty = jsonTagValues[1] == "omitempty"
			}
			if jsonTagName == "" {
				jsonTagName = valueType.Name
			}
			tag := valueType.Tag.Get("encrypt")
			encryptionRequired := tag == "true"
			underlyingFieldValueStr := simpleEncode(value.Field(index))
			// removing below check because what if there is empty map or slice
			// if underlyingFieldValueStr == "{}" || underlyingFieldValueStr == "[]" {
			// 	underlyingFieldValueStr = "null" // this is for same behavior as encoding/json package
			// }
			if isOmitEmpty && (underlyingFieldValueStr == "null") {
				continue
			}
			if encryptionRequired {
				underlyingFieldValueStr = encrypt.Encrypt(underlyingFieldValueStr)
			}
			if index > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(`"`)
			sb.WriteString(jsonTagName)
			sb.WriteString(`"`)
			sb.WriteString(":")
			sb.WriteString(underlyingFieldValueStr)
		}
		sb.WriteString("}")
		return sb.String()
	case reflect.Ptr:
		return simpleEncode(value.Elem())
	default:
		return "null"
	}
}

func getMapKey(keyValue reflect.Value) string {
	keyValueStr := keyValue.String()
	switch keyValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		keyValueStr = fmt.Sprintf("%d", keyValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		keyValueStr = fmt.Sprintf("%d", keyValue.Uint())
	case reflect.Float32, reflect.Float64:
		keyValueStr = fmt.Sprintf("%f", keyValue.Float())
	}
	return keyValueStr
}
