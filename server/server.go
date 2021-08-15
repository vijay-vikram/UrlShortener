package server

import (
	"github.com/UrlShortener/logger"
	"github.com/UrlShortener/server/config"
	"github.com/gorilla/mux"

	"net/http"
)

var (
	router = mux.NewRouter()
)

func Serve() {
	config.LoadConfig()
	mapUrls()
	BeforeStart()

	serv := &http.Server{
		Addr:    config.AppConfig.Server.Address + ":" + config.AppConfig.Server.Port,
		Handler: router,
	}
	logger.Handler.Info(logTag, "Starting Application...")
	err := serv.ListenAndServe()
	if err != nil {
		logger.Handler.Fatal(logTag, "Error while starting server %s", err.Error())
	}
}
