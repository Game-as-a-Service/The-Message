// database.go

package main

import (
	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root_password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	// 自動創建資料表
	db.AutoMigrate(&Todo{})

	return db, nil
}
