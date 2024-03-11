package base62

import "testing"

func TestBase62Decode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base62Decode(tt.str); got != tt.want {
				t.Errorf("Base62Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase62Encode(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base62Encode(tt.num); got != tt.want {
				t.Errorf("Base62Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
