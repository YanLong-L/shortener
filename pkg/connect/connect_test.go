package connect

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{
			"可正常访问的url",
			"http://www.baidu.com",
			true,
		},
		{
			"可正常访问的url",
			"xxx",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Connect(tt.url)
			assert.Equal(t, got, tt.want)
		})
	}
}
