package database

import (
	"fmt"
	"time"

	"github.com/lolo35/dispatch_manele/database/models"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user string
var password string
var db string
var host string
var port string
var charset string
var dbConn *gorm.DB

func init() {
	user = env.Env("DB_USERNAME")
	password = env.Env("DB_PASSWORD")
	db = env.Env("DB_DATABASE")
	host = env.Env("DB_HOST")
	port = env.Env("DB_PORT")
	charset = env.Env("DB_CHARSET")
}

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, db, charset)
}

func CreateDBConnection() error {
	if dbConn != nil {
		CloseDBConnection(dbConn)
	}
	db, err := gorm.Open(mysql.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		logger.Err(err.Error())
		return err
	}

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbConn = db
	return err
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		logger.Err(err.Error())
		return dbConn, err
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Err(err.Error())
		return dbConn, err
	}

	return dbConn, nil
}

func CloseDBConnection(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		logger.Err(err.Error())
	}
	defer sqlDB.Close()
}

func Migrate() error {
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		logger.Err(connErr.Error())
		return connErr
	}

	err := db.AutoMigrate(&models.Dispatch{}, &models.DeletedDispatches{}, &models.DispatchDescriptions{})

	return err
}
