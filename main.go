package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ledepede1/url-shortener/pkg/config"
	"github.com/ledepede1/url-shortener/pkg/create"
	deleteurl "github.com/ledepede1/url-shortener/pkg/delete"
	"github.com/ledepede1/url-shortener/pkg/list"

	gotoURL "github.com/ledepede1/url-shortener/pkg/goto"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/backend/delete", func(w http.ResponseWriter, r *http.Request) {
		deleteurl.DeleteUrl(w, r)
	})
	r.HandleFunc("/backend/getlist", list.GetAllUrls)

	// Adding the create handler
	r.HandleFunc("/backend/create", create.CreateShortUrl)

	// Making the Shorturl handler
	r.HandleFunc("/{shorturl}", gotoURL.GotoURL)

	fmt.Printf("\n\nCreating listener on Port %s", config.Port)
	server := http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
