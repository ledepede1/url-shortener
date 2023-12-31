package create

import (
	"net/http"
	"net/url"
)

// Function to check if the given url has https:// or http:// in it if not it will be added
func checkString(urlSent string) string {

	url, err := url.ParseRequestURI(urlSent)
	if err != nil { // Invalid url
		resp, err := http.Get("http://" + urlSent)
		if err != nil {
			return "http://" + urlSent
		}
		defer resp.Body.Close()

		return "https://" + urlSent
	}

	return url.String()
}
