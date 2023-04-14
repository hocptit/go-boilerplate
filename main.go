package main

import (
	"go-server/src"
	"go-server/src/configs"
	"go-server/src/share/constant"
	getLogger "go-server/src/share/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := configs.LoadConfig(".env")
	logger := getLogger.GetNewLogger(config.AppIsWriteLog).Logging
	logger.Info("Prepare start server..")
	db := configs.Init(config)
	if config.Env == constant.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		configs.Migrate(db)
	}
	app := gin.New()
	src.Router(app, config)
	logger.Infof("Server start on port %s", config.Port)
	app.Run(":" + config.Port).Error()
}
