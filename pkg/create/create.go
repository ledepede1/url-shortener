// Not in use rightnow
package create

import (
	"math/rand"

	database "github.com/ledepede1/url-shortener/pkg/db"
)

func generateUrl() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	shortenUrl := make([]byte, 10)
	for i := range shortenUrl {
		shortenUrl[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(shortenUrl)
}

func CreateNewUrl(url string) {
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
	}
}
