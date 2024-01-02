package create

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	database "github.com/ledepede1/url-shortener/pkg/db"
	"github.com/ledepede1/url-shortener/pkg/middleware"
)

type Url struct {
	Url string `json:"url"`
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var url Url

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		log.Fatal("Error in handling this shit", err)
	}

	urlChecked, isValid := checkUrl(url.Url)
	if !isValid {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	shortUrl := createNewUrl(urlChecked)

	// Send a JSON response back to the client
	response := map[string]interface{}{"result": shortUrl}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResponse)

	// Creating the newly created handler
	CreateNewHandler(shortUrl, urlChecked)
}

func generateUrl() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	shortenUrl := make([]byte, 10)
	for i := range shortenUrl {
		shortenUrl[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(shortenUrl)
}

func createNewUrl(url string) string {
	if len(url) >= 1 {
		shortenedUrl := generateUrl()
		var fetchedUrl string

		db, _ := database.EstablishDBCon()
		defer db.Close()

		db.QueryRow("SELECT shorturl FROM urls WHERE shorturl = ?", shortenedUrl).Scan(&fetchedUrl) // Check if the generated shorten url already exist

		for i := shortenedUrl; i == fetchedUrl; {
			shortenedUrl = generateUrl()

			db.QueryRow("").Scan(&fetchedUrl)

			if fetchedUrl != shortenedUrl {
				break
			}
		}

		db.Query("INSERT INTO urls (url, shorturl) VALUES (?, ?)", url, shortenedUrl)

		return shortenedUrl
	}

	return "Error"
}
