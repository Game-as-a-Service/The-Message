package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	DSN := GetDSN(dbUser, dbPwd, dbHost, dbPort, dbDatabase)
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
		return nil
	}
	return db
}

func GetDSN(user string, password string, host string, port string, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}
