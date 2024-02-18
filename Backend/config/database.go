package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := DefaultDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	return db
}

func DefaultDSN() string {
	dsn := BaseDSN()

	val := url.Values{}
	val.Add("parseTime", "true")
	val.Add("loc", "Local")

	dsn = fmt.Sprintf("%s?%s", dsn, val.Encode())
	return dsn
}

func BaseDSN() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	return dsn
}
