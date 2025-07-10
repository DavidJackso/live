package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPServer `mapstructure:"http_server"`
	DBConfig   `mapstructure:"db"`
}

type HTTPServer struct {
	Address      string        `mapstructure:"address"`
	ReadTimeout  time.Duration `mapstructure:"read_t"`
	WriteTimeout time.Duration `mapstructure:"write_t"`
}

type DBConfig struct {
	Address  string `mapstructure:"address"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func InitializationConfig() (*Config, error) {
	viper.SetConfigFile("../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("error reading configuration")
		return &Config{
			HTTPServer: HTTPServer{},
			DBConfig:   DBConfig{},
		}, err
	}
	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logrus.WithError(err).Error("failed unmarshal config")
		return &Config{}, err
	}

	return &cfg, nil
}
