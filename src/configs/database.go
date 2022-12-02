package configs

import (
	models "go-boilerplate/src/models/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.BookEntity{}, &models.UserEntity{})
	if err != nil {
		panic(err)
	}

}
func Init(url string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			//LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
			// todo: depend color
			Colorful: true,
		},
	)

	// url config
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
