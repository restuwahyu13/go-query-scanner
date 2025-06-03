package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func Condition(field reflect.StructField, fieldValue reflect.Value, val interface{}, stag string) error {
	storeValue := fmt.Sprintf("%s", val)

	switch field.Type.Kind() {
	case reflect.String:
		fieldValue.SetString(storeValue)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		toInt, err := strconv.Atoi(storeValue)
		if err != nil {
			return Exception("invalid_integer_format", stag, nil, reflect.StructField{})
		}

		fieldValue.SetInt(int64(toInt))

	case reflect.Bool:
		toBool, err := strconv.ParseBool(storeValue)
		if err != nil {
			return Exception("invalid_boolean_format", stag, nil, reflect.StructField{})
		}

		fieldValue.SetBool(toBool)

	case reflect.Float32, reflect.Float64:
		toFloat, err := strconv.ParseFloat(storeValue, 64)
		if err != nil {
			return Exception("invalid_float_format", stag, nil, reflect.StructField{})
		}

		fieldValue.SetFloat(toFloat)

	case reflect.Map:
		if field.Type != reflect.TypeOf(map[string]interface{}{}) {
			return Exception("invalid_map_format", stag, nil, reflect.StructField{})
		}

		store := make(map[string]interface{})
		storeValueByte := []byte(storeValue)

		if err := json.Unmarshal(storeValueByte, &store); err != nil {
			return Exception("invalid_json_format", stag, nil, reflect.StructField{})
		}

		fieldValue.Set(reflect.ValueOf(store))

	default:
		return Exception("unsupported_type", stag, nil, reflect.StructField{})
	}

	return nil
}
