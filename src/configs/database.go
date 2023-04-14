package configs

import (
	getLogger "go-server/src/share/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		panic(err)
	}
}
func Init(config Config) *gorm.DB {
	var dbConfig gorm.Config

	if config.DBIsWriteLog == "true" {
		dbConfig.Logger = getLogger.GetLogger().DatabaseLogging
	}
	db, err := gorm.Open(postgres.Open(config.DBUrl), &dbConfig)
	if err != nil {
		panic(err)
	}
	sqlCfg, _ := db.DB()
	maxConn := 100
	sqlCfg.SetMaxOpenConns(maxConn)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
