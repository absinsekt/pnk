package strings

import (
	"fmt"
	"net/url"
	"path"
)

// GetPermalink todo
func GetPermalink(baseURL string, relLink string) string {
	parsedURL, _ := url.Parse(baseURL)
	return fmt.Sprintf("%s://%s", parsedURL.Scheme, path.Join(parsedURL.Host, relLink))
}
