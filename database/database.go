package database

import (
	"github.com/cjaewon/echo-gorm-example/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect : Database connect
func Connect() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo_01?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	models.Migrate(db)

	return db
}
