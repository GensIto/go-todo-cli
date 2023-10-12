package usecase

import (
	"go-todo-api/model"
	"go-todo-api/repository"
)

type ITodoUsecase interface {
	Create(todo *model.Todo) (int, error)
	Read() ([]*model.Todo, error)
	Update(id int, todo model.Todo) error
	Delete(id int) error
}

type todoUsecase struct {
	tr repository.ITodoRepository
}

func NewTodoUsecase(tr repository.ITodoRepository) ITodoUsecase {
	return &todoUsecase{tr}
}

func (tu *todoUsecase) Create(todo *model.Todo) (int, error) {
	id, err := tu.tr.Create(todo)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (tu *todoUsecase) Read() ([]*model.Todo, error) {
	todos, err := tu.tr.Read()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tu *todoUsecase) Update(id int, todo model.Todo) error {
	err := tu.tr.Update(id, todo)
	if err != nil {
		return err
	}
	return nil
}

func (tu *todoUsecase) Delete(id int) error {
	err := tu.tr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
