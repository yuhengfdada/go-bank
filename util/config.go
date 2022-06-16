package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"db_driver"`
	DBCredentials       string        `mapstructure:"db_credentials"`
	ServerAddr          string        `mapstructure:"server_addr"`
	TokenSymmetricKey   string        `mapstructure:"token_symmetric_key"`
	AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
}

func ReadConfigFromPath(path string) (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := Config{}
	viper.Unmarshal(&config)
	return &config, nil
}
