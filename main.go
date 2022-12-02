package main

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/src"
	"go-boilerplate/src/configs"
	getLogger "go-boilerplate/src/shared/logger"
)

func main() {
	config, _ := configs.LoadConfig(".env")
	getLogger.GetNewLogger()
	db := configs.Init(config.DbUrl)
	// todo: migrate only dev env
	configs.Migrate(db)
	// todo: config mode
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	src.Router(app)
	app.Run(":" + config.Port).Error()
}
