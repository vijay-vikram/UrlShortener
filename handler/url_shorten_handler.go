package handler

import (
	"encoding/json"
	"net/http"
)

//URLShortenHandler ...
type URLShortenHandler interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Ping(http.ResponseWriter, *http.Request)
}

type urlShortenHandler struct {
}

// NewURLShortenHandler ...
func NewURLShortenHandler() URLShortenHandler {
	return &urlShortenHandler{}
}

//Get ...
func (handler *urlShortenHandler) Get(w http.ResponseWriter, r *http.Request) {
}

//Create ...
func (handler *urlShortenHandler) Create(w http.ResponseWriter, r *http.Request) {
}

//Ping ...
func (handler *urlShortenHandler) Ping(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Pong...")
}
