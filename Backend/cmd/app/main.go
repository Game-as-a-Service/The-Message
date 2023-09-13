package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Id   int    `gorm:"primaryKey"`
	Name string `json:"name"`
}

type PetRepo struct {
	db *sql.DB
}

func NewPetRepository(db *sql.DB) *PetRepo {
	return &PetRepo{
		db: db,
	}
}

func (p *PetRepo) GetPetById(ctx context.Context, id string) (*Pet, error) {
	row := p.db.QueryRow("SELECT * FROM pet WHERE id = " + id)
	var pet Pet
	err := row.Scan(&pet.Id, &pet.Name)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %v\n", err)
	}
}

func main() {
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPwd := viper.GetString("DB_PASSWORD")
	dbPort := viper.GetString("DB_PORT")

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbDatabase)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.Table("pets").AutoMigrate(&Pet{})

	engine := gin.Default()

	engine.GET("/v2/pet/:petId", func(c *gin.Context) {
		petId, _ := strconv.Atoi(c.Param("petId"))

		pet := &Pet{}

		db.Table("pets").Create(&Pet{Id: petId, Name: "Jack"})

		// db.Table("pets").Where("id = ?", petId).Select("Id", "Name").Find(pet)
		db.Table("pets").Select("Id", "Name").First(pet, "id = ?", petId)

		fmt.Println(pet)

		c.JSON(http.StatusOK, Pet{
			Id:   pet.Id,
			Name: pet.Name,
		})
	})

	engine.Run(":8080")
}
