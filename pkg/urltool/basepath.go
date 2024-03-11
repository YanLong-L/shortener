package urltool

import (
	"net/url"
	"path/filepath"
)

func BasePath(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return filepath.Base(parsedURL.Path)
}
