package main

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/src"
	"go-boilerplate/src/configs"
)

func main() {
	config, _ := configs.LoadConfig(".env")
	db := configs.Init(config.DbUrl)
	// todo: migrate only dev env
	configs.Migrate(db)
	// todo: config mode
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	src.Router(app)
	app.Run(":" + config.Port).Error()
}
