package helper

import (
	"reflect"
	"testing"
)

func TestException(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		format string
		stag   any
		field  reflect.StructField
		want   string
	}{
		{
			name:   "invalid query string error",
			key:    "invalid_query_string",
			format: "",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "query must be a valid query string",
		},
		{
			name:   "invalid struct pointer error",
			key:    "invalid_struct_pointer",
			format: "",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "dest must be a struct pointer",
		},
		{
			name:   "invalid struct error",
			key:    "invalid_struct",
			format: "",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "dest must be a struct",
		},
		{
			name:   "invalid escape key error",
			key:    "invalid_escape_key",
			format: "test%2Z",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "unsupported value for a key: test%2Z",
		},
		{
			name:   "invalid tag format error",
			key:    "invalid_tag_format",
			format: "UserName",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "field UserName must have a query tag",
		},
		{
			name:   "invalid integer format error",
			key:    "invalid_integer_format",
			format: "Age",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "field Age: invalid integer value",
		},
		{
			name:   "invalid boolean format error",
			key:    "invalid_boolean_format",
			format: "IsActive",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "field IsActive: invalid boolean value",
		},
		{
			name:   "invalid float format error",
			key:    "invalid_float_format",
			format: "Price",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "field Price: invalid float value",
		},
		{
			name:   "invalid json format error",
			key:    "invalid_json_format",
			format: "Config",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "field Config: invalid json value",
		},
		{
			name:   "invalid map format error",
			key:    "invalid_map_format",
			format: "Metadata",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "field Metadata: must be a set to map[string]any",
		},
		{
			name:   "unsupported type error",
			key:    "unsupported_type",
			format: "",
			stag:   "email",
			field:  reflect.StructField{Name: "Email", Type: reflect.TypeOf(complex128(0))},
			want:   "unsupported type for field email: complex128",
		},
		{
			name:   "non-existent error key",
			key:    "non_existent_error",
			format: "",
			stag:   nil,
			field:  reflect.StructField{},
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Exception(tt.key, tt.format, tt.stag, tt.field)
			if got == nil && tt.want != "" {
				t.Errorf("Exception() = nil, want %v", tt.want)
			}
			if got != nil && got.Error() != tt.want {
				t.Errorf("Exception() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}
