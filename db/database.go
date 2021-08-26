package db

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func GetConnection() *pg.DB {
	godotenv.Load()
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})
	return db
}
