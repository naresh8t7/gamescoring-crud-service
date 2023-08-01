package config

import "github.com/spf13/viper"

type Config struct {
	HttpPort   string `mapstructure:"HTTP_PORT"`
	Store      string `mapstructure:"STORE"`
	DbHost     string `mapstructure:"LOCAL_DB_HOST"`
	DbName     string `mapstructure:"LOCAL_DB_NAME"`
	DbPort     string `mapstructure:"LOCAL_DB_PORT"`
	DbUser     string `mapstructure:"LOCAL_DB_USER"`
	DbPassword string `mapstructure:"LOCAL_DB_PASSWORD"`
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
