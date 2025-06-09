package main

import (
	"github.com/github.com/AbbyGikunju/todo-app-1/db"
	"github.com/github.com/AbbyGikunju/todo-app-1/router"
)

func main() {
	db.INITPostgresDB()
	router.InitRouter().Run()
	
}