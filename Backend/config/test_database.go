package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitTestDB() *gorm.DB {
	_, err := LoadEnvPath()
	if err != nil {
		return nil
	}

	db := InitDB()
	return db
}

func LoadEnvPath() (string, error) {
	envPath, err := getEnvPath()

	err = godotenv.Load(envPath)
	if err != nil {
		return "", err
	}

	return envPath, err
}

func getEnvPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	envPath := cwd + "/../../.env"
	return envPath, err
}
