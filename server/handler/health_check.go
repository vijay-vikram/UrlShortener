package handler

import (
	"encoding/json"
	"net/http"
)

//Ping ...
func (handler *urlShortenHandler) Ping(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Pong...")
}

