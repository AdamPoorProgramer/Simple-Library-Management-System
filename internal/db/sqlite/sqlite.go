package sqlite

import (
	"LIBRARY-API-SERVER/configs"
	"LIBRARY-API-SERVER/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteOrPanic(log *logger.Logger) *gorm.DB {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return nil
		// panic(err) // Uncomment for debugging purposes only. In production, replace with log.Fatal() instead.
	}
	db, err := gorm.Open(sqlite.Open(config.Database.Name), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
