package server

import (
	"github.com/UrlShortener/server/handler"
	"net/http"
)

func mapUrls() {
	urlShortenHandler := handler.NewURLShortenHandler()
	router.HandleFunc("/{key}", urlShortenHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/shorten", urlShortenHandler.CreateShortUrl).Methods(http.MethodPost)
	router.HandleFunc("/ping", urlShortenHandler.Ping).Methods(http.MethodGet)
}
