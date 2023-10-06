package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	dbHost := "127.0.0.1"   //os.Getenv("DB_HOST")
	dbDatabase := "message" //os.Getenv("DB_DATABASE")
	dbUser := "root"        //os.Getenv("DB_USER")
	dbPwd := ""             //os.Getenv("DB_PASSWORD")
	dbPort := "3306"        //os.Getenv("DB_PORT")

	var err error

	db, err := gorm.Open(mysql.Open(GetDSN(dbUser, dbPwd, dbHost, dbPort, dbDatabase)), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database: %v", err)
	}
	return db
}

func GetDSN(user string, password string, host string, port string, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
}
