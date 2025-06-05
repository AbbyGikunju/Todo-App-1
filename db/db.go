package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Todo struct{
	gorm.Model 
	Title		string
	Description	string
	Status		bool

}

func INITPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	var(
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbName = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")

	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s sslmode=disable",
	host,
	port,
	dbUser,
	dbName,
	password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	DB.AutoMigrate(Todo{})
}