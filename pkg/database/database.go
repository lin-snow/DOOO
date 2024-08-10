/*
database package is responsible for handling the database connection and operations.
*/
package database

import (
	"time"

	"github.com/lin-snow/dooo/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects to the database and returns the connection.
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:041107@L1nSn0w@tcp(127.0.0.1:3306)/dooo?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name                                                                                  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

// MigrateDB migrates the database schema.
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
	db.AutoMigrate(&model.User{})
}
