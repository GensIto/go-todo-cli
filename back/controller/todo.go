package controller

import (
	"go-todo-api/model"
	"go-todo-api/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITodoController interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

func (tc *todoController) Create(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	id, err := tc.tu.Create(&todo)
	if err != nil {
		return c.JSON(500, err)
	}
	todo.ID = id
	return c.JSON(200, todo)
}

func (tc *todoController) Read(c echo.Context) error {
	todos, err := tc.tu.Read()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, todos)
}

func (tc *todoController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	todo := model.Todo{ID: id}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(400, err)
	}

	err = tc.tu.Update(id, todo)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, todo)
}

func (tc *todoController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	err = tc.tu.Delete(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "deleted successfully")
}
