/*
database package is responsible for handling the database connection and operations.
*/
package database

import (
	"fmt"
	"time"

	"github.com/lin-snow/dooo/config"
	"github.com/lin-snow/dooo/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects to the database and returns the connection.
func ConnectDB() (*gorm.DB, error) {
	// Load the configuration
	Conf := config.LoadConfig()
	dbtype := Conf.Database.DBtype
	user := Conf.Database.User
	password := Conf.Database.Password
	host := Conf.Database.Host
	port := Conf.Database.Port
	dbname := Conf.Database.DBname

	fmt.Println("@@@"+user, password, host, port, dbname)

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	// fmt.Println(dsn)
	if dbtype == "mysql" {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	}

	return db, nil
}

// MigrateDB migrates the database schema.
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
}
