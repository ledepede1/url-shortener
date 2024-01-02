package create

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/bombsimon/tld-validator"
	urlParser "github.com/jpillora/go-tld"
)

// Checking the url and assigning transfer proctol if it needs it
func checkUrl(urlSent string) (string, bool) {

	var transferProtocol string

	urlParsed, err := urlParser.Parse("https://" + urlSent) // only reason for the https is because else it will not work for some shitty reason
	if err != nil {
		return "error", false
	}

	isTLDValid := tld.IsValid(urlParsed.TLD)
	if !isTLDValid {
		fmt.Println("Invalid TLD:", urlParsed.TLD)
		return "error", false
	}

	_, err = url.ParseRequestURI(urlSent)
	if err != nil {
		resp, err := http.Get("http://" + urlSent)
		if err != nil {
			transferProtocol = "http://"
		} else {
			defer resp.Body.Close()
			transferProtocol = "https://"
		}
	}

	finalURL := transferProtocol + urlSent

	return finalURL, true
}
