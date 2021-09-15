package handler

import (
	"encoding/json"
	"github.com/UrlShortener/dto"
	"github.com/UrlShortener/internal/url"
	"github.com/UrlShortener/logger"
	"github.com/UrlShortener/utils/resterrors"
	"net/http"
	"strings"
)

func (handler *urlShortenHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var req *dto.CreateShortUrlRequest
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		_ = json.NewEncoder(w).Encode(resterrors.NewBadRequestError(err.Error()))
		return
	}

	if req == nil {
		logger.Handler.Warn(logTag, "Request is empty")
		_ = json.NewEncoder(w).Encode(resterrors.NewBadRequestError("request_empty"))
		return
	}

	if len(strings.TrimSpace(req.Url)) == 0 {
		logger.Handler.Warn(logTag, "Input URL cannot be empty")
		_ = json.NewEncoder(w).Encode(resterrors.NewBadRequestError("Input URL cannot be empty"))
		return
	}

	// TODO:// get shortenEligibleLength from config
	if len(strings.TrimSpace(req.Url)) <= shortenEligibleLength {
		logger.Handler.Warn(logTag, "Url is already short enough")
		_ = json.NewEncoder(w).Encode(resterrors.NewBadRequestError("Url is already short enough"))
		return
	}

	url.CreateShortUrl(req)
}
