package configs

import (
	models "go-boilerplate/src/models/entities"
	getLogger "go-boilerplate/src/shared/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.BookEntity{}, &models.UserEntity{})
	if err != nil {
		panic(err)
	}

}
func Init(url string) *gorm.DB {
	// url config
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		// todo: only for dev
		Logger: getLogger.GetLogger().DatabaseLogging,
	})
	if err != nil {
		panic("Failed to connect database")
	}
	sqlCfg, _ := db.DB()
	sqlCfg.SetMaxOpenConns(100)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
