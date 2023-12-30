package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ledepede1/url-shortener/pkg/config"
	database "github.com/ledepede1/url-shortener/pkg/db"
	gotoURL "github.com/ledepede1/url-shortener/pkg/goto"
)

func main() {
	db, _ := database.EstablishDBCon()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM urls")

	for rows.Next() {
		var url, shorturl string

		err := rows.Scan(&url, &shorturl)
		if err != nil {
			log.Fatal(err)
		}

		http.HandleFunc("/"+shorturl, func(u, s string) func(http.ResponseWriter, *http.Request) {
			return func(w http.ResponseWriter, r *http.Request) {
				db.QueryRow("SELECT url FROM urls WHERE shorturl = ?", shorturl).Scan(&url)
				gotoURL.GotoURL(url)
			}
		}(url, shorturl))

		fmt.Printf("\nCreating URL localhost%s/%s", config.Port, shorturl)
	}

	fmt.Printf("\n\nCreating listener on Port %s", config.Port)
	http.ListenAndServe(config.Port, nil)
}
