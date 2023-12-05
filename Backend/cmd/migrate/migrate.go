//go:build migrate

package main

import (
	"fmt"
	"github.com/Game-as-a-Service/The-Message/config"
	"net/url"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	sourceURL := "file://" + dir + "/database/migrations"

	dsn := config.BaseDSN()
	val := url.Values{}
	val.Add("multiStatements", "true")
	dsn = fmt.Sprintf("%s?%s", dsn, val.Encode())

	m, err := config.NewMigration(dsn, sourceURL)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("no change")
			return
		}
		panic(err)
	}
}
