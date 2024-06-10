package repository

import (
	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ToDoList interface {
	Create(list entities.ToDoList) (int, error)
}
type ToDoItem interface{}

type Repository struct {
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ToDoList: NewTodoListPostgres(db),
	}
}
