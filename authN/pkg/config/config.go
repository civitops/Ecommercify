package config

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthNConfig struct {
	PORT     string `mapstructure:"PORT"`
	Mode     string `mapstructure:"MODE"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
	Encoding string `mapstructure:"ENCODING"`
	// DatabseURI string `mapstructure:"DATABASE_URI"`
	ArgonPassword string `mapstructure:"ARGON_PASSWORD"`
}

var defaultsValue = map[string]string{
	"PORT":         "6969",
	"MODE":         Development,
	"DATABASE_URI": "postgres://postgres:mypass@localhost:5432/ecommercify",
}

func LoadConfig(path string) (*AuthNConfig, error) {
	// "" -> loads timezone as UTC:
	loc, err := time.LoadLocation("")
	if err != nil {
		return nil, err
	}

	time.Local = loc
	//Checks the defaults Value map and sets the default
	for key, val := range defaultsValue {
		viper.SetDefault(key, val)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg AuthNConfig

	err = viper.Unmarshal(&cfg)

	if cfg.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	return &cfg, err
}
