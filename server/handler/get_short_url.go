package handler

import (
	"encoding/json"
	"github.com/UrlShortener/storage/redis"
	"github.com/UrlShortener/utils/resterrors"
	"github.com/gorilla/mux"
	"net/http"
)

//Get ...
func (handler *urlShortenHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrlKey := vars["key"]
	url, err := redis.GetRedisStorage().GetByShortUrl(shortUrlKey)
	if err != nil {
		_ = json.NewEncoder(w).Encode(resterrors.NewInternalServerError(err.Error()))
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}
