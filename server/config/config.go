package config

import (
	"fmt"
	"github.com/UrlShortener/logger"
	"io/ioutil"

	"gitlab.myteksi.net/gophers/go/food/common/encoding/json"
)

var (
	// AppConfig makes app config available globally
	AppConfig *config
)

type config struct {
	Server    server    `json:"server"`
	LogConfig LogConfig `json:"log_config"`
}

// LoadConfig ...
func LoadConfig() {

	// TODO:// get config file path from env var instead of fixed path
	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("Error while loading configs %s", err.Error()))
	}

	err = json.Unmarshal(configJSON, &AppConfig)
	if err != nil {
		panic(fmt.Sprintf("Error while unmarshaling configs %v", err))
	}

	logger.Initialize(AppConfig.LogConfig.LogLevel)
	logger.Handler.Info(logTag, "Config Initialized")
}
