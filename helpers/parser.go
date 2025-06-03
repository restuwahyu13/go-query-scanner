package helper

import (
	"net/url"
	"reflect"
)

func Parser(query string, dest interface{}) (reflect.Value, map[string]interface{}, error) {
	if ok := CheckValidQuery(query); !ok {
		return reflect.Value{}, nil, Exception("invalid_query_string", "", nil, reflect.StructField{})
	}

	parsed, err := url.ParseQuery(query)
	if err != nil {
		return reflect.Value{}, nil, Exception("invalid_query_string", "", nil, reflect.StructField{})
	}

	valueof := reflect.ValueOf(dest)
	if valueof.Kind() != reflect.Pointer {
		return reflect.Value{}, nil, Exception("invalid_struct_pointer", "", nil, reflect.StructField{})
	}

	if valueof.Elem().Kind() != reflect.Struct {
		return reflect.Value{}, nil, Exception("invalid_struct", "", nil, reflect.StructField{})
	}

	store := make(map[string]interface{}, 1)

	for key, values := range parsed {
		if len(values) > 0 {
			unescaped, err := url.QueryUnescape(values[0])
			if err != nil {
				return reflect.Value{}, nil, Exception("invalid_escape_key", key, nil, reflect.StructField{})
			}

			store[key] = unescaped
		}
	}

	return valueof, store, nil
}
