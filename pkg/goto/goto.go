package gotoURL

import (
	"net/http"
)

func GotoURL(url string, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
