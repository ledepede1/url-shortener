package create

import (
	"net/http"

	database "github.com/ledepede1/url-shortener/pkg/db"
	gotoURL "github.com/ledepede1/url-shortener/pkg/goto"
)

func CreateNewHandler(shortUrl string, urlChecked string) {
	db, _ := database.EstablishDBCon()
	defer db.Close()

	http.HandleFunc("/"+shortUrl, func(w http.ResponseWriter, r *http.Request) {
		gotoURL.GotoURL(urlChecked, w, r)
	})
}
