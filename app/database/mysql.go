package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DBConn *gorm.DB
	once   sync.Once
)

func ConnectDbMysql() {
	// Update these values with your local MySQL credentials
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", username, password, host, port, dbName)

	db, err := gorm.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}

	// Set up connection pool
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Hour)

	log.Println("Connected to local MySQL")

	DBConn = db
}
