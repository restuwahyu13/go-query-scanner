package helper

import "testing"

func TestCheckValidQuery(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  bool
	}{
		{
			name:  "simple valid query",
			query: "param=value",
			want:  true,
		},
		{
			name:  "multiple parameters",
			query: "first=value1&second=value2&third=value3",
			want:  true,
		},
		{
			name:  "empty query",
			query: "",
			want:  true,
		},
		{
			name:  "query with hash",
			query: "param=value#",
			want:  true,
		},
		{
			name:  "invalid query missing value",
			query: "param=&second=value",
			want:  true,
		},
		{
			name:  "invalid format with multiple equals",
			query: "param=value=another",
			want:  true,
		},
		{
			name:  "invalid format with question mark",
			query: "param=value?another=value",
			want:  true,
		},
		{
			name:  "invalid format starting with ampersand",
			query: "&param=value",
			want:  true,
		},
		{
			name:  "invalid format with multiple ampersands",
			query: "param=value&&another=value",
			want:  false,
		},
		{
			name:  "special characters in value",
			query: "param=value+with%20spaces",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidQuery(tt.query); got != tt.want {
				t.Errorf("CheckValidQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
