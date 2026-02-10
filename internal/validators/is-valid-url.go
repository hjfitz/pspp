package validators

import "net/url"

func IsValidUrl(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	return err == nil && parsed.Scheme != "" && parsed.Host != ""
}
