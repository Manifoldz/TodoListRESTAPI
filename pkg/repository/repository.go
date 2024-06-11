package repository

import (
	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ToDoList interface {
	Create(list entities.ToDoList) (int, error)
	GetAll() ([]entities.ToDoList, error)
	GetById(id int) (entities.ToDoList, error)
	DeleteById(id int) error
	UpdateById(id int, input entities.UpdateListInput) error
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
