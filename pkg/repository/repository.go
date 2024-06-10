package repository

import "github.com/jmoiron/sqlx"

type ToDoList interface{}
type ToDoItem interface{}

type Repository struct {
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
