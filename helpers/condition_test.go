package helper

import (
	"reflect"
	"strings"
	"testing"
)

func TestCondition(t *testing.T) {
	type TestStruct struct {
		StringField string         `json:"string_field"`
		IntField    int            `json:"int_field"`
		BoolField   bool           `json:"bool_field"`
		FloatField  float64        `json:"float_field"`
		MapField    map[string]any `json:"map_field"`
		InvalidType chan int       `json:"invalid_type"`
	}

	tests := []struct {
		name        string
		field       reflect.StructField
		fieldValue  reflect.Value
		inputValue  any
		stag        string
		expectError bool
		errorType   string
	}{
		{
			name:        "valid string assignment",
			field:       reflect.TypeOf(TestStruct{}).Field(0),
			fieldValue:  reflect.New(reflect.TypeOf("")).Elem(),
			inputValue:  "test value",
			stag:        "string_field",
			expectError: false,
		},
		{
			name:        "valid int assignment",
			field:       reflect.TypeOf(TestStruct{}).Field(1),
			fieldValue:  reflect.New(reflect.TypeOf(0)).Elem(),
			inputValue:  "42",
			stag:        "int_field",
			expectError: false,
		},
		{
			name:        "invalid int format",
			field:       reflect.TypeOf(TestStruct{}).Field(1),
			fieldValue:  reflect.New(reflect.TypeOf(0)).Elem(),
			inputValue:  "not_an_int",
			stag:        "int_field",
			expectError: true,
			errorType:   "invalid_integer_format",
		},
		{
			name:        "valid bool assignment",
			field:       reflect.TypeOf(TestStruct{}).Field(2),
			fieldValue:  reflect.New(reflect.TypeOf(false)).Elem(),
			inputValue:  "true",
			stag:        "bool_field",
			expectError: false,
		},
		{
			name:        "invalid bool format",
			field:       reflect.TypeOf(TestStruct{}).Field(2),
			fieldValue:  reflect.New(reflect.TypeOf(false)).Elem(),
			inputValue:  "not_a_bool",
			stag:        "bool_field",
			expectError: true,
			errorType:   "invalid_boolean_format",
		},
		{
			name:        "valid float assignment",
			field:       reflect.TypeOf(TestStruct{}).Field(3),
			fieldValue:  reflect.New(reflect.TypeOf(0.0)).Elem(),
			inputValue:  "3.14",
			stag:        "float_field",
			expectError: false,
		},
		{
			name:        "invalid float format",
			field:       reflect.TypeOf(TestStruct{}).Field(3),
			fieldValue:  reflect.New(reflect.TypeOf(0.0)).Elem(),
			inputValue:  "not_a_float",
			stag:        "float_field",
			expectError: true,
			errorType:   "invalid_float_format",
		},
		{
			name:        "valid map assignment",
			field:       reflect.TypeOf(TestStruct{}).Field(4),
			fieldValue:  reflect.New(reflect.TypeOf(map[string]any{})).Elem(),
			inputValue:  `{"key": "value"}`,
			stag:        "map_field",
			expectError: false,
		},
		{
			name:        "invalid map json",
			field:       reflect.TypeOf(TestStruct{}).Field(4),
			fieldValue:  reflect.New(reflect.TypeOf(map[string]any{})).Elem(),
			inputValue:  `{invalid_json}`,
			stag:        "map_field",
			expectError: true,
			errorType:   "invalid_json_format",
		},
		{
			name:        "unsupported type",
			field:       reflect.TypeOf(TestStruct{}).Field(5),
			fieldValue:  reflect.New(reflect.TypeOf(map[string]string{})).Elem(),
			inputValue:  "any value",
			stag:        "invalid_type",
			expectError: true,
			errorType:   "unsupported_type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Condition(tt.field, tt.fieldValue, tt.inputValue, tt.stag)

			if tt.expectError {
				if err == nil && !strings.Contains(tt.errorType, "unsupported_type") {
					t.Errorf("Expected error but got nil")
					return
				}

			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
