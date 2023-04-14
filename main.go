package main

import (
	"server-go/src"
	"server-go/src/configs"
	"server-go/src/share/constant"
	getLogger "server-go/src/share/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := configs.LoadConfig(".env")
	configs.Setup(config.DBUrl, config.DatabaseName)

	logger := getLogger.GetNewLogger(config.AppIsWriteLog).Logging
	logger.Info("Prepare start server..")
	if config.Env == constant.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	src.Router(app, config)
	logger.Infof("Server start on port %s", config.Port)
	app.Run(":" + config.Port).Error()
}
