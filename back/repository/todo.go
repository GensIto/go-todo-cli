package repository

import (
	"go-todo-api/model"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(todo *model.Todo) (int, error)
	Read() ([]*model.Todo, error)
	Update(id int, todo model.Todo) error
	Delete(id int) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) Create(todo *model.Todo) (int, error) {
	result := tr.db.Create(&todo)
	if result.Error != nil {
		return 0, result.Error
	}
	return todo.ID, nil
}

func (tr *todoRepository) Read() ([]*model.Todo, error) {
	todos := []*model.Todo{}
	result := tr.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (tr *todoRepository) Update(id int, todo model.Todo) error {
	result := tr.db.Save(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *todoRepository) Delete(id int) error {
	result := tr.db.Delete(&model.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
