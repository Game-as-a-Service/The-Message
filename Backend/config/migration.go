package config

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/joho/godotenv/autoload"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMigration() (*migrate.Migrate, error) {
	dir, _ := os.Getwd()
	sourceURL := "file://" + dir + "/database/migrations"

	dsn := BaseDSN()

	val := url.Values{}
	val.Add("multiStatements", "true")

	dsn = fmt.Sprintf("%s?%s", dsn, val.Encode())

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
	m, err := NewMigration()
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
