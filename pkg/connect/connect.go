package connect

import (
	"net/http"
	"time"
)

func Connect(url string) bool {
	timeout := 5 * time.Second

	if isUrlAccessible(url, timeout) {
		return true
	} else {
		return false
	}
}

func isUrlAccessible(url string, timeout time.Duration) bool {
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	}

	return false
}
