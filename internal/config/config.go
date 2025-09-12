package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func LocalConfig() {
	Config = viper.New()
	Config.AddConfigPath("./config")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	if err := Config.ReadInConfig(); err != nil {
		panic("load config.yaml failed: " + err.Error())
	}
}
