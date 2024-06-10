package repository

import (
	"fmt"

	"github.com/Manifoldz/TodoListRESTAPI/internal/entities"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(list entities.ToDoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", toDoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll() ([]entities.ToDoList, error) {
	var list []entities.ToDoList
	query := fmt.Sprintf("SELECT * FROM %s", toDoListsTable)
	err := r.db.Select(&list, query)

	return list, err
}

func (r *TodoListPostgres) GetById(id int) (entities.ToDoList, error) {
	var list entities.ToDoList
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", toDoListsTable)
	err := r.db.Get(&list, query, id)

	return list, err

}
