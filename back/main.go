package main

import (
	"go-todo-api/controller"
	"go-todo-api/db"
	"go-todo-api/repository"
	"go-todo-api/usecase"
	"go-todo-api/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	validate "github.com/go-playground/validator"
)

func main() {
	e := echo.New()
	e.Validator = &validator.CustomValidator{Validator: validate.New()}
	e.Use(middleware.Logger())

	db := db.NewDB()
	tr := repository.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)
	tc := controller.NewTodoController(tu)

	e.POST("/todos", tc.Create)
	e.GET("/todos", tc.Read)
	e.PUT("/todos/:id", tc.Update)
	e.DELETE("/todos/:id", tc.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
