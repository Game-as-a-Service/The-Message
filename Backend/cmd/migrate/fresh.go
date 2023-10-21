//go:build migrate

package main

import (
	"database/sql"
	"fmt"
	"github.com/Game-as-a-Service/The-Message/config"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	db := config.InitDB()
	tableNames := GetTableNames(db)
	TruncateTables(db, tableNames)
}

func GetTableNames(db *gorm.DB) []string {
	var tableNames []string
	var rows *sql.Rows
	var err error

	rows, err = db.Raw("SHOW TABLES").Rows()
	if err != nil {
		fmt.Println("Error fetching table names:", err)
		return tableNames
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		tableNames = append(tableNames, tableName)
	}

	return tableNames
}

func TruncateTables(db *gorm.DB, tableNames []string) {
	for _, tableName := range tableNames {
		db.Exec(fmt.Sprintf("SET FOREIGN_KEY_CHECKS = 0"))
		db.Exec(fmt.Sprintf("DROP TABLE `%s`", tableName))
		db.Exec(fmt.Sprintf("SET FOREIGN_KEY_CHECKS = 1"))
	}
	fmt.Println("All specified tables truncated successfully.")
}
