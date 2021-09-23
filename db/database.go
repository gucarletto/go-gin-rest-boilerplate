package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gucarletto/go-gin-rest-boilerplate/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Manager *gorm.DB

func init() {
	if Manager == nil {
		godotenv.Load()

		dsn := fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_PORT"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		Manager = db
	}
}

func GetConnectionManager() *gorm.DB {
	return Manager
}

func Migrate() {
	Manager.AutoMigrate(&models.User{})
	Manager.AutoMigrate(&models.Transaction{})
}
