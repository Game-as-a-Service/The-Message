package config

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

func NewMigration(dsn string, sourceURL string) (*migrate.Migrate, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		"mysql",
		driver,
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func RunRefresh() {
	sourceURL := GetSourceURL()

	dsn := BaseDSN()
	val := url.Values{}
	val.Add("multiStatements", "true")
	dsn = fmt.Sprintf("%s?%s", dsn, val.Encode())

	m, err := NewMigration(dsn, sourceURL)
	if err != nil {
		panic(err)
	}

	err = m.Down()
	if err != nil {
		if err.Error() == "no change" {
		} else {
			panic(err)
		}
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}
}

func GetSourceURL() string {
	dir, _ := os.Getwd()
	dir = strings.SplitAfter(dir, "Backend")[0]
	dir = strings.ReplaceAll(dir, "\\", "/")

	sourceURL := "file://" + dir + "/database/migrations"

	return sourceURL
}
