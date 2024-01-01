package list

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "github.com/ledepede1/url-shortener/pkg/db"
	"github.com/ledepede1/url-shortener/pkg/middleware"
)

// just having it ready not in use right now
type RowsJSON struct {
	Realurl  string `json:"url"`
	Shorturl string `json:"shorturl"`
}

func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	db, _ := database.EstablishDBCon()
	defer db.Close()

	var rowsJson RowsJSON
	var allUrls []RowsJSON
	rows, _ := db.Query("SELECT * FROM urls")

	for rows.Next() {
		var url, shorturl string

		err := rows.Scan(&url, &shorturl)
		if err != nil {
			log.Fatal(err)
		}

		rowsJson.Realurl = url
		rowsJson.Shorturl = shorturl

		allUrls = append(allUrls, rowsJson)
	}

	jsonUrls, err := json.Marshal(allUrls)
	if err != nil {
		fmt.Println("Couldnt marshal for some shitty reason")
	}

	//fmt.Println(jsonUrls)

	w.Write(jsonUrls)

	//return jsonUrls
}
