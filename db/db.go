package db

import (
	"errors"
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


//Create todo operation
func CreateTodo (todo *Todo) (*Todo, error) {
	res := DB.Create(&todo)
	if res.Error != nil {
		return nil, res.Error
	}
	return todo, nil
}

//Read all todos operation
func ReadTodos() ([]*Todo, error) {
	var todos []*Todo
	res  := DB.Find(&todos)
	if res.Error != nil {
		return nil, errors.New("no todo list found")
	}
	return todos, nil
}

//Read a todo operation
func ReadTodo(id string) (*Todo,  error) {
	var todo Todo
	res := DB.First(&todo, "id= ?", id)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("todo of id %s not found", id)
	}
	return &todo, nil
}

//Update todo operation
func UpdateTodo(id string, todo *Todo) (*Todo, error) {
	var todoToUpdate Todo
	result := DB.Model(&todoToUpdate).Where("id = ?", id).Updates(todo)
	if result.RowsAffected == 0 {
		return &todoToUpdate, errors.New("todo not updates")
	}
	return todo, nil
}


//Delete Todo Operation
func DeleteTodo(id string) error{
	var deletedTodo Todo
	result := DB.Where("id = ?", id).Delete(&deletedTodo)
	if result.RowsAffected == 0 {
		return errors.New("todo not deleted")
	}
	return nil
}