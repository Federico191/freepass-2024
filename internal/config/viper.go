package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("app")
	config.SetConfigType("env")
	config.AddConfigPath("../..")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s", err))
	}
	return config
}
