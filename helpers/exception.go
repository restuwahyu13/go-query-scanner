package helper

import (
	"errors"
	"fmt"
	"reflect"
)

func Exception(key, format string, stag interface{}, field reflect.StructField) error {
	msg := make(map[string]error)

	msg["invalid_query_string"] = errors.New("query must be a valid query string")
	msg["invalid_struct_pointer"] = errors.New("dest must be a struct pointer")
	msg["invalid_struct"] = errors.New("dest must be a struct")
	msg["invalid_escape_key"] = fmt.Errorf("unsupported value for a key: %s", format)
	msg["invalid_tag_format"] = fmt.Errorf("field %s must have a query tag", format)

	msg["invalid_integer_format"] = fmt.Errorf("field %s: invalid integer value", format)
	msg["invalid_boolean_format"] = fmt.Errorf("field %s: invalid boolean value", format)
	msg["invalid_float_format"] = fmt.Errorf("field %s: invalid float value", format)
	msg["invalid_json_format"] = fmt.Errorf("field %s: invalid json value", format)
	msg["invalid_map_format"] = fmt.Errorf("field %s: must be a set to map[string]any", format)

	if stag != nil && !reflect.DeepEqual(field, reflect.StructField{}) {
		msg["unsupported_type"] = fmt.Errorf("unsupported type for field %s: %s", stag, field.Type.Kind())
	}

	return msg[key]
}
