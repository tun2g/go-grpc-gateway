package config

import (
	"app/src/lib/logger"

	"github.com/spf13/viper"
)

type Config struct {
	// Application
	AppPort int `mapstructure:"APP_PORT"`

	// Services
	AuthServiceUri string `mapstructure:"AUTH_SERVICE_URI"`
}

var log = logger.NewLogger("AppConfiguration")

func loadConfiguration(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Load config failed %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("Load config failed %s", err)
	}

	return config
}

var AppConfiguration = loadConfiguration(".")
