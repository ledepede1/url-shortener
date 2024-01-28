package gotoURL

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	database "github.com/ledepede1/url-shortener/pkg/db"
)

func GotoURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shorturl := vars["shorturl"]

	db, _ := database.EstablishDBCon()
	defer db.Close()

	var url string
	db.QueryRow("SELECT url FROM urls WHERE shorturl=?", shorturl).Scan(&url)

	if len(url) >= 1 {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	} else {
		invalidLink := `<style>p {
				font-family: Arial, Helvetica, sans-serif;
				text-align: center;
			}
			h1 {
				font-family: Arial, Helvetica, sans-serif;
				text-align: center;
			}</style>
			
			<h1>404 - Page not found</h1>
			<p>Invalid URL</p>
		`

		fmt.Fprint(w, invalidLink)
	}
}
