package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ledepede1/url-shortener/pkg/config"
	"github.com/ledepede1/url-shortener/pkg/create"
	deleteurl "github.com/ledepede1/url-shortener/pkg/delete"
	"github.com/ledepede1/url-shortener/pkg/list"

	"github.com/ledepede1/url-shortener/pkg/db"
	gotoURL "github.com/ledepede1/url-shortener/pkg/goto"
)

func main() {
	CreateHandlers()

	http.HandleFunc("/backend/delete", func(w http.ResponseWriter, r *http.Request) {
		deleteurl.DeleteUrl(w, r)
	})
	http.HandleFunc("/backend/getlist", list.GetAllUrls)

	// Adding the create handler
	http.HandleFunc("/backend/create", create.CreateShortUrl)

	fmt.Printf("\n\nCreating listener on Port %s", config.Port)
	http.ListenAndServe(config.Port, nil)
}

func CreateHandlers() {
	db, _ := db.EstablishDBCon()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM urls")

	for rows.Next() {
		var url, shorturl string

		err := rows.Scan(&url, &shorturl)
		if err != nil {
			log.Fatal(err)
		}

		_, err = http.Get("localhost:8080/" + shorturl)
		if err != nil {
			http.HandleFunc("/"+shorturl, func(w http.ResponseWriter, r *http.Request) {
				gotoURL.GotoURL(url, w, r)
			})

			fmt.Printf("\nCreating URL localhost%s/%s", config.Port, shorturl)
		}
	}
}
