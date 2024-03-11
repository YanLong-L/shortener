package urltool

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		wantRes string
	}{
		{
			name:    "case1",
			rawURL:  "https://www.example.com/path/res",
			wantRes: "res",
		},
		{
			name:    "case2",
			rawURL:  "https://www.example.com/path?name=123&pwd=456",
			wantRes: "path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GetBasePath(tt.rawURL)
			assert.Equal(t, res, tt.wantRes)
		})
	}
}
