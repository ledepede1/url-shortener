package deleteurl

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/ledepede1/url-shortener/pkg/db"
	"github.com/ledepede1/url-shortener/pkg/middleware"
)

type ShortUrl struct {
	Url string `json:"shorturl"`
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	if r.Method != http.MethodDelete {
		http.Error(w, "Error", http.StatusMethodNotAllowed)
	}

	var shorturl ShortUrl

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&shorturl)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	db, _ := database.EstablishDBCon()

	_, err = db.Exec("DELETE FROM urls WHERE shorturl=?", shorturl.Url)
	if err != nil {
		log.Printf("Shit error: %s: %v", shorturl.Url, err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL deleted successfully"))
}
