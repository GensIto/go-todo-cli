package main

import (
	"fmt"
	"go-todo-api/db"
	"go-todo-api/model"
)

func main() {
	db := db.NewDB()
	db.AutoMigrate(&model.Todo{})
	db.Create(&model.Todo{Title: "Test", Description: "Test description"})
	defer fmt.Println("Migration completed")
}
