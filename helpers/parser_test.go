package helper

import "testing"

func TestParser(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Email string
	}

	tests := []struct {
		name    string
		query   string
		dest    any
		wantErr bool
	}{
		{
			name:    "non-pointer destination",
			query:   "name=test",
			dest:    TestStruct{},
			wantErr: true,
		},
		{
			name:    "non-struct pointer destination",
			query:   "name=test",
			dest:    new(string),
			wantErr: true,
		},
		{
			name:    "invalid percent encoding",
			query:   "name=test%ZZ",
			dest:    &TestStruct{},
			wantErr: true,
		},
		{
			name:    "empty values in query",
			query:   "name=&age=&email=",
			dest:    &TestStruct{},
			wantErr: false,
		},
		{
			name:    "multiple values for same key",
			query:   "name=test1&name=test2",
			dest:    &TestStruct{},
			wantErr: false,
		},
		{
			name:    "special characters in key",
			query:   "user%20name=test",
			dest:    &TestStruct{},
			wantErr: false,
		},
		{
			name:    "nil destination",
			query:   "name=test",
			dest:    nil,
			wantErr: true,
		},
		{
			name:    "unicode characters in value",
			query:   "name=测试",
			dest:    &TestStruct{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := Parser(tt.query, tt.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
