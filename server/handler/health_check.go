package handler

import (
	"encoding/json"
	"github.com/UrlShortener/internal/keyprovider"
	"net/http"
	"strconv"
)

//Ping ...
func (handler *urlShortenHandler) Ping(w http.ResponseWriter, r *http.Request) {
	uniqueKey := keyprovider.GetUniqueKeyProvider().GetUniqueKey()
	_ = json.NewEncoder(w).Encode("Pong..." + strconv.FormatInt(uniqueKey, 10))
}
