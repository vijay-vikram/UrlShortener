package config

import (
	"encoding/json"
	"fmt"
	"github.com/UrlShortener/logger"
	"io/ioutil"
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
	configFile := "config.json"
	configJSON, err := ioutil.ReadFile(configFile)
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
