package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	DBUsername string `mapstructure:"DATABASE_USERNAME"`
	DBPassword string `mapstructure:"DATABASE_PASSWORD"`
	DBHost     string `mapstructure:"DATABASE_HOST"`
	DBPort     string `mapstructure:"DATABASE_PORT"`
	DBName     string `mapstructure:"DATABASE_NAME"`
}

func NewEnv(viper viper.Viper, log *logrus.Logger) *Env {
	env := Env{}
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config:%s", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("environment cannot be loaded:%s", err)
	}

	return &env
}
