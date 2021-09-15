package handler

import "net/http"

//URLShortenHandler ...
type URLShortenHandler interface {
	CreateShortUrl(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Ping(http.ResponseWriter, *http.Request)
}

type urlShortenHandler struct {
}

// NewURLShortenHandler ...
func NewURLShortenHandler() URLShortenHandler {
	return &urlShortenHandler{}
}
