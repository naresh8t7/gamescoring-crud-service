package config

import "github.com/spf13/viper"

type Config struct {
	HttpPort string `mapstructure:"HTTP_PORT"`
	Store    string `mapstructure:"STORE"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
