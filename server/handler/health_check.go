package handler

import (
	"encoding/json"
	"github.com/UrlShortener/storage/redis"
	"net/http"
	"strconv"
)

//Ping ...
func (handler *urlShortenHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ab := redis.GetRedisStorage().Set()
	_ = json.NewEncoder(w).Encode("Pong..." + strconv.FormatInt(ab, 10))
}
