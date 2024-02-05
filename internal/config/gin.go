package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewGin(log *logrus.Logger) *gin.Engine {
	app := gin.New()
	gin.DefaultWriter = log.Out
	return app
}
